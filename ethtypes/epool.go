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

const EPoolABI = "[{\"inputs\":[{\"internalType\":\"contractIController\",\"name\":\"_controller\",\"type\":\"address\"},{\"internalType\":\"contractIETokenFactory\",\"name\":\"_eTokenFactory\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"inverseRate\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"}],\"name\":\"AddedTranche\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cumulativeFeeA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cumulativeFeeB\",\"type\":\"uint256\"}],\"name\":\"CollectedFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"IssuedEToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deltaA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deltaB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rChange\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rDiv\",\"type\":\"uint256\"}],\"name\":\"RebalancedTranches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoveredToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"RedeemedEToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"inverseRate\",\"type\":\"bool\"}],\"name\":\"SetAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"SetController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"}],\"name\":\"SetFeeRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minRDiv\",\"type\":\"uint256\"}],\"name\":\"SetMinRDiv\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interval\",\"type\":\"uint256\"}],\"name\":\"SetRebalanceInterval\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE_RATE_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANCHE_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targetRatio\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"eTokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"eTokenSymbol\",\"type\":\"string\"}],\"name\":\"addTranche\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cumulativeFeeA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cumulativeFeeB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eTokenFactory\",\"outputs\":[{\"internalType\":\"contractIETokenFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"}],\"name\":\"getTranche\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIEToken\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sFactorE\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetRatio\",\"type\":\"uint256\"}],\"internalType\":\"structIEPool.Tranche\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTranches\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIEToken\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sFactorE\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetRatio\",\"type\":\"uint256\"}],\"internalType\":\"structIEPool.Tranche[]\",\"name\":\"_tranches\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issueExact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRebalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fracDelta\",\"type\":\"uint256\"}],\"name\":\"rebalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"deltaA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deltaB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rChange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rDiv\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebalanceInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebalanceMinRDiv\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeemExact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sFactorA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sFactorB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"inverseRate\",\"type\":\"bool\"}],\"name\":\"setAggregator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeRate\",\"type\":\"uint256\"}],\"name\":\"setFeeRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minRDiv\",\"type\":\"uint256\"}],\"name\":\"setMinRDiv\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"interval\",\"type\":\"uint256\"}],\"name\":\"setRebalanceInterval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenA\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenB\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tranches\",\"outputs\":[{\"internalType\":\"contractIEToken\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sFactorE\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetRatio\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tranchesByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

var EPool = NewEPoolDecoder()

type EPoolDecoder struct {
	*ethgen.Decoder
}

func NewEPoolDecoder() *EPoolDecoder {
	dec := ethgen.NewDecoder(EPoolABI)
	return &EPoolDecoder{
		dec,
	}
}

type EPoolSetMinRDivEvent struct {
	MinRDiv *big.Int
	Raw     types.Log
}

func (e *EPoolSetMinRDivEvent) MinRDivDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.MinRDiv, exp)
}

func (d *EPoolDecoder) SetMinRDivEventID() common.Hash {
	return common.HexToHash("0x4a7d6cd4901b6056e935ae8117764092378eea4896b4f247039c613b42c15c05")
}

func (d *EPoolDecoder) IsSetMinRDivEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetMinRDivEventID()
}

func (d *EPoolDecoder) IsSetMinRDivEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetMinRDivEventID().String()
}

func (d *EPoolDecoder) SetMinRDivEventW3(w3l web3types.Log) (EPoolSetMinRDivEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EPoolSetMinRDivEvent{}, err
	}

	return d.SetMinRDivEvent(l)
}

func (d *EPoolDecoder) SetMinRDivEvent(l types.Log) (EPoolSetMinRDivEvent, error) {
	var out EPoolSetMinRDivEvent
	if !d.IsSetMinRDivEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SetMinRDiv", l)
	out.Raw = l
	return out, err
}

type EPoolAddedTrancheEvent struct {
	EToken common.Address
	Raw    types.Log
}

func (d *EPoolDecoder) AddedTrancheEventID() common.Hash {
	return common.HexToHash("0x4f07ccfd1b8dd69c100ce0f0a3f368aa28cadc543706f2fa14f813177703a1a6")
}

