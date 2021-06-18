package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableEtokenTransfers, downCreateTableEtokenTransfers)
}

func upCreateTableEtokenTransfers(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table etoken_transfers
		(
            e_token_address   text    not null,
			sender            text    not null,
			receiver          text    not null,
			value             numeric(78),
		
			tx_hash           text    not null,
			tx_index          integer not null,
			log_index         integer not null,
		
			block_timestamp   bigint  not null,
			included_in_block bigint  not null,
		
			created_at        timestamp default now()
		);
	`)
	return err
}

func downCreateTableEtokenTransfers(tx *sql.Tx) error {
	_, err := tx.Exec("drop table if exists etoken_transfers")
	return err
}
