create table smart_yield.controller_harvests
(
    protocol_id           text    not null,

    controller_address    text    not null,
    caller_address        text    not null,
    comp_reward_total     numeric(78),
    comp_reward_sold      numeric(78),
    underlying_pool_share numeric(78),
    underlying_reward     numeric(78),
    harvest_cost          numeric(78),

    block_timestamp       bigint  not null,
    included_in_block     bigint  not null,
    tx_hash               text    not null,
    tx_index              integer not null,
    log_index             integer not null,
    created_at            timestamp default now()
);
