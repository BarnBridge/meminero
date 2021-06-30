package types

import (
	"math/big"
)

type ERC20Transfer struct {
	TokenAddress     string
	From             string
	To               string
	Value            *big.Int
	TransactionHash  string
	TransactionIndex int64
	LogIndex         int64
}

