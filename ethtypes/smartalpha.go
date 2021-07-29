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

const SmartAlphaABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"juniorProfits\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seniorProfits\",\"type\":\"uint256\"}],\"name\":\"EpochEnd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"underlyingIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentQueueBalance\",\"type\":\"uint256\"}],\"name\":\"JuniorJoinEntryQueue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentQueueBalance\",\"type\":\"uint256\"}],\"name\":\"JuniorJoinExitQueue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensOut\",\"type\":\"uint256\"}],\"name\":\"JuniorRedeemTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"underlyingOut\",\"type\":\"uint256\"}],\"name\":\"JuniorRedeemUnderlying\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PauseSystem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ResumeSystem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"underlyingIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentQueueBalance\",\"type\":\"uint256\"}],\"name\":\"SeniorJoinEntryQueue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentQueueBalance\",\"type\":\"uint256\"}],\"name\":\"SeniorJoinExitQueue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensOut\",\"type\":\"uint256\"}],\"name\":\"SeniorRedeemTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"underlyingOut\",\"type\":\"uint256\"}],\"name\":\"SeniorRedeemUnderlying\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldModel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newModel\",\"type\":\"address\"}],\"name\":\"SetAccountingModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"SetFeesOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPercentage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPercentage\",\"type\":\"uint256\"}],\"name\":\"SetFeesPercentage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"SetPriceOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldModel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newModel\",\"type\":\"address\"}],\"name\":\"SetSeniorRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldDAO\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDAO\",\"type\":\"address\"}],\"name\":\"TransferDAO\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newGuardian\",\"type\":\"address\"}],\"name\":\"TransferGuardian\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_FEES_PERCENTAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountingModel\",\"outputs\":[{\"internalType\":\"contractIAccountingModel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"advanceEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositJunior\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositSenior\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch1Start\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDownsideProtectionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochEntryPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochJuniorLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochSeniorLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochUpsideExposureRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimateCurrentJuniorLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimateCurrentJuniorTokenPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimateCurrentSeniorLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimateCurrentSeniorTokenPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountJuniorTokens\",\"type\":\"uint256\"}],\"name\":\"exitJunior\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountSeniorTokens\",\"type\":\"uint256\"}],\"name\":\"exitSenior\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feesAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feesOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feesPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentJuniorProfits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentSeniorProfits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEpochJuniorTokenPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEpochSeniorTokenPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"history_epochJuniorTokenPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"history_epochSeniorTokenPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracleAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seniorRateModelAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"accountingModelAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"juniorTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seniorTokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_epoch1Start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epochDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"juniorEntryQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"juniorExitQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"juniorToken\",\"outputs\":[{\"internalType\":\"contractOwnableERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceOracle\",\"outputs\":[{\"internalType\":\"contractIPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queuedJuniorTokensBurn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queuedJuniorsUnderlyingIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queuedJuniorsUnderlyingOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queuedSeniorTokensBurn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queuedSeniorsUnderlyingIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queuedSeniorsUnderlyingOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemJuniorTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemJuniorUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemSeniorTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemSeniorUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resumeSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scaleFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"seniorEntryQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"seniorExitQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seniorRateModel\",\"outputs\":[{\"internalType\":\"contractISeniorRateModel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seniorToken\",\"outputs\":[{\"internalType\":\"contractOwnableERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newModel\",\"type\":\"address\"}],\"name\":\"setAccountingModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setFeesOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"}],\"name\":\"setFeesPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newModel\",\"type\":\"address\"}],\"name\":\"setSeniorRateModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDAO\",\"type\":\"address\"}],\"name\":\"transferDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGuardian\",\"type\":\"address\"}],\"name\":\"transferGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var SmartAlpha = NewSmartAlphaDecoder()

type SmartAlphaDecoder struct {
	*ethgen.Decoder
}

func NewSmartAlphaDecoder() *SmartAlphaDecoder {
	dec := ethgen.NewDecoder(SmartAlphaABI)
	return &SmartAlphaDecoder{
		dec,
	}
}

type SmartAlphaSeniorRedeemUnderlyingEvent struct {
	User          common.Address
	EpochId       *big.Int
	UnderlyingOut *big.Int
	Raw           types.Log
}

func (e *SmartAlphaSeniorRedeemUnderlyingEvent) EpochIdDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.EpochId, exp)
}

