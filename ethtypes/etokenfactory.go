// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethtypes

import (
	web3types "github.com/alethio/web3-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lacasian/ethwheels/ethgen"
)

const EtokenfactoryABI = "[{\"inputs\":[{\"internalType\":\"contractIController\",\"name\":\"_controller\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ePool\",\"type\":\"address\"}],\"name\":\"CreatedEToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"contractIController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"createEToken\",\"outputs\":[{\"internalType\":\"contractIEToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var Etokenfactory = NewEtokenfactoryDecoder()

type EtokenfactoryDecoder struct {
	*ethgen.Decoder
}

func NewEtokenfactoryDecoder() *EtokenfactoryDecoder {
	dec := ethgen.NewDecoder(EtokenfactoryABI)
	return &EtokenfactoryDecoder{
		dec,
	}
}

type EtokenfactoryCreatedETokenEvent struct {
	EToken common.Address
	EPool  common.Address
	Raw    types.Log
}

func (d *EtokenfactoryDecoder) EtokenfactoryCreatedETokenEventID() common.Hash {
	return common.HexToHash("0x98d5a18d1ecc5924c3270fc708d83f1413759b0f8bd8e9b9353e5434747b271d")
}

func (d *EtokenfactoryDecoder) IsEtokenfactoryCreatedETokenEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EtokenfactoryCreatedETokenEventID()
}

func (d *EtokenfactoryDecoder) EtokenfactoryCreatedETokenEventW3(w3l web3types.Log) (EtokenfactoryCreatedETokenEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EtokenfactoryCreatedETokenEvent{}, err
	}

	return d.EtokenfactoryCreatedETokenEvent(l)
}

func (d *EtokenfactoryDecoder) EtokenfactoryCreatedETokenEvent(l types.Log) (EtokenfactoryCreatedETokenEvent, error) {
	var out EtokenfactoryCreatedETokenEvent
	if !d.IsEtokenfactoryCreatedETokenEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "CreatedEToken", l)
	out.Raw = l
	return out, err
}
