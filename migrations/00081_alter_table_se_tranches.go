package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upAlterTableSeTranches, downAlterTableSeTranches)
}

func upAlterTableSeTranches(tx *sql.Tx) error {
	_, err := tx.Exec(`
		alter table smart_exposure_tranches add column etoken_symbol text;
	`)
	return err
}

func downAlterTableSeTranches(tx *sql.Tx) error {
	_, err := tx.Exec(`
		alter table smart_exposure_tranches drop column if exists etoken_symbol;
	`)
	return err
}