func (e *SmartAlphaSeniorRedeemUnderlyingEvent) UnderlyingOutDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.UnderlyingOut, exp)
}

func (d *SmartAlphaDecoder) SeniorRedeemUnderlyingEventID() common.Hash {
	return common.HexToHash("0x014d40541bf53e076bad44aa8bfef33f82c777c0f04196e5582587bad9ee5b36")
}

func (d *SmartAlphaDecoder) IsSeniorRedeemUnderlyingEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SeniorRedeemUnderlyingEventID()
}

func (d *SmartAlphaDecoder) IsSeniorRedeemUnderlyingEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SeniorRedeemUnderlyingEventID().String()
}

func (d *SmartAlphaDecoder) SeniorRedeemUnderlyingEventW3(w3l web3types.Log) (SmartAlphaSeniorRedeemUnderlyingEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaSeniorRedeemUnderlyingEvent{}, err
	}

	return d.SeniorRedeemUnderlyingEvent(l)
}

func (d *SmartAlphaDecoder) SeniorRedeemUnderlyingEvent(l types.Log) (SmartAlphaSeniorRedeemUnderlyingEvent, error) {
	var out SmartAlphaSeniorRedeemUnderlyingEvent
	if !d.IsSeniorRedeemUnderlyingEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SeniorRedeemUnderlying", l)
	out.Raw = l
	return out, err
}

type SmartAlphaTransferGuardianEvent struct {
	OldGuardian common.Address
	NewGuardian common.Address
	Raw         types.Log
}

func (d *SmartAlphaDecoder) TransferGuardianEventID() common.Hash {
	return common.HexToHash("0x19e3cbfd9b25c12fb88132e5887dc3a2a4f52979bd0af17f78e6fea817addf4b")
}

func (d *SmartAlphaDecoder) IsTransferGuardianEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.TransferGuardianEventID()
}

func (d *SmartAlphaDecoder) IsTransferGuardianEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.TransferGuardianEventID().String()
}

func (d *SmartAlphaDecoder) TransferGuardianEventW3(w3l web3types.Log) (SmartAlphaTransferGuardianEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaTransferGuardianEvent{}, err
	}

	return d.TransferGuardianEvent(l)
}

func (d *SmartAlphaDecoder) TransferGuardianEvent(l types.Log) (SmartAlphaTransferGuardianEvent, error) {
	var out SmartAlphaTransferGuardianEvent
	if !d.IsTransferGuardianEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "TransferGuardian", l)
	out.Raw = l
	return out, err
}

type SmartAlphaResumeSystemEvent struct {
	Raw types.Log
}

func (d *SmartAlphaDecoder) ResumeSystemEventID() common.Hash {
	return common.HexToHash("0x3008289779e367cc52c347e8bb0a995361494ac17d2ccb8b09a75a77f8c58b94")
}

func (d *SmartAlphaDecoder) IsResumeSystemEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.ResumeSystemEventID()
}

func (d *SmartAlphaDecoder) IsResumeSystemEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.ResumeSystemEventID().String()
}

func (d *SmartAlphaDecoder) ResumeSystemEventW3(w3l web3types.Log) (SmartAlphaResumeSystemEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaResumeSystemEvent{}, err
	}

	return d.ResumeSystemEvent(l)
}

func (d *SmartAlphaDecoder) ResumeSystemEvent(l types.Log) (SmartAlphaResumeSystemEvent, error) {
	var out SmartAlphaResumeSystemEvent
	if !d.IsResumeSystemEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "ResumeSystem", l)
	out.Raw = l
	return out, err
}

type SmartAlphaSetFeesOwnerEvent struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log
}

func (d *SmartAlphaDecoder) SetFeesOwnerEventID() common.Hash {
	return common.HexToHash("0x461364f084b7657c2380660ebd35dd6c4560dc78cc6e9795919e53b4e257de83")
}

func (d *SmartAlphaDecoder) IsSetFeesOwnerEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetFeesOwnerEventID()
}

func (d *SmartAlphaDecoder) IsSetFeesOwnerEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetFeesOwnerEventID().String()
}

func (d *SmartAlphaDecoder) SetFeesOwnerEventW3(w3l web3types.Log) (SmartAlphaSetFeesOwnerEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaSetFeesOwnerEvent{}, err
	}

	return d.SetFeesOwnerEvent(l)
}

