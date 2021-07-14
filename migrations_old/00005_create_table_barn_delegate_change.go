package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableBarnDelegateChange, downCreateTableBarnDelegateChange)
}

func upCreateTableBarnDelegateChange(tx *sql.Tx) error {
	_, err := tx.Exec(`
	create type delegate_change_type as enum('INCREASE','DECREASE');

	create table barn_delegate_changes
	(
		tx_hash                    text    not null,
		tx_index                   integer not null,
		log_index                  integer not null,
		logged_by                  text not null,
		
		action_type                   delegate_change_type not null,
		sender                        text not null,
		receiver                      text not null,
		amount                        numeric(78) not null,
		receiver_new_delegated_power  numeric(78) not null,
		timestamp                     bigint,
		
		included_in_block          bigint  not null,
		created_at                 timestamp default now()
	);
	
	`)
	return err
}

func downCreateTableBarnDelegateChange(tx *sql.Tx) error {
	_, err := tx.Exec("drop table barn_delegate_changes")
	return err
}
