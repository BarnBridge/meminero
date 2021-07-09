package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableSmartYieldSeniorBuy, downCreateTableSmartYieldSeniorBuy)
}

func upCreateTableSmartYieldSeniorBuy(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table smart_yield_senior_buy
		(
			sy_address        text    not null,
			buyer_address     text    not null,
			senior_bond_address text  not null,
			senior_bond_id    numeric(78)  not null,
			underlying_in     numeric(78),
			gain              numeric(78),
			for_days          bigint,
		
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

func downCreateTableSmartYieldSeniorBuy(tx *sql.Tx) error {
	_, err := tx.Exec("drop table if exists smart_yield_senior_buy")
	return err
}
