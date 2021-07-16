create table smart_yield.pools
(
    protocol_id           text                not null,
    pool_address          text                not null,

    controller_address    text                not null,
    model_address         text                not null,
    provider_address      text                not null,
    oracle_address        text                not null,
    junior_bond_address   text                not null,
    senior_bond_address   text                not null,

    receipt_token_address text                not null,
    underlying_address    text                not null,
    underlying_symbol     text                not null,
    underlying_decimals   integer             not null,

    start_at_block        bigint    default 0 not null,

    created_at            timestamp default now()
);

---- create above / drop below ----

drop table if exists smart_yield.pools;
