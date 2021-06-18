package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upAlterTableBlocks, downAlterTableBlocks)
}

func upAlterTableBlocks(tx *sql.Tx) error {
	_, err := tx.Exec(`
		alter table blocks drop column block_gas_limit;
		alter table blocks drop column block_gas_used;
		alter table blocks drop column block_difficulty;
		alter table blocks drop column total_block_difficulty;
		alter table blocks drop column block_extra_data;
		alter table blocks drop column block_mix_hash;
		alter table blocks drop column block_nonce;
		alter table blocks drop column block_size;
		alter table blocks drop column block_logs_bloom;
		alter table blocks drop column includes_uncle;
		alter table blocks drop column has_beneficiary;
		alter table blocks drop column has_receipts_trie;
		alter table blocks drop column has_tx_trie;
		alter table blocks drop column sha3_uncles;
		alter table blocks drop column number_of_uncles;
		alter table blocks drop column number_of_txs;
	`)
	return err
}

func downAlterTableBlocks(tx *sql.Tx) error {
	return nil
}
