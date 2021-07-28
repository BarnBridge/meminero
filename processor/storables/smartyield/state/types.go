package state

import (
	"math/big"

	"github.com/shopspring/decimal"
)

type State struct {
	PoolAddress string

	TotalLiquidity  decimal.Decimal
	JuniorLiquidity decimal.Decimal
	JTokenPrice     decimal.Decimal

	AbondAPY         float64
	SeniorAPY        float64
	JuniorAPY        float64
	OriginatorApy    float64
	OriginatorNetApy float64

	Abond Abond
}

type Abond struct {
	Principal  *big.Int
	Gain       *big.Int
	MaturesAt  *big.Int
	IssuedAt   *big.Int
	Liquidated bool
}
