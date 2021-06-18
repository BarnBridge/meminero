package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateFunctionCountNftHolders, downCreateFunctionCountNftHolders)
}

func upCreateFunctionCountNftHolders(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create or replace function number_of_nft_holders(addr text) returns bigint
    		language plpgsql as
		$$
		declare
			value bigint;
		begin
			select into value count(*)
			from ( with transfers as ( select sender as address, -1 as amount
									   from erc721_transfers
									   where token_address = addr
									   union all
									   select receiver as address, 1 as amount
									   from erc721_transfers
									   where token_address = addr )
				   select address, sum(amount) as balance
				   from transfers
				   group by address ) x
			where balance > 0;
		
			return value;
		end;
		$$;

		create or replace function number_of_seniors(pool_address text) returns bigint
			language plpgsql as
		$$
		declare
			value bigint;
		begin
			select into value number_of_nft_holders(
									  ( select senior_bond_address from smart_yield_pools where sy_address = pool_address ));
		
			return value;
		end;
		$$;

		create or replace function number_of_juniors_locked(pool_address text) returns bigint
			language plpgsql as
		$$
		declare
			value bigint;
		begin
			select into value number_of_nft_holders(
									  ( select junior_bond_address from smart_yield_pools where sy_address = pool_address ));
		
			return value;
		end;
		$$;

		create or replace function number_of_jtoken_holders(addr text) returns bigint
			language plpgsql as
		$$
		declare
			holdersCount bigint;
		begin
			select into holdersCount count(*)
			from ( with transfers as ( select sender as address, -value as amount
									   from jtoken_transfers
									   where sy_address = addr
									   union all
									   select receiver as address, value as amount
									   from jtoken_transfers
									   where sy_address = addr )
				   select address, sum(amount) as balance
				   from transfers
				   group by address ) x
			where balance > 0;
		
			return holdersCount;
		end;
		$$;

		create or replace function junior_liquidity_locked(pool_address text) returns numeric(78)
			language plpgsql as
		$$
		declare
			value numeric(78);
		begin
			select into value sum(case when ( select count(*)
											  from smart_yield_junior_redeem as r
											  where r.junior_bond_address = b.junior_bond_address
												and r.junior_bond_id = b.junior_bond_id ) = 0 then tokens_in
									   else 0 end)
			from smart_yield_junior_buy as b
			where b.sy_address = pool_address;
		
			return value;
		end
		$$;
	`)
	return err
}

func downCreateFunctionCountNftHolders(tx *sql.Tx) error {
	_, err := tx.Exec(`
		drop function if exists number_of_nft_holders;
		drop function if exists number_of_seniors;
		drop function if exists number_of_juniors_locked;
		drop function if exists number_of_jtoken_holders;
		drop function if exists junior_liquidity_locked;
	`)
	return err
}
