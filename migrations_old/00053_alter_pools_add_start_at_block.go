package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upAlterPoolsAddStartAtBlock, downAlterPoolsAddStartAtBlock)
}

func upAlterPoolsAddStartAtBlock(tx *sql.Tx) error {
	_, err := tx.Exec(`
		alter table smart_yield_pools add column start_at_block bigint not null default 0;
		alter table smart_yield_reward_pools add column start_at_block bigint not null default 0;
	`)
	return err
}

func downAlterPoolsAddStartAtBlock(tx *sql.Tx) error {
	_, err := tx.Exec(`
		alter table smart_yield_pools drop column if exists start_at_block;
		alter table smart_yield_reward_pools drop column if exists start_at_block;
	`)
	return err
}
