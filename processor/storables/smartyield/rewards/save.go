package rewards

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.Debug("storing")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Debug("done storing")
	}()

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
