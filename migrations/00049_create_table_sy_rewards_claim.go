package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableSyRewardsClaim, downCreateTableSyRewardsClaim)
}

func upCreateTableSyRewardsClaim(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table smart_yield_rewards_claims
		(
		    user_address 		text not null,
		    amount		 		numeric(78),
		    pool_address		text not null,
		    
		    tx_hash				text not null,
		    tx_index          integer not null,
			log_index         integer not null,
		
			block_timestamp   bigint  not null,
			included_in_block bigint  not null
		)
	`)
	return err
}

func downCreateTableSyRewardsClaim(tx *sql.Tx) error {
	_, err := tx.Exec(`drop table if exists smart_yield_rewards_claims;`)
	return err
}
