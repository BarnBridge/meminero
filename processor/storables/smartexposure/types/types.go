package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	globalTypes "github.com/barnbridge/meminero/types"
)

type Pool struct {
	EPoolAddress string
	ProtocolId   string

	TokenA globalTypes.Token
	TokenB globalTypes.Token

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
	Etoken      common.Address `json:"eToken"`
	SFactorE    *big.Int       `json:"sFactorE"`
	ReserveA    *big.Int       `json:"reserveA"`
	ReserveB    *big.Int       `json:"reserveB"`
	TargetRatio *big.Int       `json:"targetRatio"`
}
