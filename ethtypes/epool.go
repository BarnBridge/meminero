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

const EpoolABI = "[{\"inputs\":[{\"internalType\":\"contractIController\",\"name\":\"_controller\",\"type\":\"address\"},{\"internalType\":\"contractIETokenFactory\",\"name\":\"_eTokenFactory\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"inverseRate\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"}],\"name\":\"AddedTranche\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cumulativeFeeA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cumulativeFeeB\",\"type\":\"uint256\"}],\"name\":\"CollectedFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"IssuedEToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deltaA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deltaB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rChange\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rDiv\",\"type\":\"uint256\"}],\"name\":\"RebalancedTranches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoveredToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"RedeemedEToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"inverseRate\",\"type\":\"bool\"}],\"name\":\"SetAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"SetController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"}],\"name\":\"SetFeeRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minRDiv\",\"type\":\"uint256\"}],\"name\":\"SetMinRDiv\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interval\",\"type\":\"uint256\"}],\"name\":\"SetRebalanceInterval\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE_RATE_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANCHE_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targetRatio\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"eTokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"eTokenSymbol\",\"type\":\"string\"}],\"name\":\"addTranche\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cumulativeFeeA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cumulativeFeeB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eTokenFactory\",\"outputs\":[{\"internalType\":\"contractIETokenFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"}],\"name\":\"getTranche\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIEToken\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sFactorE\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetRatio\",\"type\":\"uint256\"}],\"internalType\":\"structIEPool.Tranche\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTranches\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIEToken\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sFactorE\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetRatio\",\"type\":\"uint256\"}],\"internalType\":\"structIEPool.Tranche[]\",\"name\":\"_tranches\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issueExact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRebalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fracDelta\",\"type\":\"uint256\"}],\"name\":\"rebalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"deltaA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deltaB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rChange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rDiv\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebalanceInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebalanceMinRDiv\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeemExact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sFactorA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sFactorB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"inverseRate\",\"type\":\"bool\"}],\"name\":\"setAggregator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeRate\",\"type\":\"uint256\"}],\"name\":\"setFeeRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minRDiv\",\"type\":\"uint256\"}],\"name\":\"setMinRDiv\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"interval\",\"type\":\"uint256\"}],\"name\":\"setRebalanceInterval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenA\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenB\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tranches\",\"outputs\":[{\"internalType\":\"contractIEToken\",\"name\":\"eToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sFactorE\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetRatio\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tranchesByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

var Epool = NewEpoolDecoder()

type EpoolDecoder struct {
	*ethgen.Decoder
}

func NewEpoolDecoder() *EpoolDecoder {
	dec := ethgen.NewDecoder(EpoolABI)
	return &EpoolDecoder{
		dec,
	}
}

type EpoolSetMinRDivEvent struct {
	MinRDiv *big.Int
	Raw     types.Log
}

func (d *EpoolDecoder) EpoolSetMinRDivEventID() common.Hash {
	return common.HexToHash("0x4a7d6cd4901b6056e935ae8117764092378eea4896b4f247039c613b42c15c05")
}

func (d *EpoolDecoder) IsEpoolSetMinRDivEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EpoolSetMinRDivEventID()
}

func (d *EpoolDecoder) EpoolSetMinRDivEventW3(w3l web3types.Log) (EpoolSetMinRDivEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EpoolSetMinRDivEvent{}, err
	}

	return d.EpoolSetMinRDivEvent(l)
}

func (d *EpoolDecoder) EpoolSetMinRDivEvent(l types.Log) (EpoolSetMinRDivEvent, error) {
	var out EpoolSetMinRDivEvent
	if !d.IsEpoolSetMinRDivEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SetMinRDiv", l)
	out.Raw = l
	return out, err
}

type EpoolAddedTrancheEvent struct {
	EToken common.Address
	Raw    types.Log
}

func (d *EpoolDecoder) EpoolAddedTrancheEventID() common.Hash {
	return common.HexToHash("0x4f07ccfd1b8dd69c100ce0f0a3f368aa28cadc543706f2fa14f813177703a1a6")
}

func (d *EpoolDecoder) IsEpoolAddedTrancheEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EpoolAddedTrancheEventID()
}

func (d *EpoolDecoder) EpoolAddedTrancheEventW3(w3l web3types.Log) (EpoolAddedTrancheEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EpoolAddedTrancheEvent{}, err
	}

	return d.EpoolAddedTrancheEvent(l)
}

func (d *EpoolDecoder) EpoolAddedTrancheEvent(l types.Log) (EpoolAddedTrancheEvent, error) {
	var out EpoolAddedTrancheEvent
	if !d.IsEpoolAddedTrancheEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "AddedTranche", l)
	out.Raw = l
	return out, err
}

