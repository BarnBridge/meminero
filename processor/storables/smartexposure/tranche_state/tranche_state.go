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
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type Storable struct {
	block *types.Block

	state  *state.Manager
	logger *logrus.Entry

	processed struct {
		trancheState map[string]*TrancheState
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
	spew.Dump(config.Store.Storable.SmartExposure)
	s.processed.trancheState = make(map[string]*TrancheState)
	wg, _ := errgroup.WithContext(ctx)
	var mu = &sync.Mutex{}
	for trancheAddress, tranche := range s.state.SETranches() {

		var t smartexposure.TrancheFromChain
		a := ethtypes.Epool.ABI
		var currentRatio, amountAConversion, amountBConversion *big.Int
		wg.Go(eth.CallContractFunction(*a, tranche.EPoolAddress, "getTranche", []interface{}{common.HexToAddress(trancheAddress)}, &t))
		wg.Go(eth.CallContractFunction(*a, config.Store.Storable.SmartExposure.EPoolHelperAddress, "currentRatio", []interface{}{common.HexToAddress(tranche.EPoolAddress), common.HexToAddress(trancheAddress)}, &currentRatio))
		wg.Go(eth.CallContractFunction(*a, config.Store.Storable.SmartExposure.EPoolHelperAddress, "tokenATokenBForEToken", []interface{}{common.HexToAddress(tranche.EPoolAddress), common.HexToAddress(trancheAddress), tranche.SFactorE}, &amountAConversion, &amountBConversion))

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
	}

	err := wg.Wait()
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
