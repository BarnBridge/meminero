create function smart_exposure.active_position_at_ts(user_address text, ts bigint)
    returns TABLE
            (
                etoken_address text,
                balance        numeric
            )
    language plpgsql
as
$$
begin
    return query with transfers as ( select token_address as address, -value as amount
                                     from public.erc20_transfers
                                     where sender = user_address
                                       and block_timestamp <= ts
                                     union all
                                     select etoken_address as address, value as amount
                                     from public.erc20_transfers
                                     where receiver = user_address
                                       and block_timestamp <= ts )
                 select address, sum(amount) as balance
                 from transfers
                 group by address;
end;
$$;

create or replace function smart_exposure.get_ratio_deviation(_etoken_address text, start bigint, _date_trunc text)
    returns TABLE
            (
                point     timestamp without time zone,
                deviation numeric
            )
    language plpgsql
as
$$
declare
    target_ratio numeric(78, 18);
begin
    select into target_ratio t.target_ratio::numeric(78, 18)
    from smart_exposure.tranches t
    where t.etoken_address = _etoken_address;

    return query select date_trunc(_date_trunc, to_timestamp(block_timestamp))::timestamp          as point,
                        avg(1 - ((ts.current_ratio::numeric(78, 18)) / target_ratio)) as deviation
                 from smart_exposure.tranche_state ts
                 where ts.etoken_address = _etoken_address
                   and ts.block_timestamp > start
                 group by point
                 order by point;
end
$$;

create function smart_exposure.get_token_price_chart(_token_address text, start bigint, _date_trunc text)
    returns TABLE
            (
                point       timestamp without time zone,
                token_price double precision
            )
    language plpgsql
as
$$
begin
    return query select date_trunc(_date_trunc, to_timestamp(block_timestamp)::date)::timestamp as point,
                        avg(price_usd)                                                          as token_price
                 from public.token_prices
                 where token_address = _token_address
                   and block_timestamp > start
                 group by point
                 order by point;
end
$$;

create function smart_exposure.get_tranche_details(_etoken_address text)
    returns TABLE
            (
                s_factor_e                          numeric,
                target_ratio                        numeric,
                token_a_ratio                       double precision,
                token_a_address                     text,
                token_a_symbol                      text,
                token_a_decimals                    bigint,
                token_a_price_usd                   double precision,
                token_a_included_in_block           bigint,
                token_a_block_timestamp             bigint,
                token_b_address                     text,
                token_b_price_usd                   double precision,
                token_b_included_in_block           bigint,
                token_b_block_timestamp             bigint,
                token_b_ratio                       double precision,
                token_b_symbol                      text,
                token_b_decimals                    bigint,
                pool_state_rebalancing_interval     bigint,
                pool_state_rebalancing_condition    numeric,
                pool_state_last_rebalance           bigint,
                tranche_state_token_a_liquidity     double precision,
                tranche_state_token_b_liquidity     double precision,
                tranche_state_e_token_price         double precision,
                tranche_state_current_ratio         double precision,
                tranche_state_token_a_current_ratio double precision,
                tranche_state_token_b_current_ratio double precision,
                tranche_state_included_in_block     bigint,
                tranche_state_block_timestamp       bigint
            )
    language plpgsql
as
$$
declare
    _pool_address                       text;
    s_factor_e                          numeric(78);
    target_ratio                        numeric(78);
    token_a_ratio                       double precision;
    token_b_ratio                       double precision;
    token_a_address                     text;
    token_a_symbol                      text;
    token_a_decimals                    bigint;
    token_a_price_usd                   double precision;
    token_a_included_in_block           bigint;
    token_a_block_timestamp             bigint;
    token_b_address                     text;
    token_b_symbol                      text;
    token_b_decimals                    bigint;
    token_b_price_usd                   double precision;
    token_b_included_in_block           bigint;
    token_b_block_timestamp             bigint;
    pool_state_rebalancing_interval     bigint;
    pool_state_rebalancing_condition    numeric(78);
    pool_state_last_rebalance           bigint;
    tranche_state_token_a_liquidity     double precision;
    tranche_state_token_b_liquidity     double precision;
    tranche_state_e_token_price         double precision;
    tranche_state_current_ratio         double precision;
    tranche_state_token_a_current_ratio double precision;
    tranche_state_token_b_current_ratio double precision;
    tranche_state_included_in_block     bigint;
    tranche_state_block_timestamp       bigint;
