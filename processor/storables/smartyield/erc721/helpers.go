package erc721

import (
	"math/big"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"

	"github.com/barnbridge/meminero/processor/storables/smartyield"
	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) txHistory(user string, amount *big.Int, tranche string, txType smartyield.TxType, raw gethtypes.Log) []interface{} {
	p := s.state.SmartYield.PoolBySeniorBondAddress(raw.Address.String())
	if p == nil {
		p = s.state.SmartYield.PoolByJuniorBondAddress(raw.Address.String())
	}

	return []interface{}{
		p.ProtocolId,
		p.PoolAddress,
		p.UnderlyingAddress,
		utils.NormalizeAddress(user),
		decimal.NewFromBigInt(amount, 0),
		tranche,
		txType,
		s.block.BlockCreationTime,
		s.block.Number,
		utils.NormalizeAddress(raw.TxHash.String()),
		raw.TxIndex,
		raw.Index,
	}
}
