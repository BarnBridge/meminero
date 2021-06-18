package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableTokensPrice, downCreateTableTokensPrice)
}

func upCreateTableTokensPrice(tx *sql.Tx) error {
	_, err := tx.Exec(`
	create table tokens
	(
		address             text   not null,
		symbol              text   not null,
		decimals            bigint not null,
		aggregator_address  text   not null,
		price_provider_type text   not null,
		created_at          timestamp default now()
	)
	`)
	return err
}

func downCreateTableTokensPrice(tx *sql.Tx) error {
	_, err := tx.Exec("drop table if exists tokens")
	return err
}
