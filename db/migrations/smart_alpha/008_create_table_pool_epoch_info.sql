create table if not exists smart_alpha.pool_epoch_info
(
    pool_address             text             not null,
    epoch_id                 bigint           not null,
    senior_liquidity         numeric(78)      not null,
    junior_liquidity         numeric(78)      not null,
    upside_exposure_rate     double precision not null,
    downside_protection_rate double precision not null,
    junior_token_price_start numeric(78)      not null,
    senior_token_price_Start numeric(78)      not null,

    block_timestamp          bigint           not null,
    included_in_block        bigint           not null,
    created_at               timestamp default now()
);

create unique index if not exists pool_address_epoch_id_idx on smart_alpha.pool_epoch_info (pool_address, epoch_id desc);
