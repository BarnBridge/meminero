package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateSyIndexes, downCreateSyIndexes)
}

func upCreateSyIndexes(tx *sql.Tx) error {
	_, err := tx.Exec(`
create index if not exists prices_protocol_id_token_address_idx on smart_yield_prices (protocol_id, token_address, block_timestamp desc);
create index if not exists smart_yield_state_pool_address_idx on smart_yield_state(pool_address, block_timestamp desc);

create index if not exists jtoken_transfers_sender_idx on jtoken_transfers (sender, block_timestamp desc);
create index if not exists jtoken_transfers_receiver_idx on jtoken_transfers (receiver, block_timestamp desc);

create index if not exists smart_yield_junior_buy_junior_bond_address_id_idx on smart_yield_junior_buy (junior_bond_address, junior_bond_id);
create index if not exists smart_yield_junior_redeem_junior_bond_address_id_idx on smart_yield_junior_redeem (junior_bond_address, junior_bond_id, block_timestamp desc);
create index if not exists sy_junior_redeem_user_address_idx on smart_yield_junior_redeem (owner_address);

create index if not exists smart_yield_senior_buy_junior_bond_address_id_idx on smart_yield_senior_buy (senior_bond_address, senior_bond_id);
create index if not exists smart_yield_senior_redeem_junior_bond_address_id_idx on smart_yield_senior_redeem (senior_bond_address, senior_bond_id, block_timestamp desc);
create index if not exists sy_senior_redeem_user_address_idx on smart_yield_senior_redeem (owner_address);

create index if not exists erc721_transfers_token_address_id_idx on erc721_transfers (token_address, token_id, block_timestamp desc);
create index if not exists erc721_transfers_token_type_receiver_idx on erc721_transfers (token_type, receiver, block_timestamp desc);

create index if not exists sy_tx_history_user_address_idx on smart_yield_transaction_history (user_address, block_timestamp desc, tx_index desc, log_index desc);

create index if not exists sy_state_apy_trend_idx on smart_yield_state (pool_address, date_trunc('day', block_timestamp));
	`)
	return err
}

func downCreateSyIndexes(tx *sql.Tx) error {
	_, err := tx.Exec(`
		drop index if exists prices_protocol_id_token_address_idx;
		drop index if exists smart_yield_state_pool_address_idx;
		drop index if exists jtoken_transfers_sender_idx;
		drop index if exists jtoken_transfers_receiver_idx;
		drop index if exists smart_yield_junior_buy_junior_bond_address_id_idx;
		drop index if exists smart_yield_junior_redeem_junior_bond_address_id_idx;
		drop index if exists smart_yield_senior_buy_junior_bond_address_id_idx;
		drop index if exists smart_yield_senior_redeem_junior_bond_address_id_idx;
		drop index if exists erc721_transfers_token_address_id_idx;
		drop index if exists erc721_transfers_token_type_receiver_idx;
		drop index if exists sy_tx_history_user_address_idx;
		drop index if exists sy_junior_redeem_user_address_idx;
		drop index if exists sy_senior_redeem_user_address_idx;
		drop index if exists sy_state_apy_trend_idx;
	`)
	return err
}
