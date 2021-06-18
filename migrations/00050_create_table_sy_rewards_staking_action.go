package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableSyRewardsStakingAction, downCreateTableSyRewardsStakingAction)
}

func upCreateTableSyRewardsStakingAction(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create type reward_action as enum('JUNIOR_STAKE','JUNIOR_UNSTAKE');

		create table smart_yield_rewards_staking_actions
		(
		    user_address 		text not null,
		    amount		 		numeric(78),
		    balance_after		numeric(78),
		    action_type			reward_action not null,
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

func downCreateTableSyRewardsStakingAction(tx *sql.Tx) error {
	_, err := tx.Exec(`drop table if exists smart_yield_rewards_staking_action;`)
	return err
}
