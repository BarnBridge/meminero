create function yield_farming.yf_sum_actions_by_token(a_type staking_action_type, _token_address text, _start_time bigint, _end_time bigint, scale text)
    returns TABLE(amount numeric, point timestamp with time zone)
    language plpgsql
as
$$
begin
    return query select sum(yf.amount),
                        date_trunc(scale, to_timestamp(block_timestamp)) as wk
                 from yield_farming.yield_farming_actions yf
                 where action_type = a_type
                   and token_address = _token_address
                   and block_timestamp between _start_time and _end_time
                 group by wk
                 order by wk;
end;
$$;

create function yield_farming.yf_stats_by_token(_token_address text, _start_time bigint, _end_time bigint, scale text)
    returns TABLE(point timestamp with time zone, sum_deposits numeric, sum_withdrawals numeric)
    language plpgsql
as
$$
begin
    return query select coalesce(d.point, w.point), coalesce(d.amount, 0) as sum_deposits, coalesce(w.amount, 0) as sum_withdrawals
                 from yield_farming.yf_sum_actions_by_token('DEPOSIT', _token_address, _start_time,_end_time,scale ) d
                          full outer join yield_farming.yf_sum_actions_by_token('WITHDRAW', _token_address,  _start_time,_end_time,scale ) w
                                          on d.point = w.point;
end
$$;
---- create above / drop below ----

drop function if exists yield_farming.yf_sum_actions_by_token(a_type staking_action_type, _token_address text, _start_time bigint, _end_time bigint, scale text);
drop function if exists yield_farming.yf_stats_by_token(_token_address text, _start_time bigint, _end_time bigint, scale text);