func (d *SmartAlphaDecoder) SetFeesOwnerEvent(l types.Log) (SmartAlphaSetFeesOwnerEvent, error) {
	var out SmartAlphaSetFeesOwnerEvent
	if !d.IsSetFeesOwnerEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SetFeesOwner", l)
	out.Raw = l
	return out, err
}

type SmartAlphaFeesTransferEvent struct {
	Caller      common.Address
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log
}

func (e *SmartAlphaFeesTransferEvent) AmountDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.Amount, exp)
}

func (d *SmartAlphaDecoder) FeesTransferEventID() common.Hash {
	return common.HexToHash("0x6d86532cebce232743f74710d18164bfb38cbebf1ca203cb83382b9cf54cd9ac")
}

func (d *SmartAlphaDecoder) IsFeesTransferEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.FeesTransferEventID()
}

func (d *SmartAlphaDecoder) IsFeesTransferEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.FeesTransferEventID().String()
}

func (d *SmartAlphaDecoder) FeesTransferEventW3(w3l web3types.Log) (SmartAlphaFeesTransferEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaFeesTransferEvent{}, err
	}

	return d.FeesTransferEvent(l)
}

func (d *SmartAlphaDecoder) FeesTransferEvent(l types.Log) (SmartAlphaFeesTransferEvent, error) {
	var out SmartAlphaFeesTransferEvent
	if !d.IsFeesTransferEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "FeesTransfer", l)
	out.Raw = l
	return out, err
}

type SmartAlphaJuniorJoinEntryQueueEvent struct {
	User                common.Address
	EpochId             *big.Int
	UnderlyingIn        *big.Int
	CurrentQueueBalance *big.Int
	Raw                 types.Log
}

func (e *SmartAlphaJuniorJoinEntryQueueEvent) EpochIdDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.EpochId, exp)
}

func (e *SmartAlphaJuniorJoinEntryQueueEvent) UnderlyingInDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.UnderlyingIn, exp)
}

func (e *SmartAlphaJuniorJoinEntryQueueEvent) CurrentQueueBalanceDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.CurrentQueueBalance, exp)
}

func (d *SmartAlphaDecoder) JuniorJoinEntryQueueEventID() common.Hash {
	return common.HexToHash("0x81b0ac0867278845eee541a0351e48acc52b65ba1a469e840ad5a4f8a8650a2e")
}

func (d *SmartAlphaDecoder) IsJuniorJoinEntryQueueEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.JuniorJoinEntryQueueEventID()
}

func (d *SmartAlphaDecoder) IsJuniorJoinEntryQueueEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.JuniorJoinEntryQueueEventID().String()
}

func (d *SmartAlphaDecoder) JuniorJoinEntryQueueEventW3(w3l web3types.Log) (SmartAlphaJuniorJoinEntryQueueEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaJuniorJoinEntryQueueEvent{}, err
	}

	return d.JuniorJoinEntryQueueEvent(l)
}

func (d *SmartAlphaDecoder) JuniorJoinEntryQueueEvent(l types.Log) (SmartAlphaJuniorJoinEntryQueueEvent, error) {
	var out SmartAlphaJuniorJoinEntryQueueEvent
	if !d.IsJuniorJoinEntryQueueEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "JuniorJoinEntryQueue", l)
	out.Raw = l
	return out, err
}

type SmartAlphaJuniorRedeemTokensEvent struct {
	User      common.Address
	EpochId   *big.Int
	TokensOut *big.Int
	Raw       types.Log
}

func (e *SmartAlphaJuniorRedeemTokensEvent) EpochIdDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.EpochId, exp)
}

func (e *SmartAlphaJuniorRedeemTokensEvent) TokensOutDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.TokensOut, exp)
}

func (d *SmartAlphaDecoder) JuniorRedeemTokensEventID() common.Hash {
	return common.HexToHash("0x822bdec5f81295b087914a2de8160ebad0523c35c96ab4ab599564d86c588ced")
}

func (d *SmartAlphaDecoder) IsJuniorRedeemTokensEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.JuniorRedeemTokensEventID()
}

func (d *SmartAlphaDecoder) IsJuniorRedeemTokensEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.JuniorRedeemTokensEventID().String()
}

func (d *SmartAlphaDecoder) JuniorRedeemTokensEventW3(w3l web3types.Log) (SmartAlphaJuniorRedeemTokensEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaJuniorRedeemTokensEvent{}, err
	}

	return d.JuniorRedeemTokensEvent(l)
}

