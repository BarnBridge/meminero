package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableLabels, downCreateTableLabels)
}

func upCreateTableLabels(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table labels(
		    address text not null,
		    label	text not null
		)
	`)
	return err
}

func downCreateTableLabels(tx *sql.Tx) error {
	_, err := tx.Exec(`
		drop table if exists labels;
		`)
	return err
}
