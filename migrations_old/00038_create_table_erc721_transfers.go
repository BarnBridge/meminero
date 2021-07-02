package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableERC721Transfers, downCreateTableERC721Transfers)
}

func upCreateTableERC721Transfers(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table erc721_transfers
		(
			token_address     text    not null,
			token_type 		  text    not null,
			sender            text    not null,
			receiver          text    not null,
			token_id          bigint  not null,
		
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

func downCreateTableERC721Transfers(tx *sql.Tx) error {
	_, err := tx.Exec("drop table if exists erc721_transfers")
	return err
}
