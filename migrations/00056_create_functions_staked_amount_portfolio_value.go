package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateFunctionsStakedAmountPortfolioValue, downCreateFunctionsStakedAmountPortfolioValue)
}

func upCreateFunctionsStakedAmountPortfolioValue(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create or replace function staked_amount_at_ts_by_reward_pool(pool text, address text, ts bigint) returns numeric(78)
			language plpgsql as
		$$
		declare
			value numeric(78);
		begin
			select into value balance_after
			from smart_yield_rewards_staking_actions as a
			where a.pool_address = pool
			  and a.user_address = address
			  and a.block_timestamp <= ts
			order by block_timestamp desc
			limit 1;
		
			return value;
		end;
		$$;

		create or replace function jtoken_price_scaled_at_ts(sy_address text, ts bigint) returns double precision
			language plpgsql as
		$$
		declare
			value double precision;
		begin
			select into value jtoken_price / pow(10, 18)
			from smart_yield_state
			where pool_address = sy_address
			  and block_timestamp <= to_timestamp(ts)
			order by block_timestamp desc
			limit 1;
		
			return value;
		end;
		$$;

		create or replace function junior_staked_balance_at_ts(user_address text, ts bigint) returns double precision
			language plpgsql as
		$$
		declare
			value double precision;
		begin
			select into value sum(
				staked_amount_at_ts_by_reward_pool(pool_address, user_address, ts)::numeric(78, 18) /
				pow(10, ( select underlying_decimals from smart_yield_pools as p where p.sy_address = rp.pool_token_address )) *
				jtoken_price_scaled_at_ts(pool_token_address, ts) *
				pool_underlying_price_at_ts(pool_token_address, ts))
			from smart_yield_reward_pools as rp;
		
			return value;
		end;
		$$;

create or replace function junior_portfolio_value_at_ts(addr text, ts bigint) returns double precision
    language plpgsql as
$$
declare
    value double precision;
begin
    select into value coalesce(junior_locked_balance_at_ts(addr, ts), 0) + 
                      coalesce(junior_active_balance_at_ts(addr, ts), 0) + 
                      coalesce(junior_staked_balance_at_ts(addr, ts), 0);

    return value;
end;
$$;

`)
	return err
}

func downCreateFunctionsStakedAmountPortfolioValue(tx *sql.Tx) error {
	_, err := tx.Exec(`
		drop function if exists staked_amount_at_ts_by_reward_pool;
		drop function if exists jtoken_price_scaled_at_ts;
		drop function if exists junior_staked_value_at_ts;

		create or replace function junior_portfolio_value_at_ts(addr text, ts bigint) returns double precision
			language plpgsql as
		$$
		declare
			value double precision;
		begin
			select into value coalesce(junior_locked_balance_at_ts(addr, ts), 0) + 
							  coalesce(junior_active_balance_at_ts(addr, ts), 0);
			
			return value;
		end;
		$$;
	`)
	return err
}
