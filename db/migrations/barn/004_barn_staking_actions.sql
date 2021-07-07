create type action_type as enum('DEPOSIT','WITHDRAW');
create table barn.barn_staking_actions
(
    tx_hash text not null,
    tx_index integer not null,
    log_index integer not null,
    address text not null,
    user_address text not null,
    action_type action_type not null,
    amount numeric(78) not null,
    balance_after numeric(78) not null,
    included_in_block bigint not null,
    created_at timestamp default now()
);

create index user_balance_idx
    on barn.barn_staking_actions (user_address asc, included_in_block desc, log_index desc);

create index barn_staking_actions_included_in_block_idx
    on barn.barn_staking_actions (included_in_block desc);

create trigger refresh_barn_users
    after insert or update or delete or truncate
    on barn.barn_staking_actions
execute procedure  barn.refresh_barn_users();

---- create above / drop below ----

drop table if exists barn.barn_staking_actions;
drop index if exists barn.user_balance_idx;
drop index if exists barn.barn_staking_actions_included_in_block_idx;
