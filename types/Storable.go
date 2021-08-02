package types

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// Storable
// role: a Storable serves as a means of transforming raw data and inserting it into the database
// input: raw Ethereum data + a database transaction
// output: processed/derived/enhanced data stored directly to the db
type Storable interface {
	Execute(ctx context.Context) error
	Rollback(ctx context.Context, pgx pgx.Tx) error
	SaveToDatabase(ctx context.Context, tx pgx.Tx) error
	Result() interface{}
}