func (d *EPoolDecoder) IsAddedTrancheEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.AddedTrancheEventID()
}

func (d *EPoolDecoder) IsAddedTrancheEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.AddedTrancheEventID().String()
}

func (d *EPoolDecoder) AddedTrancheEventW3(w3l web3types.Log) (EPoolAddedTrancheEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EPoolAddedTrancheEvent{}, err
	}

	return d.AddedTrancheEvent(l)
}

func (d *EPoolDecoder) AddedTrancheEvent(l types.Log) (EPoolAddedTrancheEvent, error) {
	var out EPoolAddedTrancheEvent
	if !d.IsAddedTrancheEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "AddedTranche", l)
	out.Raw = l
	return out, err
}

type EPoolSetControllerEvent struct {
	Controller common.Address
	Raw        types.Log
}

func (d *EPoolDecoder) SetControllerEventID() common.Hash {
	return common.HexToHash("0x4ff638452bbf33c012645d18ae6f05515ff5f2d1dfb0cece8cbf018c60903f70")
}

func (d *EPoolDecoder) IsSetControllerEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetControllerEventID()
}

func (d *EPoolDecoder) IsSetControllerEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetControllerEventID().String()
}

func (d *EPoolDecoder) SetControllerEventW3(w3l web3types.Log) (EPoolSetControllerEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EPoolSetControllerEvent{}, err
	}

	return d.SetControllerEvent(l)
}

func (d *EPoolDecoder) SetControllerEvent(l types.Log) (EPoolSetControllerEvent, error) {
	var out EPoolSetControllerEvent
	if !d.IsSetControllerEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SetController", l)
	out.Raw = l
	return out, err
}

type EPoolSetFeeRateEvent struct {
	FeeRate *big.Int
	Raw     types.Log
}

func (e *EPoolSetFeeRateEvent) FeeRateDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.FeeRate, exp)
}

func (d *EPoolDecoder) SetFeeRateEventID() common.Hash {
	return common.HexToHash("0x6717373928cccf59cc9912055cfa8db86e7085b95c94c15862b121114aa333be")
}

func (d *EPoolDecoder) IsSetFeeRateEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetFeeRateEventID()
}

func (d *EPoolDecoder) IsSetFeeRateEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetFeeRateEventID().String()
}

func (d *EPoolDecoder) SetFeeRateEventW3(w3l web3types.Log) (EPoolSetFeeRateEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EPoolSetFeeRateEvent{}, err
	}

	return d.SetFeeRateEvent(l)
}

func (d *EPoolDecoder) SetFeeRateEvent(l types.Log) (EPoolSetFeeRateEvent, error) {
	var out EPoolSetFeeRateEvent
	if !d.IsSetFeeRateEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SetFeeRate", l)
	out.Raw = l
	return out, err
}

type EPoolRedeemedETokenEvent struct {
	EToken  common.Address
	Amount  *big.Int
	AmountA *big.Int
	AmountB *big.Int
	User    common.Address
	Raw     types.Log
}

func (e *EPoolRedeemedETokenEvent) AmountDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.Amount, exp)
}

func (e *EPoolRedeemedETokenEvent) AmountADecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.AmountA, exp)
}

func (e *EPoolRedeemedETokenEvent) AmountBDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.AmountB, exp)
}

func (d *EPoolDecoder) RedeemedETokenEventID() common.Hash {
	return common.HexToHash("0x6ccf4b3c348e324c7a3cc286369614139a347bbff3f2315520c87ce795c50dde")
}

func (d *EPoolDecoder) IsRedeemedETokenEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.RedeemedETokenEventID()
}

func (d *EPoolDecoder) IsRedeemedETokenEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.RedeemedETokenEventID().String()
}

func (d *EPoolDecoder) RedeemedETokenEventW3(w3l web3types.Log) (EPoolRedeemedETokenEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EPoolRedeemedETokenEvent{}, err
	}

	return d.RedeemedETokenEvent(l)
}

