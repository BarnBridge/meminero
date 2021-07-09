package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableAccountErc20Transfers, downCreateTableAccountErc20Transfers)
}

func upCreateTableAccountErc20Transfers(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table account_erc20_transfers
		(
		    token_address		text not null,
		    sender            	text    not null,
			receiver          	text    not null,
			value             	numeric(78),
		    
		    tx_hash				text not null,
		    tx_index          	integer not null,
			log_index         	integer not null,
		
			block_timestamp   bigint  not null,
			included_in_block bigint  not null
		);
	`)
	return err
}

func downCreateTableAccountErc20Transfers(tx *sql.Tx) error {
	_, err := tx.Exec(`drop table if exists account_erc20_transfers;`)
	return err
}
