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

const SmartYieldCompoundProviderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feesOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"name\":\"TransferFees\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EXP_SCALE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_UINT256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlyingAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takeFees_\",\"type\":\"uint256\"}],\"name\":\"_depositProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingAmount_\",\"type\":\"uint256\"}],\"name\":\"_sendUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_setup\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingAmount_\",\"type\":\"uint256\"}],\"name\":\"_takeUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlyingAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takeFees_\",\"type\":\"uint256\"}],\"name\":\"_withdrawProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateCurrentCached\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateCurrentCachedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newController_\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"smartYield_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"controller_\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"smartYield\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateAllowances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var SmartYieldCompoundProvider = NewSmartYieldCompoundProviderDecoder()

type SmartYieldCompoundProviderDecoder struct {
	*ethgen.Decoder
}

func NewSmartYieldCompoundProviderDecoder() *SmartYieldCompoundProviderDecoder {
	dec := ethgen.NewDecoder(SmartYieldCompoundProviderABI)
	return &SmartYieldCompoundProviderDecoder{
		dec,
	}
}

type SmartYieldCompoundProviderTransferFeesEvent struct {
	Caller    common.Address
	FeesOwner common.Address
	Fees      *big.Int
	Raw       types.Log
}

func (e *SmartYieldCompoundProviderTransferFeesEvent) FeesDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.Fees, exp)
}

func (d *SmartYieldCompoundProviderDecoder) TransferFeesEventID() common.Hash {
	return common.HexToHash("0x9f87918b63a7c3c287a77288ad6cb96549e9af160ef33eb03880fd127ed46deb")
}

func (d *SmartYieldCompoundProviderDecoder) IsTransferFeesEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.TransferFeesEventID()
}

func (d *SmartYieldCompoundProviderDecoder) IsTransferFeesEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.TransferFeesEventID().String()
}

func (d *SmartYieldCompoundProviderDecoder) TransferFeesEventW3(w3l web3types.Log) (SmartYieldCompoundProviderTransferFeesEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartYieldCompoundProviderTransferFeesEvent{}, err
	}

	return d.TransferFeesEvent(l)
}

func (d *SmartYieldCompoundProviderDecoder) TransferFeesEvent(l types.Log) (SmartYieldCompoundProviderTransferFeesEvent, error) {
	var out SmartYieldCompoundProviderTransferFeesEvent
	if !d.IsTransferFeesEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "TransferFees", l)
	out.Raw = l
	return out, err
}
