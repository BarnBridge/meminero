create type event_type as enum ('CREATED','QUEUED','EXECUTED','CANCELED');
create table governance.governance_events
(
    proposal_id       bigint     not null,
    caller            text       not null,
    event_type        event_type not null,
    block_timestamp   bigint     not null,
    tx_hash           text       not null,
    tx_index          integer    not null,
    log_index         integer    not null,
    logged_by         text       not null,
    event_data        jsonb,
    included_in_block bigint     not null,
    created_at        timestamp default now()
);


create index governance_votes_proposal_id_event_type_idx
    on governance.governance_events (proposal_id, event_type);

---- create above / drop below ----

drop table if exists governance.governance_events;
drop index if exists governance.governance_votes_proposal_id_event_type_idx;