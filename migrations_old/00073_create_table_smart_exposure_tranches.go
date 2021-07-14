package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableSmartExposureTranche, downCreateTableSmartExposureTranche)
}

func upCreateTableSmartExposureTranche(tx *sql.Tx) error {
	_, err := tx.Exec(`
	create table smart_exposure_tranches
	(
		pool_address         text   not null,
		etoken_address       text   not null,

		s_factor_e           numeric(78) not null,	
		target_ratio         numeric(78) not null,
		token_a_ratio        double precision not null,
		token_b_ratio        double precision not null,
	
		start_at_block       bigint not null,
		created_at           timestamp default now()
	)
	`)
	return err
}

func downCreateTableSmartExposureTranche(tx *sql.Tx) error {
	_, err := tx.Exec("drop table if exists smart_exposure_tranches")
	return err
}