begin
    select into _pool_address,s_factor_e,target_ratio,token_a_ratio,token_b_ratio t.pool_address,
                                                                                  t.s_factor_e,
                                                                                  t.target_ratio,
                                                                                  t.token_a_ratio,
                                                                                  t.token_b_ratio
    from smart_exposure.tranches t
    where t.etoken_address = _etoken_address;

    select into token_a_address,token_a_symbol,token_a_decimals,token_b_address,token_b_symbol,token_b_decimals p.token_a_address,
                                                                                                                p.token_a_symbol,
                                                                                                                p.token_a_decimals,
                                                                                                                p.token_b_address,
                                                                                                                p.token_b_symbol,
                                                                                                                p.token_b_decimals
    from smart_exposure.pools p
    where p.pool_address = _pool_address;

    select into token_a_price_usd,token_a_included_in_block,token_a_block_timestamp price.price_usd,
                                                                                    price.included_in_block,
                                                                                    price.block_timestamp
    from public.token_prices price
    where price.token_address = token_a_address
    order by block_timestamp desc
    limit 1;

    select into token_b_price_usd,token_b_included_in_block,token_b_block_timestamp price.price_usd,
                                                                                    price.included_in_block,
                                                                                    price.block_timestamp
    from public.token_prices price
    where price.token_address = token_b_address
    order by block_timestamp desc
    limit 1;


    select into pool_state_rebalancing_interval,pool_state_rebalancing_condition,pool_state_last_rebalance ps.rebalancing_interval,
                                                                                                           ps.rebalancing_condition,
                                                                                                           ps.last_rebalance
    from smart_exposure.pool_state ps
    where ps.pool_address = _pool_address
    order by block_timestamp desc
    limit 1;

    select into tranche_state_token_a_liquidity, tranche_state_token_b_liquidity,tranche_state_e_token_price,tranche_state_current_ratio, tranche_state_token_a_current_ratio,tranche_state_token_b_current_ratio,tranche_state_included_in_block,tranche_state_block_timestamp ts.token_a_liquidity,
                                                                                                                                                                                                                                                                                ts.token_b_liquidity,
                                                                                                                                                                                                                                                                                ts.etoken_price,
                                                                                                                                                                                                                                                                                ts.current_ratio,
                                                                                                                                                                                                                                                                                ts.token_a_current_ratio,
                                                                                                                                                                                                                                                                                ts.token_b_current_ratio,
                                                                                                                                                                                                                                                                                ts.included_in_block,
                                                                                                                                                                                                                                                                                ts.block_timestamp
    from smart_exposure.tranche_state ts
    where ts.etoken_address = _etoken_address
    order by block_timestamp desc
    limit 1;
    return query select s_factor_e,
                        target_ratio,
                        token_a_ratio,
                        token_a_address,
                        token_a_symbol,
                        token_a_decimals,
                        token_a_price_usd,
                        token_a_included_in_block,
                        token_a_block_timestamp,
                        token_b_address,
                        token_b_price_usd,
                        token_b_included_in_block,
                        token_b_block_timestamp,
                        token_b_ratio,
                        token_b_symbol,
                        token_b_decimals,
                        pool_state_rebalancing_interval,
                        pool_state_rebalancing_condition,
                        pool_state_last_rebalance,
                        tranche_state_token_a_liquidity,
                        tranche_state_token_b_liquidity,
                        tranche_state_e_token_price,
                        tranche_state_current_ratio,
                        tranche_state_token_a_current_ratio,
                        tranche_state_token_b_current_ratio,
                        tranche_state_included_in_block,
                        tranche_state_block_timestamp;
end
$$;

create function smart_exposure.user_portfolio_value(addr text, ts bigint) returns double precision
    language plpgsql as
$$
declare
    value double precision;
begin
    select into value sum(coalesce(a.balance / pow(10, 18) * ( select etoken_price
                                                               from smart_exposure.tranche_state s
                                                               where s.etoken_address = a.etoken_address
                                                                 and s.block_timestamp <= to_timestamp(ts)
                                                               order by s.block_timestamp desc
                                                               limit 1 ), 0))
    from smart_exposure.active_position_at_ts(addr, ts) a;

    return value;
end;
$$;

create function smart_exposure.user_portfolio_value_by_pool(addr text, ts bigint, _pool_address text) returns double precision
    language plpgsql as
$$
declare
    value double precision;
begin
    select into value sum(coalesce(a.balance / pow(10, 18) * ( select etoken_price
                                                               from smart_exposure.tranche_state s
                                                               where s.etoken_address = a.etoken_address
                                                                 and s.block_timestamp <= to_timestamp(ts)
                                                               order by s.block_timestamp desc
                                                               limit 1 ), 0))
    from smart_exposure.active_position_at_ts(addr, ts) a
    where a.etoken_address in ( select etoken_address from smart_exposure.tranches where pool_address = _pool_address );
    return value;
end;
$$;
