package types

import (
	"github.com/barnbridge/smartbackend/types"
)

type LogEntry struct {
	TxHash            string                 `json:"txHash"`
	LogIndex          int32                  `json:"logIndex"`
	LogData           types.ByteArray        `json:"logData"`
	LoggedBy          string                 `json:"loggedBy"`
	HasLogTopics      []string               `json:"hasLogTopics"`
	EventDecoded      map[string]interface{} `json:"eventDecoded"`
	EventDecodedError string                 `json:"eventDecodedError"`
}
