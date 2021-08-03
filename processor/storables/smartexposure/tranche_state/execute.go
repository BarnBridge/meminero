package tranche_state

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/processor/storables/smartexposure"
	seTypes "github.com/barnbridge/meminero/processor/storables/smartexposure/types"
	"github.com/barnbridge/meminero/processor/storables/tokenprices"
)

func (s *Storable) Execute(ctx context.Context) error {
	s.processed.trancheState = make(map[string]TrancheState)
	tokens, err := smartexposure.BuildTokensSliceForSE(s.state)
	if err != nil {
		return err
	}

	s.processed.tokenPrices, err = tokenprices.GetTokensPrices(ctx, tokens, s.block.Number)
	if err != nil {
		return err
	}

	wg, _ := errgroup.WithContext(ctx)
	var mu = &sync.Mutex{}

	for trancheAddress, tranche := range s.state.SmartExposure.Tranches {
		if s.block.Number < tranche.StartAtBlock {
			s.logger.WithField("tranche", trancheAddress).Info("skipping tranche due to StartAtBlock property")
			continue
		}

		trancheAddress := trancheAddress
		tranche := tranche

		wg.Go(func() error {
			var t seTypes.TrancheFromChain
			var currentRatio *big.Int
			var conversionRate struct {
				AmountA *big.Int `json:"amountA"`
				AmountB *big.Int `json:"amountB"`
			}

			subwg, _ := errgroup.WithContext(ctx)
			subwg.Go(eth.CallContractFunction(*ethtypes.EPool.ABI, tranche.EPoolAddress, "getTranche", []interface{}{common.HexToAddress(trancheAddress)}, &t, s.block.Number))
			subwg.Go(eth.CallContractFunction(*ethtypes.EPoolHelper.ABI, config.Store.Storable.SmartExposure.EPoolHelperAddress, "currentRatio", []interface{}{common.HexToAddress(tranche.EPoolAddress), common.HexToAddress(trancheAddress)}, &currentRatio, s.block.Number))
			subwg.Go(eth.CallContractFunction(*ethtypes.EPoolHelper.ABI, config.Store.Storable.SmartExposure.EPoolHelperAddress, "tokenATokenBForEToken", []interface{}{common.HexToAddress(tranche.EPoolAddress), common.HexToAddress(trancheAddress), tranche.SFactorE.BigInt()}, &conversionRate, s.block.Number))

			err := subwg.Wait()
			if err != nil {
				return err
			}

			mu.Lock()
			s.processed.trancheState[trancheAddress] = TrancheState{
				EPoolAddress:    tranche.EPoolAddress,
				CurrentRatio:    decimal.NewFromBigInt(currentRatio, 0),
				TokenALiquidity: decimal.NewFromBigInt(t.ReserveA, 0),
				TokenBLiquidity: decimal.NewFromBigInt(t.ReserveB, 0),
				ConversionRate: ConversionRate{
					AmountAConversion: decimal.NewFromBigInt(conversionRate.AmountA, 0),
					AmountBConversion: decimal.NewFromBigInt(conversionRate.AmountB, 0),
				},
			}
			mu.Unlock()

			return nil
		})
	}

	err = wg.Wait()
	if err != nil {
		return errors.Wrap(err, "could not get data from chain")
	}

	return nil
}
