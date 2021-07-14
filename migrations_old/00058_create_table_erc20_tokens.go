package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableTokens, downCreateTableTokens)
}

func upCreateTableTokens(tx *sql.Tx) error {
	_, err := tx.Exec(`
	create table erc20_tokens
	(
	    token_address text,
	    symbol text,
	    decimals bigint,
	
		created_at            timestamp default now()
	);
	`)
	return err
}

func downCreateTableTokens(tx *sql.Tx) error {
	_, err := tx.Exec("drop table if exists erc20_tokens")
	return err
}
