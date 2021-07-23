package yieldfarming

import (
	"math/big"
)

type StakingAction struct {
	UserAddress      string
	TokenAddress     string
	Amount           *big.Int
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