type EpoolSetControllerEvent struct {
	Controller common.Address
	Raw        types.Log
}

func (d *EpoolDecoder) EpoolSetControllerEventID() common.Hash {
	return common.HexToHash("0x4ff638452bbf33c012645d18ae6f05515ff5f2d1dfb0cece8cbf018c60903f70")
}

func (d *EpoolDecoder) IsEpoolSetControllerEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EpoolSetControllerEventID()
}

func (d *EpoolDecoder) EpoolSetControllerEventW3(w3l web3types.Log) (EpoolSetControllerEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EpoolSetControllerEvent{}, err
	}

	return d.EpoolSetControllerEvent(l)
}

func (d *EpoolDecoder) EpoolSetControllerEvent(l types.Log) (EpoolSetControllerEvent, error) {
	var out EpoolSetControllerEvent
	if !d.IsEpoolSetControllerEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SetController", l)
	out.Raw = l
	return out, err
}

type EpoolSetFeeRateEvent struct {
	FeeRate *big.Int
	Raw     types.Log
}

func (d *EpoolDecoder) EpoolSetFeeRateEventID() common.Hash {
	return common.HexToHash("0x6717373928cccf59cc9912055cfa8db86e7085b95c94c15862b121114aa333be")
}

func (d *EpoolDecoder) IsEpoolSetFeeRateEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EpoolSetFeeRateEventID()
}

func (d *EpoolDecoder) EpoolSetFeeRateEventW3(w3l web3types.Log) (EpoolSetFeeRateEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EpoolSetFeeRateEvent{}, err
	}

	return d.EpoolSetFeeRateEvent(l)
}

func (d *EpoolDecoder) EpoolSetFeeRateEvent(l types.Log) (EpoolSetFeeRateEvent, error) {
	var out EpoolSetFeeRateEvent
	if !d.IsEpoolSetFeeRateEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SetFeeRate", l)
	out.Raw = l
	return out, err
}

type EpoolRedeemedETokenEvent struct {
	EToken  common.Address
	Amount  *big.Int
	AmountA *big.Int
	AmountB *big.Int
	User    common.Address
	Raw     types.Log
}

func (d *EpoolDecoder) EpoolRedeemedETokenEventID() common.Hash {
	return common.HexToHash("0x6ccf4b3c348e324c7a3cc286369614139a347bbff3f2315520c87ce795c50dde")
}

func (d *EpoolDecoder) IsEpoolRedeemedETokenEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EpoolRedeemedETokenEventID()
}

func (d *EpoolDecoder) EpoolRedeemedETokenEventW3(w3l web3types.Log) (EpoolRedeemedETokenEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EpoolRedeemedETokenEvent{}, err
	}

	return d.EpoolRedeemedETokenEvent(l)
}

func (d *EpoolDecoder) EpoolRedeemedETokenEvent(l types.Log) (EpoolRedeemedETokenEvent, error) {
	var out EpoolRedeemedETokenEvent
	if !d.IsEpoolRedeemedETokenEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "RedeemedEToken", l)
	out.Raw = l
	return out, err
}

type EpoolRecoveredTokenEvent struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log
}

func (d *EpoolDecoder) EpoolRecoveredTokenEventID() common.Hash {
	return common.HexToHash("0x6de8b63479ce07cf2dfc515e20a5c88a3a5bab6cbd76f753388b77e244ca7071")
}

func (d *EpoolDecoder) IsEpoolRecoveredTokenEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EpoolRecoveredTokenEventID()
}

func (d *EpoolDecoder) EpoolRecoveredTokenEventW3(w3l web3types.Log) (EpoolRecoveredTokenEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EpoolRecoveredTokenEvent{}, err
	}

	return d.EpoolRecoveredTokenEvent(l)
}

func (d *EpoolDecoder) EpoolRecoveredTokenEvent(l types.Log) (EpoolRecoveredTokenEvent, error) {
	var out EpoolRecoveredTokenEvent
	if !d.IsEpoolRecoveredTokenEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "RecoveredToken", l)
	out.Raw = l
	return out, err
}

type EpoolIssuedETokenEvent struct {
	EToken  common.Address
	Amount  *big.Int
	AmountA *big.Int
	AmountB *big.Int
	User    common.Address
	Raw     types.Log
}

func (d *EpoolDecoder) EpoolIssuedETokenEventID() common.Hash {
	return common.HexToHash("0x99b554b7dd396926e9ca4dc2f8349b638f196fb693daf374c850139debc63447")
}

func (d *EpoolDecoder) IsEpoolIssuedETokenEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EpoolIssuedETokenEventID()
}

func (d *EpoolDecoder) EpoolIssuedETokenEventW3(w3l web3types.Log) (EpoolIssuedETokenEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EpoolIssuedETokenEvent{}, err
	}

	return d.EpoolIssuedETokenEvent(l)
}