func (d *EPoolDecoder) RedeemedETokenEvent(l types.Log) (EPoolRedeemedETokenEvent, error) {
	var out EPoolRedeemedETokenEvent
	if !d.IsRedeemedETokenEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "RedeemedEToken", l)
	out.Raw = l
	return out, err
}

type EPoolRecoveredTokenEvent struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log
}

func (e *EPoolRecoveredTokenEvent) AmountDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.Amount, exp)
}

func (d *EPoolDecoder) RecoveredTokenEventID() common.Hash {
	return common.HexToHash("0x6de8b63479ce07cf2dfc515e20a5c88a3a5bab6cbd76f753388b77e244ca7071")
}

func (d *EPoolDecoder) IsRecoveredTokenEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.RecoveredTokenEventID()
}

func (d *EPoolDecoder) IsRecoveredTokenEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.RecoveredTokenEventID().String()
}

func (d *EPoolDecoder) RecoveredTokenEventW3(w3l web3types.Log) (EPoolRecoveredTokenEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EPoolRecoveredTokenEvent{}, err
	}

	return d.RecoveredTokenEvent(l)
}

func (d *EPoolDecoder) RecoveredTokenEvent(l types.Log) (EPoolRecoveredTokenEvent, error) {
	var out EPoolRecoveredTokenEvent
	if !d.IsRecoveredTokenEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "RecoveredToken", l)
	out.Raw = l
	return out, err
}

type EPoolIssuedETokenEvent struct {
	EToken  common.Address
	Amount  *big.Int
	AmountA *big.Int
	AmountB *big.Int
	User    common.Address
	Raw     types.Log
}

func (e *EPoolIssuedETokenEvent) AmountDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.Amount, exp)
}

func (e *EPoolIssuedETokenEvent) AmountADecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.AmountA, exp)
}

func (e *EPoolIssuedETokenEvent) AmountBDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.AmountB, exp)
}

func (d *EPoolDecoder) IssuedETokenEventID() common.Hash {
	return common.HexToHash("0x99b554b7dd396926e9ca4dc2f8349b638f196fb693daf374c850139debc63447")
}

func (d *EPoolDecoder) IsIssuedETokenEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.IssuedETokenEventID()
}

func (d *EPoolDecoder) IsIssuedETokenEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.IssuedETokenEventID().String()
}

func (d *EPoolDecoder) IssuedETokenEventW3(w3l web3types.Log) (EPoolIssuedETokenEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EPoolIssuedETokenEvent{}, err
	}

	return d.IssuedETokenEvent(l)
}

func (d *EPoolDecoder) IssuedETokenEvent(l types.Log) (EPoolIssuedETokenEvent, error) {
	var out EPoolIssuedETokenEvent
	if !d.IsIssuedETokenEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "IssuedEToken", l)
	out.Raw = l
	return out, err
}

type EPoolSetAggregatorEvent struct {
	Aggregator  common.Address
	InverseRate bool
	Raw         types.Log
}

func (d *EPoolDecoder) SetAggregatorEventID() common.Hash {
	return common.HexToHash("0x9aaad5d73fc4de1befd3e790b855dfdc6363f068e93abfdf01ad70681d31d0ce")
}

func (d *EPoolDecoder) IsSetAggregatorEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetAggregatorEventID()
}

func (d *EPoolDecoder) IsSetAggregatorEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetAggregatorEventID().String()
}

func (d *EPoolDecoder) SetAggregatorEventW3(w3l web3types.Log) (EPoolSetAggregatorEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EPoolSetAggregatorEvent{}, err
	}

	return d.SetAggregatorEvent(l)
}

func (d *EPoolDecoder) SetAggregatorEvent(l types.Log) (EPoolSetAggregatorEvent, error) {
	var out EPoolSetAggregatorEvent
	if !d.IsSetAggregatorEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SetAggregator", l)
	out.Raw = l
	return out, err
}

type EPoolCollectedFeesEvent struct {
	CumulativeFeeA *big.Int
	CumulativeFeeB *big.Int
	Raw            types.Log
}

