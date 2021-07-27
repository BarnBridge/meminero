package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type Pool struct {
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

type Tranche struct {
	EPoolAddress  string
	ETokenAddress string
	ETokenSymbol  string
	SFactorE      decimal.Decimal
	TargetRatio   decimal.Decimal
	TokenARatio   decimal.Decimal
	TokenBRatio   decimal.Decimal
}

type TrancheFromChain struct {
	Etoken      common.Address  `json:"eToken"`
	SFactorE    decimal.Decimal `json:"sFactorE"`
	ReserveA    decimal.Decimal `json:"reserveA"`
	ReserveB    decimal.Decimal `json:"reserveB"`
	TargetRatio decimal.Decimal `json:"targetRatio"`
}
