package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableNotificationJobs, downCreateTableNotificationJobs)
}

func upCreateTableNotificationJobs(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table "notification_jobs"
		(
		    "id" 				 bigserial primary key,
			"type"               text      	not null,
			"execute_on"         bigint,
			"metadata"           jsonb,
			"included_in_block"  bigint,
			"deleted"			 boolean 	default FALSE,
			"created_on"		 timestamp 	default now()
		)
		;
		
		create index "notification_jobs_execute_on_index"
			on "notification_jobs" ("execute_on" desc)
		;

		create index "notification_jobs_deleted_on_index"
			on "notification_jobs" ("deleted" desc)
		;

		create index "notification_jobs_included_in_block_index"
			on "notification_jobs" ("included_in_block" desc)
		;
	`)

	return err
}

func downCreateTableNotificationJobs(tx *sql.Tx) error {
	_, err := tx.Exec(`drop table if exists "notification_jobs";`)
	return err
}
