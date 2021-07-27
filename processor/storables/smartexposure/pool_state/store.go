package pool_state

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func (s *Storable) storePoolsState(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.poolStates) == 0 {

		return nil
	}
	var rows [][]interface{}
	for _, p := range s.processed.poolStates {
		liq, _ := p.PoolLiquidity.Float64()
		rows = append(rows, []interface{}{
			p.PoolAddress,
			liq,
			p.LastRebalance.BigInt().Int64(),
			p.RebalancingInterval.BigInt().Int64(),
			p.RebalancingCondition,
			s.block.BlockCreationTime,
			s.block.Number,
		})
	}
	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_exposure", "pool_state"},
		[]string{"pool_address", "pool_liquidity", "last_rebalance", "rebalancing_interval", "rebalancing_condition", "block_timestamp", "included_in_block"},
		pgx.CopyFromRows(rows),
	)
	return err
}
