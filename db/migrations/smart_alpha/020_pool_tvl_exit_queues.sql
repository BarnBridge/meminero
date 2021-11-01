create or replace function smart_alpha.pool_tvl_v2(pool text)
    returns table
            (
                epoch_junior_tvl       double precision,
                epoch_senior_tvl       double precision,
                junior_entry_queue_tvl double precision,
                senior_entry_queue_tvl double precision,
                junior_exit_queue_tvl  double precision,
                senior_exit_queue_tvl  double precision,
                junior_exited_tvl      double precision,
                senior_exited_tvl      double precision
            )
    language plpgsql
as
$$
declare
    token_price    double precision;
    token_address  text;
    token_decimals integer;
    epoch_jl       numeric(78, 18);
    epoch_sl       numeric(78, 18);
    j_entry        numeric(78, 18);
    jtokens_burn   numeric(78, 18);
    j_exit         numeric(78, 18);
    s_entry        numeric(78, 18);
    stokens_burn   numeric(78, 18);
    s_exit         numeric(78, 18);
begin
    select into token_address, token_decimals p.pool_token_address, p.pool_token_decimals
    from smart_alpha.pools p
    where pool_address = pool;

    select into token_price public.token_usd_price_at_ts(token_address, ( select extract(epoch from now())::bigint ));

    select into epoch_jl, epoch_sl junior_liquidity, senior_liquidity
    from smart_alpha.pool_epoch_info
    where pool_address = pool
    order by epoch_id desc
    limit 1;

    select into j_entry, jtokens_burn, j_exit, s_entry, stokens_burn, s_exit queued_juniors_underlying_in,
                                                                             queued_junior_tokens_burn / pow(10, 18) * estimated_junior_token_price,
                                                                             queued_juniors_underlying_out,
                                                                             queued_seniors_underlying_in,
                                                                             queued_senior_tokens_burn / pow(10, 18) * estimated_senior_token_price,
                                                                             queued_seniors_underlying_out
    from smart_alpha.pool_state
    where pool_address = pool
    order by block_timestamp desc
    limit 1;

    return query select coalesce(epoch_jl, 0) / pow(10, token_decimals) * token_price     as epoch_junior_tvl,
                        coalesce(epoch_sl, 0) / pow(10, token_decimals) * token_price     as epoch_senior_tvl,
                        coalesce(j_entry, 0) / pow(10, token_decimals) * token_price      as junior_entry_queue_tvl,
                        coalesce(s_entry, 0) / pow(10, token_decimals) * token_price      as senior_entry_queue_tvl,
                        coalesce(jtokens_burn, 0) / pow(10, token_decimals) * token_price as junior_exit_queue_tvl,
                        coalesce(stokens_burn, 0) / pow(10, token_decimals) * token_price as senior_exit_queue_tvl,
                        coalesce(j_exit, 0) / pow(10, token_decimals) * token_price       as junior_exited_tvl,
                        coalesce(s_exit, 0) / pow(10, token_decimals) * token_price       as senior_exited_tvl;
end
$$;
