create table if not exists smart_alpha.pools
(
    pool_name            text    not null,
    pool_address         text    not null,
    pool_token_address   text    not null,
    pool_token_symbol    text    not null,
    pool_token_decimals  integer not null,
    junior_token_address text    not null,
    senior_token_address text    not null,
    oracle_address       text    not null,
    oracle_asset_symbol  text    not null,
    epoch1_start         bigint  not null,
    epoch_duration       bigint  not null,
    start_at_block       bigint  not null
);

create unique index if not exists pools_pool_address_idx on smart_alpha.pools (pool_address);
