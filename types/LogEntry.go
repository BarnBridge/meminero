package types

import (
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

type LogEntry struct {
	TxHash          string
	LogIndex        int32
	LogData         ByteArray
	LoggedBy        string
	Topic0          string
	Topic1          string
	Topic2          string
	Topic3          string
	IncludedInBlock int64
}

type LogEntries []gethtypes.Log

func (l LogEntries) Len() int           { return len(l) }
func (l LogEntries) Less(i, j int) bool { return l[i].Index < l[j].Index }
func (l LogEntries) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
