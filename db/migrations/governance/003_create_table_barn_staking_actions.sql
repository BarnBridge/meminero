create type governance.action_type as enum ('DEPOSIT','WITHDRAW');
create table governance.barn_staking_actions
(
    user_address      text                   not null,
    action_type       governance.action_type not null,
    amount            numeric(78)            not null,
    balance_after     numeric(78)            not null,

    block_timestamp   bigint                 not null,
    included_in_block bigint                 not null,
    tx_hash           text                   not null,
    tx_index          integer                not null,
    log_index         integer                not null,
    created_at        timestamp default now()
);

create index user_balance_idx on governance.barn_staking_actions (user_address asc, included_in_block desc, log_index desc);

create index barn_staking_actions_included_in_block_idx on governance.barn_staking_actions (included_in_block desc);

---- create above / drop below ----

drop table if exists governance.barn_staking_actions;
drop index if exists governance.user_balance_idx;
drop index if exists governance.barn_staking_actions_included_in_block_idx;
drop type if exists governance.action_type;
