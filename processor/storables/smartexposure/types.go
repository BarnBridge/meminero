package smartexposure

import (
	"math/big"
)

type SETransaction struct {
	ETokenAddress   string
	UserAddress     string
	Amount          *big.Int
	AmountA         *big.Int
	AmountB         *big.Int
	TransactionType string

	TxHash   string
	TxIndex  int64
	LogIndex int64
}
