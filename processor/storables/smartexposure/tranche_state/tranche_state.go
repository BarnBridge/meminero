package tranche_state

import (
	"context"
	"math/big"
	"sync"
	"time"

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
			var currentRatio, amountAConversion, amountBConversion *big.Int
			subwg.Go(eth.CallContractFunction(*a, tranche.EPoolAddress, "getTranche", []interface{}{common.HexToAddress(trancheAddress)}, &t, s.block.Number))

			//wg.Go(eth.CallContractFunction(*a, config.Store.Storable.SmartExposure.EPoolHelperAddress, "currentRatio", []interface{}{common.HexToAddress(tranche.EPoolAddress), common.HexToAddress(trancheAddress)}, &currentRatio))
			//wg.Go(eth.CallContractFunction(*a, config.Store.Storable.SmartExposure.EPoolHelperAddress, "tokenATokenBForEToken", []interface{}{common.HexToAddress(tranche.EPoolAddress), common.HexToAddress(trancheAddress), tranche.SFactorE}, &amountAConversion, &amountBConversion))
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
					AmountAConversion: amountAConversion,
					AmountBConversion: amountBConversion,
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
	return nil
}

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	err := s.storeTranchesState(ctx, tx)
	return err
}

func (s *Storable) Result() interface{} {
	return s.processed

}
