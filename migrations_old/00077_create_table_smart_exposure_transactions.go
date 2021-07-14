package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableSmartExposureTransactions, downCreateTableSmartExposureTransactions)
}

func upCreateTableSmartExposureTransactions(tx *sql.Tx) error {
	_, err := tx.Exec(`
	create table smart_exposure_transactions
	(
	    user_address               text,
		tx_hash                    text    not null,
		tx_index 				   integer not null,
		log_index                  integer not null,
		
		e_token_address			   text not null ,
		amount 					   numeric (78),
		amount_a                   numeric(78),
		amount_b                   numeric(78),
		transaction_type           staking_action_type not null,
		
		block_timestamp			   bigint not null,
		included_in_block          bigint  not null,
		created_at                 timestamp default now()
	);
	
	`)
	return err
}

func downCreateTableSmartExposureTransactions(tx *sql.Tx) error {
	_, err := tx.Exec("drop table if exists smart_exposure_transactions")
	return err
}
