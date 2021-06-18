package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableBarnDelegateActions, downCreateTableBarnDelegateActions)
}

func upCreateTableBarnDelegateActions(tx *sql.Tx) error {
	_, err := tx.Exec(`
	create type delegate_action_type as enum('START','STOP');

	create table barn_delegate_actions
	(
		tx_hash                    text    not null,
		tx_index                   integer not null,
		log_index                  integer not null,
		logged_by                  text    not null,
		
		sender                        text                 not null,
		receiver                      text                 not null,
		action_type                   delegate_action_type not null,
		timestamp                     bigint               not null,
		
		included_in_block          bigint  not null,
		created_at                 timestamp default now()
	);
	
	`)
	return err
}

func downCreateTableBarnDelegateActions(tx *sql.Tx) error {
	_, err := tx.Exec("drop table barn_delegate_actions")
	return err
}
