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

const SmartYieldPoolSingleABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poolToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceAfter\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceAfter\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ackFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceBefore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSoftPullTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pullRewardFromSource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardNotTransferred\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRatePerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardSource\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"setRewardRatePerSecond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"}],\"name\":\"setRewardsSource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"softPullReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAndClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var SmartYieldPoolSingle = NewSmartYieldPoolSingleDecoder()

type SmartYieldPoolSingleDecoder struct {
	*ethgen.Decoder
}

func NewSmartYieldPoolSingleDecoder() *SmartYieldPoolSingleDecoder {
	dec := ethgen.NewDecoder(SmartYieldPoolSingleABI)
	return &SmartYieldPoolSingleDecoder{
		dec,
	}
}

type SmartYieldPoolSingleClaimEvent struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log
}

func (d *SmartYieldPoolSingleDecoder) ClaimEventID() common.Hash {
	return common.HexToHash("0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4")
}

func (d *SmartYieldPoolSingleDecoder) IsClaimEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.ClaimEventID()
}

func (d *SmartYieldPoolSingleDecoder) IsClaimEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.ClaimEventID().String()
}

func (d *SmartYieldPoolSingleDecoder) ClaimEventW3(w3l web3types.Log) (SmartYieldPoolSingleClaimEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartYieldPoolSingleClaimEvent{}, err
	}

	return d.ClaimEvent(l)
}

func (d *SmartYieldPoolSingleDecoder) ClaimEvent(l types.Log) (SmartYieldPoolSingleClaimEvent, error) {
	var out SmartYieldPoolSingleClaimEvent
	if !d.IsClaimEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "Claim", l)
	out.Raw = l
	return out, err
}

type SmartYieldPoolSingleOwnershipTransferredEvent struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log
}

func (d *SmartYieldPoolSingleDecoder) OwnershipTransferredEventID() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (d *SmartYieldPoolSingleDecoder) IsOwnershipTransferredEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.OwnershipTransferredEventID()
}

func (d *SmartYieldPoolSingleDecoder) IsOwnershipTransferredEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.OwnershipTransferredEventID().String()
}

func (d *SmartYieldPoolSingleDecoder) OwnershipTransferredEventW3(w3l web3types.Log) (SmartYieldPoolSingleOwnershipTransferredEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartYieldPoolSingleOwnershipTransferredEvent{}, err
	}

	return d.OwnershipTransferredEvent(l)
}

func (d *SmartYieldPoolSingleDecoder) OwnershipTransferredEvent(l types.Log) (SmartYieldPoolSingleOwnershipTransferredEvent, error) {
	var out SmartYieldPoolSingleOwnershipTransferredEvent
	if !d.IsOwnershipTransferredEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "OwnershipTransferred", l)
	out.Raw = l
	return out, err
}

type SmartYieldPoolSingleDepositEvent struct {
	User         common.Address
	Amount       *big.Int
	BalanceAfter *big.Int
	Raw          types.Log
}

func (d *SmartYieldPoolSingleDecoder) DepositEventID() common.Hash {
	return common.HexToHash("0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15")
}

func (d *SmartYieldPoolSingleDecoder) IsDepositEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.DepositEventID()
}

func (d *SmartYieldPoolSingleDecoder) IsDepositEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.DepositEventID().String()
}

func (d *SmartYieldPoolSingleDecoder) DepositEventW3(w3l web3types.Log) (SmartYieldPoolSingleDepositEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartYieldPoolSingleDepositEvent{}, err
	}

	return d.DepositEvent(l)
}

func (d *SmartYieldPoolSingleDecoder) DepositEvent(l types.Log) (SmartYieldPoolSingleDepositEvent, error) {
	var out SmartYieldPoolSingleDepositEvent
	if !d.IsDepositEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "Deposit", l)
	out.Raw = l
	return out, err
}

type SmartYieldPoolSingleWithdrawEvent struct {
	User         common.Address
	Amount       *big.Int
	BalanceAfter *big.Int
	Raw          types.Log
}

func (d *SmartYieldPoolSingleDecoder) WithdrawEventID() common.Hash {
	return common.HexToHash("0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568")
}

func (d *SmartYieldPoolSingleDecoder) IsWithdrawEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.WithdrawEventID()
}

func (d *SmartYieldPoolSingleDecoder) IsWithdrawEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.WithdrawEventID().String()
}

func (d *SmartYieldPoolSingleDecoder) WithdrawEventW3(w3l web3types.Log) (SmartYieldPoolSingleWithdrawEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartYieldPoolSingleWithdrawEvent{}, err
	}

	return d.WithdrawEvent(l)
}

func (d *SmartYieldPoolSingleDecoder) WithdrawEvent(l types.Log) (SmartYieldPoolSingleWithdrawEvent, error) {
	var out SmartYieldPoolSingleWithdrawEvent
	if !d.IsWithdrawEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "Withdraw", l)
	out.Raw = l
	return out, err
}
