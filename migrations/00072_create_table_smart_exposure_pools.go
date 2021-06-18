package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableSmartExposurePools, downCreateTableSmartExposurePools)
}

func upCreateTableSmartExposurePools(tx *sql.Tx) error {
	_, err := tx.Exec(`
	create table smart_exposure_pools
	(
		pool_address    text   not null,
		pool_name       text   not null,
	
		token_a_address  text   not null,
		token_a_symbol   text   not null,
		token_a_decimals bigint not null,
	
		token_b_address  text   not null,
		token_b_symbol   text   not null,
		token_b_decimals bigint not null,
	
		start_at_block  bigint not null,
		created_at      timestamp default now()
	)
	`)
	return err
}

func downCreateTableSmartExposurePools(tx *sql.Tx) error {
	_, err := tx.Exec("drop table if exists smart_exposure_pools")
	return err
}
