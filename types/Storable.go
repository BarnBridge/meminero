package types

import (
	"database/sql"
)

// Storable
// role: a Storable serves as a means of transforming raw data and inserting it into the database
// input: raw Ethereum data + a database transaction
// output: processed/derived/enhanced data stored directly to the db
type Storable interface {
	Execute() error
	Rollback(tx *sql.Tx) error
	SaveToDatabase(tx *sql.Tx) error
	Result() interface{}
}
