package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upAddIndexes, downAddIndexes)
}

func upAddIndexes(tx *sql.Tx) error {
	_, err := tx.Exec(`
		-- user voting power
		create index user_balance_idx on barn_staking_actions (user_address, included_in_block desc, log_index desc);
		create index user_delegation_idx on barn_delegate_actions (sender, included_in_block desc, log_index desc);
		create index user_locked_until_idx on barn_locks (user_address, included_in_block desc, log_index desc);
		create index user_delegated_power_idx on barn_delegate_changes (receiver, included_in_block desc, log_index desc);
		
		-- proposal
		create index governance_proposals_proposal_id_idx on governance_proposals(proposal_id desc);
		create index governance_votes_proposal_id_idx on governance_votes(proposal_id desc);
		create index governance_votes_proposal_id_composed_idx on governance_votes(proposal_id, user_id, block_timestamp desc);
		create index governance_votes_canceled_idx on governance_votes_canceled(proposal_id, user_id, block_timestamp desc);
		create index governance_abrogation_votes_proposal_id_idx on governance_abrogation_votes(proposal_id desc);
		create index governance_abrogation_votes_proposal_id_composed_idx on governance_abrogation_votes(proposal_id, user_id, block_timestamp desc);
		create index governance_abrogation_votes_canceled_idx on governance_abrogation_votes_canceled(proposal_id, user_id, block_timestamp desc);

		-- voters
		create index governance_votes_user_id_idx on governance_votes (lower(user_id));
		create index governance_abrogation_votes_user_id_idx on governance_abrogation_votes (lower(user_id));
		create index governance_proposals_proposer_idx on governance_proposals (lower(proposer));
		
		-- bond_staked_at_ts
		create index barn_staking_actions_included_in_block_idx on barn_staking_actions (included_in_block desc);

		-- proposal state
		create index governance_votes_proposal_id_event_type_idx on governance_events(proposal_id, event_type);
		create index governance_abrogation_proposals_proposal_id_idx on governance_abrogation_proposals(proposal_id desc);
	`)
	return err
}

func downAddIndexes(tx *sql.Tx) error {
	_, err := tx.Exec(`
		drop index if exists user_balance_idx;
		drop index if exists user_delegation_idx;
		drop index if exists user_locked_until_idx;
		drop index if exists user_delegated_power_idx;
		drop index if exists governance_proposals_proposal_id_idx;
		drop index if exists governance_votes_proposal_id_idx;
		drop index if exists governance_votes_proposal_id_composed_idx;
		drop index if exists governance_votes_canceled_idx;
		drop index if exists governance_abrogation_votes_proposal_id_idx;
		drop index if exists governance_abrogation_votes_proposal_id_composed_idx;
		drop index if exists governance_abrogation_votes_canceled_idx;
		drop index if exists governance_votes_user_id_idx;
		drop index if exists governance_abrogation_votes_user_id_idx;
		drop index if exists governance_proposals_proposer_idx;
		drop index if exists barn_staking_actions_included_in_block_idx;
		drop index if exists governance_votes_proposal_id_event_type_idx;
		drop index if exists governance_abrogation_proposals_proposal_id_idx;
	`)
	return err
}
