create or replace function governance.balance_of(addr text) returns numeric(78)
    language plpgsql as
$$
declare
    result numeric(78);
begin
    select balance_after
    into result
    from governance.barn_staking_actions
    where user_address = addr
    order by included_in_block desc, log_index desc
    limit 1;

    return coalesce(result, 0);
end;
$$;

create or replace function governance.has_active_delegation(addr text) returns bool
    language plpgsql as
$$
declare
    action text;
begin
    select action_type
    into action
    from governance.barn_delegate_actions
    where sender = addr
    order by included_in_block desc, log_index desc
    limit 1;

    if action = 'START' then return true; end if;

    return false;
end;
$$;

create or replace function governance.user_multiplier(addr text) returns numeric(78)
    language plpgsql as
$$
declare
    locked_until_ts bigint;
    time_now        bigint;
    time_left       bigint;
    multiplier      numeric(78);
begin
    multiplier = 1 * 10 ^ 18;

    select locked_until
    into locked_until_ts
    from governance.barn_locks
    where user_address = addr
    order by included_in_block desc, log_index desc
    limit 1;

    if not found then return multiplier; end if;

    select floor(extract(epoch from now())) into time_now;

    if locked_until_ts <= time_now then return multiplier; end if;

    time_left = locked_until_ts - time_now;

    multiplier = multiplier + (time_left::numeric * 10 ^ 18 / 31536000::numeric);

    return multiplier;
end;
$$;

create or replace function governance.delegated_power(addr text) returns numeric(78)
    language plpgsql as
$$
declare
    result numeric(78);
begin
    select receiver_new_delegated_power
    into result
    from governance.barn_delegate_changes
    where receiver = addr
    order by included_in_block desc, log_index desc;

    return coalesce(result, 0);
end;
$$;

create or replace function governance.voting_power(addr text) returns numeric(78)
    language plpgsql as
$$
declare
    is_delegating   bool;
    delegated_power numeric(78);
    self_power      numeric(78);
begin
    select governance.has_active_delegation(addr) into is_delegating;

    if is_delegating then
        self_power = 0;
    else
        select governance.balance_of(addr) * governance.user_multiplier(addr) / 10 ^ 18 into self_power;
    end if;

    select governance.delegated_power(addr) into delegated_power;

    return self_power + delegated_power;
end;
$$;

create function governance.bond_staked_at_ts(ts timestamp with time zone) returns numeric
    language plpgsql
as
$$
declare
    value numeric(78);
begin
    with values as ( select action_type, sum(amount) as amount
                     from governance.barn_staking_actions
                     where included_in_block < ( select number
                                                 from blocks
                                                 where block_creation_time < ts
                                                 order by block_creation_time desc
                                                 limit 1 )
                     group by action_type )
    select into value coalesce(( select amount from values where action_type = 'DEPOSIT' ),0) -
                      coalesce(( select amount from values where action_type = 'WITHDRAW' ),0);

    return value;
end;

$$;

create or replace function governance.refresh_barn_users() returns TRIGGER
    language plpgsql as
$$
begin
    refresh materialized view concurrently governance.barn_users;
    return null;
end
$$;

---- create above / drop below ----

drop function governance.delegated_power(addr text);
drop function governance.voting_power(addr text);
drop function governance.balance_of(addr text);
drop function governance.user_multiplier(addr text);
drop function governance.has_active_delegation(addr text);
drop function governance.bond_staked_at_ts(ts timestamp with time zone);
drop function governance.refresh_barn_users();
