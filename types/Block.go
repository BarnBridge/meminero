package types

type Block struct {
	Number               int64
	BlockHash            string
	ParentBlockHash      string
	BlockCreationTime    DatetimeToJSONUnix
	BlockGasLimit        string
	BlockGasUsed         string
	BlockDifficulty      string
	TotalBlockDifficulty string
	BlockExtraData       ByteArray
	BlockMixHash         ByteArray
	BlockNonce           ByteArray
	BlockSize            int64
	BlockLogsBloom       ByteArray
	IncludesUncle        JSONStringArray
	HasBeneficiary       ByteArray
	HasReceiptsTrie      ByteArray
	HasTxTrie            ByteArray
	Sha3Uncles           ByteArray
	NumberOfUncles       int32

	Txs Txs
}
