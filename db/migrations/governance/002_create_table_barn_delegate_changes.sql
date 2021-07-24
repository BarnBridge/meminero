create type delegate_change_type as enum ('INCREASE','DECREASE');
create table governance.barn_delegate_changes
(
    action_type                  delegate_change_type not null,
    sender                       text                 not null,
    receiver                     text                 not null,
    amount                       numeric(78)          not null,
    receiver_new_delegated_power numeric(78)          not null,

    block_timestamp              bigint,
    included_in_block            bigint               not null,
    tx_hash                      text                 not null,
    tx_index                     integer              not null,
    log_index                    integer              not null,
    created_at                   timestamp default now()
);

create index user_delegated_power_idx on governance.barn_delegate_changes (receiver asc, included_in_block desc, log_index desc);
