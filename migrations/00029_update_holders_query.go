package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upUpdateHoldersQuery, downUpdateHoldersQuery)
}

func upUpdateHoldersQuery(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create or replace view bond_users_with_balance_no_staking as
		with transfers as ( select sender as address, -value as amount
							from bond_transfers
							where sender not in ('0xb0fa2beee3cf36a7ac7e99b885b48538ab364853', '0x10e138877df69ca44fdc68655f86c88cde142d7f')
							  and receiver not in ('0xb0fa2beee3cf36a7ac7e99b885b48538ab364853', '0x10e138877df69ca44fdc68655f86c88cde142d7f')
							union all
							select receiver as addrress, value as amount
							from bond_transfers
							where sender not in ('0xb0fa2beee3cf36a7ac7e99b885b48538ab364853', '0x10e138877df69ca44fdc68655f86c88cde142d7f')
							  and receiver not in ('0xb0fa2beee3cf36a7ac7e99b885b48538ab364853', '0x10e138877df69ca44fdc68655f86c88cde142d7f') )
		select address, sum(amount) as balance
		from transfers
		group by address;
	`)
	// This code is executed when the migration is applied.
	return err
}

func downUpdateHoldersQuery(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
