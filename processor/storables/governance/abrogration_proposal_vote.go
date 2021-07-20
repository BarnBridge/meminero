package governance

import (
	"context"

	"github.com/barnbridge/smartbackend/ethtypes"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func (g *GovStorable) handleAbrogationProposalVotes(logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.Governance.IsGovernanceAbrogationProposalVoteEvent(&log) {
			vote, err := ethtypes.Governance.GovernanceAbrogationProposalVoteEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode abrogation proposal event")
			}
			g.Processed.abrogationVotes = append(g.Processed.abrogationVotes, vote)
		}

		if ethtypes.Governance.IsGovernanceAbrogationProposalVoteEvent(&log) {
			vote, err := ethtypes.Governance.GovernanceAbrogationProposalVoteCancelledEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode abrogation proposal event")
			}

			g.Processed.abrogationCanceledVotes = append(g.Processed.abrogationCanceledVotes, vote)
		}
	}

	return nil
}

func (g *GovStorable) storeProposalAbrogationVotes(ctx context.Context, tx pgx.Tx) error {
	if len(g.Processed.abrogationVotes) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, v := range g.Processed.abrogationVotes {
		power := decimal.NewFromBigInt(v.Power, 0)
		rows = append(rows, []interface{}{
			v.ProposalId.Int64(),
			v.User.String(),
			v.Support,
			power,
			g.block.BlockCreationTime,
			v.Raw.BlockNumber,
			v.Raw.TxHash.String(),
			v.Raw.TxIndex,
			v.Raw.Index,
		})
	}
	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"governance", "abrogation_votes"},
		[]string{"proposal_id", "user_id", "support", "power", "block_timestamp", "included_in_block", "tx_hash", "tx_index", "log_index"},
		pgx.CopyFromSlice(len(g.Processed.abrogationVotes), func(i int) ([]interface{}, error) {
			return []interface{}{rows}, nil
		}),
	)
	if err != nil {
		return errors.Wrap(err, "could not store abrogation proposal  votes")
	}

	return nil
}

func (g *GovStorable) storeAbrogationProposalCanceledVotes(ctx context.Context, tx pgx.Tx) error {
	if len(g.Processed.abrogationCanceledVotes) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, v := range g.Processed.abrogationCanceledVotes {
		rows = append(rows, []interface{}{
			v.ProposalId.Int64(),
			v.User.String(),
			g.block.BlockCreationTime,
			v.Raw.BlockNumber,
			v.Raw.TxHash.String(),
			v.Raw.TxIndex,
			v.Raw.Index,
		})
	}
	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"governance", "abrogation_votes_canceled"},
		[]string{"proposal_id", "user_id", "block_timestamp", "included_in_block", "tx_hash", "tx_index", "log_index"},
		pgx.CopyFromSlice(len(g.Processed.abrogationCanceledVotes), func(i int) ([]interface{}, error) {
			return []interface{}{rows}, nil
		}))
	if err != nil {
		return errors.Wrap(err, "could not store abrogation proposal canceled votes")
	}

	return nil
}
