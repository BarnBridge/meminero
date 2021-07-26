create table public.integrity_checkpoints
(
    number     bigint,
    created_at timestamp default now()
);

create index integrity_checkpoints_created_at_idx on public.integrity_checkpoints (created_at desc);
