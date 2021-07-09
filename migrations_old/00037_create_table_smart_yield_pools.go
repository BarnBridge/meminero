package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableSmartYieldPools, downCreateTableSmartYieldPools)
}

func upCreateTableSmartYieldPools(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table smart_yield_pools
		(
			protocol_id           text    not null,
		
			controller_address    text    not null,
			model_address         text    not null,
			provider_address      text    not null,
			sy_address            text    not null,
			oracle_address        text    not null,
		
			junior_bond_address   text    not null,
			senior_bond_address   text    not null,
		
			receipt_token_address text    not null,
			underlying_address    text    not null,
			underlying_symbol     text    not null,
			underlying_decimals   integer not null,
		
			created_at            timestamp default now()
		)
	`)
	return err
}

func downCreateTableSmartYieldPools(tx *sql.Tx) error {
	_, err := tx.Exec("drop table if exists smart_yield_pools")
	return err
}
