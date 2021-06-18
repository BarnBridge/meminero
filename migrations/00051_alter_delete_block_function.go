package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(UpAlterDeleteBlock2Function, DownAlterDeleteBlock2Function)
}

func UpAlterDeleteBlock2Function(tx *sql.Tx) error {
	_, err := tx.Exec(`

	create or replace function delete_block(in block_number bigint) returns void as
	$body$
	declare
		tables varchar[];
		tbl    varchar;
	begin
		tables := array [
			'barn_delegate_actions',
			'barn_delegate_changes',
			'barn_locks',
			'barn_staking_actions',
			'bond_transfers',
		    'compound_controller_harvests',
		    'compound_provider_transfer_fees',
		    'erc721_transfers',
			'governance_abrogation_proposals',
			'governance_abrogation_votes',
			'governance_abrogation_votes_canceled',
			'governance_events',
			'governance_proposals',
			'governance_votes',
			'governance_votes_canceled',
		    'jtoken_transfers',
		    'smart_yield_junior_buy',
		    'smart_yield_junior_redeem',
		    'smart_yield_prices',
		    'smart_yield_rewards_claims',
		    'smart_yield_rewards_staking_actions',
		    'smart_yield_senior_buy',
		    'smart_yield_senior_redeem',
		    'smart_yield_state',
		    'smart_yield_token_buy',
		    'smart_yield_token_sell',
		    'smart_yield_transaction_history',
		    'yield_farming_actions'
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

func DownAlterDeleteBlock2Function(tx *sql.Tx) error {
	return nil
}
