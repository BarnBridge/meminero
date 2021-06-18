package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableNotifications, downCreateTableNotifications)
}

func upCreateTableNotifications(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table "notifications"
		(
		    "id" 				 bigserial primary key,
			"target"             text,
			"type"               text      	not null,
			"starts_on"          bigint,
			"expires_on"         bigint 	not null,
			"message"            text,
			"metadata"           jsonb,
			"included_in_block"  bigint,
			"created_on"		 timestamp 	default now()
		)
		;
		
		create index "notifications_target_starts_on_index"
			on "notifications" ("target" asc, "starts_on" desc)
		;
		
		create index "notifications_included_in_block_index"
			on "notifications" ("included_in_block" desc)
		;
	`)

	return err
}

func downCreateTableNotifications(tx *sql.Tx) error {
	_, err := tx.Exec(`drop table if exists "notifications";`)
	return err
}
