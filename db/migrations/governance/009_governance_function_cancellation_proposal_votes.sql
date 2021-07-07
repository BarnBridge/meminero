create or replace function governance.abrogation_proposal_votes(id bigint)
    returns table
            (
                user_id         text,
                support         boolean,
                block_timestamp bigint,
                power           numeric(78)
            )
    language plpgsql
as
$$
begin
    return query
        select distinct v.user_id,
                        first_value(v.support) over (partition by v.user_id order by v.block_timestamp desc)         as support,
                        first_value(v.block_timestamp)
                        over (partition by v.user_id order by v.block_timestamp desc)                                as block_timestamp,
                        v.power
        from governance.governance_abrogation_votes v
        where proposal_id = id
          and (select count(*)
               from governance.governance_abrogation_votes_canceled vc
               where vc.proposal_id = v.proposal_id
                 and vc.user_id = v.user_id
                 and vc.block_timestamp > v.block_timestamp) = 0;
end;
$$;

---- create above / drop below ----

drop function if exists abrogation_proposal_votes;