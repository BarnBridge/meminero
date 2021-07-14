package types

type Block struct {
	Number            int64
	BlockHash         string
	ParentBlockHash   string
	BlockCreationTime DatetimeToJSONUnix

	Txs Txs
}
