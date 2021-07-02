package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateFunctionProposalState, downCreateFunctionProposalState)
}

func upCreateFunctionProposalState(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create or replace function bond_staked_at_ts(ts timestamp with time zone) returns numeric(78)
			language plpgsql as
		$$
		declare
			value numeric(78);
		begin
			with values as ( select action_type, sum(amount) as amount
							 from barn_staking_actions
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

		create or replace function proposal_state(id bigint) returns text
			language plpgsql as
		$$
		declare
			createTime                 bigint;
			warmUpDuration             bigint;
			activeDuration             bigint;
			queueDuration              bigint;
			gracePeriodDuration        bigint;
			acceptanceThreshold        bigint;
			minQuorum                  bigint;
			bondStaked                 numeric(78);
			forVotes                   numeric(78);
			againstVotes               numeric(78);
			eta                        bigint;
			abrogationProposalQuorum numeric(78);
		begin
			if ( select count(*) from governance_events where proposal_id = id and event_type = 'CANCELED' ) > 0 then
				return 'CANCELED';
			end if;
		
			if ( select count(*) from governance_events where proposal_id = id and event_type = 'EXECUTED' ) > 0 then
				return 'EXECUTED';
			end if;
		
			select into createTime, warmUpDuration, activeDuration, queueDuration, gracePeriodDuration, acceptanceThreshold, minQuorum create_time,
																																	   warm_up_duration,
																																	   active_duration,
																																	   queue_duration,
																																	   grace_period_duration,
																																	   acceptance_threshold,
																																	   min_quorum
			from governance_proposals
			where proposal_id = id;
		
			if ( select extract(epoch from now()) ) <= (createTime + warmUpDuration) then return 'WARMUP'; end if;
			if ( select extract(epoch from now()) ) <= (createTime + warmUpDuration + activeDuration) then
				return 'ACTIVE';
			end if;
		
			select into bondStaked bond_staked_at_ts(to_timestamp(createTime + warmUpDuration));
		
			with total_votes as ( select support, sum(power) as power from proposal_votes(id) group by support )
			select into forVotes, againstVotes coalesce(( select coalesce(power, 0) from total_votes where support = true ), 0),
											   coalesce(( select coalesce(power, 0) from total_votes where support = false ), 0);
		
			-- check if quorum is met
			if (forVotes + againstVotes < minQuorum / 100 * bondStaked) then return 'FAILED'; end if;
		
			-- check if votes met the acceptance threshold
			if (forVotes < ((forVotes + againstVotes) * acceptanceThreshold / 100)) then return 'FAILED'; end if;
		
			if ( select count(*) from governance_events where proposal_id = id and event_type = 'QUEUED' ) = 0 then
				return 'ACCEPTED';
			end if;
		
			select into eta event_data -> 'eta' as eta from governance_events where proposal_id = id and event_type = 'QUEUED';
		
			if ( select extract(epoch from now()) ) < eta then return 'QUEUED'; end if;
		
			-- check if there's a abrogation proposal that passed
			if ( select count(*) from governance_abrogation_proposals where proposal_id = id ) > 0 then
				select into abrogationProposalQuorum bond_staked_at_ts(to_timestamp(( select create_time - 1
																		   from governance_abrogation_proposals
																		   where proposal_id = id ))) / 2;
		
				if coalesce(( select sum(power) from abrogation_proposal_votes(id) where support = true ), 0) >= abrogationProposalQuorum then
					return 'ABROGATED';
				end if;
			end if;
		
			if ( select extract(epoch from now()) ) <= eta + gracePeriodDuration then return 'GRACE'; end if;

			return 'EXPIRED';
		end;
		$$;
	`)

	return err
}

func downCreateFunctionProposalState(tx *sql.Tx) error {
	_, err := tx.Exec(`
		drop function if exists proposal_state;
		drop function if exists bond_staked_at_ts;
	`)

	return err
}
