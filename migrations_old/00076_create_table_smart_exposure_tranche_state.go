package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableSmartExposureTrancheState, downCreateTableSmartExposureTrancheState)
}

func upCreateTableSmartExposureTrancheState(tx *sql.Tx) error {
	_, err := tx.Exec(`
	create table smart_exposure_tranche_state
	(
		included_in_block     bigint    not null,
		block_timestamp       timestamp not null,
	
		pool_address          text      not null,
		e_token_address       text      not null,
	
		token_a_liquidity     double precision,
		token_b_liquidity     double precision,
		current_ratio         numeric(78),
		amount_a_conversion   numeric(78),
		amount_b_conversion   numeric(78),
		etoken_price          double precision,
		token_a_current_ratio double precision,
		token_b_current_ratio double precision
	)
	`)
	return err
}

func downCreateTableSmartExposureTrancheState(tx *sql.Tx) error {
	_, err := tx.Exec(`drop table if exists smart_exposure_tranche_state;`)
	return err
}
