package events

import (
	"context"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/ethtypes"
)

func (s *Storable) Execute(ctx context.Context) error {
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

func (s *Storable) processUserEvent(log gethtypes.Log) error {
	sa := ethtypes.SmartAlpha

	if sa.IsJuniorJoinEntryQueueEvent(&log) {
		e, err := sa.JuniorJoinEntryQueueEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode JuniorJoinEntryQueue event")
		}

		s.processed.JuniorJoinEntryQueueEvents = append(s.processed.JuniorJoinEntryQueueEvents, e)

		return nil
	}

	if sa.IsJuniorRedeemTokensEvent(&log) {
		e, err := sa.JuniorRedeemTokensEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode JuniorRedeemTokens event")
		}

		s.processed.JuniorRedeemTokensEvents = append(s.processed.JuniorRedeemTokensEvents, e)

		return nil
	}

	if sa.IsJuniorJoinExitQueueEvent(&log) {
		e, err := sa.JuniorJoinExitQueueEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode JuniorJoinExitQueue event")
		}

		s.processed.JuniorJoinExitQueueEvents = append(s.processed.JuniorJoinExitQueueEvents, e)

		return nil
	}

	if sa.IsJuniorRedeemUnderlyingEvent(&log) {
		e, err := sa.JuniorRedeemUnderlyingEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode JuniorRedeemUnderlying event")
		}

		s.processed.JuniorRedeemUnderlyingEvents = append(s.processed.JuniorRedeemUnderlyingEvents, e)

		return nil
	}

	if sa.IsSeniorJoinEntryQueueEvent(&log) {
		e, err := sa.SeniorJoinEntryQueueEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode SeniorJoinEntryQueue event")
		}

		s.processed.SeniorJoinEntryQueueEvents = append(s.processed.SeniorJoinEntryQueueEvents, e)

		return nil
	}

	if sa.IsSeniorRedeemTokensEvent(&log) {
		e, err := sa.SeniorRedeemTokensEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode SeniorRedeemTokens event")
		}

		s.processed.SeniorRedeemTokensEvents = append(s.processed.SeniorRedeemTokensEvents, e)

		return nil
	}

	if sa.IsSeniorJoinExitQueueEvent(&log) {
		e, err := sa.SeniorJoinExitQueueEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode SeniorJoinExitQueue event")
		}

		s.processed.SeniorJoinExitQueueEvents = append(s.processed.SeniorJoinExitQueueEvents, e)

		return nil
	}

	if sa.IsSeniorRedeemUnderlyingEvent(&log) {
		e, err := sa.SeniorRedeemUnderlyingEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode SeniorRedeemUnderlying event")
		}

		s.processed.SeniorRedeemUnderlyingEvents = append(s.processed.SeniorRedeemUnderlyingEvents, e)

		return nil
	}

	return nil
}
