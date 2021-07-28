package governance

import (
	"context"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/utils"
)

func (s *GovStorable) handleAbrogationProposalVotes(logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.Governance.IsAbrogationProposalVoteEvent(&log) {
			vote, err := ethtypes.Governance.AbrogationProposalVoteEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode abrogation proposal event")
			}

			s.Processed.abrogationVotes = append(s.Processed.abrogationVotes, vote)
		}

		if ethtypes.Governance.IsAbrogationProposalVoteEvent(&log) {
			vote, err := ethtypes.Governance.AbrogationProposalVoteCancelledEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode abrogation proposal event")
			}

			s.Processed.abrogationCanceledVotes = append(s.Processed.abrogationCanceledVotes, vote)
		}
	}

	return nil
}

func (s *GovStorable) storeProposalAbrogationVotes(ctx context.Context, tx pgx.Tx) error {
	if len(s.Processed.abrogationVotes) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, v := range s.Processed.abrogationVotes {
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
		pgx.Identifier{"governance", "abrogation_votes"},
		[]string{"proposal_id", "user_id", "support", "power", "block_timestamp", "included_in_block", "tx_hash", "tx_index", "log_index"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not store abrogation proposal  votes")
	}

	return nil
}

func (s *GovStorable) storeAbrogationProposalCanceledVotes(ctx context.Context, tx pgx.Tx) error {
	if len(s.Processed.abrogationCanceledVotes) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, v := range s.Processed.abrogationCanceledVotes {
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
		pgx.Identifier{"governance", "abrogation_votes_canceled"},
		[]string{"proposal_id", "user_id", "block_timestamp", "included_in_block", "tx_hash", "tx_index", "log_index"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not store abrogation proposal canceled votes")
	}

	return nil
}
