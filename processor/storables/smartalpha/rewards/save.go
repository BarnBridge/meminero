package rewards

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	err := s.saveStakingActions(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save staking actions")
	}

	err = s.saveClaimEvents(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save claim events")
	}

	return nil
}

func (s *Storable) saveStakingActions(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.Deposits) == 0 && len(s.processed.Withdrawals) == 0 {
		return nil
	}

	var rows [][]interface{}

	for _, e := range s.processed.Deposits {
		p := s.state.SmartAlpha.RewardPoolByAddress(e.Raw.Address.String())

		rows = append(rows, []interface{}{
			utils.NormalizeAddress(e.User.String()),
			e.AmountDecimal(0),
			e.BalanceAfterDecimal(0),
			"DEPOSIT",
			p.PoolAddress,
			s.block.BlockCreationTime,
			s.block.Number,
			utils.NormalizeAddress(e.Raw.TxHash.String()),
			e.Raw.TxIndex,
			e.Raw.Index,
		})
	}

	for _, e := range s.processed.Withdrawals {
		p := s.state.SmartAlpha.RewardPoolByAddress(e.Raw.Address.String())

		rows = append(rows, []interface{}{
			utils.NormalizeAddress(e.User.String()),
			e.AmountDecimal(0),
			e.BalanceAfterDecimal(0),
			"WITHDRAW",
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
		pgx.Identifier{"smart_alpha", "rewards_staking_actions"},
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

func (s *Storable) saveClaimEvents(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.Claims) == 0 && len(s.processed.ClaimsMulti) == 0 {
		return nil
	}

	var rows [][]interface{}

	for _, e := range s.processed.Claims {
		rp := s.state.SmartAlpha.RewardPoolByAddress(e.Raw.Address.String())

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
		rp := s.state.SmartAlpha.RewardPoolByAddress(e.Raw.Address.String())

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
		pgx.Identifier{"smart_alpha", "rewards_claims"},
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