func (d *SmartAlphaDecoder) JuniorRedeemTokensEvent(l types.Log) (SmartAlphaJuniorRedeemTokensEvent, error) {
	var out SmartAlphaJuniorRedeemTokensEvent
	if !d.IsJuniorRedeemTokensEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "JuniorRedeemTokens", l)
	out.Raw = l
	return out, err
}

type SmartAlphaSeniorRedeemTokensEvent struct {
	User      common.Address
	EpochId   *big.Int
	TokensOut *big.Int
	Raw       types.Log
}

func (e *SmartAlphaSeniorRedeemTokensEvent) EpochIdDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.EpochId, exp)
}

func (e *SmartAlphaSeniorRedeemTokensEvent) TokensOutDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.TokensOut, exp)
}

func (d *SmartAlphaDecoder) SeniorRedeemTokensEventID() common.Hash {
	return common.HexToHash("0x8b9d5ffabdef5875b86c23ad0939273948c7611ebc419a3c23c8f8532cbe20c8")
}

func (d *SmartAlphaDecoder) IsSeniorRedeemTokensEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SeniorRedeemTokensEventID()
}

func (d *SmartAlphaDecoder) IsSeniorRedeemTokensEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SeniorRedeemTokensEventID().String()
}

func (d *SmartAlphaDecoder) SeniorRedeemTokensEventW3(w3l web3types.Log) (SmartAlphaSeniorRedeemTokensEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaSeniorRedeemTokensEvent{}, err
	}

	return d.SeniorRedeemTokensEvent(l)
}

func (d *SmartAlphaDecoder) SeniorRedeemTokensEvent(l types.Log) (SmartAlphaSeniorRedeemTokensEvent, error) {
	var out SmartAlphaSeniorRedeemTokensEvent
	if !d.IsSeniorRedeemTokensEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SeniorRedeemTokens", l)
	out.Raw = l
	return out, err
}

type SmartAlphaSeniorJoinExitQueueEvent struct {
	User                common.Address
	EpochId             *big.Int
	TokensIn            *big.Int
	CurrentQueueBalance *big.Int
	Raw                 types.Log
}

func (e *SmartAlphaSeniorJoinExitQueueEvent) EpochIdDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.EpochId, exp)
}

func (e *SmartAlphaSeniorJoinExitQueueEvent) TokensInDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.TokensIn, exp)
}

func (e *SmartAlphaSeniorJoinExitQueueEvent) CurrentQueueBalanceDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.CurrentQueueBalance, exp)
}

func (d *SmartAlphaDecoder) SeniorJoinExitQueueEventID() common.Hash {
	return common.HexToHash("0x8fe0fbd18893ea6a3597f5ccec3e4494f92cc4cea7276a9a868c6a9bd1504e61")
}

func (d *SmartAlphaDecoder) IsSeniorJoinExitQueueEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SeniorJoinExitQueueEventID()
}

func (d *SmartAlphaDecoder) IsSeniorJoinExitQueueEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SeniorJoinExitQueueEventID().String()
}

func (d *SmartAlphaDecoder) SeniorJoinExitQueueEventW3(w3l web3types.Log) (SmartAlphaSeniorJoinExitQueueEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaSeniorJoinExitQueueEvent{}, err
	}

	return d.SeniorJoinExitQueueEvent(l)
}

func (d *SmartAlphaDecoder) SeniorJoinExitQueueEvent(l types.Log) (SmartAlphaSeniorJoinExitQueueEvent, error) {
	var out SmartAlphaSeniorJoinExitQueueEvent
	if !d.IsSeniorJoinExitQueueEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SeniorJoinExitQueue", l)
	out.Raw = l
	return out, err
}

type SmartAlphaSetPriceOracleEvent struct {
	OldOracle common.Address
	NewOracle common.Address
	Raw       types.Log
}

func (d *SmartAlphaDecoder) SetPriceOracleEventID() common.Hash {
	return common.HexToHash("0x944f260fa5475db9151efe54fc9fb6df2b40f73533d9483011e94522398cb3c1")
}

func (d *SmartAlphaDecoder) IsSetPriceOracleEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetPriceOracleEventID()
}

func (d *SmartAlphaDecoder) IsSetPriceOracleEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetPriceOracleEventID().String()
}

