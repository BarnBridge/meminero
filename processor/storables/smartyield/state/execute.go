package state

import (
	"context"
	"math"
	"math/big"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
)

func (s *Storable) Execute(ctx context.Context) error {
	syAbi := *ethtypes.SmartYield.ABI
	controllerAbi := *ethtypes.SmartYieldCompoundController.ABI

	wg, ctx1 := errgroup.WithContext(ctx)

	for _, p := range s.state.SmartYield.Pools {
		if s.block.Number < p.StartAtBlock {
			continue
		}

		p := p

		wg.Go(func() error {
			var mu = new(sync.Mutex)

			poolState := &State{
				PoolAddress: p.PoolAddress,
			}

			var underlyingTotal, underlyingJuniors, jtokenPrice *big.Int
			var abond Abond

			subwg, _ := errgroup.WithContext(ctx1)
			subwg.Go(func() error {
				err := eth.CallContractFunction(syAbi, p.PoolAddress, "underlyingTotal", []interface{}{}, &underlyingTotal, s.block.Number)()
				if err != nil && strings.Contains(err.Error(), "execution reverted") {
					underlyingTotal = big.NewInt(0)
					return nil
				}

				return err
			})
			subwg.Go(func() error {
				err := eth.CallContractFunction(syAbi, p.PoolAddress, "underlyingJuniors", []interface{}{}, &underlyingJuniors, s.block.Number)()
				if err != nil && strings.Contains(err.Error(), "execution reverted") {
					underlyingJuniors = big.NewInt(0)
					return nil
				}

				return err
			})
			subwg.Go(func() error {
				err := eth.CallContractFunction(syAbi, p.PoolAddress, "price", []interface{}{}, &jtokenPrice, s.block.Number)()
				if err != nil && strings.Contains(err.Error(), "execution reverted") {
					jtokenPrice = big.NewInt(0)
					return nil
				}

				return err
			})
			subwg.Go(eth.CallContractFunction(syAbi, p.PoolAddress, "abond", []interface{}{}, &abond, s.block.Number))
			subwg.Go(func() error {
				var maxBondDailyRate = big.NewInt(0)

				err := eth.CallContractFunction(syAbi, p.PoolAddress, "maxBondDailyRate", []interface{}{}, &maxBondDailyRate, s.block.Number)()
				if err != nil && !strings.Contains(err.Error(), "Reverted") && !strings.Contains(err.Error(), "execution reverted") {
					return errors.Wrap(err, "could not get maxBondDailyRate")
				}

				rate, _ := decimal.NewFromBigInt(maxBondDailyRate, -18).Float64()

				apy := math.Pow(rate+1, 365) - 1
				mu.Lock()
				poolState.SeniorAPY = apy
				mu.Unlock()

				return nil
			})

			if p.ProtocolId == "compound/v2" || p.ProtocolId == "cream/v2" {
				subwg.Go(func() error {
					var rate = big.NewInt(0)

					err := eth.CallContractFunction(controllerAbi, p.ControllerAddress, "spotDailySupplyRateProvider", []interface{}{}, &rate, s.block.Number)()
					if err != nil && !strings.Contains(err.Error(), "Reverted") {
						return errors.Wrap(err, "could not get originator apy")
					}

					r := decimal.NewFromBigInt(rate, -18)
					rf, _ := r.Float64()

					mu.Lock()
					poolState.OriginatorApy = math.Pow(rf+1, 365) - 1
					mu.Unlock()

					return nil
				})

				subwg.Go(func() error {
					var rate = big.NewInt(0)

					err := eth.CallContractFunction(controllerAbi, p.ControllerAddress, "spotDailyRate", []interface{}{}, &rate, s.block.Number)()
					if err != nil && !strings.Contains(err.Error(), "Reverted") {
						return errors.Wrap(err, "could not get originator NET apy")
					}

					r := decimal.NewFromBigInt(rate, -18)
					rf, _ := r.Float64()

					mu.Lock()
					poolState.OriginatorNetApy = math.Pow(rf+1, 365) - 1
					mu.Unlock()

					return nil
				})
			} else if p.ProtocolId == "aave/v2" {
				subwg.Go(func() error {
					var rate = big.NewInt(0)

					err := eth.CallContractFunction(controllerAbi, p.ControllerAddress, "spotDailySupplyRateProvider", []interface{}{}, &rate, s.block.Number)()
					if err != nil && !strings.Contains(err.Error(), "Reverted") {
						return errors.Wrap(err, "could not get originator apy")
					}

					r := decimal.NewFromBigInt(rate, -18)
					rf, _ := r.Float64()

					apy := rf * 365
					mu.Lock()
					poolState.OriginatorApy = apy
					poolState.OriginatorNetApy = apy
					mu.Unlock()

					return nil
				})
			}

			err := subwg.Wait()
			if err != nil {
				return errors.Wrapf(err, "could not get pool state (%s)", p.PoolAddress)
			}

			poolState.TotalLiquidity = decimal.NewFromBigInt(underlyingTotal, 0)
			poolState.JuniorLiquidity = decimal.NewFromBigInt(underlyingJuniors, 0)
			poolState.JTokenPrice = decimal.NewFromBigInt(jtokenPrice, 0)
			poolState.Abond = abond

			abondGain := decimal.NewFromBigInt(abond.Gain, -int32(p.UnderlyingDecimals))
			abondPrincipal := decimal.NewFromBigInt(abond.Principal, -int32(p.UnderlyingDecimals))
			abondIssuedAt := decimal.NewFromBigInt(abond.IssuedAt, -18)
			abondMaturesAt := decimal.NewFromBigInt(abond.MaturesAt, -18)

			var abondAPY float64
			if !abondPrincipal.Equal(decimal.NewFromInt(0)) {
				abondAPY, _ = abondGain.Div(abondPrincipal).
					Div(abondMaturesAt.Sub(abondIssuedAt)).
					Mul(decimal.NewFromInt(365 * 24 * 60 * 60)).
					Float64()
			}

			seniorLiq := poolState.TotalLiquidity.Sub(poolState.JuniorLiquidity)

			if poolState.JuniorLiquidity.Equal(decimal.Zero) {
				poolState.JuniorAPY = poolState.OriginatorNetApy
			} else {
				juniorApy := decimal.NewFromFloat(poolState.OriginatorNetApy).Add(
					seniorLiq.
						Div(poolState.JuniorLiquidity).
						Mul(decimal.NewFromFloat(poolState.OriginatorNetApy - abondAPY)),
				)

				poolState.JuniorAPY, _ = juniorApy.Float64()
			}

			poolState.AbondAPY = abondAPY

			s.processedMu.Lock()
			s.processed.PoolStates = append(s.processed.PoolStates, poolState)
			s.processedMu.Unlock()

			return nil
		})
	}

	return wg.Wait()
}
