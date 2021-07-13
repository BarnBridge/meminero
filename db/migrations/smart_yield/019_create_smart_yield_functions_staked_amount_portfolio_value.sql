create or replace function smart_yield.staked_amount_at_ts_by_reward_pool(pool text, address text, ts bigint) returns numeric(78)
    language plpgsql as
$$
declare
    value numeric(78);
begin
    select into value balance_after
    from smart_yield.rewards_staking_actions as a
    where a.pool_address = pool
      and a.user_address = address
      and a.block_timestamp <= ts
    order by block_timestamp desc
    limit 1;

    return value;
end;
$$;

create or replace function smart_yield.jtoken_price_scaled_at_ts(sy_address text, ts bigint) returns double precision
    language plpgsql as
$$
declare
    value double precision;
begin
    select into value jtoken_price / pow(10, 18)
    from smart_yield.state
    where pool_address = sy_address
      and block_timestamp <= to_timestamp(ts)
    order by block_timestamp desc
    limit 1;

    return value;
end;
$$;

create function junior_staked_balance_at_ts(user_address text, ts bigint) returns double precision
    language plpgsql
as
$$
declare
    value double precision;
begin
    select into value sum(
                                      smart_yield.staked_amount_at_ts_by_reward_pool(pool_address, user_address,
                                                                                     ts)::numeric(78, 18) /
                                      pow(10, (select underlying_decimals
                                               from smart_yield.pools as p
                                               where p.sy_address = rp.pool_token_address)) *
                                      smart_yield.jtoken_price_scaled_at_ts(pool_token_address, ts) *
                                      smart_yield.pool_underlying_price_at_ts(pool_token_address, ts))
    from smart_yield.reward_pools as rp;

    return value;
end;

$$;

---- create above / drop below ----

drop function if exists smart_yield.staked_amount_at_ts_by_reward_pool(pool text, address text, ts bigint);
drop function if exists smart_yield.jtoken_price_scaled_at_ts(sy_address text, ts bigint);
drop function if exists smart_yield.junior_staked_balance_at_ts(user_address text, ts bigint);
