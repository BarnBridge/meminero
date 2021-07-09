package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableBarnStakingActions, downCreateTableBarnStakingActions)
}

func upCreateTableBarnStakingActions(tx *sql.Tx) error {
	_, err := tx.Exec(`
	create type action_type as enum('DEPOSIT','WITHDRAW');
	create table barn_staking_actions
	(
		tx_hash                    text    not null,
		tx_index 				   integer not null,
		log_index                  integer not null,
		address					   text not null,
		user_address			   text not null,
		action_type				   action_type not null,
		amount					   numeric(78) not null,
		balance_after			   numeric(78) not null,
		included_in_block          bigint  not null,
		created_at                 timestamp default now()
	);
	
	`)
	return err
}

func downCreateTableBarnStakingActions(tx *sql.Tx) error {
	_, err := tx.Exec("drop table barn_staking_actions")
	return err
}
