package governance

import (
	"context"

	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/utils"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (s *GovStorable) handleVotes(logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.Governance.IsVoteEvent(&log) {
			vote, err := ethtypes.Governance.VoteEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal vote event")
			}

			s.Processed.votes = append(s.Processed.votes, vote)
		}

		if ethtypes.Governance.IsVoteCanceledEvent(&log) {
			vote, err := ethtypes.Governance.VoteCanceledEvent(log)
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
		return nil
	}

	var rows [][]interface{}
	for _, v := range s.Processed.votes {
		rows = append(rows, []interface{}{
			v.ProposalId.Int64(),
			utils.NormalizeAddress(v.User.String()),
			v.Support,
			v.PowerDecimal(0),
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
