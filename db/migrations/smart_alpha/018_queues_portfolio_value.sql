-- there are 2 cases to consider:
-- 1. the user has a queue position for the current epoch -> we multiply the underlying in with the underlying token price to usd
-- 2. the user has a queue position in a previous epoch that was not redeemed yet -> we must convert the underlying to tokens at the epoch price and back into underlying at the estimated price
create or replace function smart_alpha.pool_active_epoch_at_ts(addr text, ts bigint) returns bigint
    language plpgsql as
$$
begin
    return ( select epoch_id
             from smart_alpha.pool_epoch_info
             where pool_address = addr
               and block_timestamp < ts
             order by block_timestamp desc
             limit 1 );
end;
$$;

create or replace function smart_alpha.underlying_to_junior_tokens_at_epoch(pool text, amount numeric(78, 18), _epoch bigint) returns double precision
    language plpgsql as
$$
begin
    return ( select amount / pow(10, p.pool_token_decimals) /
                    (pei.junior_token_price_start::numeric(78, 18) / pow(10, 18))
             from smart_alpha.pool_epoch_info pei
                      inner join smart_alpha.pools p on p.pool_address = pool
             where pei.pool_address = pool
               and pei.epoch_id > _epoch
             order by pei.epoch_id
             limit 1 );
end;
$$;

create or replace function smart_alpha.underlying_to_senior_tokens_at_epoch(pool text, amount numeric(78, 18), _epoch bigint) returns double precision
    language plpgsql as
$$
begin
    return ( select amount / pow(10, p.pool_token_decimals) /
                    (pei.senior_token_price_start::numeric(78, 18) / pow(10, 18))
             from smart_alpha.pool_epoch_info pei
                      inner join smart_alpha.pools p on p.pool_address = pool
             where pei.pool_address = pool
               and pei.epoch_id > _epoch
             order by pei.epoch_id
             limit 1 );
end;
$$;

create or replace function smart_alpha.estimated_junior_token_price_at_ts(pool text, ts bigint) returns double precision
    language plpgsql as
$$
begin
    return (( select estimated_junior_token_price
              from smart_alpha.pool_state ps
              where ps.pool_address = pool
                and ps.block_timestamp < ts
              order by ps.block_timestamp desc
              limit 1 ) / pow(10, 18));
end;
$$;

create or replace function smart_alpha.estimated_senior_token_price_at_ts(pool text, ts bigint) returns double precision
    language plpgsql as
$$
begin
    return (( select estimated_senior_token_price
              from smart_alpha.pool_state ps
              where ps.pool_address = pool
                and ps.block_timestamp < ts
              order by ps.block_timestamp desc
              limit 1 ) / pow(10, 18));
end;
$$;

create or replace function smart_alpha.junior_tokens_to_underlying_at_epoch(pool text, amount numeric(78, 18), _epoch bigint) returns double precision
    language plpgsql as
$$
begin
    return ( select amount / pow(10, p.pool_token_decimals) *
                    (pei.junior_token_price_start::numeric(78, 18) / pow(10, 18))
             from smart_alpha.pool_epoch_info pei
                      inner join smart_alpha.pools p on p.pool_address = pool
             where pei.pool_address = pool
               and pei.epoch_id > _epoch
             order by pei.epoch_id
             limit 1 );
end;
$$;

create or replace function smart_alpha.senior_tokens_to_underlying_at_epoch(pool text, amount numeric(78, 18), _epoch bigint) returns double precision
    language plpgsql as
$$
begin
    return ( select amount / pow(10, p.pool_token_decimals) *
                    (pei.senior_token_price_start::numeric(78, 18) / pow(10, 18))
             from smart_alpha.pool_epoch_info pei
                      inner join smart_alpha.pools p on p.pool_address = pool
             where pei.pool_address = pool
               and pei.epoch_id > _epoch
             order by pei.epoch_id
             limit 1 );
end;
$$;

create or replace function smart_alpha.entry_queue_portfolio_value_at_ts(addr text, ts bigint) returns double precision
    language plpgsql as
