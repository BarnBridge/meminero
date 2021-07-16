package types

type Tx struct {
	TxHash            string
	IncludedInBlock   int64
	TxIndex           int64
	From              string
	To                string
	Value             string
	TxNonce           int64
	MsgGasLimit       string
	TxGasUsed         string
	TxGasPrice        string
	CumulativeGasUsed string
	MsgPayload        ByteArray
	MsgStatus         string
	Creates           string
	TxLogsBloom       ByteArray
	BlockCreationTime int64

	LogEntries LogEntries
}

type Txs []Tx

func (t Txs) Len() int           { return len(t) }
func (t Txs) Less(i, j int) bool { return t[i].TxIndex < t[j].TxIndex }
func (t Txs) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
