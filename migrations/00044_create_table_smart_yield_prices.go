package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableSmartYieldPrices, downCreateTableSmartYieldPrices)
}

func upCreateTableSmartYieldPrices(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table smart_yield_prices
		(
			protocol_id       text             not null,
			token_address     text             not null,
			token_symbol      text,
			price_usd         double precision not null,
		
			block_timestamp   bigint           not null,
			included_in_block bigint           not null,
		
			created_at        timestamp default now()
		);
	`)

	return err
}

func downCreateTableSmartYieldPrices(tx *sql.Tx) error {
	_, err := tx.Exec("drop table if exists smart_yield_prices")
	return err
}
