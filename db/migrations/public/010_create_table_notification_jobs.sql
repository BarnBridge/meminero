create table notification_jobs
(
    id bigserial
        constraint notification_jobs_pkey
            primary key,
    type text not null,
    execute_on bigint,
    metadata jsonb,
    included_in_block bigint,
    deleted boolean default false,
    created_on timestamp default now()
);

create index notification_jobs_execute_on_index
    on notification_jobs (execute_on desc);

create index notification_jobs_deleted_on_index
    on notification_jobs (deleted desc);

create index notification_jobs_included_in_block_index
    on notification_jobs (included_in_block desc);

---- create above / drop below ----

drop table if exists notification_jobs;
