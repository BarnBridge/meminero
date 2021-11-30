create materialized view if not exists smart_yield.junior_apy_30d_avg as
select pool_address, avg(junior_apy) as avg_junior_apy, now() as last_updated
from smart_yield.pool_state
where block_timestamp > ( select extract(epoch from now() - interval '30 days') )::bigint
group by pool_address;

create unique index if not exists smart_yield_junior_apy_30d_avg_pool_address_idx_unique on smart_yield.junior_apy_30d_avg (pool_address);

create index if not exists smart_yield_pool_state_block_timestamp_idx on smart_yield.pool_state (block_timestamp);
create index if not exists erc20_transfers_token_address_sender_idx on public.erc20_transfers (token_address, sender);
create index if not exists erc20_transfers_token_address_receiver_idx on public.erc20_transfers (token_address, receiver);
create index if not exists smart_yield_junior_2step_withdraw_events_pool_address_idx on smart_yield.junior_2step_withdraw_events (pool_address);

