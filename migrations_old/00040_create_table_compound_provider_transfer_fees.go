package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableCompoundProviderTransferFees, downCreateTableCompoundProviderTransferFees)
}

func upCreateTableCompoundProviderTransferFees(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table compound_provider_transfer_fees
		(
			provider_address  text    not null,
			caller_address    text    not null,
			fees_owner        text    not null,
			fees              numeric(78),
		
			tx_hash           text    not null,
			tx_index          integer not null,
			log_index         integer not null,
		
			block_timestamp   bigint  not null,
			included_in_block bigint  not null,
		
			created_at        timestamp default now()
		);
	`)

	return err
}

func downCreateTableCompoundProviderTransferFees(tx *sql.Tx) error {
	_, err := tx.Exec("drop table if exists compound_provider_transfer_fees")
	return err
}
