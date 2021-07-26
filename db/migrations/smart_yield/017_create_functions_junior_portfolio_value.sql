create or replace function smart_yield.junior_active_positions_at_ts(user_address text, ts bigint)
    returns table
            (
                pool    text,
                balance numeric(78)
            )
    language plpgsql
as
$$
begin
    return query with transfers as ( select token_address as address, -value as amount
                                     from erc20_transfers
                                     where sender = user_address
                                       and block_timestamp <= ts

                                     union all
                                     select token_address as address, value as amount
                                     from erc20_transfers
                                     where receiver = user_address
                                       and block_timestamp <= ts )
                 select address, sum(amount) as balance
                 from transfers
                 where ( select count(*) from smart_yield.pools where pool_address = address ) > 0
                 group by address;

end;
$$;

create or replace function smart_yield.pool_underlying_price_at_ts(addr text, ts bigint) returns double precision
    language plpgsql as
$$
declare
    price double precision;
begin
    select into price price_usd
    from public.token_prices p
    where p.token_address = ( select underlying_address from smart_yield.pools where pool_address = addr )
      and block_timestamp <= ts
    order by block_timestamp desc
    limit 1;

    return price;
end;
$$;

create function smart_yield.junior_active_balance_at_ts(user_address text, ts bigint) returns double precision
    language plpgsql as
$$
declare
    total_balance double precision;
begin
    select into total_balance sum(balance::numeric(78, 18) / pow(10, ( select underlying_decimals
                                                                       from smart_yield.pools
                                                                       where pool_address = pool
                                                                       limit 1 )) * ( select jtoken_price / pow(10, 18)
                                                                                      from smart_yield.state
                                                                                      where pool_address = pool
                                                                                        and block_timestamp <= to_timestamp(ts)
                                                                                      order by block_timestamp desc
                                                                                      limit 1 ) *
                                  ( select smart_yield.pool_underlying_price_at_ts(pool, ts) ))
    from smart_yield.junior_active_positions_at_ts(user_address, ts);

    return total_balance;
end;
$$;

create or replace function smart_yield.junior_bond_value_at_ts(token_address text, token_id bigint, ts bigint) returns numeric(78)
    language plpgsql as
$$
declare
    value numeric(78);
begin
    select into value tokens_in
    from smart_yield.junior_2step_withdraw_events
    where junior_bond_address = token_address
      and junior_bond_id = token_id;

    return value;
end;
$$;

create or replace function smart_yield.junior_bond_redeemed_at_ts(token_address text, token_id bigint, ts bigint) returns boolean
    language plpgsql as
$$
declare
    redeemed boolean;
begin
    select into redeemed count(*) > 0
    from smart_yield.junior_2step_redeem_events
    where junior_bond_address = token_address
      and junior_bond_id = token_id
      and block_timestamp <= ts;

    return redeemed;
end;
$$;

create or replace function smart_yield.junior_underlying_price_at_ts(addr text, ts bigint) returns double precision
    language plpgsql as
$$
declare
    price double precision;
begin
    select into price price_usd
    from token_prices p
    where p.token_address = ( select underlying_address from smart_yield.pools where junior_bond_address = addr )
      and block_timestamp <= ts
    order by block_timestamp desc
    limit 1;

    return price;
end;
$$;

create or replace function smart_yield.junior_locked_positions_at_ts(addr text, ts bigint)
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
                 from smart_yield.erc721_transfers t
                 where token_type = 'junior'
                   and receiver = addr
                   and block_timestamp <= ts
                   and smart_yield.current_owner_of_token_at_ts(t.token_address, t.token_id, ts) = receiver
                   and not smart_yield.junior_bond_redeemed_at_ts(t.token_address, t.token_id, ts);
end;
$$;

create or replace function smart_yield.junior_locked_balance_at_ts(addr text, ts bigint) returns double precision
    language plpgsql as
$$
declare
    value double precision;
begin
    select into value coalesce(sum(smart_yield.junior_bond_value_at_ts(token_address, token_id, ts)::numeric(78, 18) /
                                   pow(10, ( select underlying_decimals
                                             from smart_yield.pools
                                             where junior_bond_address = token_address )) *
                                   ( select jtoken_price / pow(10, 18)
                                     from smart_yield.state
                                     where pool_address = ( select pool_address
                                                            from smart_yield.pools
                                                            where junior_bond_address = token_address )
                                       and block_timestamp <= to_timestamp(ts)
                                     order by block_timestamp desc
                                     limit 1 ) * smart_yield.junior_underlying_price_at_ts(token_address, ts)), 0)
    from smart_yield.junior_locked_positions_at_ts(addr, ts);

    return value;
end;
$$;

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
    language plpgsql as
$$
declare
    value double precision;
begin
    select into value sum(smart_yield.staked_amount_at_ts_by_reward_pool(pool_address, user_address,
                                                                         ts)::numeric(78, 18) / pow(10,
                                                                                                    ( select underlying_decimals
                                                                                                      from smart_yield.pools as p
                                                                                                      where p.sy_address = rp.pool_token_address )) *
                          smart_yield.jtoken_price_scaled_at_ts(pool_token_address, ts) *
                          smart_yield.pool_underlying_price_at_ts(pool_token_address, ts))
    from smart_yield.reward_pools as rp;

    return value;
end;

$$;

create function junior_portfolio_value_at_ts(addr text, ts bigint) returns double precision
    language plpgsql as
$$
declare
    value double precision;
begin
    select into value coalesce(smart_yield.junior_locked_balance_at_ts(addr, ts), 0) +
                      coalesce(smart_yield.junior_active_balance_at_ts(addr, ts), 0) +
                      coalesce(smart_yield.junior_staked_balance_at_ts(addr, ts), 0);

    return value;
end;
$$;
