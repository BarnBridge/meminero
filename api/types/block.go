package types

import (
	"github.com/barnbridge/smartbackend/types"
)

type Block struct {
	Number               int64                    `json:"number"`
	BlockHash            string                   `json:"blockHash"`
	ParentBlockHash      string                   `json:"parentBlockHash"`
	BlockCreationTime    types.DatetimeToJSONUnix `json:"blockCreationTime"`
	BlockGasLimit        string                   `json:"blockGasLimit"`
	BlockGasUsed         string                   `json:"blockGasUsed"`
	BlockDifficulty      string                   `json:"blockDifficulty"`
	TotalBlockDifficulty string                   `json:"totalBlockDifficulty"`
	BlockExtraData       types.ByteArray          `json:"blockExtraData"`
	BlockMixHash         types.ByteArray          `json:"blockMixHash"`
	BlockNonce           types.ByteArray          `json:"blockNonce"`
	BlockSize            int64                    `json:"blockSize"`
	BlockLogsBloom       types.ByteArray          `json:"blockLogsBloom"`
	IncludesUncle        types.JSONStringArray    `json:"includesUncle"`
	HasBeneficiary       types.ByteArray          `json:"hasBeneficiary"`
	HasReceiptsTrie      types.ByteArray          `json:"hasReceiptsTrie"`
	HasTxTrie            types.ByteArray          `json:"hasTxTrie"`
	Sha3Uncles           types.ByteArray          `json:"sha3Uncles"`
	NumberOfUncles       int32                    `json:"numberOfUncles"`
	NumberOfTxs          int32                    `json:"numberOfTxs"`

	Txs []Tx `json:"txs"`
}
