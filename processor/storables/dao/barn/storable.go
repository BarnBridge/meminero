package barn

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/types"
)

type Storable struct {
	block *types.Block

	logger *logrus.Entry

	processed struct {
		delegateActions []DelegateAction
		delegateChanges []DelegateChange
		locks           []ethtypes.BarnLockEvent
		stakingActions  []StakingAction
	}
}

const storableID = "dao.barn"

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
