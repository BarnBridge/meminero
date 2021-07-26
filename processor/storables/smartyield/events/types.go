package events

import (
	"math/big"
)

type JuniorEntryEvent struct {
	SYAddress              string
	ProtocolId             string
	UnderlyingTokenAddress string
	BuyerAddress           string
	UnderlyingIn           *big.Int
	TokensOut              *big.Int
	Fee                    *big.Int
}
