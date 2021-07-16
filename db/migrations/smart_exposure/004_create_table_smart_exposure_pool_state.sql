create table smart_exposure.pool_state
(
    pool_address          text   not null,
    pool_liquidity        double precision,
    last_rebalance        bigint,
    rebalancing_interval  bigint,
    rebalancing_condition numeric(78),

    block_timestamp       bigint not null,
    included_in_block     bigint not null
);

create index pool_state_pool_address_idx
    on smart_exposure.pool_state (pool_address asc, block_timestamp desc);

---- create above / drop below ----

drop table if exists smart_exposure.pool_state;
