create table smart_yield.state
(
    included_in_block  bigint           not null,
    block_timestamp    timestamp        not null,
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
    originator_net_apy double precision not null
);

create index state_pool_address_idx
    on smart_yield.state (pool_address asc, block_timestamp desc);

create index state_apy_trend_idx
    on smart_yield.state (pool_address, date_trunc('day'::text, block_timestamp));


---- create above / drop below ----

drop table if exists smart_yield.state;
drop index if exists smart_yield.state_pool_address_idx;
drop index if exists smart_yield.state_apy_trend_idx;