func (d *EpoolDecoder) EpoolIssuedETokenEvent(l types.Log) (EpoolIssuedETokenEvent, error) {
	var out EpoolIssuedETokenEvent
	if !d.IsEpoolIssuedETokenEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "IssuedEToken", l)
	out.Raw = l
	return out, err
}

type EpoolSetAggregatorEvent struct {
	Aggregator  common.Address
	InverseRate bool
	Raw         types.Log
}

func (d *EpoolDecoder) EpoolSetAggregatorEventID() common.Hash {
	return common.HexToHash("0x9aaad5d73fc4de1befd3e790b855dfdc6363f068e93abfdf01ad70681d31d0ce")
}

func (d *EpoolDecoder) IsEpoolSetAggregatorEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EpoolSetAggregatorEventID()
}

func (d *EpoolDecoder) EpoolSetAggregatorEventW3(w3l web3types.Log) (EpoolSetAggregatorEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EpoolSetAggregatorEvent{}, err
	}

	return d.EpoolSetAggregatorEvent(l)
}

func (d *EpoolDecoder) EpoolSetAggregatorEvent(l types.Log) (EpoolSetAggregatorEvent, error) {
	var out EpoolSetAggregatorEvent
	if !d.IsEpoolSetAggregatorEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SetAggregator", l)
	out.Raw = l
	return out, err
}

type EpoolCollectedFeesEvent struct {
	CumulativeFeeA *big.Int
	CumulativeFeeB *big.Int
	Raw            types.Log
}

func (d *EpoolDecoder) EpoolCollectedFeesEventID() common.Hash {
	return common.HexToHash("0xaeb342f3c261bafef9d4f2cccd5ada643628fa7f7fadb7035ee1e91c2385b873")
}

func (d *EpoolDecoder) IsEpoolCollectedFeesEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EpoolCollectedFeesEventID()
}

func (d *EpoolDecoder) EpoolCollectedFeesEventW3(w3l web3types.Log) (EpoolCollectedFeesEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EpoolCollectedFeesEvent{}, err
	}

	return d.EpoolCollectedFeesEvent(l)
}

func (d *EpoolDecoder) EpoolCollectedFeesEvent(l types.Log) (EpoolCollectedFeesEvent, error) {
	var out EpoolCollectedFeesEvent
	if !d.IsEpoolCollectedFeesEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "CollectedFees", l)
	out.Raw = l
	return out, err
}

type EpoolRebalancedTranchesEvent struct {
	DeltaA  *big.Int
	DeltaB  *big.Int
	RChange *big.Int
	RDiv    *big.Int
	Raw     types.Log
}

func (d *EpoolDecoder) EpoolRebalancedTranchesEventID() common.Hash {
	return common.HexToHash("0xe219e81e936fbe5bc0195b0cc0755ef3e79c6910fc4398345d8b4c6c267fd40f")
}

func (d *EpoolDecoder) IsEpoolRebalancedTranchesEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EpoolRebalancedTranchesEventID()
}

func (d *EpoolDecoder) EpoolRebalancedTranchesEventW3(w3l web3types.Log) (EpoolRebalancedTranchesEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EpoolRebalancedTranchesEvent{}, err
	}

	return d.EpoolRebalancedTranchesEvent(l)
}

func (d *EpoolDecoder) EpoolRebalancedTranchesEvent(l types.Log) (EpoolRebalancedTranchesEvent, error) {
	var out EpoolRebalancedTranchesEvent
	if !d.IsEpoolRebalancedTranchesEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "RebalancedTranches", l)
	out.Raw = l
	return out, err
}

type EpoolSetRebalanceIntervalEvent struct {
	Interval *big.Int
	Raw      types.Log
}

func (d *EpoolDecoder) EpoolSetRebalanceIntervalEventID() common.Hash {
	return common.HexToHash("0xe92aa3ac048565d1668fe6ffad28e03b8cbeed2210cd1fdef353d88d7f8e694b")
}

func (d *EpoolDecoder) IsEpoolSetRebalanceIntervalEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EpoolSetRebalanceIntervalEventID()
}

func (d *EpoolDecoder) EpoolSetRebalanceIntervalEventW3(w3l web3types.Log) (EpoolSetRebalanceIntervalEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EpoolSetRebalanceIntervalEvent{}, err
	}

	return d.EpoolSetRebalanceIntervalEvent(l)
}

func (d *EpoolDecoder) EpoolSetRebalanceIntervalEvent(l types.Log) (EpoolSetRebalanceIntervalEvent, error) {
	var out EpoolSetRebalanceIntervalEvent
	if !d.IsEpoolSetRebalanceIntervalEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "SetRebalanceInterval", l)
	out.Raw = l
	return out, err
}
