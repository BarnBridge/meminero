package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableTokensPrices, downCreateTableTokensPrices)
}

func upCreateTableTokensPrices(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table tokens_prices
		(
			token_address     text             not null,
			price_usd         double precision not null,
		
			block_timestamp   bigint           not null,
			included_in_block bigint           not null,
		
			created_at        timestamp default now()
		);
	`)

	return err
}

func downCreateTableTokensPrices(tx *sql.Tx) error {
	_, err := tx.Exec("drop table if exists tokens_prices")
	return err
}
