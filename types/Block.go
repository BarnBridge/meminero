package types

type Block struct {
	Number            int64
	BlockHash         string
	ParentBlockHash   string
	BlockCreationTime int64

	Txs Txs
}
