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

const SmartYieldABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"juniorBondId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maturesAt\",\"type\":\"uint256\"}],\"name\":\"BuyJuniorBond\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"seniorBondId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"underlyingIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gain\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forDays\",\"type\":\"uint256\"}],\"name\":\"BuySeniorBond\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"underlyingIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"BuyTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"juniorBondId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"underlyingOut\",\"type\":\"uint256\"}],\"name\":\"RedeemJuniorBond\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"seniorBondId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"RedeemSeniorBond\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"underlyingOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forfeits\",\"type\":\"uint256\"}],\"name\":\"SellTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EXP_SCALE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_UINT256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_setup\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"abond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"issuedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturesAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"liquidated\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"abondDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"abondGain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"abondPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"principalAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"forDays_\",\"type\":\"uint16\"}],\"name\":\"bondGain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"principalAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minGain_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline_\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"forDays_\",\"type\":\"uint16\"}],\"name\":\"buyBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMaturesAt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline_\",\"type\":\"uint256\"}],\"name\":\"buyJuniorBond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlyingAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTokens_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline_\",\"type\":\"uint256\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"juniorBond\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"juniorBondId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"juniorBonds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturesAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"juniorBondsMaturingAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"juniorBondsMaturities\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"juniorBondsMaturitiesPrev\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"upUntilTimestamp_\",\"type\":\"uint256\"}],\"name\":\"liquidateJuniorBonds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBondDailyRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bondId_\",\"type\":\"uint256\"}],\"name\":\"redeemBond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jBondId_\",\"type\":\"uint256\"}],\"name\":\"redeemJuniorBond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minUnderlying_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline_\",\"type\":\"uint256\"}],\"name\":\"sellTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seniorBond\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seniorBondId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"seniorBonds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"issuedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturesAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"liquidated\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newController_\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seniorBond_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"juniorBond_\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensInJuniorBonds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"bondIds_\",\"type\":\"uint256[]\"}],\"name\":\"unaccountBonds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingJuniors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingLiquidatedJuniors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingLoanable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var SmartYield = NewSmartYieldDecoder()

type SmartYieldDecoder struct {
	*ethgen.Decoder
}

func NewSmartYieldDecoder() *SmartYieldDecoder {
	dec := ethgen.NewDecoder(SmartYieldABI)
	return &SmartYieldDecoder{
		dec,
	}
}

type SmartYieldBuySeniorBondEvent struct {
	Buyer        common.Address
	SeniorBondId *big.Int
	UnderlyingIn *big.Int
	Gain         *big.Int
	ForDays      *big.Int
	Raw          types.Log
}

func (d *SmartYieldDecoder) SmartYieldBuySeniorBondEventID() common.Hash {
	return common.HexToHash("0x2108cac1cef8ac37dbe265869f744fd7ae754c3ed949530f24fb62a47f71b23c")
}

func (d *SmartYieldDecoder) IsSmartYieldBuySeniorBondEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SmartYieldBuySeniorBondEventID()
}

func (d *SmartYieldDecoder) SmartYieldBuySeniorBondEventW3(w3l web3types.Log) (SmartYieldBuySeniorBondEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartYieldBuySeniorBondEvent{}, err
	}

	return d.SmartYieldBuySeniorBondEvent(l)
}

func (d *SmartYieldDecoder) SmartYieldBuySeniorBondEvent(l types.Log) (SmartYieldBuySeniorBondEvent, error) {
	var out SmartYieldBuySeniorBondEvent
	if !d.IsSmartYieldBuySeniorBondEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "BuySeniorBond", l)
	out.Raw = l
	return out, err
}

type SmartYieldApprovalEvent struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (d *SmartYieldDecoder) SmartYieldApprovalEventID() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (d *SmartYieldDecoder) IsSmartYieldApprovalEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SmartYieldApprovalEventID()
}

func (d *SmartYieldDecoder) SmartYieldApprovalEventW3(w3l web3types.Log) (SmartYieldApprovalEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartYieldApprovalEvent{}, err
	}

	return d.SmartYieldApprovalEvent(l)
}

func (d *SmartYieldDecoder) SmartYieldApprovalEvent(l types.Log) (SmartYieldApprovalEvent, error) {
	var out SmartYieldApprovalEvent
	if !d.IsSmartYieldApprovalEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "Approval", l)
	out.Raw = l
	return out, err
}

type SmartYieldBuyTokensEvent struct {
	Buyer        common.Address
	UnderlyingIn *big.Int
	TokensOut    *big.Int
	Fee          *big.Int
	Raw          types.Log
}

func (d *SmartYieldDecoder) SmartYieldBuyTokensEventID() common.Hash {
	return common.HexToHash("0x90d8b08a6c17cc6733ded05f205dd10dd0538fb7890449f561eedef38c91a6fa")
}

func (d *SmartYieldDecoder) IsSmartYieldBuyTokensEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SmartYieldBuyTokensEventID()
}

func (d *SmartYieldDecoder) SmartYieldBuyTokensEventW3(w3l web3types.Log) (SmartYieldBuyTokensEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartYieldBuyTokensEvent{}, err
	}

	return d.SmartYieldBuyTokensEvent(l)
}

func (d *SmartYieldDecoder) SmartYieldBuyTokensEvent(l types.Log) (SmartYieldBuyTokensEvent, error) {
	var out SmartYieldBuyTokensEvent
	if !d.IsSmartYieldBuyTokensEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "BuyTokens", l)
	out.Raw = l
	return out, err
}

