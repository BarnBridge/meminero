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

func (s *GovStorable) handleVotes(logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.Governance.IsGovernanceVoteEvent(&log) {
			vote, err := ethtypes.Governance.GovernanceVoteEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal vote event")
			}

			s.Processed.votes = append(s.Processed.votes, vote)
		}

		if ethtypes.Governance.IsGovernanceVoteCanceledEvent(&log) {
			vote, err := ethtypes.Governance.GovernanceVoteCanceledEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal vote canceled event")
			}

			s.Processed.canceledVotes = append(s.Processed.canceledVotes, vote)
		}
	}

	return nil
}

func (s *GovStorable) storeProposalVotes(ctx context.Context, tx pgx.Tx) error {
	if len(s.Processed.votes) == 0 {
		s.logger.WithField("handler", "votes").Debug("no events found")

		return nil
	}

	var rows [][]interface{}
	for _, v := range s.Processed.votes {
		power := decimal.NewFromBigInt(v.Power, 0)
		rows = append(rows, []interface{}{
			v.ProposalId.Int64(),
			utils.NormalizeAddress(v.User.String()),
			v.Support,
			power,
			s.block.BlockCreationTime,
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
		pgx.CopyFromSlice(len(s.Processed.votes), func(i int) ([]interface{}, error) {
			return rows[i], nil
		}),
	)
	if err != nil {
		return errors.Wrap(err, "could not store proposal votes")
	}

	return nil
}

func (s *GovStorable) storeProposalCanceledVotes(ctx context.Context, tx pgx.Tx) error {
	if len(s.Processed.canceledVotes) == 0 {
		s.logger.WithField("handler", "canceled votes").Debug("no events found")
		return nil
	}

	var rows [][]interface{}
	for _, v := range s.Processed.canceledVotes {
		rows = append(rows, []interface{}{
			v.ProposalId.Int64(),
			utils.NormalizeAddress(v.User.String()),
			s.block.BlockCreationTime,
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
