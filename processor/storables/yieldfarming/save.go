package yieldfarming

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.stakingActions) == 0 {
		return nil
	}

	var rows [][]interface{}

	for _, t := range s.processed.stakingActions {
		rows = append(rows, []interface{}{
			t.UserAddress,
			t.TokenAddress,
			t.Amount,
			t.ActionType,
			s.block.BlockCreationTime,
			s.block.Number,
			t.TransactionHash,
			t.TransactionIndex,
			t.LogIndex,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"yield_farming", "transactions"},
		[]string{"user_address", "token_address", "amount", "action_type", "block_timestamp", "included_in_block", "tx_hash", "tx_index", "log_index"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not execute copy")
	}

	return nil
}
