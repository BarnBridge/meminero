create function smart_alpha.lp_token_epoch_start_price_at_ts(pool text, ts bigint)
    returns table
            (
                junior_price_start double precision,
                senior_price_start double precision
            )
    language plpgsql
as
$$
begin

    return query (select junior_token_price_start::numeric(78, 18) / pow(10, 18) as junior_price_start,
                         senior_token_price_start::numeric(78, 18) / pow(10, 18) as senior_price_start
                  from smart_alpha.pool_epoch_info
                  where pool_address = pool
                    and block_timestamp <= ts
                  order by block_timestamp desc
                  limit 1);
end;
$$;


create function smart_alpha.format_lp_token(amount numeric, pool text) returns double precision
    language plpgsql
as
$$
declare
    pool_token_decimals integer;
begin
    select into pool_token_decimals p.pool_token_decimals from smart_alpha.pools p where p.pool_address = pool;
    return amount::numeric(78, 18) / pow(10, pool_token_decimals);
end;
$$;

create function smart_alpha.junior_position_performance_at_ts(addr text, pool text, ts bigint)
    returns table
            (
                junior_performance_with_sa    double precision,
                junior_performance_without_sa double precision
            )
    language plpgsql
as
$$
declare
    junior_balance                numeric;
    junior_token_address          text;
    pool_token_address            text;
    junior_performance_with_sa    double precision;
    junior_performance_without_sa double precision;
begin

    select into junior_token_address,pool_token_address p.junior_token_address, p.pool_token_address
    from smart_alpha.pools p
    where p.pool_address = pool;

    select into junior_balance balance
    from public.erc20_balances_at_ts(addr, (select array_agg(junior_token_address)), ts);


    select into junior_performance_with_sa smart_alpha.junior_token_to_usd_at_ts(junior_token_address, junior_balance,
                                                                                 ts);
    select into junior_performance_without_sa (select smart_alpha.format_lp_token(junior_balance, pool) *
                                                      (select lp.junior_price_start
                                                       from smart_alpha.lp_token_epoch_start_price_at_ts(pool, ts) lp
                                                       limit 1) *
                                                      (select public.token_usd_price_at_ts(pool_token_address, ts)));

    return query select  coalesce(junior_performance_with_sa, 0), coalesce(junior_performance_without_sa, 0);
end;
$$;

create function smart_alpha.senior_position_performance_at_ts(addr text, pool text, ts bigint)
    returns table
            (
                senior_performance_with_sa    double precision,
                senior_performance_without_sa double precision
            )
    language plpgsql
as
$$
declare
    senior_balance                numeric;
    senior_token_address          text;
    pool_token_address            text;
    senior_performance_with_sa    double precision;
    senior_performance_without_sa double precision;
begin

    select into senior_token_address,pool_token_address p.senior_token_address, p.pool_token_address
    from smart_alpha.pools p
    where p.pool_address = pool;

    select into senior_balance balance
    from public.erc20_balances_at_ts(addr, (select array_agg(senior_token_address)), ts);


    select into senior_performance_with_sa smart_alpha.senior_token_to_usd_at_ts(senior_token_address, senior_balance,
                                                                                 ts);
    select into senior_performance_without_sa (select smart_alpha.format_lp_token(senior_balance, pool) *
                                                      (select lp.senior_price_start
                                                       from smart_alpha.lp_token_epoch_start_price_at_ts(pool, ts) lp) *
                                                      (select public.token_usd_price_at_ts(pool_token_address, ts)));
    return query select coalesce(senior_performance_with_sa, 0), coalesce(senior_performance_without_sa, 0);
end;
$$;