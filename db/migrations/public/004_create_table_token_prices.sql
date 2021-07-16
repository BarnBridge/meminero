create table public.token_prices
(
    token_address     text             not null,
    token_symbol      text,
    price_usd         double precision not null,

    block_timestamp   bigint           not null,
    included_in_block bigint           not null,
    created_at        timestamp default now()
);


---- create above / drop below ----

