create table notifications
(
    id bigserial
        constraint notifications_pkey
            primary key,
    target text,
    type text not null,
    starts_on bigint,
    expires_on bigint not null,
    message text,
    metadata jsonb,
    included_in_block bigint,
    created_on timestamp default now()
);

create index notifications_target_starts_on_index
    on notifications (target asc, starts_on desc);

create index notifications_included_in_block_index
    on notifications (included_in_block desc);

---- create above / drop below ----

drop table if exists notifications;
