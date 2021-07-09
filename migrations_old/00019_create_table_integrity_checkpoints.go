package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableIntegrityCheckpoints, downCreateTableIntegrityCheckpoints)
}

func upCreateTableIntegrityCheckpoints(tx *sql.Tx) error {
	_, err := tx.Exec(`
	create table integrity_checkpoints (
	    number bigint,
	    created_at timestamp default now() 
	);
	
	create index on integrity_checkpoints (created_at desc); 
	`)

	return err
}

func downCreateTableIntegrityCheckpoints(tx *sql.Tx) error {
	_, err := tx.Exec(`
	drop table if exists integrity_checkpoints;
	`)

	return err
}
