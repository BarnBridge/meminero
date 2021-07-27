package pool_state

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/processor/storables/smartexposure"
	smartexposure2 "github.com/barnbridge/meminero/state/smartexposure"

	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type Storable struct {
	block *types.Block

	state  *state.Manager
	logger *logrus.Entry

	processed struct {
		poolStates  map[string]*PoolState
		tokenPrices map[string]decimal.Decimal
	}
}

func New(block *types.Block, state *state.Manager) *Storable {
	return &Storable{
		block:  block,
		state:  state,
		logger: logrus.WithField("module", "storable(smart_exposure_pool_state)"),
	}
}

func (s *Storable) Execute(ctx context.Context) error {
	s.logger.Trace("executing")
	start := time.Now()
	defer func() {
		s.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	s.processed.poolStates = make(map[string]*PoolState)
	var err error
	s.processed.tokenPrices, err = smartexposure.GetTokensPrice(ctx, s.state, s.block.Number)
	if err != nil {
		return err
	}

	wg, _ := errgroup.WithContext(ctx)
	var mu = &sync.Mutex{}
	a := ethtypes.Epool.ABI

	for address, pool := range s.state.SmartExposure.SEPools() {
		if s.block.Number < pool.StartAtBlock {
			s.logger.WithField("pool", address).Info("skipping pool due to StartAtBlock property")
			continue
		}

		address := address
		pool := pool
		wg.Go(func() error {
			subwg, _ := errgroup.WithContext(ctx)
			var _tranches []smartexposure2.TrancheFromChain
			var lastRebalance, rebalancingInterval, rebalancingCondition *big.Int

			subwg.Go(eth.CallContractFunction(*a, address, "getTranches", []interface{}{}, &_tranches))
			subwg.Go(eth.CallContractFunction(*a, address, "lastRebalance", []interface{}{}, &lastRebalance))
			subwg.Go(eth.CallContractFunction(*a, address, "rebalanceInterval", []interface{}{}, &rebalancingInterval))
			subwg.Go(eth.CallContractFunction(*a, address, "rebalanceMinRDiv", []interface{}{}, &rebalancingCondition))

			err := subwg.Wait()
			if err != nil {
				return err
			}

			mu.Lock()
			var liqA, liqB decimal.Decimal

			for _, t := range _tranches {
				liqA = liqA.Add(t.ReserveA.Shift(-int32(pool.ATokenDecimals)))
				liqB = liqB.Add(t.ReserveB.Shift(-int32(pool.BTokenDecimals)))
			}

			liqA = liqA.Mul(s.processed.tokenPrices[pool.ATokenAddress])
			liqB = liqB.Mul(s.processed.tokenPrices[pool.BTokenAddress])
			s.processed.poolStates[address] = &PoolState{
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
	err = wg.Wait()
	return err

}

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.WithField("block", s.block.Number).Debug("rolling back block")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Debug("done rolling back block")
	}()

	_, err := tx.Exec(ctx, `delete from smart_exposure.pool_state where included_in_block = $1`, s.block.Number)

	return err
}

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	err := s.storePoolsState(ctx, tx)

	return err
}

func (s *Storable) Result() interface{} {
	return s.processed
}
