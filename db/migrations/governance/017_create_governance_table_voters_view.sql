create or replace view governance.voters as
select user_address,
       governance.balance_of(user_address)                      as bond_staked,
       coalesce((select locked_until
                 from governance.barn_locks
                 where user_address = barn_users.user_address
                 order by included_in_block desc, log_index desc
                 limit 1),
                0)                                              as locked_until,
       governance.delegated_power(user_address),
       (select count(*) from governance.votes where lower(user_id) = lower(barn_users.user_address)) +
       (select count(*)
        from governance.abrogation_votes
        where lower(user_id) = lower(barn_users.user_address))  as votes,
       (select count(*)
        from governance.proposals
        where lower(proposer) = lower(barn_users.user_address)) as proposals,
       governance.voting_power(user_address)                    as voting_power,
       governance.has_active_delegation(user_address)
from governance.barn_users;

---- create above / drop below ----

drop view governance.voters