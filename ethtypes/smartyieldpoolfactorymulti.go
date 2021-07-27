// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethtypes

import (
	"math/big"

	web3types "github.com/alethio/web3-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lacasian/ethwheels/ethgen"
)

// Reference imports to suppress errors
var (
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = web3types.Log{}
)

const SmartYieldPoolFactoryMultiABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolMultiCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poolToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardSource\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"}],\"internalType\":\"structPoolFactoryMulti.RewardToken[]\",\"name\":\"rewardTokens\",\"type\":\"tuple[]\"}],\"name\":\"deployPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfPools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"contractPoolMulti\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var SmartYieldPoolFactoryMulti = NewSmartYieldPoolFactoryMultiDecoder()

type SmartYieldPoolFactoryMultiDecoder struct {
	*ethgen.Decoder
}

func NewSmartYieldPoolFactoryMultiDecoder() *SmartYieldPoolFactoryMultiDecoder {
	dec := ethgen.NewDecoder(SmartYieldPoolFactoryMultiABI)
	return &SmartYieldPoolFactoryMultiDecoder{
		dec,
	}
}

type SmartYieldPoolFactoryMultiPoolMultiCreatedEvent struct {
	Pool common.Address
	Raw  types.Log
}

func (d *SmartYieldPoolFactoryMultiDecoder) PoolMultiCreatedEventID() common.Hash {
	return common.HexToHash("0x4cdb7db74a6c2bc668aa4de2e5c57339c91a49d4770ce81dab3e64e328763bfb")
}

func (d *SmartYieldPoolFactoryMultiDecoder) IsPoolMultiCreatedEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.PoolMultiCreatedEventID()
}

func (d *SmartYieldPoolFactoryMultiDecoder) IsPoolMultiCreatedEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.PoolMultiCreatedEventID().String()
}

func (d *SmartYieldPoolFactoryMultiDecoder) PoolMultiCreatedEventW3(w3l web3types.Log) (SmartYieldPoolFactoryMultiPoolMultiCreatedEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartYieldPoolFactoryMultiPoolMultiCreatedEvent{}, err
	}

	return d.PoolMultiCreatedEvent(l)
}

func (d *SmartYieldPoolFactoryMultiDecoder) PoolMultiCreatedEvent(l types.Log) (SmartYieldPoolFactoryMultiPoolMultiCreatedEvent, error) {
	var out SmartYieldPoolFactoryMultiPoolMultiCreatedEvent
	if !d.IsPoolMultiCreatedEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "PoolMultiCreated", l)
	out.Raw = l
	return out, err
}

type SmartYieldPoolFactoryMultiOwnershipTransferredEvent struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log
}

func (d *SmartYieldPoolFactoryMultiDecoder) OwnershipTransferredEventID() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (d *SmartYieldPoolFactoryMultiDecoder) IsOwnershipTransferredEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.OwnershipTransferredEventID()
}

func (d *SmartYieldPoolFactoryMultiDecoder) IsOwnershipTransferredEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.OwnershipTransferredEventID().String()
}

func (d *SmartYieldPoolFactoryMultiDecoder) OwnershipTransferredEventW3(w3l web3types.Log) (SmartYieldPoolFactoryMultiOwnershipTransferredEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartYieldPoolFactoryMultiOwnershipTransferredEvent{}, err
	}

	return d.OwnershipTransferredEvent(l)
}

func (d *SmartYieldPoolFactoryMultiDecoder) OwnershipTransferredEvent(l types.Log) (SmartYieldPoolFactoryMultiOwnershipTransferredEvent, error) {
	var out SmartYieldPoolFactoryMultiOwnershipTransferredEvent
	if !d.IsOwnershipTransferredEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "OwnershipTransferred", l)
	out.Raw = l
	return out, err
}
