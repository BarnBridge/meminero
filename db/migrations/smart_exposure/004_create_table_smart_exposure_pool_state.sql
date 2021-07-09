create table smart_exposure.pool_state
(
    included_in_block     bigint not null,
    block_timestamp       bigint not null,
    pool_address          text   not null,
    pool_liquidity        double precision,
    last_rebalance        bigint,
    rebalancing_interval  bigint,
    rebalancing_condition numeric(78)
);

---- create above / drop below ----

drop table if exists smart_exposure.pool_state;