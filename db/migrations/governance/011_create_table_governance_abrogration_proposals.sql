create table governance.abrogation_proposals
(
<<<<<<< HEAD
    proposal_id       numeric(78) not null,
    creator           text        not null,
    create_time       bigint      not null,
    description       text        not null,
    included_in_block bigint      not null,
    tx_hash           text        not null,
    tx_index          integer     not null,
    log_index         integer     not null,
=======
    proposal_id       bigint  not null,
    creator           text    not null,
    create_time       bigint  not null,
    description       text    not null,

    included_in_block bigint  not null,
    tx_hash           text    not null,
    tx_index          integer not null,
    log_index         integer not null,
>>>>>>> origin
    created_at        timestamp default now()
);

create index abrogation_proposals_proposal_id_idx on governance.abrogation_proposals (proposal_id desc);

---- create above / drop below ----

drop table if exists governance.abrogation_proposals;
drop index if exists governance.abrogation_proposals_proposal_id_idx;
