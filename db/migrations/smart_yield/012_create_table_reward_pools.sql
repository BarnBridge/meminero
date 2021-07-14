create type reward_pool_type as enum ('SINGLE', 'MULTI');
create table smart_yield.reward_pools
(
    pool_type              reward_pool_type default 'SINGLE'::reward_pool_type,

    pool_address           text                       not null,
    pool_token_address     text                       not null,
    reward_token_addresses text[]                     not null,

    start_at_block         bigint           default 0 not null,

    created_at             timestamp        default now()
);


---- create above / drop below ----

drop table if exists smart_yield.reward_pools;
drop type if exists reward_pool_type;
