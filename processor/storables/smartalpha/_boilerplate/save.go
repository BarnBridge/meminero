package events

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
)

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.Trace("storing")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Trace("done storing")
	}()

	return nil
}