type SmartYieldBuyJuniorBondEvent struct {
	Buyer        common.Address
	JuniorBondId *big.Int
	TokensIn     *big.Int
	MaturesAt    *big.Int
	Raw          types.Log
}

func (d *SmartYieldDecoder) SmartYieldBuyJuniorBondEventID() common.Hash {
	return common.HexToHash("0x93f0f0774770973693e0c5a43673d8aef029f09f65f442397777c9af9fadc60c")
}

func (d *SmartYieldDecoder) IsSmartYieldBuyJuniorBondEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SmartYieldBuyJuniorBondEventID()
}

func (d *SmartYieldDecoder) SmartYieldBuyJuniorBondEventW3(w3l web3types.Log) (SmartYieldBuyJuniorBondEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartYieldBuyJuniorBondEvent{}, err
	}

	return d.SmartYieldBuyJuniorBondEvent(l)
}

func (d *SmartYieldDecoder) SmartYieldBuyJuniorBondEvent(l types.Log) (SmartYieldBuyJuniorBondEvent, error) {
	var out SmartYieldBuyJuniorBondEvent
	if !d.IsSmartYieldBuyJuniorBondEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "BuyJuniorBond", l)
	out.Raw = l
	return out, err
}

type SmartYieldSellTokensEvent struct {
	Seller        common.Address
	TokensIn      *big.Int
	UnderlyingOut *big.Int
	Forfeits      *big.Int
	Raw           types.Log
}

func (d *SmartYieldDecoder) SmartYieldSellTokensEventID() common.Hash {
	return common.HexToHash("0x95ff24e35ad23e93c0738cee55f0903db5c47b23968d07627a68fe23ebd11b6d")
}

func (d *SmartYieldDecoder) IsSmartYieldSellTokensEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SmartYieldSellTokensEventID()
}

func (d *SmartYieldDecoder) SmartYieldSellTokensEventW3(w3l web3types.Log) (SmartYieldSellTokensEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartYieldSellTokensEvent{}, err
	}

	return d.SmartYieldSellTokensEvent(l)
}

func (d *SmartYieldDecoder) SmartYieldSellTokensEvent(l types.Log) (SmartYieldSellTokensEvent, error) {
	var out SmartYieldSellTokensEvent
	if !d.IsSmartYieldSellTokensEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SellTokens", l)
	out.Raw = l
	return out, err
}

type SmartYieldTransferEvent struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (d *SmartYieldDecoder) SmartYieldTransferEventID() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (d *SmartYieldDecoder) IsSmartYieldTransferEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SmartYieldTransferEventID()
}

func (d *SmartYieldDecoder) SmartYieldTransferEventW3(w3l web3types.Log) (SmartYieldTransferEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartYieldTransferEvent{}, err
	}

	return d.SmartYieldTransferEvent(l)
}

func (d *SmartYieldDecoder) SmartYieldTransferEvent(l types.Log) (SmartYieldTransferEvent, error) {
	var out SmartYieldTransferEvent
	if !d.IsSmartYieldTransferEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "Transfer", l)
	out.Raw = l
	return out, err
}

type SmartYieldRedeemJuniorBondEvent struct {
	Owner         common.Address
	JuniorBondId  *big.Int
	UnderlyingOut *big.Int
	Raw           types.Log
}

func (d *SmartYieldDecoder) SmartYieldRedeemJuniorBondEventID() common.Hash {
	return common.HexToHash("0xe34274b2ac2992188914cc9b0f4cb53202d013fc7b1996edb6b8564431f7bd53")
}

func (d *SmartYieldDecoder) IsSmartYieldRedeemJuniorBondEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SmartYieldRedeemJuniorBondEventID()
}

func (d *SmartYieldDecoder) SmartYieldRedeemJuniorBondEventW3(w3l web3types.Log) (SmartYieldRedeemJuniorBondEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartYieldRedeemJuniorBondEvent{}, err
	}

	return d.SmartYieldRedeemJuniorBondEvent(l)
}

func (d *SmartYieldDecoder) SmartYieldRedeemJuniorBondEvent(l types.Log) (SmartYieldRedeemJuniorBondEvent, error) {
	var out SmartYieldRedeemJuniorBondEvent
	if !d.IsSmartYieldRedeemJuniorBondEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "RedeemJuniorBond", l)
	out.Raw = l
	return out, err
}

type SmartYieldRedeemSeniorBondEvent struct {
	Owner        common.Address
	SeniorBondId *big.Int
	Fee          *big.Int
	Raw          types.Log
}

func (d *SmartYieldDecoder) SmartYieldRedeemSeniorBondEventID() common.Hash {
	return common.HexToHash("0xfa51bdcf530ef35114732d8f7598a2938621008a16d9bb235a8c84fe82e4841e")
}

func (d *SmartYieldDecoder) IsSmartYieldRedeemSeniorBondEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SmartYieldRedeemSeniorBondEventID()
}

func (d *SmartYieldDecoder) SmartYieldRedeemSeniorBondEventW3(w3l web3types.Log) (SmartYieldRedeemSeniorBondEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartYieldRedeemSeniorBondEvent{}, err
	}

	return d.SmartYieldRedeemSeniorBondEvent(l)
}

func (d *SmartYieldDecoder) SmartYieldRedeemSeniorBondEvent(l types.Log) (SmartYieldRedeemSeniorBondEvent, error) {
	var out SmartYieldRedeemSeniorBondEvent
	if !d.IsSmartYieldRedeemSeniorBondEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "RedeemSeniorBond", l)
	out.Raw = l
	return out, err
}
