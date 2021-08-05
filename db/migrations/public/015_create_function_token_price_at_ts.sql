create or replace function public.token_price_at_ts(addr text, quote text, ts bigint) returns double precision
    language plpgsql as
$$
begin
    return ( select price
             from public.token_prices p
             where p.token_address = addr
               and quote_asset = quote
               and block_timestamp <= ts
             order by block_timestamp desc
             limit 1 );
end;
$$;
