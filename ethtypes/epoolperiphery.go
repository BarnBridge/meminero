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

const EpoolperipheryABI = "[{\"inputs\":[{\"internalType\":\"contractIController\",\"name\":\"_controller\",\"type\":\"address\"},{\"internalType\":\"contractIUniswapV2Factory\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"contractIUniswapV2Router01\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"contractIKeeperSubsidyPool\",\"name\":\"_keeperSubsidyPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxFlashSwapSlippage\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ePool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"IssuedEToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoveredToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ePool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"RedeemedEToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"SetController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ePool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approval\",\"type\":\"bool\"}],\"name\":\"SetEPoolApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxFlashSwapSlippage\",\"type\":\"uint256\"}],\"name\":\"SetMaxFlashSwapSlippage\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ePools\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Factory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxInputAmountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"issueForMaxTokenA\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxInputAmountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"issueForMaxTokenB\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keeperSubsidyPool\",\"outputs\":[{\"internalType\":\"contractIKeeperSubsidyPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxFlashSwapSlippage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"maxOutputAmountAForEToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxTokenA\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"maxOutputAmountBForEToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxTokenB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"minInputAmountAForEToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minTokenA\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"minInputAmountBForEToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minTokenB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fracDelta\",\"type\":\"uint256\"}],\"name\":\"rebalanceWithFlashSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOutputA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"redeemForMinTokenA\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOutputB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"redeemForMinTokenB\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router01\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEPool\",\"name\":\"ePool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approval\",\"type\":\"bool\"}],\"name\":\"setEPoolApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxFlashSwapSlippage\",\"type\":\"uint256\"}],\"name\":\"setMaxFlashSwapSlippage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV2Call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var Epoolperiphery = NewEpoolperipheryDecoder()

type EpoolperipheryDecoder struct {
	*ethgen.Decoder
}

func NewEpoolperipheryDecoder() *EpoolperipheryDecoder {
	dec := ethgen.NewDecoder(EpoolperipheryABI)
	return &EpoolperipheryDecoder{
		dec,
	}
}

type EpoolperipherySetMaxFlashSwapSlippageEvent struct {
	MaxFlashSwapSlippage *big.Int
	Raw                  types.Log
}

func (e *EpoolperipherySetMaxFlashSwapSlippageEvent) MaxFlashSwapSlippageDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.MaxFlashSwapSlippage, exp)
}

func (d *EpoolperipheryDecoder) SetMaxFlashSwapSlippageEventID() common.Hash {
	return common.HexToHash("0x0643d67fca12c9df31681eb40c77a7af653ff2111856476f8025530321e4a4f1")
}

func (d *EpoolperipheryDecoder) IsSetMaxFlashSwapSlippageEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetMaxFlashSwapSlippageEventID()
}

func (d *EpoolperipheryDecoder) IsSetMaxFlashSwapSlippageEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetMaxFlashSwapSlippageEventID().String()
}

func (d *EpoolperipheryDecoder) SetMaxFlashSwapSlippageEventW3(w3l web3types.Log) (EpoolperipherySetMaxFlashSwapSlippageEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EpoolperipherySetMaxFlashSwapSlippageEvent{}, err
	}

	return d.SetMaxFlashSwapSlippageEvent(l)
}

func (d *EpoolperipheryDecoder) SetMaxFlashSwapSlippageEvent(l types.Log) (EpoolperipherySetMaxFlashSwapSlippageEvent, error) {
	var out EpoolperipherySetMaxFlashSwapSlippageEvent
	if !d.IsSetMaxFlashSwapSlippageEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SetMaxFlashSwapSlippage", l)
	out.Raw = l
	return out, err
}

type EpoolperipherySetEPoolApprovalEvent struct {
	EPool    common.Address
	Approval bool
	Raw      types.Log
}

func (d *EpoolperipheryDecoder) SetEPoolApprovalEventID() common.Hash {
	return common.HexToHash("0x36110d300c0361d81d13e1816828a532a7119d19d23d5a72bb6174959090ef20")
}

func (d *EpoolperipheryDecoder) IsSetEPoolApprovalEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetEPoolApprovalEventID()
}

func (d *EpoolperipheryDecoder) IsSetEPoolApprovalEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetEPoolApprovalEventID().String()
}

func (d *EpoolperipheryDecoder) SetEPoolApprovalEventW3(w3l web3types.Log) (EpoolperipherySetEPoolApprovalEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EpoolperipherySetEPoolApprovalEvent{}, err
	}

	return d.SetEPoolApprovalEvent(l)
}

func (d *EpoolperipheryDecoder) SetEPoolApprovalEvent(l types.Log) (EpoolperipherySetEPoolApprovalEvent, error) {
	var out EpoolperipherySetEPoolApprovalEvent
	if !d.IsSetEPoolApprovalEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SetEPoolApproval", l)
	out.Raw = l
	return out, err
}

type EpoolperipheryRedeemedETokenEvent struct {
	EPool   common.Address
	EToken  common.Address
	Amount  *big.Int
	AmountA *big.Int
	AmountB *big.Int
	User    common.Address
	Raw     types.Log
}

