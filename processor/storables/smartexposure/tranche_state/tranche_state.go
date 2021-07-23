package tranche_state

import (
	"context"
	"sync"
	"time"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/processor/storables/smartexposure"
	"github.com/barnbridge/meminero/state"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
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

	s.processed.trancheState = make(map[string]*TrancheState)
	var wg = &errgroup.Group{}
	var mu = &sync.Mutex{}
	tranches := make(map[string]smartexposure.TrancheFromChain)
	for trancheAddress, tranche := range s.state.SETranches() {
		var t smartexposure.TrancheFromChain
		a := ethtypes.Epool.ABI

		wg.Go(eth.CallContractFunction(*a, tranche.EPoolAddress, "getTranche", []interface{}{common.HexToAddress(trancheAddress)}, &t))

		mu.Lock()
		tranches[t.Etoken.String()] = t
		mu.Unlock()
	}

	return nil

}

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	return nil
}

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	return nil
}

func (s *Storable) Result() interface{} {
	return s.processed

}
