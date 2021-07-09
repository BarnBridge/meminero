create table governance.votes
(
    proposal_id       bigint      not null,
    user_id           text        not null,
    support           boolean     not null,
    power             numeric(78) not null,
    block_timestamp   bigint      not null,
    tx_hash           text        not null,
    tx_index          integer     not null,
    log_index         integer     not null,
    logged_by         text        not null,
    included_in_block bigint      not null,
    created_at        timestamp default now()
);

create index votes_proposal_id_idx
    on governance.votes (proposal_id desc);

create index votes_proposal_id_composed_idx
    on governance.votes (proposal_id asc, user_id asc, block_timestamp desc);

create index votes_user_id_idx
    on governance.votes (lower(user_id));

---- create above / drop below ----

drop table if exists governance.votes;
drop index if exists governance.votes_proposal_id_idx;
drop index if exists governance.votes_proposal_id_composed_idx;
drop index if exists governance.votes_user_id_idx;

