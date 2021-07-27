package pool_state

import (
	"github.com/shopspring/decimal"
)

type PoolState struct {
	PoolAddress string

	PoolLiquidity decimal.Decimal

	LastRebalance        decimal.Decimal
	RebalancingInterval  decimal.Decimal
	RebalancingCondition decimal.Decimal
}