func (d *SmartAlphaDecoder) SetPriceOracleEventW3(w3l web3types.Log) (SmartAlphaSetPriceOracleEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaSetPriceOracleEvent{}, err
	}

	return d.SetPriceOracleEvent(l)
}

func (d *SmartAlphaDecoder) SetPriceOracleEvent(l types.Log) (SmartAlphaSetPriceOracleEvent, error) {
	var out SmartAlphaSetPriceOracleEvent
	if !d.IsSetPriceOracleEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SetPriceOracle", l)
	out.Raw = l
	return out, err
}

type SmartAlphaJuniorRedeemUnderlyingEvent struct {
	User          common.Address
	EpochId       *big.Int
	UnderlyingOut *big.Int
	Raw           types.Log
}

func (e *SmartAlphaJuniorRedeemUnderlyingEvent) EpochIdDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.EpochId, exp)
}

func (e *SmartAlphaJuniorRedeemUnderlyingEvent) UnderlyingOutDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.UnderlyingOut, exp)
}

func (d *SmartAlphaDecoder) JuniorRedeemUnderlyingEventID() common.Hash {
	return common.HexToHash("0xa362bf831ec0975732381d8a2218cc84161953de04be755d93f9474860f432ee")
}

func (d *SmartAlphaDecoder) IsJuniorRedeemUnderlyingEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.JuniorRedeemUnderlyingEventID()
}

func (d *SmartAlphaDecoder) IsJuniorRedeemUnderlyingEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.JuniorRedeemUnderlyingEventID().String()
}

func (d *SmartAlphaDecoder) JuniorRedeemUnderlyingEventW3(w3l web3types.Log) (SmartAlphaJuniorRedeemUnderlyingEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaJuniorRedeemUnderlyingEvent{}, err
	}

	return d.JuniorRedeemUnderlyingEvent(l)
}

func (d *SmartAlphaDecoder) JuniorRedeemUnderlyingEvent(l types.Log) (SmartAlphaJuniorRedeemUnderlyingEvent, error) {
	var out SmartAlphaJuniorRedeemUnderlyingEvent
	if !d.IsJuniorRedeemUnderlyingEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "JuniorRedeemUnderlying", l)
	out.Raw = l
	return out, err
}

type SmartAlphaSeniorJoinEntryQueueEvent struct {
	User                common.Address
	EpochId             *big.Int
	UnderlyingIn        *big.Int
	CurrentQueueBalance *big.Int
	Raw                 types.Log
}

func (e *SmartAlphaSeniorJoinEntryQueueEvent) EpochIdDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.EpochId, exp)
}

func (e *SmartAlphaSeniorJoinEntryQueueEvent) UnderlyingInDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.UnderlyingIn, exp)
}

func (e *SmartAlphaSeniorJoinEntryQueueEvent) CurrentQueueBalanceDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.CurrentQueueBalance, exp)
}

func (d *SmartAlphaDecoder) SeniorJoinEntryQueueEventID() common.Hash {
	return common.HexToHash("0xae031a79e52a4da207e03aac3a3217f2cff9d56d24642655fbf7844a07f0f36a")
}

func (d *SmartAlphaDecoder) IsSeniorJoinEntryQueueEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SeniorJoinEntryQueueEventID()
}

func (d *SmartAlphaDecoder) IsSeniorJoinEntryQueueEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SeniorJoinEntryQueueEventID().String()
}

func (d *SmartAlphaDecoder) SeniorJoinEntryQueueEventW3(w3l web3types.Log) (SmartAlphaSeniorJoinEntryQueueEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaSeniorJoinEntryQueueEvent{}, err
	}

	return d.SeniorJoinEntryQueueEvent(l)
}

func (d *SmartAlphaDecoder) SeniorJoinEntryQueueEvent(l types.Log) (SmartAlphaSeniorJoinEntryQueueEvent, error) {
	var out SmartAlphaSeniorJoinEntryQueueEvent
	if !d.IsSeniorJoinEntryQueueEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SeniorJoinEntryQueue", l)
	out.Raw = l
	return out, err
}

type SmartAlphaEpochEndEvent struct {
	EpochId       *big.Int
	JuniorProfits *big.Int
	SeniorProfits *big.Int
	Raw           types.Log
}

func (e *SmartAlphaEpochEndEvent) EpochIdDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.EpochId, exp)
}

func (e *SmartAlphaEpochEndEvent) JuniorProfitsDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.JuniorProfits, exp)
}

