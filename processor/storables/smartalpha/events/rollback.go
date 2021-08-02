package events

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
)

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.WithField("block", s.block.Number).Trace("rolling back block")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Trace("done rolling back block")
	}()

	b := &pgx.Batch{}
	tables := []string{
		"user_join_entry_queue_events",
		"user_join_exit_queue_events",
		"user_redeem_tokens_events",
		"user_redeem_underlying_events",
		"transaction_history",
		"epoch_end_events",
	}
	for _, t := range tables {
		query := fmt.Sprintf(`delete from smart_alpha.%s where included_in_block = $1`, t)
		b.Queue(query, s.block.Number)
	}

	br := tx.SendBatch(ctx, b)
	_, err := br.Exec()
	if err != nil {
		return err
	}

	return br.Close()
}