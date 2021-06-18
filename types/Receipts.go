package types

import (
	"strconv"

	web3types "github.com/alethio/web3-go/types"
)

type RawReceipts []web3types.Receipt

func (a RawReceipts) Len() int      { return len(a) }
func (a RawReceipts) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a RawReceipts) Less(i, j int) bool {
	iIdx, _ := strconv.ParseInt(a[i].TransactionIndex, 0, 64)
	jIdx, _ := strconv.ParseInt(a[j].TransactionIndex, 0, 64)

	return iIdx < jIdx
}
