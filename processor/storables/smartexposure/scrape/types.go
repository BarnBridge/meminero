package scrape

import (
	"github.com/shopspring/decimal"
)

type SETransaction struct {
	ETokenAddress   string
	UserAddress     string
	Amount          decimal.Decimal
	AmountA         decimal.Decimal
	AmountB         decimal.Decimal
	TransactionType string

	TxHash   string
	TxIndex  int64
	LogIndex int64
}