$$
begin
    return ( select sum(case when j.epoch_id = ( select smart_alpha.pool_active_epoch_at_ts(j.pool_address, ts) ) then
                                         underlying_in::numeric(78, 18) / pow(10, p.pool_token_decimals) *
                                         token_usd_price_at_ts(p.pool_token_address, ts)
                             else case when j.tranche = 'JUNIOR'
                                           then ( select smart_alpha.underlying_to_junior_tokens_at_epoch(
                                                                 j.pool_address, j.underlying_in::numeric(78, 18),
                                                                 j.epoch_id) *
                                                         ( select smart_alpha.estimated_junior_token_price_at_ts(j.pool_address, ts) ) *
                                                         ( select token_usd_price_at_ts(p.pool_token_address, ts) ) )
                                       when j.tranche = 'SENIOR'
                                           then ( select smart_alpha.underlying_to_senior_tokens_at_epoch(
                                                                 j.pool_address, j.underlying_in::numeric(78, 18),
                                                                 j.epoch_id) *
                                                         ( select smart_alpha.estimated_senior_token_price_at_ts(j.pool_address, ts) ) *
                                                         ( select token_usd_price_at_ts(p.pool_token_address, ts) ) ) end end)
             from smart_alpha.user_join_entry_queue_events j
                      left join smart_alpha.user_redeem_tokens_events r
                                on j.user_address = r.user_address and j.pool_address = r.pool_address and
                                   j.epoch_id = r.epoch_id and j.tranche = r.tranche and r.block_timestamp < ts
                      inner join smart_alpha.pools p on j.pool_address = p.pool_address
             where j.user_address = addr
               and j.block_timestamp < ts
               and r.user_address is null );
end;
$$;

create or replace function smart_alpha.exit_queue_portfolio_value_at_ts(addr text, ts bigint) returns double precision
    language plpgsql as
$$
begin
    return ( select sum(case when j.epoch_id = ( select smart_alpha.pool_active_epoch_at_ts(j.pool_address, ts) )
                                 then case when j.tranche = 'JUNIOR' then ( select j.tokens_in::numeric(78, 18) /
                                                                                   pow(10, p.pool_token_decimals) *
                                                                                   ( select smart_alpha.estimated_junior_token_price_at_ts(j.pool_address, ts) ) *
                                                                                   ( select token_usd_price_at_ts(p.pool_token_address, ts) ) )
                                           when j.tranche = 'SENIOR' then ( select j.tokens_in::numeric(78, 18) /
                                                                                   pow(10, p.pool_token_decimals) *
                                                                                   ( select smart_alpha.estimated_senior_token_price_at_ts(j.pool_address, ts) ) *
                                                                                   ( select token_usd_price_at_ts(p.pool_token_address, ts) ) ) end
                             else case when j.tranche = 'JUNIOR' then
                                               ( select smart_alpha.junior_tokens_to_underlying_at_epoch(j.pool_address,
                                                                                                         j.tokens_in::numeric(78, 18),
                                                                                                         j.epoch_id) ) *
                                               ( select token_usd_price_at_ts(p.pool_token_address, ts) )
                                       when j.tranche = 'SENIOR' then
                                               ( select smart_alpha.senior_tokens_to_underlying_at_epoch(j.pool_address,
                                                                                                         j.tokens_in::numeric(78, 18),
                                                                                                         j.epoch_id) ) *
                                               ( select token_usd_price_at_ts(p.pool_token_address, ts) ) end end)
             from smart_alpha.user_join_exit_queue_events j
                      left join smart_alpha.user_redeem_underlying_events r
                                on j.user_address = r.user_address and j.pool_address = r.pool_address and
                                   j.epoch_id = r.epoch_id and j.tranche = r.tranche and r.block_timestamp < ts
                      inner join smart_alpha.pools p on j.pool_address = p.pool_address
             where j.user_address = addr
               and j.block_timestamp < ts
               and r.user_address is null );
end;
$$;
