package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upAlterTableRewardPools, downAlterTableRewardPools)
}

func upAlterTableRewardPools(tx *sql.Tx) error {
	_, err := tx.Exec(`
		alter table smart_yield_reward_pools 
		    alter column reward_token_address type text[] using array[reward_token_address];
		alter table smart_yield_reward_pools
			rename column reward_token_address to reward_token_addresses;

		create type reward_pool_type as enum ('SINGLE', 'MULTI');
		alter table smart_yield_reward_pools 
			add column pool_type reward_pool_type default 'SINGLE';

		alter table smart_yield_rewards_claims
			add column reward_token_address text;
	`)
	return err
}

func downAlterTableRewardPools(tx *sql.Tx) error {
	return nil
}
