create table governance.governance_proposals
(
    proposal_id           bigint not null,
    proposer              text   not null,
    description           text   not null,
    title                 text   not null,
    create_time           bigint not null,
    targets               jsonb  not null,
    values                jsonb  not null,
    signatures            jsonb  not null,
    calldatas             jsonb  not null,
    block_timestamp       bigint not null,
    included_in_block     bigint not null,
    created_at            timestamp default now(),
    warm_up_duration      bigint,
    active_duration       bigint,
    queue_duration        bigint,
    grace_period_duration bigint,
    acceptance_threshold  bigint,
    min_quorum            bigint
);

create index governance_proposals_proposal_id_idx
    on governance.governance_proposals (proposal_id desc);

create index governance_proposals_proposer_idx
    on governance.governance_proposals (lower(proposer));


---- create above / drop below ----

drop table if exists governance.governance_proposals;
drop index if exists governance.governance_proposals_proposal_id_idx;
drop index if exists governance.governance_proposals_proposer_idx;
