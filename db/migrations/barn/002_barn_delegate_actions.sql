create type delegate_action_type as enum('START','STOP');
create table barn.barn_delegate_actions
(
    tx_hash text not null,
    tx_index integer not null,
    log_index integer not null,
    logged_by text not null,
    sender text not null,
    receiver text not null,
    action_type delegate_action_type not null,
    timestamp bigint not null,
    included_in_block bigint not null,
    created_at timestamp default now()
);

create index user_delegation_idx
    on barn.barn_delegate_actions (sender asc, included_in_block desc, log_index desc);


---- create above / drop below ----

drop table if exists barn.barn_delegate_actions;
drop index if exists barn.user_delegation_idx;
drop type if exists  delegate_action_type;
