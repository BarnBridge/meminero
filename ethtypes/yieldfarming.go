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

func (e *YieldFarmingDepositEvent) AmountDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.Amount, exp)
}

func (d *YieldFarmingDecoder) DepositEventID() common.Hash {
	return common.HexToHash("0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62")
}

func (d *YieldFarmingDecoder) IsDepositEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.DepositEventID()
}

func (d *YieldFarmingDecoder) IsDepositEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.DepositEventID().String()
}

func (d *YieldFarmingDecoder) DepositEventW3(w3l web3types.Log) (YieldFarmingDepositEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return YieldFarmingDepositEvent{}, err
	}

	return d.DepositEvent(l)
}

func (d *YieldFarmingDecoder) DepositEvent(l types.Log) (YieldFarmingDepositEvent, error) {
	var out YieldFarmingDepositEvent
	if !d.IsDepositEvent(&l) {
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

func (e *YieldFarmingWithdrawEvent) AmountDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.Amount, exp)
}

func (d *YieldFarmingDecoder) WithdrawEventID() common.Hash {
	return common.HexToHash("0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb")
}

func (d *YieldFarmingDecoder) IsWithdrawEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.WithdrawEventID()
}

func (d *YieldFarmingDecoder) IsWithdrawEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.WithdrawEventID().String()
}

func (d *YieldFarmingDecoder) WithdrawEventW3(w3l web3types.Log) (YieldFarmingWithdrawEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return YieldFarmingWithdrawEvent{}, err
	}

	return d.WithdrawEvent(l)
}

func (d *YieldFarmingDecoder) WithdrawEvent(l types.Log) (YieldFarmingWithdrawEvent, error) {
	var out YieldFarmingWithdrawEvent
	if !d.IsWithdrawEvent(&l) {
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

func (e *YieldFarmingManualEpochInitEvent) EpochIdDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.EpochId, exp)
}

func (d *YieldFarmingDecoder) ManualEpochInitEventID() common.Hash {
	return common.HexToHash("0xb85c32b8d9cecc81feba78646289584a693e9a8afea40ab2fd31efae4408429f")
}

func (d *YieldFarmingDecoder) IsManualEpochInitEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.ManualEpochInitEventID()
}

func (d *YieldFarmingDecoder) IsManualEpochInitEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.ManualEpochInitEventID().String()
}

func (d *YieldFarmingDecoder) ManualEpochInitEventW3(w3l web3types.Log) (YieldFarmingManualEpochInitEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return YieldFarmingManualEpochInitEvent{}, err
	}

	return d.ManualEpochInitEvent(l)
}

func (d *YieldFarmingDecoder) ManualEpochInitEvent(l types.Log) (YieldFarmingManualEpochInitEvent, error) {
	var out YieldFarmingManualEpochInitEvent
	if !d.IsManualEpochInitEvent(&l) {
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

func (e *YieldFarmingEmergencyWithdrawEvent) AmountDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.Amount, exp)
}

func (d *YieldFarmingDecoder) EmergencyWithdrawEventID() common.Hash {
	return common.HexToHash("0xf24ef89f38eadc1bde50701ad6e4d6d11a2dc24f7cf834a486991f3883328504")
}

func (d *YieldFarmingDecoder) IsEmergencyWithdrawEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EmergencyWithdrawEventID()
}

func (d *YieldFarmingDecoder) IsEmergencyWithdrawEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EmergencyWithdrawEventID().String()
}

func (d *YieldFarmingDecoder) EmergencyWithdrawEventW3(w3l web3types.Log) (YieldFarmingEmergencyWithdrawEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return YieldFarmingEmergencyWithdrawEvent{}, err
	}

	return d.EmergencyWithdrawEvent(l)
}

func (d *YieldFarmingDecoder) EmergencyWithdrawEvent(l types.Log) (YieldFarmingEmergencyWithdrawEvent, error) {
	var out YieldFarmingEmergencyWithdrawEvent
	if !d.IsEmergencyWithdrawEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "EmergencyWithdraw", l)
	out.Raw = l
	return out, err
}