func (e *SmartAlphaEpochEndEvent) SeniorProfitsDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.SeniorProfits, exp)
}

func (d *SmartAlphaDecoder) EpochEndEventID() common.Hash {
	return common.HexToHash("0xaed63947ee74099c60670bac488f65fecd95864d94817d4300ee62fa63afe188")
}

func (d *SmartAlphaDecoder) IsEpochEndEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EpochEndEventID()
}

func (d *SmartAlphaDecoder) IsEpochEndEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EpochEndEventID().String()
}

func (d *SmartAlphaDecoder) EpochEndEventW3(w3l web3types.Log) (SmartAlphaEpochEndEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaEpochEndEvent{}, err
	}

	return d.EpochEndEvent(l)
}

func (d *SmartAlphaDecoder) EpochEndEvent(l types.Log) (SmartAlphaEpochEndEvent, error) {
	var out SmartAlphaEpochEndEvent
	if !d.IsEpochEndEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "EpochEnd", l)
	out.Raw = l
	return out, err
}

type SmartAlphaJuniorJoinExitQueueEvent struct {
	User                common.Address
	EpochId             *big.Int
	TokensIn            *big.Int
	CurrentQueueBalance *big.Int
	Raw                 types.Log
}

func (e *SmartAlphaJuniorJoinExitQueueEvent) EpochIdDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.EpochId, exp)
}

func (e *SmartAlphaJuniorJoinExitQueueEvent) TokensInDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.TokensIn, exp)
}

func (e *SmartAlphaJuniorJoinExitQueueEvent) CurrentQueueBalanceDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.CurrentQueueBalance, exp)
}

func (d *SmartAlphaDecoder) JuniorJoinExitQueueEventID() common.Hash {
	return common.HexToHash("0xd634a1ad9b45b7e951a1f30528c0209e937d46d164fd43268090f427349205a3")
}

func (d *SmartAlphaDecoder) IsJuniorJoinExitQueueEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.JuniorJoinExitQueueEventID()
}

func (d *SmartAlphaDecoder) IsJuniorJoinExitQueueEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.JuniorJoinExitQueueEventID().String()
}

func (d *SmartAlphaDecoder) JuniorJoinExitQueueEventW3(w3l web3types.Log) (SmartAlphaJuniorJoinExitQueueEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaJuniorJoinExitQueueEvent{}, err
	}

	return d.JuniorJoinExitQueueEvent(l)
}

func (d *SmartAlphaDecoder) JuniorJoinExitQueueEvent(l types.Log) (SmartAlphaJuniorJoinExitQueueEvent, error) {
	var out SmartAlphaJuniorJoinExitQueueEvent
	if !d.IsJuniorJoinExitQueueEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "JuniorJoinExitQueue", l)
	out.Raw = l
	return out, err
}

type SmartAlphaTransferDAOEvent struct {
	OldDAO common.Address
	NewDAO common.Address
	Raw    types.Log
}

func (d *SmartAlphaDecoder) TransferDAOEventID() common.Hash {
	return common.HexToHash("0xe010a6a96c097bdc495fe0aa7b1c1343e8528bd70556da39cd47b0130555c190")
}

func (d *SmartAlphaDecoder) IsTransferDAOEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.TransferDAOEventID()
}

func (d *SmartAlphaDecoder) IsTransferDAOEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.TransferDAOEventID().String()
}

func (d *SmartAlphaDecoder) TransferDAOEventW3(w3l web3types.Log) (SmartAlphaTransferDAOEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaTransferDAOEvent{}, err
	}

	return d.TransferDAOEvent(l)
}

func (d *SmartAlphaDecoder) TransferDAOEvent(l types.Log) (SmartAlphaTransferDAOEvent, error) {
	var out SmartAlphaTransferDAOEvent
	if !d.IsTransferDAOEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "TransferDAO", l)
	out.Raw = l
	return out, err
}

type SmartAlphaSetFeesPercentageEvent struct {
	OldPercentage *big.Int
	NewPercentage *big.Int
	Raw           types.Log
}

func (e *SmartAlphaSetFeesPercentageEvent) OldPercentageDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.OldPercentage, exp)
}

func (e *SmartAlphaSetFeesPercentageEvent) NewPercentageDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.NewPercentage, exp)
}

