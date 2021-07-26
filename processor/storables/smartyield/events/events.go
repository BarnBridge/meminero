package events

import (
	"context"
	"time"

	"github.com/pkg/errors"
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

func New(block *types.Block, state *state.Manager) *Storable {
	return &Storable{
		block:  block,
		state:  state,
		logger: logrus.WithField("module", "storable(smartYield.events)"),
	}
}

func (s *Storable) Execute(ctx context.Context) error {
	s.logger.Trace("executing")
	start := time.Now()
	defer func() {
		s.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if s.state.SmartYield.PoolByAddress(log.Address.String()) != nil {
				err := s.decodePoolEvent(log)
				if err != nil {
					return errors.Wrap(err, "could not decode pool event")
				}
			}

			if s.state.SmartYield.PoolByControllerAddress(log.Address.String()) != nil {
				err := s.decodeControllerEvent(log)
				if err != nil {
					return errors.Wrap(err, "could not decode controller event")
				}
			}

			if s.state.SmartYield.PoolByProviderAddress(log.Address.String()) != nil {
				err := s.decodeProviderEvent(log)
				if err != nil {
					return errors.Wrap(err, "could not decode provider event")
				}
			}
		}
	}

	return nil
}

func (s *Storable) Result() interface{} {
	return s.processed
}
