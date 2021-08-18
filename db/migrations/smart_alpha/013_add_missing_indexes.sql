create index if not exists pool_state_pool_address_block_timestamp_idx on smart_alpha.pool_state (pool_address, block_timestamp desc);
create index if not exists pool_epoch_info_pool_address_block_timestamp_idx on smart_alpha.pool_epoch_info (pool_address, block_timestamp desc);
