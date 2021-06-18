package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableSmartYieldState, downCreateTableSmartYieldState)
}

func upCreateTableSmartYieldState(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table smart_yield_state
		(
			included_in_block  bigint           not null,
			block_timestamp    timestamp        not null,
		
			pool_address       text             not null,
		
			senior_liquidity   numeric(78),
			junior_liquidity   numeric(78),
			jtoken_price       numeric(78),
		
			abond_principal    numeric(78),
			abond_gain         numeric(78),
			abond_issued_at    bigint,
			abond_matures_at   bigint,
			abond_apy          double precision not null,

			senior_apy         double precision not null,
			junior_apy         double precision not null,
			originator_apy     double precision not null,
			originator_net_apy double precision not null
		)
	`)
	return err
}

func downCreateTableSmartYieldState(tx *sql.Tx) error {
	_, err := tx.Exec(`drop table if exists smart_yield_state;`)
	return err
}
