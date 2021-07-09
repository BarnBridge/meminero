create table governance.barn_locks
(
    tx_hash           text    not null,
    tx_index          integer not null,
    log_index         integer not null,
    logged_by         text    not null,
    user_address      text    not null,
    locked_until      bigint,
    locked_at         bigint,
    included_in_block bigint  not null,
    created_at        timestamp default now()
);
create index user_locked_until_idx
    on governance.barn_locks (user_address asc, included_in_block desc, log_index desc);

---- create above / drop below ----

drop table if exists governance.barn_locks;
drop index if exists governance.user_locked_until_idx;