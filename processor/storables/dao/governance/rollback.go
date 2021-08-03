package governance

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	b := &pgx.Batch{}
	tables := []string{"proposals", "abrogation_proposals", "proposal_events", "votes", "votes_canceled", "abrogation_votes", "abrogation_votes_canceled"}
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
