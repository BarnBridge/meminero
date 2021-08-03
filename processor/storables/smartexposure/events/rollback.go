package events

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	_, err := tx.Exec(ctx, `delete from smart_exposure.transaction_history where included_in_block = $1`, s.block.Number)
	if err != nil {
		return errors.Wrap(err, "could not execute rollback on smart_exposure.transaction_history")
	}

	_, err = tx.Exec(ctx, `delete from smart_exposure.tranches where start_at_block = $1`, s.block.Number)
	return errors.Wrap(err, "could not execute rollback on smart_exposure.tranches table")
}
