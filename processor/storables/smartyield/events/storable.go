package events

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
		JuniorEntryEvents           []ethtypes.SmartYieldBuyTokensEvent
		JuniorInstantWithdrawEvents []ethtypes.SmartYieldSellTokensEvent
		Junior2StepWithdrawEvents   []ethtypes.SmartYieldBuyJuniorBondEvent
		Junior2StepRedeemEvents     []ethtypes.SmartYieldRedeemJuniorBondEvent
		SeniorEntryEvents           []ethtypes.SmartYieldBuySeniorBondEvent
		SeniorRedeemEvents          []ethtypes.SmartYieldRedeemSeniorBondEvent
		Transfers                   []ethtypes.SmartYieldTransferEvent
		ControllerHarvests          []ethtypes.SmartYieldCompoundControllerHarvestEvent
		ProviderTransferFees        []ethtypes.SmartYieldCompoundProviderTransferFeesEvent
	}
}

const storableID = "smartYield.events"

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
