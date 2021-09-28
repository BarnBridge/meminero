package glue

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// getHighestBlock returns the highest block inserted into the database
func (g *Glue) getHighestBlock(ctx context.Context) (int64, error) {
	var block int64

	err := g.db.QueryRow(ctx, "select number from blocks order by number desc limit 1").Scan(&block)
	if err != nil && err != pgx.ErrNoRows {
		return 0, err
	}

	return block, nil
}
