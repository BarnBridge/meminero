create table smart_yield.pool_state
(
    pool_address       text             not null,
    senior_liquidity   numeric(78),
    junior_liquidity   numeric(78),
    jtoken_price       numeric(78),
    abond_principal    numeric(78),
    abond_gain         numeric(78),
    abond_issued_at    bigint,
    abond_matures_at   bigint,
    abond_apy          double precision not null,
    senior_apy         double precision not null,
    junior_apy         double precision not null,
    originator_apy     double precision not null,
    originator_net_apy double precision not null,

    block_timestamp    bigint           not null,
    included_in_block  bigint           not null
);

create index state_pool_address_idx on smart_yield.pool_state (pool_address asc, block_timestamp desc);
