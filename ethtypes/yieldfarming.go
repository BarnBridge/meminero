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

const YieldFarmingABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch1Start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epochDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"ManualEpochInit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prevBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"prevMultiplier\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"currentMultiplier\",\"type\":\"uint128\"}],\"name\":\"computeNewMultiplier\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpochMultiplier\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch1Start\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getEpochPoolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"getEpochUserBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint128\",\"name\":\"epochId\",\"type\":\"uint128\"}],\"name\":\"manualEpochInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var YieldFarming = NewYieldFarmingDecoder()

type YieldFarmingDecoder struct {
	*ethgen.Decoder
}

func NewYieldFarmingDecoder() *YieldFarmingDecoder {
	dec := ethgen.NewDecoder(YieldFarmingABI)
	return &YieldFarmingDecoder{
		dec,
	}
}

type YieldFarmingDepositEvent struct {
	User         common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log
}

func (d *YieldFarmingDecoder) YieldFarmingDepositEventID() common.Hash {
	return common.HexToHash("0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62")
}

func (d *YieldFarmingDecoder) IsYieldFarmingDepositEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.YieldFarmingDepositEventID()
}

func (d *YieldFarmingDecoder) IsYieldFarmingDepositEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.YieldFarmingDepositEventID().String()
}

func (d *YieldFarmingDecoder) YieldFarmingDepositEventW3(w3l web3types.Log) (YieldFarmingDepositEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return YieldFarmingDepositEvent{}, err
	}

	return d.YieldFarmingDepositEvent(l)
}

func (d *YieldFarmingDecoder) YieldFarmingDepositEvent(l types.Log) (YieldFarmingDepositEvent, error) {
	var out YieldFarmingDepositEvent
	if !d.IsYieldFarmingDepositEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "Deposit", l)
	out.Raw = l
	return out, err
}

type YieldFarmingWithdrawEvent struct {
	User         common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log
}

func (d *YieldFarmingDecoder) YieldFarmingWithdrawEventID() common.Hash {
	return common.HexToHash("0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb")
}

func (d *YieldFarmingDecoder) IsYieldFarmingWithdrawEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.YieldFarmingWithdrawEventID()
}

func (d *YieldFarmingDecoder) IsYieldFarmingWithdrawEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.YieldFarmingWithdrawEventID().String()
}

func (d *YieldFarmingDecoder) YieldFarmingWithdrawEventW3(w3l web3types.Log) (YieldFarmingWithdrawEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return YieldFarmingWithdrawEvent{}, err
	}

	return d.YieldFarmingWithdrawEvent(l)
}

func (d *YieldFarmingDecoder) YieldFarmingWithdrawEvent(l types.Log) (YieldFarmingWithdrawEvent, error) {
	var out YieldFarmingWithdrawEvent
	if !d.IsYieldFarmingWithdrawEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "Withdraw", l)
	out.Raw = l
	return out, err
}

type YieldFarmingManualEpochInitEvent struct {
	Caller  common.Address
	EpochId *big.Int
	Tokens  []common.Address
	Raw     types.Log
}

func (d *YieldFarmingDecoder) YieldFarmingManualEpochInitEventID() common.Hash {
	return common.HexToHash("0xb85c32b8d9cecc81feba78646289584a693e9a8afea40ab2fd31efae4408429f")
}

func (d *YieldFarmingDecoder) IsYieldFarmingManualEpochInitEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.YieldFarmingManualEpochInitEventID()
}

func (d *YieldFarmingDecoder) IsYieldFarmingManualEpochInitEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.YieldFarmingManualEpochInitEventID().String()
}

func (d *YieldFarmingDecoder) YieldFarmingManualEpochInitEventW3(w3l web3types.Log) (YieldFarmingManualEpochInitEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return YieldFarmingManualEpochInitEvent{}, err
	}

	return d.YieldFarmingManualEpochInitEvent(l)
}

func (d *YieldFarmingDecoder) YieldFarmingManualEpochInitEvent(l types.Log) (YieldFarmingManualEpochInitEvent, error) {
	var out YieldFarmingManualEpochInitEvent
	if !d.IsYieldFarmingManualEpochInitEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "ManualEpochInit", l)
	out.Raw = l
	return out, err
}

type YieldFarmingEmergencyWithdrawEvent struct {
	User         common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log
}

func (d *YieldFarmingDecoder) YieldFarmingEmergencyWithdrawEventID() common.Hash {
	return common.HexToHash("0xf24ef89f38eadc1bde50701ad6e4d6d11a2dc24f7cf834a486991f3883328504")
}

func (d *YieldFarmingDecoder) IsYieldFarmingEmergencyWithdrawEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.YieldFarmingEmergencyWithdrawEventID()
}

func (d *YieldFarmingDecoder) IsYieldFarmingEmergencyWithdrawEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.YieldFarmingEmergencyWithdrawEventID().String()
}

func (d *YieldFarmingDecoder) YieldFarmingEmergencyWithdrawEventW3(w3l web3types.Log) (YieldFarmingEmergencyWithdrawEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return YieldFarmingEmergencyWithdrawEvent{}, err
	}

	return d.YieldFarmingEmergencyWithdrawEvent(l)
}

func (d *YieldFarmingDecoder) YieldFarmingEmergencyWithdrawEvent(l types.Log) (YieldFarmingEmergencyWithdrawEvent, error) {
	var out YieldFarmingEmergencyWithdrawEvent
	if !d.IsYieldFarmingEmergencyWithdrawEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "EmergencyWithdraw", l)
	out.Raw = l
	return out, err
}
