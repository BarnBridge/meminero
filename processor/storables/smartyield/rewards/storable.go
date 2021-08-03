package rewards

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/processor/storables/smartyield"
	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
)

type Storable struct {
	block *types.Block

	state  *state.Manager
	logger *logrus.Entry

	factories []string

	processed struct {
		Pools       []smartyield.RewardPool
		Claims      []ethtypes.SmartYieldPoolSingleClaimEvent
		ClaimsMulti []ethtypes.SmartYieldPoolMultiClaimRewardTokenEvent
		Deposits    []ethtypes.SmartYieldPoolSingleDepositEvent
		Withdrawals []ethtypes.SmartYieldPoolSingleWithdrawEvent
	}
}

const storableID = "smartYield.rewards"

func New(block *types.Block, state *state.Manager) *Storable {
	return &Storable{
		block:     block,
		state:     state,
		logger:    logrus.WithField("module", fmt.Sprintf("storable(%s)", storableID)),
		factories: strings.Split(config.Store.Storable.SmartYield.Rewards.Factories, ","),
	}
}

func (s *Storable) ID() string {
	return storableID
}

func (s Storable) Result() interface{} {
	return s.processed
}
