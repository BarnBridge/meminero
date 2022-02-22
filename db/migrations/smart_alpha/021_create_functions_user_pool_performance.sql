create or replace function smart_alpha.lp_token_epoch_start_price_at_ts(pool text, ts bigint)
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


create or replace function smart_alpha.format_lp_token(amount numeric, pool text) returns double precision
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

create or replace function smart_alpha.junior_position_performance_at_ts(addr text, pool text, ts bigint)
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
    epoch_id                      integer;
begin

    select into junior_token_address,pool_token_address,epoch_id p.junior_token_address,
                                                                 p.pool_token_address,
                                                                 ep.epoch_id
    from smart_alpha.pools p
             inner join smart_alpha.pool_epoch_info ep on p.pool_address = ep.pool_address
    where p.pool_address = pool
      and ep.block_timestamp <= ts
    order by ep.epoch_id desc
    limit 1;

    select into junior_balance balance + (select
                                          from smart_alpha.underlying_to_junior_tokens_at_epoch(pool,
                                                                                                (smart_alpha.junior_tokens_not_redeemed(addr, pool, ts)),
                                                                                                epoch_id))
    from public.erc20_balances_at_ts(addr, (select array_agg(junior_token_address)), ts);


    select into junior_performance_with_sa smart_alpha.junior_token_to_usd_at_ts(junior_token_address, junior_balance,
                                                                                 ts);
    select into junior_performance_without_sa (select smart_alpha.format_lp_token(junior_balance, pool) *
                                                      (select lp.junior_price_start
                                                       from smart_alpha.lp_token_epoch_start_price_at_ts(pool, ts) lp
                                                       limit 1) *
                                                      (select public.token_usd_price_at_ts(pool_token_address, ts)));

    return query select coalesce(junior_performance_with_sa, 0), coalesce(junior_performance_without_sa, 0);
end;
$$;

create or replace function smart_alpha.senior_position_performance_at_ts(addr text, pool text, ts bigint)
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
    epoch_id                      integer;
begin

    select into senior_token_address,pool_token_address,epoch_id p.senior_token_address,
                                                                 p.pool_token_address,
                                                                 ep.epoch_id
    from smart_alpha.pools p
             inner join smart_alpha.pool_epoch_info ep on p.pool_address = ep.pool_address
    where p.pool_address = pool
      and ep.block_timestamp <= ts
    order by ep.epoch_id desc
    limit 1;

    select into senior_balance balance + (select
                                          from smart_alpha.underlying_to_senior_tokens_at_epoch(pool,
                                                                                                (select from smart_alpha.senior_tokens_not_redeemed(addr, pool, ts)),
                                                                                                epoch_id)
    )
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

create or replace function smart_alpha.junior_tokens_not_redeemed(addr text, pool text, ts bigint) returns numeric
    language plpgsql as
$$
begin
    return (select sum(j.underlying_in)
            from smart_alpha.user_join_entry_queue_events j
                     left join smart_alpha.user_redeem_tokens_events r
                               on j.user_address = r.user_address and j.pool_address = r.pool_address and
                                  j.epoch_id = r.epoch_id and j.tranche = r.tranche and r.block_timestamp <= ts
            where j.user_address = addr
              and j.pool_address = pool
              and j.tranche = 'JUNIOR'
              and j.block_timestamp <= ts
              and r.user_address is null);
end;
$$;

create or replace function smart_alpha.senior_tokens_not_redeemed(addr text, pool text, ts bigint) returns numeric
    language plpgsql as
$$
begin
    return (select sum(j.underlying_in)
            from smart_alpha.user_join_entry_queue_events j
                     left join smart_alpha.user_redeem_tokens_events r
                               on j.user_address = r.user_address and j.pool_address = r.pool_address and
                                  j.epoch_id = r.epoch_id and j.tranche = r.tranche and r.block_timestamp <= ts
            where j.user_address = addr
              and j.pool_address = pool
              and j.tranche = 'SENIOR'
              and j.block_timestamp <= ts
              and r.user_address is null);
end;
$$;