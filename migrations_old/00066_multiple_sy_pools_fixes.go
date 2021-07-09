package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upMultipleSyPoolsFixes, downMultipleSyPoolsFixes)
}

func upMultipleSyPoolsFixes(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create or replace function junior_active_balance_at_ts(user_address text, ts bigint) returns double precision
			language plpgsql as
		$$
		declare
			total_balance double precision;
		begin
			select into total_balance sum(balance::numeric(78, 18) / pow(10, ( select underlying_decimals
																		   from smart_yield_pools
																		   where sy_address = pool
																		   limit 1 )) * ( select jtoken_price / pow(10, 18)
																						  from smart_yield_state
																						  where pool_address = pool
																							and block_timestamp <= to_timestamp(ts)
																						  order by block_timestamp desc
																						  limit 1 ) *
									  ( select pool_underlying_price_at_ts(pool, ts) ))
			from junior_active_positions_at_ts(user_address, ts);
		
			return total_balance;
		end;
		$$;
	`)

	return err
}

func downMultipleSyPoolsFixes(tx *sql.Tx) error {
	return nil
}
