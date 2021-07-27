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

const SmartYieldPoolFactorySingleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poolToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardSource\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardRatePerSecond\",\"type\":\"uint256\"}],\"name\":\"deployPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfPools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"contractYieldFarmContinuous\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var SmartYieldPoolFactorySingle = NewSmartYieldPoolFactorySingleDecoder()

type SmartYieldPoolFactorySingleDecoder struct {
	*ethgen.Decoder
}

func NewSmartYieldPoolFactorySingleDecoder() *SmartYieldPoolFactorySingleDecoder {
	dec := ethgen.NewDecoder(SmartYieldPoolFactorySingleABI)
	return &SmartYieldPoolFactorySingleDecoder{
		dec,
	}
}

type SmartYieldPoolFactorySinglePoolCreatedEvent struct {
	Pool common.Address
	Raw  types.Log
}

func (d *SmartYieldPoolFactorySingleDecoder) PoolCreatedEventID() common.Hash {
	return common.HexToHash("0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc")
}

func (d *SmartYieldPoolFactorySingleDecoder) IsPoolCreatedEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.PoolCreatedEventID()
}

func (d *SmartYieldPoolFactorySingleDecoder) IsPoolCreatedEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.PoolCreatedEventID().String()
}

func (d *SmartYieldPoolFactorySingleDecoder) PoolCreatedEventW3(w3l web3types.Log) (SmartYieldPoolFactorySinglePoolCreatedEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartYieldPoolFactorySinglePoolCreatedEvent{}, err
	}

	return d.PoolCreatedEvent(l)
}

func (d *SmartYieldPoolFactorySingleDecoder) PoolCreatedEvent(l types.Log) (SmartYieldPoolFactorySinglePoolCreatedEvent, error) {
	var out SmartYieldPoolFactorySinglePoolCreatedEvent
	if !d.IsPoolCreatedEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "PoolCreated", l)
	out.Raw = l
	return out, err
}

type SmartYieldPoolFactorySingleOwnershipTransferredEvent struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log
}

func (d *SmartYieldPoolFactorySingleDecoder) OwnershipTransferredEventID() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (d *SmartYieldPoolFactorySingleDecoder) IsOwnershipTransferredEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.OwnershipTransferredEventID()
}

func (d *SmartYieldPoolFactorySingleDecoder) IsOwnershipTransferredEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.OwnershipTransferredEventID().String()
}

func (d *SmartYieldPoolFactorySingleDecoder) OwnershipTransferredEventW3(w3l web3types.Log) (SmartYieldPoolFactorySingleOwnershipTransferredEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartYieldPoolFactorySingleOwnershipTransferredEvent{}, err
	}

	return d.OwnershipTransferredEvent(l)
}

func (d *SmartYieldPoolFactorySingleDecoder) OwnershipTransferredEvent(l types.Log) (SmartYieldPoolFactorySingleOwnershipTransferredEvent, error) {
	var out SmartYieldPoolFactorySingleOwnershipTransferredEvent
	if !d.IsOwnershipTransferredEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "OwnershipTransferred", l)
	out.Raw = l
	return out, err
}
