package boilerplate

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.transfers) == 0 {
		return nil
	}

	return nil
}
