create table governance.governance_abrogation_proposals
(
    proposal_id bigint not null,
    creator text not null,
    create_time bigint not null,
    description text not null,
    tx_hash text not null,
    tx_index integer not null,
    log_index integer not null,
    logged_by text not null,
    included_in_block bigint not null,
    created_at timestamp default now()
);

create index governance_abrogation_proposals_proposal_id_idx
    on governance.governance_abrogation_proposals (proposal_id desc);


---- create above / drop below ----

drop table if exists  governance.governance_abrogation_proposals;
drop index if exists  governance.governance_abrogation_proposals_proposal_id_idx;