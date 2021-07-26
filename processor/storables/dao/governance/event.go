package governance

import (
	"context"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/types"
	"github.com/barnbridge/meminero/utils"
)

func (s *GovStorable) handleEvents(logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.Governance.IsProposalCreatedEvent(&log) {
			e, err := ethtypes.Governance.ProposalCreatedEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal created event")
			}
			s.Processed.proposalEvents = append(s.Processed.proposalEvents, ProposalEvent{
				ProposalID: e.ProposalId,
				EventType:  CREATED,
				BaseLog: BaseLog{
					TransactionHash:  utils.NormalizeAddress(e.Raw.TxHash.String()),
					TransactionIndex: int64(e.Raw.TxIndex),
					LogIndex:         int64(e.Raw.Index),
				},
			})
			continue
		}

		if ethtypes.Governance.IsProposalQueuedEvent(&log) {
			e, err := ethtypes.Governance.ProposalQueuedEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal queued event")
			}

			s.Processed.proposalEvents = append(s.Processed.proposalEvents, ProposalEvent{
				ProposalID: e.ProposalId,
				Caller:     e.Caller,
				Eta:        e.Eta,
				EventType:  QUEUED,
				BaseLog: BaseLog{
					TransactionHash:  utils.NormalizeAddress(e.Raw.TxHash.String()),
					TransactionIndex: int64(e.Raw.TxIndex),
					LogIndex:         int64(e.Raw.Index),
				},
			})
			continue
		}

		if ethtypes.Governance.IsProposalExecutedEvent(&log) {
			e, err := ethtypes.Governance.ProposalExecutedEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal executed event")
			}

			s.Processed.proposalEvents = append(s.Processed.proposalEvents, ProposalEvent{
				ProposalID: e.ProposalId,
				Caller:     e.Caller,
				EventType:  EXECUTED,
				BaseLog: BaseLog{
					TransactionHash:  utils.NormalizeAddress(e.Raw.TxHash.String()),
					TransactionIndex: int64(e.Raw.TxIndex),
					LogIndex:         int64(e.Raw.Index),
				},
			})
			continue
		}

		if ethtypes.Governance.IsProposalCanceledEvent(&log) {
			e, err := ethtypes.Governance.ProposalCanceledEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal canceled event")
			}

			s.Processed.proposalEvents = append(s.Processed.proposalEvents, ProposalEvent{
				ProposalID: e.ProposalId,
				Caller:     e.Caller,
				EventType:  CANCELED,
				BaseLog: BaseLog{
					TransactionHash:  utils.NormalizeAddress(e.Raw.TxHash.String()),
					TransactionIndex: int64(e.Raw.TxIndex),
					LogIndex:         int64(e.Raw.Index),
				},
			})
			continue
		}

	}
	return nil
}

func (s *GovStorable) storeEvents(ctx context.Context, tx pgx.Tx) error {
	if len(s.Processed.proposalEvents) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, e := range s.Processed.proposalEvents {
		var eventData types.JSONObject
		if e.Eta != nil {
			eventData = make(types.JSONObject)
			eventData["eta"] = e.Eta.Int64()
		}

		rows = append(rows, []interface{}{
			e.ProposalID.Int64(),
			utils.NormalizeAddress(e.Caller.String()),
			e.EventType,
			eventData,
			s.block.BlockCreationTime,
			s.block.Number,
			e.TransactionHash,
			e.TransactionIndex,
			e.LogIndex,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"governance", "proposal_events"},
		[]string{"proposal_id", "caller", "event_type", "event_data", "block_timestamp", "included_in_block", "tx_hash", "tx_index", "log_index"},
		pgx.CopyFromRows(rows),
	)

	if err != nil {
		return errors.Wrap(err, "could not store proposal events")
	}

	return nil
}
