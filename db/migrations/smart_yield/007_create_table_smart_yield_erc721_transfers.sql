create table smart_yield.erc721_transfers
(
    token_address     text    not null,
    token_type        text    not null,
    sender            text    not null,
    receiver          text    not null,
    token_id          bigint  not null,
    block_timestamp   bigint  not null,
    included_in_block bigint  not null,
    tx_hash           text    not null,
    tx_index          integer not null,
    log_index         integer not null,
    created_at        timestamp default now()
);

create index erc721_transfers_token_address_id_idx
    on smart_yield.erc721_transfers (token_address asc, token_id asc, block_timestamp desc);

create index erc721_transfers_token_type_receiver_idx
    on smart_yield.erc721_transfers (token_type asc, receiver asc, block_timestamp desc);


---- create above / drop below ----

drop table if exists smart_yield.erc721_transfers;
drop index if exists smart_yield.erc721_transfers_token_address_id_idx;
drop index if exists smart_yield.erc721_transfers_token_type_receiver_idx;