func (d *SmartAlphaDecoder) SetFeesPercentageEventID() common.Hash {
	return common.HexToHash("0xe5bd021bb3e7a336d89185c09eb4286717c23a651140650eeda4d774812cde12")
}

func (d *SmartAlphaDecoder) IsSetFeesPercentageEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetFeesPercentageEventID()
}

func (d *SmartAlphaDecoder) IsSetFeesPercentageEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetFeesPercentageEventID().String()
}

func (d *SmartAlphaDecoder) SetFeesPercentageEventW3(w3l web3types.Log) (SmartAlphaSetFeesPercentageEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaSetFeesPercentageEvent{}, err
	}

	return d.SetFeesPercentageEvent(l)
}

func (d *SmartAlphaDecoder) SetFeesPercentageEvent(l types.Log) (SmartAlphaSetFeesPercentageEvent, error) {
	var out SmartAlphaSetFeesPercentageEvent
	if !d.IsSetFeesPercentageEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SetFeesPercentage", l)
	out.Raw = l
	return out, err
}

type SmartAlphaPauseSystemEvent struct {
	Raw types.Log
}

func (d *SmartAlphaDecoder) PauseSystemEventID() common.Hash {
	return common.HexToHash("0xee68d332edc397b95695f3d00eb879007ec02eb6733fe32e305cc12d406d03c1")
}

func (d *SmartAlphaDecoder) IsPauseSystemEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.PauseSystemEventID()
}

func (d *SmartAlphaDecoder) IsPauseSystemEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.PauseSystemEventID().String()
}

func (d *SmartAlphaDecoder) PauseSystemEventW3(w3l web3types.Log) (SmartAlphaPauseSystemEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaPauseSystemEvent{}, err
	}

	return d.PauseSystemEvent(l)
}

func (d *SmartAlphaDecoder) PauseSystemEvent(l types.Log) (SmartAlphaPauseSystemEvent, error) {
	var out SmartAlphaPauseSystemEvent
	if !d.IsPauseSystemEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "PauseSystem", l)
	out.Raw = l
	return out, err
}

type SmartAlphaSetSeniorRateModelEvent struct {
	OldModel common.Address
	NewModel common.Address
	Raw      types.Log
}

func (d *SmartAlphaDecoder) SetSeniorRateModelEventID() common.Hash {
	return common.HexToHash("0xfbbe86207337f27b9ba95decb76b63122f31b20555b8b3737055683611443aac")
}

func (d *SmartAlphaDecoder) IsSetSeniorRateModelEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetSeniorRateModelEventID()
}

func (d *SmartAlphaDecoder) IsSetSeniorRateModelEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetSeniorRateModelEventID().String()
}

func (d *SmartAlphaDecoder) SetSeniorRateModelEventW3(w3l web3types.Log) (SmartAlphaSetSeniorRateModelEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaSetSeniorRateModelEvent{}, err
	}

	return d.SetSeniorRateModelEvent(l)
}

func (d *SmartAlphaDecoder) SetSeniorRateModelEvent(l types.Log) (SmartAlphaSetSeniorRateModelEvent, error) {
	var out SmartAlphaSetSeniorRateModelEvent
	if !d.IsSetSeniorRateModelEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SetSeniorRateModel", l)
	out.Raw = l
	return out, err
}

type SmartAlphaSetAccountingModelEvent struct {
	OldModel common.Address
	NewModel common.Address
	Raw      types.Log
}

func (d *SmartAlphaDecoder) SetAccountingModelEventID() common.Hash {
	return common.HexToHash("0xfd041dbbf45fb8b18dc79ae79199a37ddc52c7559ac44fdd5fefb9aa9c202843")
}

func (d *SmartAlphaDecoder) IsSetAccountingModelEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetAccountingModelEventID()
}

func (d *SmartAlphaDecoder) IsSetAccountingModelEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetAccountingModelEventID().String()
}

func (d *SmartAlphaDecoder) SetAccountingModelEventW3(w3l web3types.Log) (SmartAlphaSetAccountingModelEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartAlphaSetAccountingModelEvent{}, err
	}

	return d.SetAccountingModelEvent(l)
}

func (d *SmartAlphaDecoder) SetAccountingModelEvent(l types.Log) (SmartAlphaSetAccountingModelEvent, error) {
	var out SmartAlphaSetAccountingModelEvent
	if !d.IsSetAccountingModelEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SetAccountingModel", l)
	out.Raw = l
	return out, err
}
