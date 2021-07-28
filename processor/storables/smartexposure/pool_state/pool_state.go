package pool_state

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/processor/storables/smartexposure"
	seTypes "github.com/barnbridge/meminero/processor/storables/smartexposure/types"
	"github.com/barnbridge/meminero/processor/storables/tokenprices"
	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
)

type Storable struct {
	block *types.Block

	state  *state.Manager
	logger *logrus.Entry

	processed struct {
		poolStates  map[string]PoolState
		tokenPrices map[string]decimal.Decimal
	}
}

func New(block *types.Block, state *state.Manager) *Storable {
	return &Storable{
		block:  block,
		state:  state,
		logger: logrus.WithField("module", "storable(smartExposure.poolState)"),
	}
}

func (s *Storable) Execute(ctx context.Context) error {
	s.logger.Trace("executing")
	start := time.Now()
	defer func() {
		s.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

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
			liqA = liqA.Mul(s.processed.tokenPrices[pool.TokenA.Address])
			liqB = liqB.Mul(s.processed.tokenPrices[pool.TokenB.Address])
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

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.WithField("block", s.block.Number).Trace("rolling back block")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Trace("done rolling back block")
	}()

	_, err := tx.Exec(ctx, `delete from smart_exposure.pool_state where included_in_block = $1`, s.block.Number)

	return err
}

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	s.logger.Trace("executing")
	start := time.Now()
	defer func() {
		s.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	err := s.storePoolsState(ctx, tx)

	return errors.Wrap(err, "could not store pools state")
}

func (s *Storable) Result() interface{} {
	return s.processed
}
