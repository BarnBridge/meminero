package yieldfarming

import (
	"github.com/shopspring/decimal"
)

type StakingAction struct {
	UserAddress      string
	TokenAddress     string
	Amount           decimal.Decimal
	TransactionHash  string
	TransactionIndex int64
	LogIndex         int64
	ActionType       ActionType
}

type ActionType string

const (
	DEPOSIT  ActionType = "DEPOSIT"
	WITHDRAW ActionType = "WITHDRAW"
)
