package events

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/processor/storables/smartalpha"
	"github.com/barnbridge/meminero/state"
	globalTypes "github.com/barnbridge/meminero/types"
)

type Storable struct {
	block *globalTypes.Block

	state  *state.Manager
	logger *logrus.Entry

	processed struct {
		JuniorJoinEntryQueueEvents   []ethtypes.SmartAlphaJuniorJoinEntryQueueEvent
		JuniorRedeemTokensEvents     []ethtypes.SmartAlphaJuniorRedeemTokensEvent
		JuniorJoinExitQueueEvents    []ethtypes.SmartAlphaJuniorJoinExitQueueEvent
		JuniorRedeemUnderlyingEvents []ethtypes.SmartAlphaJuniorRedeemUnderlyingEvent
		SeniorJoinEntryQueueEvents   []ethtypes.SmartAlphaSeniorJoinEntryQueueEvent
		SeniorRedeemTokensEvents     []ethtypes.SmartAlphaSeniorRedeemTokensEvent
		SeniorJoinExitQueueEvents    []ethtypes.SmartAlphaSeniorJoinExitQueueEvent
		SeniorRedeemUnderlyingEvents []ethtypes.SmartAlphaSeniorRedeemUnderlyingEvent
		EpochEndEvents               []ethtypes.SmartAlphaEpochEndEvent
		TokenTransferEvents          []ethtypes.ERC20TransferEvent
		EpochInfos                   []smartalpha.EpochInfo
	}
}

const storableID = "smartAlpha.events"

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
