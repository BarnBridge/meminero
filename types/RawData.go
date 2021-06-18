package types

import (
	web3types "github.com/alethio/web3-go/types"
)

type RawData struct {
	Block    web3types.Block
	Receipts RawReceipts
	Uncles   []web3types.Block
}
