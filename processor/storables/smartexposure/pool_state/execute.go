package pool_state

import (
	"context"
	"math/big"
	"sync"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/processor/storables/smartexposure"
	seTypes "github.com/barnbridge/meminero/processor/storables/smartexposure/types"
	"github.com/barnbridge/meminero/processor/storables/tokenprices"
)

func (s *Storable) Execute(ctx context.Context) error {
	s.processed.poolStates = make(map[string]PoolState)
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

	for address, pool := range s.state.SmartExposure.Pools {
		if s.block.Number < pool.StartAtBlock {
			s.logger.WithField("pool", address).Info("skipping pool due to StartAtBlock property")
			continue
		}

		address := address
		pool := pool
		wg.Go(func() error {
			subwg, _ := errgroup.WithContext(ctx)
			var tranches []seTypes.TrancheFromChain
			var lastRebalance, rebalancingInterval, rebalancingCondition *big.Int

			subwg.Go(eth.CallContractFunction(*ethtypes.EPool.ABI, address, "getTranches", []interface{}{}, &tranches))
			subwg.Go(eth.CallContractFunction(*ethtypes.EPool.ABI, address, "lastRebalance", []interface{}{}, &lastRebalance))
			subwg.Go(eth.CallContractFunction(*ethtypes.EPool.ABI, address, "rebalanceInterval", []interface{}{}, &rebalancingInterval))
			subwg.Go(eth.CallContractFunction(*ethtypes.EPool.ABI, address, "rebalanceMinRDiv", []interface{}{}, &rebalancingCondition))
			err := subwg.Wait()
			if err != nil {
				return errors.Wrap(err, "could not get pool info from chain ")
			}

			mu.Lock()
			var liqA, liqB decimal.Decimal
			for _, t := range tranches {
				liqA = liqA.Add(decimal.NewFromBigInt(t.ReserveA, -int32(pool.TokenA.Decimals)))
				liqB = liqB.Add(decimal.NewFromBigInt(t.ReserveB, -int32(pool.TokenB.Decimals)))
			}
			liqA = liqA.Mul(s.processed.tokenPrices[pool.TokenA.Address]["USD"])
			liqB = liqB.Mul(s.processed.tokenPrices[pool.TokenB.Address]["USD"])
			s.processed.poolStates[address] = PoolState{
				PoolAddress:          address,
				PoolLiquidity:        liqA.Add(liqB),
				LastRebalance:        decimal.NewFromBigInt(lastRebalance, 0),
				RebalancingInterval:  decimal.NewFromBigInt(rebalancingInterval, 0),
				RebalancingCondition: decimal.NewFromBigInt(rebalancingCondition, 0),
			}
			mu.Unlock()

			return nil
		})
	}

	return wg.Wait()
}
