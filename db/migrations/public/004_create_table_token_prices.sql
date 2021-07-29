create table public.token_prices
(
    token_address     text             not null,
    token_symbol      text             not null,
    quote_asset       text             not null,
    price             double precision not null,

    block_timestamp   bigint           not null,
    included_in_block bigint           not null,
    created_at        timestamp default now()
);

create index token_prices_token_address_idx on public.token_prices (token_address, block_timestamp desc);
create index token_prices_token_address_quote_asset_idx on public.token_prices (token_address, quote_asset, block_timestamp desc);

create or replace function token_usd_price_at_ts(addr text, ts bigint) returns double precision
    language plpgsql as
$$
declare
    _price double precision;
begin
    select into _price price
    from public.token_prices p
    where p.token_address = addr
      and quote_asset = "USD"
      and block_timestamp <= ts
    order by block_timestamp desc
    limit 1;

    return _price;
end;
$$;
