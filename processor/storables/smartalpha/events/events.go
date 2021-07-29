package events

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/ethtypes"
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
	}
}

func New(block *globalTypes.Block, state *state.Manager) *Storable {
	return &Storable{
		block:  block,
		state:  state,
		logger: logrus.WithField("module", "storable(smartAlpha.events)"),
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
			if s.state.SmartAlpha.PoolByAddress(log.Address.String()) != nil {
				err := s.processUserEvent(log)
				if err != nil {
					return errors.Wrap(err, "could not process pool event")
				}

				if ethtypes.SmartAlpha.IsEpochEndEvent(&log) {
					e, err := ethtypes.SmartAlpha.EpochEndEvent(log)
					if err != nil {
						return errors.Wrap(err, "could not decode epoch end event")
					}

					s.processed.EpochEndEvents = append(s.processed.EpochEndEvents, e)
				}
			}

			// capture any junior/senior token Transfer events to save them to transaction history
			if s.state.SmartAlpha.IsERC20OfInterest(log.Address.String()) && len(log.Topics) == 3 && ethtypes.ERC20.IsTransferEvent(&log) {
				e, err := ethtypes.ERC20.TransferEvent(log)
				if err != nil {
					return errors.Wrap(err, "could not decode ERC20 Transfer event")
				}

				s.processed.TokenTransferEvents = append(s.processed.TokenTransferEvents, e)
			}
		}
	}

	return nil
}

func (s *Storable) Result() interface{} {
	return s.processed
}
