create table public.blocks
(
    number              bigint not null,
    block_hash          text   not null,
    parent_block_hash   text   not null,
    block_creation_time bigint not null,
    created_at          timestamp default now()
);

create index blocks_block_hash_idx on public.blocks (block_hash);

create unique index blocks_number_idx on public.blocks (number);
