package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableWithdrawals, downCreateTableWithdrawals)
}

func upCreateTableWithdrawals(tx *sql.Tx) error {
	_, err := tx.Exec(`
	create type staking_action_type as enum('DEPOSIT','WITHDRAW');
	create table yield_farming_actions
	(
		tx_hash                    text    not null,
		tx_index 				   integer not null,
		log_index                  integer not null,
		user_address 			   text not null ,
		token_address			   text not null,
		amount 					   numeric (78),
		action_type				   staking_action_type not null,
		
		block_timestamp			   bigint not null,
		included_in_block          bigint  not null,
		created_at                 timestamp default now()
	);
	
	`)
	return err
}

func downCreateTableWithdrawals(tx *sql.Tx) error {
	_, err := tx.Exec("drop table if exists yield_farming_actions")
	return err
}
