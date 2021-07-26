package tranche_state

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/processor/storables/smartexposure"
	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type Storable struct {
	block *types.Block

	state  *state.Manager
	logger *logrus.Entry

	processed struct {
		trancheState map[string]*TrancheState
		tokenPrices  map[string]decimal.Decimal
	}
}

func New(block *types.Block, state *state.Manager) *Storable {
	return &Storable{
		block:  block,
		state:  state,
		logger: logrus.WithField("module", "storable(smart_exposure_tranche_state)"),
	}
}

func (s *Storable) Execute(ctx context.Context) error {
	s.logger.Trace("executing")
	start := time.Now()
	defer func() {
		s.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	s.processed.trancheState = make(map[string]*TrancheState)
	s.processed.tokenPrices = make(map[string]decimal.Decimal)

	err := s.getTokensPrice(ctx)
	if err != nil {
		return err
	}
	wg, _ := errgroup.WithContext(ctx)
	var mu = &sync.Mutex{}
	a := ethtypes.Epool.ABI

	for trancheAddress, tranche := range s.state.SETranches() {
		trancheAddress := trancheAddress
		tranche := tranche

		wg.Go(func() error {
			subwg, _ := errgroup.WithContext(ctx)
			var t smartexposure.TrancheFromChain
			var currentRatio *big.Int

			var conversionRate struct {
				AmountA *big.Int `json:"amountA"`
				AmountB *big.Int `json:"amountB"`
			}

			subwg.Go(eth.CallContractFunction(*a, tranche.EPoolAddress, "getTranche", []interface{}{common.HexToAddress(trancheAddress)}, &t, s.block.Number))

			wg.Go(eth.CallContractFunction(*ethtypes.Epoolhelper2.ABI, config.Store.Storable.SmartExposure.EPoolHelperAddress, "currentRatio", []interface{}{common.HexToAddress(tranche.EPoolAddress), common.HexToAddress(trancheAddress)}, &currentRatio, s.block.Number))
			wg.Go(eth.CallContractFunction(*ethtypes.Epoolhelper2.ABI, config.Store.Storable.SmartExposure.EPoolHelperAddress, "tokenATokenBForEToken", []interface{}{common.HexToAddress(tranche.EPoolAddress), common.HexToAddress(trancheAddress), tranche.SFactorE}, &conversionRate, s.block.Number))
			err := subwg.Wait()
			if err != nil {
				return err
			}

			mu.Lock()
			s.processed.trancheState[trancheAddress] = &TrancheState{
				EPoolAddress:    tranche.EPoolAddress,
				CurrentRatio:    currentRatio,
				TokenALiquidity: t.ReserveA,
				TokenBLiquidity: t.ReserveB,
				ConversionRate: ConversionRate{
					AmountAConversion: conversionRate.AmountA,
					AmountBConversion: conversionRate.AmountB,
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

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.WithField("block", s.block.Number).Debug("rolling back block")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Debug("done rolling back block")
	}()

	_, err := tx.Exec(ctx, `delete from smart_exposure.tranche_state where included_in_block = $1`, s.block.Number)

	return err
}

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	err := s.storeTranchesState(ctx, tx)
	return err
}

func (s *Storable) Result() interface{} {
	return s.processed

}
