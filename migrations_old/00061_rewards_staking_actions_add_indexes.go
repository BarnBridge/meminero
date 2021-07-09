package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upRewardsStakingActionsAddIndexes, downRewardsStakingActionsAddIndexes)
}

func upRewardsStakingActionsAddIndexes(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create index if not exists smart_yield_rewards_staking_actions_pool_addr_idx on smart_yield_rewards_staking_actions (pool_address, user_address, included_in_block desc, tx_index desc, log_index desc);
	`)
	return err
}

func downRewardsStakingActionsAddIndexes(tx *sql.Tx) error {
	_, err := tx.Exec(`
		drop index if exists smart_yield_rewards_staking_actions_pool_addr_idx;
	`)
	return err
}
