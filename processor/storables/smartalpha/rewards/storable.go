package rewards

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
)

type Storable struct {
	block *types.Block

	state  *state.Manager
	logger *logrus.Entry

	processed struct {
		Claims      []ethtypes.RewardPoolSingleClaimEvent
		ClaimsMulti []ethtypes.RewardPoolMultiClaimRewardTokenEvent
		Deposits    []ethtypes.RewardPoolSingleDepositEvent
		Withdrawals []ethtypes.RewardPoolSingleWithdrawEvent
	}
}

const storableID = "smartAlpha.rewards"

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

func (s Storable) Result() interface{} {
	return s.processed
}
