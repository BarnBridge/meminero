package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableSmartYieldJuniorBuy, downCreateTableSmartYieldJuniorBuy)
}

func upCreateTableSmartYieldJuniorBuy(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table smart_yield_junior_buy
		(
			sy_address        text    not null,
			buyer_address     text    not null,
			junior_bond_address text not null,
			junior_bond_id    bigint  not null,
			tokens_in         numeric(78),
			matures_at        bigint,
		
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

func downCreateTableSmartYieldJuniorBuy(tx *sql.Tx) error {
	_, err := tx.Exec("drop table if exists smart_yield_junior_buy")
	return err
}
