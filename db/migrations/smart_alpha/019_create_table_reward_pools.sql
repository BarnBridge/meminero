create type smart_alpha.reward_pool_type as enum ('SINGLE', 'MULTI');
create table smart_alpha.reward_pools
(
    pool_type              smart_alpha.reward_pool_type default 'SINGLE'::smart_alpha.reward_pool_type,

    pool_address           text                                   not null,
    pool_token_address     text                                   not null,
    reward_token_addresses text[]                                 not null,

    start_at_block         bigint                       default 0 not null,

    created_at             timestamp                    default now()
);

create unique index reward_pools_pool_address_idx on smart_alpha.reward_pools (pool_address);

create table smart_alpha.rewards_claims
(
    user_address         text    not null,
    amount               numeric(78),
    pool_address         text    not null,
    reward_token_address text,

    block_timestamp      bigint  not null,
    included_in_block    bigint  not null,
    tx_hash              text    not null,
    tx_index             integer not null,
    log_index            integer not null,
    created_at           timestamp default now()
);

create type smart_alpha.reward_action as enum ('DEPOSIT','WITHDRAW');
create table smart_alpha.rewards_staking_actions
(
    user_address      text                      not null,
    amount            numeric(78),
    balance_after     numeric(78),
    action_type       smart_alpha.reward_action not null,
    pool_address      text                      not null,

    block_timestamp   bigint                    not null,
    included_in_block bigint                    not null,
    tx_hash           text                      not null,
    tx_index          integer                   not null,
    log_index         integer                   not null,
    created_at        timestamp default now()
);

create index rewards_staking_actions_pool_addr_idx on smart_alpha.rewards_staking_actions (pool_address asc,
                                                                                           user_address asc,
                                                                                           included_in_block desc,
                                                                                           tx_index desc, log_index
                                                                                           desc);

