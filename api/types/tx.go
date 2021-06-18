package types

import (
	"github.com/barnbridge/smartbackend/types"
)

type Tx struct {
	TxHash              *string                   `json:"txHash,omitempty"`
	IncludedInBlock     *int64                    `json:"includedInBlock,omitempty"`
	TxIndex             *int32                    `json:"txIndex,omitempty"`
	From                *types.ByteArray          `json:"from,omitempty"`
	To                  *types.ByteArray          `json:"to,omitempty"`
	Value               *string                   `json:"value,omitempty"`
	TxNonce             *int64                    `json:"txNonce,omitempty"`
	MsgGasLimit         *string                   `json:"msgGasLimit,omitempty"`
	TxGasUsed           *string                   `json:"txGasUsed,omitempty"`
	TxGasPrice          *string                   `json:"txGasPrice,omitempty"`
	CumulativeGasUsed   *string                   `json:"cumulativeGasUsed,omitempty"`
	MsgPayload          *types.ByteArray          `json:"msgPayload,omitempty"`
	MsgStatus           *string                   `json:"msgStatus,omitempty"`
	MsgError            *bool                     `json:"msgError,omitempty"`
	MsgErrorString      *string                   `json:"msgErrorString,omitempty"`
	Creates             *types.ByteArray          `json:"creates,omitempty"`
	TxLogsBloom         *types.ByteArray          `json:"txLogsBloom,omitempty"`
	BlockCreationTime   *types.DatetimeToJSONUnix `json:"blockCreationTime,omitempty"`
	LogEntriesTriggered *int32                    `json:"logEntriesTriggered,omitempty"`
}
