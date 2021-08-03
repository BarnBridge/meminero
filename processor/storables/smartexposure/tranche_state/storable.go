package tranche_state

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
		trancheState map[string]TrancheState
		tokenPrices  map[string]map[string]decimal.Decimal
	}
}

const storableID = "smartExposure.trancheState"

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
