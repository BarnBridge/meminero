create type reward_action as enum ('JUNIOR_STAKE','JUNIOR_UNSTAKE');
create table smart_yield.rewards_staking_actions
(
    user_address      text          not null,
    amount            numeric(78),
    balance_after     numeric(78),
    action_type       reward_action not null,
    pool_address      text          not null,

    block_timestamp   bigint        not null,
    included_in_block bigint        not null,
    tx_hash           text          not null,
    tx_index          integer       not null,
    log_index         integer       not null,
    created_at        timestamp default now()
);

create index rewards_staking_actions_pool_addr_idx on smart_yield.rewards_staking_actions (pool_address asc,
                                                                                           user_address asc,
                                                                                           included_in_block desc,
                                                                                           tx_index desc, log_index
                                                                                           desc);


---- create above / drop below ----

drop table if exists smart_yield.rewards_staking_actions;
drop index if exists smart_yield.rewards_staking_actions_pool_addr_idx;
drop type if exists reward_action;
