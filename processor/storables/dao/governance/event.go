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

func (g *GovStorable) handleEvents(logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.Governance.IsGovernanceProposalCreatedEvent(&log) {
			e, err := ethtypes.Governance.GovernanceProposalCreatedEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal created event")
			}
			g.Processed.proposalEvents = append(g.Processed.proposalEvents, ProposalEvent{
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

		if ethtypes.Governance.IsGovernanceProposalQueuedEvent(&log) {
			e, err := ethtypes.Governance.GovernanceProposalQueuedEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal queued event")
			}

			g.Processed.proposalEvents = append(g.Processed.proposalEvents, ProposalEvent{
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

		if ethtypes.Governance.IsGovernanceProposalExecutedEvent(&log) {
			e, err := ethtypes.Governance.GovernanceProposalExecutedEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal executed event")
			}

			g.Processed.proposalEvents = append(g.Processed.proposalEvents, ProposalEvent{
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

		if ethtypes.Governance.IsGovernanceProposalCanceledEvent(&log) {
			e, err := ethtypes.Governance.GovernanceProposalCanceledEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal canceled event")
			}

			g.Processed.proposalEvents = append(g.Processed.proposalEvents, ProposalEvent{
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

func (g *GovStorable) storeEvents(ctx context.Context, tx pgx.Tx) error {
	if len(g.Processed.proposalEvents) == 0 {
		g.logger.WithField("handler", "proposal events").Debug("no events found")
		return nil
	}

	var rows [][]interface{}
	for _, e := range g.Processed.proposalEvents {
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
			g.block.BlockCreationTime,
			g.block.Number,
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
