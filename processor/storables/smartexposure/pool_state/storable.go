package pool_state

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
)

type Storable struct {
	block *types.Block

	state  *state.Manager
	logger *logrus.Entry

	processed struct {
		poolStates  map[string]PoolState
		tokenPrices map[string]map[string]decimal.Decimal
	}
}

const storableID = "smartExposure.poolState"

func New(block *types.Block, state *state.Manager) *Storable {
	return &Storable{
		block:  block,
		state:  state,
		logger: logrus.WithField("module", fmt.Sprintf("storable(%s)", storableID)),
	}
}

func (s *Storable) ID() string {
	return storableID
}

func (s *Storable) Result() interface{} {
	return s.processed
}
