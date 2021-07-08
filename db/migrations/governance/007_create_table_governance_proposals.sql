create table governance.proposals
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

create index proposals_proposal_id_idx
    on governance.proposals (proposal_id desc);

create index proposals_proposer_idx
    on governance.proposals (lower(proposer));


---- create above / drop below ----

drop table if exists governance.proposals;
drop index if exists governance.proposals_proposal_id_idx;
drop index if exists governance.proposals_proposer_idx;
