// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethtypes

import (
	"math/big"

	web3types "github.com/alethio/web3-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lacasian/ethwheels/ethgen"
	"github.com/shopspring/decimal"
)

// Reference imports to suppress errors
var (
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = web3types.Log{}
	_ = decimal.NewFromBigInt
)

const ETokenFactoryABI = "[{\"inputs\":[{\"internalType\":\"contractIController\",\"name\":\"_controller\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ePool\",\"type\":\"address\"}],\"name\":\"CreatedEToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"contractIController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"createEToken\",\"outputs\":[{\"internalType\":\"contractIEToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var ETokenFactory = NewETokenFactoryDecoder()

type ETokenFactoryDecoder struct {
	*ethgen.Decoder
}

func NewETokenFactoryDecoder() *ETokenFactoryDecoder {
	dec := ethgen.NewDecoder(ETokenFactoryABI)
	return &ETokenFactoryDecoder{
		dec,
	}
}

type ETokenFactoryCreatedETokenEvent struct {
	EToken common.Address
	EPool  common.Address
	Raw    types.Log
}

func (d *ETokenFactoryDecoder) CreatedETokenEventID() common.Hash {
	return common.HexToHash("0x98d5a18d1ecc5924c3270fc708d83f1413759b0f8bd8e9b9353e5434747b271d")
}

func (d *ETokenFactoryDecoder) IsCreatedETokenEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.CreatedETokenEventID()
}

func (d *ETokenFactoryDecoder) IsCreatedETokenEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.CreatedETokenEventID().String()
}

func (d *ETokenFactoryDecoder) CreatedETokenEventW3(w3l web3types.Log) (ETokenFactoryCreatedETokenEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return ETokenFactoryCreatedETokenEvent{}, err
	}

	return d.CreatedETokenEvent(l)
}

func (d *ETokenFactoryDecoder) CreatedETokenEvent(l types.Log) (ETokenFactoryCreatedETokenEvent, error) {
	var out ETokenFactoryCreatedETokenEvent
	if !d.IsCreatedETokenEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "CreatedEToken", l)
	out.Raw = l
	return out, err
}
