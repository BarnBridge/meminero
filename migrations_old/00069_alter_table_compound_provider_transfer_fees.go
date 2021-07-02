package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upAlterTableCompoundProviderTransferFees, downAlterTableCompoundProviderTransferFees)
}

func upAlterTableCompoundProviderTransferFees(tx *sql.Tx) error {
	_, err := tx.Exec(`
		alter table compound_provider_transfer_fees add column protocol_id text not null default 'compound/v2';
		alter table compound_provider_transfer_fees rename to provider_transfer_fees;
	`)
	return err

}

func downAlterTableCompoundProviderTransferFees(tx *sql.Tx) error {
	_, err := tx.Exec(`
		alter table provider_transfer_fees drop column if exists protocol_id;
		alter table provider_transfer_fees rename to compound_provider_transfer_fees;
	`)
	return err

}
