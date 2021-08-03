package accounterc20transfers

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	_, err := tx.Exec(ctx, `delete from account_erc20_transfers where included_in_block = $1`, s.block.Number)

	return err
}
