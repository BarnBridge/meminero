package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(UpAlterDeleteBlockFunction, DownAlterDeleteBlockFunction)
}

func UpAlterDeleteBlockFunction(tx *sql.Tx) error {
	_, err := tx.Exec(`

	create or replace function delete_block(in block_number bigint) returns void as
	$body$
	declare
		tables varchar[];
		tbl    varchar;
	begin
		tables := array [
			'bond_transfers',
			'barn_staking_actions',
			'barn_locks',
			'barn_delegate_changes',
			'barn_delegate_actions',
			'governance_proposals',
			'governance_events',
			'governance_votes',
			'governance_votes_canceled',
			'governance_abrogation_proposals',
			'governance_abrogation_votes',
			'governance_abrogation_votes_canceled',
		    'yield_farming_actions',
		    'compound_provider_harvest',
		    'compound_provider_transfer_fees',
		    'erc721_transfers',
		    'jtoken_transfers',
		    'smart_yield_junior_buy',
		    'smart_yield_junior_redeem',
		    'smart_yield_prices',
		    'smart_yield_senior_buy',
		    'smart_yield_senior_redeem',
		    'smart_yield_state',
		    'smart_yield_token_buy',
		    'smart_yield_token_sell',
		    'smart_yield_transaction_history'
			];
	
		foreach tbl in array tables
			loop
				perform __delete_entity(tbl, block_number);
			end loop;

		delete from blocks where number = block_number;
	end;
	$body$ language 'plpgsql';
	`)
	return err
}

func DownAlterDeleteBlockFunction(tx *sql.Tx) error {
	return nil
}
