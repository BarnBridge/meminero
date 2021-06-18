package types

import (
	"github.com/barnbridge/smartbackend/types"
)

type Uncle struct {
	BlockHash         string                   `json:"blockHash"`
	IncludedInBlock   int64                    `json:"includedInBlock"`
	Number            int64                    `json:"number"`
	BlockCreationTime types.DatetimeToJSONUnix `json:"blockCreationTime"`
	UncleIndex        int32                    `json:"uncleIndex"`
	BlockGasLimit     string                   `json:"blockGasLimit"`
	BlockGasUsed      string                   `json:"blockGasUsed"`
	HasBeneficiary    types.ByteArray          `json:"hasBeneficiary"`
	BlockDifficulty   string                   `json:"blockDifficulty"`
	BlockExtraData    types.ByteArray          `json:"blockExtraData"`
	BlockMixHash      types.ByteArray          `json:"blockMixHash"`
	BlockNonce        types.ByteArray          `json:"blockNonce"`
	Sha3Uncles        types.ByteArray          `json:"sha3Uncles"`
}
