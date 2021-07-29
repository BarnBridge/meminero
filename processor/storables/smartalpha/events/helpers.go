package events

import (
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"

	"github.com/barnbridge/meminero/processor/storables/smartalpha"
	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) generateEventData(raw gethtypes.Log) []interface{} {
	return []interface{}{
		s.block.BlockCreationTime,
		s.block.Number,
		utils.NormalizeAddress(raw.TxHash.String()),
		raw.TxIndex,
		raw.Index,
	}
}

func (s *Storable) txHistory(poolAddress string, tranche string, transactionType smartalpha.TxType, userAddress string, amount decimal.Decimal, raw gethtypes.Log) []interface{} {
	row := []interface{}{
		poolAddress,
		tranche,
		transactionType,
		utils.NormalizeAddress(userAddress),
		amount,
	}

	return append(row, s.generateEventData(raw)...)
}
