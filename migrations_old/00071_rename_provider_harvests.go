package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upRenameProviderHarvests, downRenameProviderHarvests)
}

func upRenameProviderHarvests(tx *sql.Tx) error {
	_, err := tx.Exec(`
		alter table compound_controller_harvests add column protocol_id text not null default 'compound/v2';
		alter table compound_controller_harvests rename to sy_controller_harvests;
	`)
	return err
}

func downRenameProviderHarvests(tx *sql.Tx) error {
	_, err := tx.Exec(`
		alter table sy_controller_harvests drop column if exists protocol_id;
		alter table sy_controller_harvests rename to compound_controller_harvests;
	`)
	return err
}
