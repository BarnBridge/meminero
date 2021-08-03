package events

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	b := &pgx.Batch{}
	tables := []string{"junior_2step_redeem_events", "junior_2step_withdraw_events", "junior_entry_events", "junior_instant_withdraw_events", "senior_entry_events", "senior_redeem_events", "transaction_history", "controller_harvests", "provider_transfer_fees"}
	for _, t := range tables {
		query := fmt.Sprintf(`delete from smart_yield.%s where included_in_block = $1`, t)
		b.Queue(query, s.block.Number)
	}

	br := tx.SendBatch(ctx, b)
	_, err := br.Exec()
	if err != nil {
		return err
	}

	return br.Close()
}
