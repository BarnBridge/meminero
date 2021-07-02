package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableSmartExposurePoolState, downCreateTableSmartExposurePoolState)
}

func upCreateTableSmartExposurePoolState(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table smart_exposure_pool_state
		(
			included_in_block     bigint    not null,
			block_timestamp       timestamp not null,
		
			pool_address          text      not null,
		
			pool_liquidity        double precision,
			last_rebalance        bigint,
			rebalancing_interval  bigint,
			rebalancing_condition numeric(78)
		)
	`)
	return err
}

func downCreateTableSmartExposurePoolState(tx *sql.Tx) error {
	_, err := tx.Exec(`drop table if exists smart_exposure_pool_state;`)
	return err
}
