package processor

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"

)

// checkBlockExists verifies if the current block matches any other block in the database by hash
func (p *Processor) checkBlockExists(ctx context.Context, db *pgxpool.Pool) (bool, error) {
	var count int
	err := db.QueryRow(ctx, `select count(*) from blocks where block_hash = $1`, p.Block.BlockHash).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "could not query database for block with same hash")
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

// checkBlockReorged verifies if the current block matches any block in the database on number
// this is meant to be used in order to detect if the database contains a blocks with the same number
// but different hash if the checkBlockExists function returns false
func (p *Processor) checkBlockReorged(ctx context.Context, db *pgxpool.Pool) (bool, error) {
	var count int
	err := db.QueryRow(ctx, `select count(*) from blocks where number = $1`, p.Block.Number).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "could not query database for block with same number")
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
