package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upAddIndexesSmartExposure, downAddIndexesSmartExposure)
}

func upAddIndexesSmartExposure(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create index if not exists smart_exposure_pool_state_pool_address_idx on smart_exposure_pool_state (pool_address, block_timestamp desc);	
		create index if not exists tokens_prices_token_address_idx on tokens_prices (token_address, block_timestamp desc);
		create index if not exists smart_exposure_tranche_state_e_token_address_idx on smart_exposure_tranche_state (e_token_address, block_timestamp desc);
		create index if not exists smart_exposure_transactions_user_address_chronological_idx on smart_exposure_transactions (user_address, included_in_block desc, tx_index desc, log_index desc);
		create index if not exists smart_exposure_transactions_e_token_address_chronological_idx on smart_exposure_transactions (e_token_address, included_in_block desc, tx_index desc, log_index desc);
		create index if not exists etoken_transfers_sender_idx on etoken_transfers (sender, block_timestamp desc);
		create index if not exists etoken_transfers_receiver_idx on etoken_transfers (receiver, block_timestamp desc);
	`)
	return err
}

func downAddIndexesSmartExposure(tx *sql.Tx) error {
	_, err := tx.Exec(`
		drop index if exists smart_exposure_pool_state_pool_address_idx;
		drop index if exists tokens_prices_token_address_idx;
		drop index if exists smart_exposure_tranche_state_e_token_address_idx;
		drop index if exists smart_exposure_transactions_user_address_chronological_idx;
		drop index if exists smart_exposure_transactions_e_token_address_chronological_idx;
		drop index if exists etoken_transfers_sender_idx;
		drop index if exists etoken_transfers_receiver_idx;
	`)
	return err
}
