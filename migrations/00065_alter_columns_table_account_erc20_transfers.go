package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upAlterColumnsTableAccountErc20Transfers, downAlterColumnsTableAccountErc20Transfers)
}

func upAlterColumnsTableAccountErc20Transfers(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create type transfer_type as enum('IN','OUT');
		alter table account_erc20_transfers rename column sender to account;
		alter table account_erc20_transfers rename column receiver to counterparty;
		alter table account_erc20_transfers rename column "value" to amount;

		alter table account_erc20_transfers add column tx_direction transfer_type;
		
		create index if not exists account_erc20_transfers_account_addr_idx on  account_erc20_transfers (account, included_in_block desc, tx_index desc, log_index desc);
`)
	return err
}

func downAlterColumnsTableAccountErc20Transfers(tx *sql.Tx) error {
	_, err := tx.Exec(`
		alter table account_erc20_transfers rename column account to sender;
		alter table account_erc20_transfers rename column counterparty to  receiver;
		alter table account_erc20_transfers rename column amount  to "value";
		alter table account_erc20_transfers  drop column if exists tx_direction;
		drop type if exists transfer_type;
		drop index if exists account_erc20_transfers_account_addr_idx;
	`)
	return err
}
