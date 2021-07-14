create type event_type as enum ('CREATED','QUEUED','EXECUTED','CANCELED');

create table governance.proposal_events
(
    proposal_id       bigint     not null,
    caller            text       not null,
    event_type        event_type not null,
    event_data        jsonb,

    block_timestamp   bigint     not null,
    included_in_block bigint     not null,
    tx_hash           text       not null,
    tx_index          integer    not null,
    log_index         integer    not null,
    created_at        timestamp default now()
);


create index proposal_events_id_event_type_idx on governance.proposal_events (proposal_id, event_type);

---- create above / drop below ----

drop table if exists governance.proposal_events;
drop index if exists governance.proposal_events_id_event_type_idx;
