package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableHarvest, downCreateTableHarvest)
}

func upCreateTableHarvest(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table compound_provider_harvest
		(
			provider_address     text    not null,
			caller_address       text    not null,
			underlying_got       numeric(78),
			reward_expected      numeric(78),
			underlying_deposited numeric(78),
			fees                 numeric(78),
			reward               numeric(78),
		
			tx_hash              text    not null,
			tx_index             integer not null,
			log_index            integer not null,
		
			block_timestamp      bigint  not null,
			included_in_block    bigint  not null,
		
			created_at           timestamp default now()
		);

	`)

	return err
}

func downCreateTableHarvest(tx *sql.Tx) error {
	_, err := tx.Exec("drop table if exists compound_provider_harvest")
	return err
}
