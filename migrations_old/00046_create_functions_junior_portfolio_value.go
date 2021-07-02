package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateFunctionsJuniorPortfolioValue, downCreateFunctionsJuniorPortfolioValue)
}

func upCreateFunctionsJuniorPortfolioValue(tx *sql.Tx) error {
	_, err := tx.Exec(`
create or replace function junior_active_positions_at_ts(user_address text, ts bigint)
    returns table
            (
                pool    text,
                balance numeric(78)
            )
    language plpgsql
as
$$
begin
    return query with transfers as ( select sy_address as address, -value as amount
                                     from jtoken_transfers
                                     where sender = user_address
                                       and block_timestamp <= ts
                                     union all
                                     select sy_address as address, value as amount
                                     from jtoken_transfers
                                     where receiver = user_address
                                       and block_timestamp <= ts )
                 select address, sum(amount) as balance
                 from transfers
                 group by address;
end;
$$;

create or replace function pool_underlying_price_at_ts(addr text, ts bigint) returns double precision
    language plpgsql as
$$
declare
    price double precision;
begin
    select into price price_usd
    from smart_yield_prices p
    where p.protocol_id = ( select protocol_id from smart_yield_pools where sy_address = addr )
      and p.token_address = ( select underlying_address from smart_yield_pools where sy_address = addr )
      and block_timestamp <= ts
    order by block_timestamp desc
    limit 1;

    return price;
end;
$$;

create or replace function junior_active_balance_at_ts(user_address text, ts bigint) returns double precision
    language plpgsql as
$$
declare
    total_balance double precision;
begin
    select into total_balance balance::numeric(78, 18) / pow(10, ( select underlying_decimals
                                                                   from smart_yield_pools
                                                                   where sy_address = pool
                                                                   limit 1 )) * ( select jtoken_price / pow(10, 18)
                                                                                  from smart_yield_state
                                                                                  where pool_address = pool
                                                                                    and block_timestamp <= to_timestamp(ts)
                                                                                  order by block_timestamp desc
                                                                                  limit 1 ) *
                              ( select pool_underlying_price_at_ts(pool, ts) )
    from junior_active_positions_at_ts(user_address, ts);

    return total_balance;
end;
$$;

create or replace function junior_bond_value_at_ts(token_address text, token_id bigint, ts bigint) returns numeric(78)
    language plpgsql as
$$
declare
    value numeric(78);
begin
    select into value tokens_in
    from smart_yield_junior_buy
    where junior_bond_address = token_address
      and junior_bond_id = token_id;

    return value;
end;
$$;

create or replace function junior_bond_redeemed_at_ts(token_address text, token_id bigint, ts bigint) returns boolean
    language plpgsql as
$$
declare
    redeemed boolean;
begin
    select into redeemed count(*) > 0
    from smart_yield_junior_redeem
    where junior_bond_address = token_address
      and junior_bond_id = token_id
      and block_timestamp <= ts;

    return redeemed;
end;
$$;

create or replace function junior_underlying_price_at_ts(addr text, ts bigint) returns double precision
    language plpgsql as
$$
declare
    price double precision;
begin
    select into price price_usd
    from smart_yield_prices p
    where p.protocol_id = ( select protocol_id from smart_yield_pools where junior_bond_address = addr )
      and p.token_address = ( select underlying_address from smart_yield_pools where junior_bond_address = addr )
      and block_timestamp <= ts
    order by block_timestamp desc
    limit 1;

    return price;
end;
$$;

create or replace function junior_locked_positions_at_ts(addr text, ts bigint)
	returns table
			(
				token_address text,
				token_id      bigint
			)
	language plpgsql
as
$$
begin
	return query select distinct t.token_address, t.token_id
	from erc721_transfers t
	where token_type = 'junior'
      and receiver = addr
      and block_timestamp <= ts
      and current_owner_of_token_at_ts(t.token_address, t.token_id, ts) = receiver
      and not junior_bond_redeemed_at_ts(t.token_address, t.token_id, ts);
end;
$$;

create or replace function junior_locked_balance_at_ts(addr text, ts bigint) returns double precision
    language plpgsql as
$$
declare
    value double precision;
begin
    select into value coalesce(
        sum(
            junior_bond_value_at_ts(token_address, token_id, ts)::numeric(78, 18) /
            pow(10,( select underlying_decimals from smart_yield_pools where junior_bond_address = token_address )) *
            ( select jtoken_price / pow(10, 18)
              from smart_yield_state
              where pool_address = ( select sy_address from smart_yield_pools where junior_bond_address = token_address )
                and block_timestamp <= to_timestamp(ts)
             order by block_timestamp desc
             limit 1 ) *
            junior_underlying_price_at_ts(token_address, ts)),
        0
    )
    from junior_locked_positions_at_ts(addr, ts);

    return value;
end;
$$;

create or replace function junior_portfolio_value_at_ts(addr text, ts bigint) returns double precision
    language plpgsql as
$$
declare
    value double precision;
begin
    select into value coalesce(junior_locked_balance_at_ts(addr, ts), 0) + coalesce(junior_active_balance_at_ts(addr, ts), 0);

    return value;
end;
$$;
	`)
	return err
}

func downCreateFunctionsJuniorPortfolioValue(tx *sql.Tx) error {
	_, err := tx.Exec(`
		drop function if exists junior_active_positions_at_ts;
		drop function if exists pool_underlying_price_at_ts;
		drop function if exists junior_active_balance_at_ts;
		drop function if exists junior_bond_value_at_ts;
		drop function if exists junior_bond_redeemed_at_ts;
		drop function if exists junior_underlying_price_at_ts;
		drop function if exists junior_locked_positions_at_ts;
		drop function if exists junior_locked_balance_at_ts;
		drop function if exists junior_portfolio_value_at_ts;
	`)
	return err
}