func (e *EpoolperipheryRedeemedETokenEvent) AmountDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.Amount, exp)
}

func (e *EpoolperipheryRedeemedETokenEvent) AmountADecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.AmountA, exp)
}

func (e *EpoolperipheryRedeemedETokenEvent) AmountBDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.AmountB, exp)
}

func (d *EpoolperipheryDecoder) RedeemedETokenEventID() common.Hash {
	return common.HexToHash("0x4a113d20458ea4526216176713e9b53b289c4b76d33427c850f8279c012b57e3")
}

func (d *EpoolperipheryDecoder) IsRedeemedETokenEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.RedeemedETokenEventID()
}

func (d *EpoolperipheryDecoder) IsRedeemedETokenEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.RedeemedETokenEventID().String()
}

func (d *EpoolperipheryDecoder) RedeemedETokenEventW3(w3l web3types.Log) (EpoolperipheryRedeemedETokenEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EpoolperipheryRedeemedETokenEvent{}, err
	}

	return d.RedeemedETokenEvent(l)
}

func (d *EpoolperipheryDecoder) RedeemedETokenEvent(l types.Log) (EpoolperipheryRedeemedETokenEvent, error) {
	var out EpoolperipheryRedeemedETokenEvent
	if !d.IsRedeemedETokenEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "RedeemedEToken", l)
	out.Raw = l
	return out, err
}

type EpoolperipherySetControllerEvent struct {
	Controller common.Address
	Raw        types.Log
}

func (d *EpoolperipheryDecoder) SetControllerEventID() common.Hash {
	return common.HexToHash("0x4ff638452bbf33c012645d18ae6f05515ff5f2d1dfb0cece8cbf018c60903f70")
}

func (d *EpoolperipheryDecoder) IsSetControllerEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetControllerEventID()
}

func (d *EpoolperipheryDecoder) IsSetControllerEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetControllerEventID().String()
}

func (d *EpoolperipheryDecoder) SetControllerEventW3(w3l web3types.Log) (EpoolperipherySetControllerEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EpoolperipherySetControllerEvent{}, err
	}

	return d.SetControllerEvent(l)
}

func (d *EpoolperipheryDecoder) SetControllerEvent(l types.Log) (EpoolperipherySetControllerEvent, error) {
	var out EpoolperipherySetControllerEvent
	if !d.IsSetControllerEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SetController", l)
	out.Raw = l
	return out, err
}

type EpoolperipheryRecoveredTokenEvent struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log
}

func (e *EpoolperipheryRecoveredTokenEvent) AmountDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.Amount, exp)
}

func (d *EpoolperipheryDecoder) RecoveredTokenEventID() common.Hash {
	return common.HexToHash("0x6de8b63479ce07cf2dfc515e20a5c88a3a5bab6cbd76f753388b77e244ca7071")
}

func (d *EpoolperipheryDecoder) IsRecoveredTokenEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.RecoveredTokenEventID()
}

func (d *EpoolperipheryDecoder) IsRecoveredTokenEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.RecoveredTokenEventID().String()
}

func (d *EpoolperipheryDecoder) RecoveredTokenEventW3(w3l web3types.Log) (EpoolperipheryRecoveredTokenEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EpoolperipheryRecoveredTokenEvent{}, err
	}

	return d.RecoveredTokenEvent(l)
}

func (d *EpoolperipheryDecoder) RecoveredTokenEvent(l types.Log) (EpoolperipheryRecoveredTokenEvent, error) {
	var out EpoolperipheryRecoveredTokenEvent
	if !d.IsRecoveredTokenEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "RecoveredToken", l)
	out.Raw = l
	return out, err
}

type EpoolperipheryIssuedETokenEvent struct {
	EPool   common.Address
	EToken  common.Address
	Amount  *big.Int
	AmountA *big.Int
	AmountB *big.Int
	User    common.Address
	Raw     types.Log
}

func (e *EpoolperipheryIssuedETokenEvent) AmountDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.Amount, exp)
}

func (e *EpoolperipheryIssuedETokenEvent) AmountADecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.AmountA, exp)
}

func (e *EpoolperipheryIssuedETokenEvent) AmountBDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.AmountB, exp)
}

func (d *EpoolperipheryDecoder) IssuedETokenEventID() common.Hash {
	return common.HexToHash("0xd858d811a788f94d2dd517266f495a17fe2bf01222ec01ef5a6b9863062db999")
}

func (d *EpoolperipheryDecoder) IsIssuedETokenEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.IssuedETokenEventID()
}

func (d *EpoolperipheryDecoder) IsIssuedETokenEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.IssuedETokenEventID().String()
}

func (d *EpoolperipheryDecoder) IssuedETokenEventW3(w3l web3types.Log) (EpoolperipheryIssuedETokenEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EpoolperipheryIssuedETokenEvent{}, err
	}

	return d.IssuedETokenEvent(l)
}

func (d *EpoolperipheryDecoder) IssuedETokenEvent(l types.Log) (EpoolperipheryIssuedETokenEvent, error) {
	var out EpoolperipheryIssuedETokenEvent
	if !d.IsIssuedETokenEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "IssuedEToken", l)
	out.Raw = l
	return out, err
}
