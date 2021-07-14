create table governance.barn_locks
(
    user_address      text    not null,
    locked_until      bigint,

    block_timestamp   bigint,
    included_in_block bigint  not null,
    tx_hash           text    not null,
    tx_index          integer not null,
    log_index         integer not null,
    created_at        timestamp default now()
);

create index user_locked_until_idx on governance.barn_locks (user_address asc, included_in_block desc, log_index desc);

---- create above / drop below ----

drop table if exists governance.barn_locks;
drop index if exists governance.user_locked_until_idx;
