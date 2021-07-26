package smartexposure

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type SEPool struct {
	EPoolAddress string
	ProtocolId   string

	ATokenAddress  string
	ATokenSymbol   string
	ATokenDecimals int64

	BTokenAddress  string
	BTokenSymbol   string
	BTokenDecimals int64

	StartAtBlock int64
}

type SETranche struct {
	EPoolAddress  string
	ETokenAddress string
	ETokenSymbol  string
	SFactorE      *big.Int
	TargetRatio   *big.Int
	TokenARatio   decimal.Decimal
	TokenBRatio   decimal.Decimal
}

type TrancheFromChain struct {
	Etoken      common.Address `json:"eToken"`
	SFactorE    *big.Int       `json:"sFactorE"`
	ReserveA    *big.Int       `json:"reserveA"`
	ReserveB    *big.Int       `json:"reserveB"`
	TargetRatio *big.Int       `json:"targetRatio"`
}
