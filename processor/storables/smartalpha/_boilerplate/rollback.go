package events

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
)

func (s *Storable) Rollback(ctx context.Context, pgx pgx.Tx) error {
	start := time.Now()
	s.logger.WithField("block", s.block.Number).Trace("rolling back block")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Trace("done rolling back block")
	}()

	return nil
}
