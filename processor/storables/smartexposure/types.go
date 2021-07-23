package smartexposure

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TrancheFromChain struct {
	Etoken      common.Address `json:"eToken"`
	SFactorE    *big.Int       `json:"sFactorE"`
	ReserveA    *big.Int       `json:"reserveA"`
	ReserveB    *big.Int       `json:"reserveB"`
	TargetRatio *big.Int       `json:"targetRatio"`
}
