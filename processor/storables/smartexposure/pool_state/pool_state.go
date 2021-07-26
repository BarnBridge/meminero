package pool_state

import (
	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type Storable struct {
	block *types.Block

	state  *state.Manager
	logger *logrus.Entry

	processed struct {
		poolState   map[string]*PoolState
		tokenPrices map[string]decimal.Decimal
	}
}

func New(block *types.Block, state *state.Manager) *Storable {
	return &Storable{
		block:  block,
		state:  state,
		logger: logrus.WithField("module", "storable(smart_exposure_tranche_state)"),
	}
}
