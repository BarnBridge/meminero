package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableSmartYieldSeniorRedeem, downCreateTableSmartYieldSeniorRedeem)
}

func upCreateTableSmartYieldSeniorRedeem(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table smart_yield_senior_redeem
		(
			sy_address        text    not null,
			owner_address     text    not null,
			senior_bond_address text  not null,
			senior_bond_id    bigint  not null,
			fee               numeric(78)  not null,
		
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

func downCreateTableSmartYieldSeniorRedeem(tx *sql.Tx) error {
	_, err := tx.Exec("drop table if exists smart_yield_senior_redeem")
	return err
}
