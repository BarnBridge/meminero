create type delegate_change_type as enum('INCREASE','DECREASE');
create table barn.barn_delegate_changes
(
    tx_hash text not null,
    tx_index integer not null,
    log_index integer not null,
    logged_by text not null,
    action_type delegate_change_type not null,
    sender text not null,
    receiver text not null,
    amount numeric(78) not null,
    receiver_new_delegated_power numeric(78) not null,
    timestamp bigint,
    included_in_block bigint not null,
    created_at timestamp default now()
);

create index user_delegated_power_idx
    on barn.barn_delegate_changes (receiver asc, included_in_block desc, log_index desc);

create trigger refresh_barn_users
    after insert or update or delete or truncate
    on barn.barn_delegate_changes
execute procedure barn.refresh_barn_users();


---- create above / drop below ----

drop table if exists barn.barn_delegate_changes;
drop index if exists barn.user_delegated_power_idx;