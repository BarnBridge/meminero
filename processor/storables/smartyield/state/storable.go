package state

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/state"
	globalTypes "github.com/barnbridge/meminero/types"
)

type Storable struct {
	block *globalTypes.Block

	state  *state.Manager
	logger *logrus.Entry

	processed struct {
		PoolStates []*State
	}
	processedMu sync.Mutex
}

const storableID = "smartYield.state"

func New(block *globalTypes.Block, state *state.Manager) *Storable {
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
