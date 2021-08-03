package state

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	_, err := tx.Exec(ctx, `delete from smart_yield.pool_state where included_in_block = $1`, s.block.Number)
	if err != nil {
		return errors.Wrap(err, "could not delete from smart_yield.pool_state")
	}

	return nil
}
