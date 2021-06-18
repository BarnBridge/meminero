package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableMonitoredAccounts, downCreateTableMonitoredAccounts)
}

func upCreateTableMonitoredAccounts(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table monitored_accounts(
		    address text not null,
		    
		    created_at            timestamp default now()
		)
	`)
	return err
}

func downCreateTableMonitoredAccounts(tx *sql.Tx) error {
	_, err := tx.Exec(`
		drop table if exists monitored_accounts;
		`)
	return err
}
