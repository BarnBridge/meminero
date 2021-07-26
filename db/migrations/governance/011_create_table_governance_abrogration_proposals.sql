create table governance.abrogation_proposals
(
    proposal_id       bigint  not null,
    creator           text    not null,
    create_time       bigint  not null,
    description       text    not null,

    included_in_block bigint  not null,
    tx_hash           text    not null,
    tx_index          integer not null,
    log_index         integer not null,
    created_at        timestamp default now()
);

create index abrogation_proposals_proposal_id_idx on governance.abrogation_proposals (proposal_id desc);
