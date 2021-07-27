package events

import (
	"github.com/shopspring/decimal"
)

type JuniorEntryEvent struct {
	SYAddress              string
	ProtocolId             string
	UnderlyingTokenAddress string
	BuyerAddress           string
	UnderlyingIn           decimal.Decimal
	TokensOut              decimal.Decimal
	Fee                    decimal.Decimal
}
