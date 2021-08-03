package barn

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	b := &pgx.Batch{}
	tables := []string{"barn_delegate_actions", "barn_delegate_changes", "barn_locks", "barn_staking_actions"}
	for _, t := range tables {
		query := fmt.Sprintf(`delete from governance.%s where included_in_block = $1`, t)
		b.Queue(query, s.block.Number)
	}

	br := tx.SendBatch(ctx, b)
	_, err := br.Exec()
	if err != nil {
		return err
	}

	err = br.Close()
	return err
}
