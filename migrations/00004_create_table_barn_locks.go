package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableBarnLocks, downCreateTableBarnLocks)
}

func upCreateTableBarnLocks(tx *sql.Tx) error {
	_, err := tx.Exec(`
	create table barn_locks
	(
		tx_hash                    text    not null,
		tx_index                   integer not null,
		log_index                  integer not null,
		logged_by                  text not null,
		user_address               text not null,
		locked_until               bigint,
		locked_at                  bigint,
		included_in_block          bigint  not null,
		created_at                 timestamp default now()
	);
	
	`)
	return err
}

func downCreateTableBarnLocks(tx *sql.Tx) error {
	_, err := tx.Exec("drop table barn_locks")
	return err
}
