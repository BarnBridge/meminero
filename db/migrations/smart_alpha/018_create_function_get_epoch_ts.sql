create or replace function smart_alpha.get_epoch_ts(_pool_address text, epoch_id bigint)
    returns table
            (
                start_ts bigint,
                end_ts   bigint
            )
    language plpgsql
as
$$

begin
    return query
        select epoch1_start + (epoch_id - 1) * epoch_duration + 1 as start_ts,
               epoch1_start + epoch_id * epoch_duration           as end_ts
        from smart_alpha.pools p
        where p.pool_address = _pool_address;
end
$$;