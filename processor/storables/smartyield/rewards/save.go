package rewards

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/processor/storables/smartyield"
	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	err := s.saveStakingActions(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save staking actions")
	}

	err = s.saveTxHistory(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save transaction history")
	}

	err = s.saveClaimEvents(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save claim events")
	}

	err = s.saveRewardPools(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save reward pools")
	}

	return nil
}

func (s *Storable) saveStakingActions(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.Deposits) == 0 && len(s.processed.Withdrawals) == 0 {
		return nil
	}

	var rows [][]interface{}

	for _, e := range s.processed.Deposits {
		p := s.state.SmartYield.RewardPoolByAddress(e.Raw.Address.String())

		rows = append(rows, []interface{}{
			utils.NormalizeAddress(e.User.String()),
			e.AmountDecimal(0),
			e.BalanceAfterDecimal(0),
			smartyield.JuniorStake,
			p.PoolAddress,
			s.block.BlockCreationTime,
			s.block.Number,
			utils.NormalizeAddress(e.Raw.TxHash.String()),
			e.Raw.TxIndex,
			e.Raw.Index,
		})
	}

	for _, e := range s.processed.Withdrawals {
		p := s.state.SmartYield.RewardPoolByAddress(e.Raw.Address.String())

		rows = append(rows, []interface{}{
			utils.NormalizeAddress(e.User.String()),
			e.AmountDecimal(0),
			e.BalanceAfterDecimal(0),
			smartyield.JuniorUnstake,
			p.PoolAddress,
			s.block.BlockCreationTime,
			s.block.Number,
			utils.NormalizeAddress(e.Raw.TxHash.String()),
			e.Raw.TxIndex,
			e.Raw.Index,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_yield", "rewards_staking_actions"},
		[]string{
			"user_address",
			"amount",
			"balance_after",
			"action_type",
			"pool_address",
			"block_timestamp",
			"included_in_block",
			"tx_hash",
			"tx_index",
			"log_index",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into rewards_staking_actions")
	}

	return nil
}

func (s *Storable) saveTxHistory(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.Deposits) == 0 && len(s.processed.Withdrawals) == 0 {
		return nil
	}

	var rows [][]interface{}

	for _, e := range s.processed.Deposits {
		rows = append(rows, s.txHistory(e.User.String(), e.AmountDecimal(0), smartyield.JuniorTranche, smartyield.JuniorStake, e.Raw))
	}

	for _, e := range s.processed.Withdrawals {
		rows = append(rows, s.txHistory(e.User.String(), e.AmountDecimal(0), smartyield.JuniorTranche, smartyield.JuniorUnstake, e.Raw))
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_yield", "transaction_history"},
		[]string{
			"protocol_id",
			"pool_address",
			"underlying_token_address",
			"user_address",
			"amount",
			"tranche",
			"transaction_type",
			"block_timestamp",
			"included_in_block",
			"tx_hash",
			"tx_index",
			"log_index",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into transaction_history")
	}

	return nil
}

func (s *Storable) saveClaimEvents(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.Claims) == 0 && len(s.processed.ClaimsMulti) == 0 {
		return nil
	}

	var rows [][]interface{}

	for _, e := range s.processed.Claims {
		rp := s.state.SmartYield.RewardPoolByAddress(e.Raw.Address.String())

		rows = append(rows, []interface{}{
			utils.NormalizeAddress(e.User.String()),
			e.AmountDecimal(0),
			rp.PoolAddress,
			rp.RewardTokenAddresses[0],
			s.block.BlockCreationTime,
			s.block.Number,
			utils.NormalizeAddress(e.Raw.TxHash.String()),
			e.Raw.TxIndex,
			e.Raw.Index,
		})
	}

	for _, e := range s.processed.ClaimsMulti {
		rp := s.state.SmartYield.RewardPoolByAddress(e.Raw.Address.String())

		rows = append(rows, []interface{}{
			utils.NormalizeAddress(e.User.String()),
			e.AmountDecimal(0),
			rp.PoolAddress,
			utils.NormalizeAddress(e.Token.String()),
			s.block.BlockCreationTime,
			s.block.Number,
			utils.NormalizeAddress(e.Raw.TxHash.String()),
			e.Raw.TxIndex,
			e.Raw.Index,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_yield", "rewards_claims"},
		[]string{
			"user_address",
			"amount",
			"pool_address",
			"reward_token_address",
			"block_timestamp",
			"included_in_block",
			"tx_hash",
			"tx_index",
			"log_index",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into rewards_claims")
	}

	return nil
}

func (s *Storable) saveRewardPools(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.Pools) == 0 {
		return nil
	}

	for _, p := range s.processed.Pools {
		_, err := tx.Exec(ctx, `
		insert into smart_yield.reward_pools 
		(pool_type, pool_address, pool_token_address, reward_token_addresses, start_at_block)
		values ($1, $2, $3, $4, $5)
		on conflict do nothing
		`, p.PoolType, p.PoolAddress, p.PoolTokenAddress, p.RewardTokenAddresses, p.StartAtBlock)
		if err != nil {
			return errors.Wrap(err, "could not insert pool into db")
		}
	}

	return nil
}
