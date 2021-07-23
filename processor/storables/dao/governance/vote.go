package governance

import (
	"context"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/utils"
)

func (g *GovStorable) handleVotes(logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.Governance.IsGovernanceVoteEvent(&log) {
			vote, err := ethtypes.Governance.GovernanceVoteEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal vote event")
			}

			g.Processed.votes = append(g.Processed.votes, vote)
		}

		if ethtypes.Governance.IsGovernanceVoteCanceledEvent(&log) {
			vote, err := ethtypes.Governance.GovernanceVoteCanceledEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal vote canceled event")
			}

			g.Processed.canceledVotes = append(g.Processed.canceledVotes, vote)
		}
	}

	return nil
}

func (g *GovStorable) storeProposalVotes(ctx context.Context, tx pgx.Tx) error {
	if len(g.Processed.votes) == 0 {
		g.logger.WithField("handler", "votes").Debug("no events found")

		return nil
	}

	var rows [][]interface{}
	for _, v := range g.Processed.votes {
		power := decimal.NewFromBigInt(v.Power, 0)
		rows = append(rows, []interface{}{
			v.ProposalId.Int64(),
			utils.NormalizeAddress(v.User.String()),
			v.Support,
			power,
			g.block.BlockCreationTime,
			v.Raw.BlockNumber,
			utils.NormalizeAddress(v.Raw.TxHash.String()),
			v.Raw.TxIndex,
			v.Raw.Index,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"governance", "votes"},
		[]string{"proposal_id", "user_id", "support", "power", "block_timestamp", "included_in_block", "tx_hash", "tx_index", "log_index"},
		pgx.CopyFromSlice(len(g.Processed.votes), func(i int) ([]interface{}, error) {
			return rows[i], nil
		}),
	)
	if err != nil {
		return errors.Wrap(err, "could not store proposal votes")
	}

	return nil
}

func (g *GovStorable) storeProposalCanceledVotes(ctx context.Context, tx pgx.Tx) error {
	if len(g.Processed.canceledVotes) == 0 {
		g.logger.WithField("handler", "canceled votes").Debug("no events found")
		return nil
	}

	var rows [][]interface{}
	for _, v := range g.Processed.canceledVotes {
		rows = append(rows, []interface{}{
			v.ProposalId.Int64(),
			utils.NormalizeAddress(v.User.String()),
			g.block.BlockCreationTime,
			v.Raw.BlockNumber,
			utils.NormalizeAddress(v.Raw.TxHash.String()),
			v.Raw.TxIndex,
			v.Raw.Index,
		})
	}
	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"governance", "votes_canceled"},
		[]string{"proposal_id", "user_id", "block_timestamp", "included_in_block", "tx_hash", "tx_index", "log_index"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not store proposal canceled votes")
	}

	return nil
}