func (e *EPoolCollectedFeesEvent) CumulativeFeeADecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.CumulativeFeeA, exp)
}

func (e *EPoolCollectedFeesEvent) CumulativeFeeBDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.CumulativeFeeB, exp)
}

func (d *EPoolDecoder) CollectedFeesEventID() common.Hash {
	return common.HexToHash("0xaeb342f3c261bafef9d4f2cccd5ada643628fa7f7fadb7035ee1e91c2385b873")
}

func (d *EPoolDecoder) IsCollectedFeesEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.CollectedFeesEventID()
}

func (d *EPoolDecoder) IsCollectedFeesEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.CollectedFeesEventID().String()
}

func (d *EPoolDecoder) CollectedFeesEventW3(w3l web3types.Log) (EPoolCollectedFeesEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EPoolCollectedFeesEvent{}, err
	}

	return d.CollectedFeesEvent(l)
}

func (d *EPoolDecoder) CollectedFeesEvent(l types.Log) (EPoolCollectedFeesEvent, error) {
	var out EPoolCollectedFeesEvent
	if !d.IsCollectedFeesEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "CollectedFees", l)
	out.Raw = l
	return out, err
}

type EPoolRebalancedTranchesEvent struct {
	DeltaA  *big.Int
	DeltaB  *big.Int
	RChange *big.Int
	RDiv    *big.Int
	Raw     types.Log
}

func (e *EPoolRebalancedTranchesEvent) DeltaADecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.DeltaA, exp)
}

func (e *EPoolRebalancedTranchesEvent) DeltaBDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.DeltaB, exp)
}

func (e *EPoolRebalancedTranchesEvent) RChangeDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.RChange, exp)
}

func (e *EPoolRebalancedTranchesEvent) RDivDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.RDiv, exp)
}

func (d *EPoolDecoder) RebalancedTranchesEventID() common.Hash {
	return common.HexToHash("0xe219e81e936fbe5bc0195b0cc0755ef3e79c6910fc4398345d8b4c6c267fd40f")
}

func (d *EPoolDecoder) IsRebalancedTranchesEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.RebalancedTranchesEventID()
}

func (d *EPoolDecoder) IsRebalancedTranchesEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.RebalancedTranchesEventID().String()
}

func (d *EPoolDecoder) RebalancedTranchesEventW3(w3l web3types.Log) (EPoolRebalancedTranchesEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EPoolRebalancedTranchesEvent{}, err
	}

	return d.RebalancedTranchesEvent(l)
}

func (d *EPoolDecoder) RebalancedTranchesEvent(l types.Log) (EPoolRebalancedTranchesEvent, error) {
	var out EPoolRebalancedTranchesEvent
	if !d.IsRebalancedTranchesEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "RebalancedTranches", l)
	out.Raw = l
	return out, err
}

type EPoolSetRebalanceIntervalEvent struct {
	Interval *big.Int
	Raw      types.Log
}

func (e *EPoolSetRebalanceIntervalEvent) IntervalDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.Interval, exp)
}

func (d *EPoolDecoder) SetRebalanceIntervalEventID() common.Hash {
	return common.HexToHash("0xe92aa3ac048565d1668fe6ffad28e03b8cbeed2210cd1fdef353d88d7f8e694b")
}

func (d *EPoolDecoder) IsSetRebalanceIntervalEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetRebalanceIntervalEventID()
}

func (d *EPoolDecoder) IsSetRebalanceIntervalEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.SetRebalanceIntervalEventID().String()
}

func (d *EPoolDecoder) SetRebalanceIntervalEventW3(w3l web3types.Log) (EPoolSetRebalanceIntervalEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EPoolSetRebalanceIntervalEvent{}, err
	}

	return d.SetRebalanceIntervalEvent(l)
}

func (d *EPoolDecoder) SetRebalanceIntervalEvent(l types.Log) (EPoolSetRebalanceIntervalEvent, error) {
	var out EPoolSetRebalanceIntervalEvent
	if !d.IsSetRebalanceIntervalEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SetRebalanceInterval", l)
	out.Raw = l
	return out, err
}
