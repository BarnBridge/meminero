create or replace function smart_alpha.entry_queue_portfolio_value_at_ts(addr text, ts bigint) returns double precision
    language plpgsql as
$$
begin
    return ( select sum(underlying_in::numeric(78, 18) / pow(10, p.pool_token_decimals) *
                        token_usd_price_at_ts(p.pool_token_address, ts))
             from smart_alpha.user_join_entry_queue_events j
                      left join smart_alpha.user_redeem_tokens_events r
                                on j.user_address = r.user_address and j.pool_address = r.pool_address and
                                   j.epoch_id = r.epoch_id and r.block_timestamp < ts
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
    return ( select sum(case when j.tranche = 'JUNIOR' then ( select j.tokens_in::numeric(78, 18) /
                                                                     pow(10, p.pool_token_decimals) *
                                                                     ( select estimated_junior_token_price::numeric(78, 18) / pow(10, 18)
                                                                       from smart_alpha.pool_state ps
                                                                       where ps.pool_address = p.pool_address
                                                                         and ps.block_timestamp < ts
                                                                       order by ps.block_timestamp desc
                                                                       limit 1 ) *
                                                                     ( select token_usd_price_at_ts(p.pool_token_address, ts) ) )
                             when j.tranche = 'SENIOR' then ( select j.tokens_in::numeric(78, 18) /
                                                                     pow(10, p.pool_token_decimals) *
                                                                     ( select estimated_senior_token_price::numeric(78, 18) / pow(10, 18)
                                                                       from smart_alpha.pool_state ps
                                                                       where ps.pool_address = p.pool_address
                                                                         and ps.block_timestamp < ts
                                                                       order by ps.block_timestamp desc
                                                                       limit 1 ) *
                                                                     ( select token_usd_price_at_ts(p.pool_token_address, ts) ) ) end)
             from smart_alpha.user_join_exit_queue_events j
                      left join smart_alpha.user_redeem_underlying_events r
                                on j.user_address = r.user_address and j.pool_address = r.pool_address and
                                   j.epoch_id = r.epoch_id and r.block_timestamp < ts
                      inner join smart_alpha.pools p on j.pool_address = p.pool_address
             where j.user_address = addr
               and j.block_timestamp < ts
               and r.user_address is null );
end;
$$;
