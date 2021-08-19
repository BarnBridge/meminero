alter table smart_alpha.pool_epoch_info
    add column if not exists epoch_entry_price numeric(78);
