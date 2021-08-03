package yieldfarming

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/types"
)

type Storable struct {
	block     *types.Block
	logger    *logrus.Entry
	processed struct {
		stakingActions []StakingAction
	}
}

const storableID = "yieldFarming"

func New(block *types.Block) *Storable {
	return &Storable{
		block:  block,
		logger: logrus.WithField("module", fmt.Sprintf("storable(%s)", storableID)),
	}
}

func (s *Storable) ID() string {
	return storableID
}

func (s *Storable) Result() interface{} {
	return s.processed
}
