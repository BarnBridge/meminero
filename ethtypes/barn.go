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

const BarnABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Delegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"to_newDelegatedPower\",\"type\":\"uint256\"}],\"name\":\"DelegatedPowerDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"to_newDelegatedPower\",\"type\":\"uint256\"}],\"name\":\"DelegatedPowerIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Lock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWithdrew\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLeft\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_LOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"balanceAtTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"bondStakedAtTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"delegatedPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"delegatedPowerAtTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"depositAndLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bond\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"}],\"name\":\"initBarn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"multiplierAtTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"multiplierOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"stakeAtTs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"}],\"internalType\":\"structLibBarnStorage.Stake\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userDelegatedTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userLockedUntil\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"votingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"votingPowerAtTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"diamondCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_functionSelector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facetAddress_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facetAddresses_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"_facetFunctionSelectors\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondLoupe.Facet[]\",\"name\":\"facets_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var Barn = NewBarnDecoder()

type BarnDecoder struct {
	*ethgen.Decoder
}

func NewBarnDecoder() *BarnDecoder {
	dec := ethgen.NewDecoder(BarnABI)
	return &BarnDecoder{
		dec,
	}
}

type IDiamondCutFacetCut struct {
	FacetAddress      common.Address
	Action            uint8
	FunctionSelectors [][4]byte
}

type BarnLockEvent struct {
	User      common.Address
	Timestamp *big.Int
	Raw       types.Log
}

func (d *BarnDecoder) BarnLockEventID() common.Hash {
	return common.HexToHash("0x625fed9875dada8643f2418b838ae0bc78d9a148a18eee4ee1979ff0f3f5d427")
}

func (d *BarnDecoder) IsBarnLockEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.BarnLockEventID()
}

func (d *BarnDecoder) IsBarnLockEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.BarnLockEventID().String()
}

func (d *BarnDecoder) BarnLockEventW3(w3l web3types.Log) (BarnLockEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return BarnLockEvent{}, err
	}

	return d.BarnLockEvent(l)
}

func (d *BarnDecoder) BarnLockEvent(l types.Log) (BarnLockEvent, error) {
	var out BarnLockEvent
	if !d.IsBarnLockEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "Lock", l)
	out.Raw = l
	return out, err
}

type BarnOwnershipTransferredEvent struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log
}

func (d *BarnDecoder) BarnOwnershipTransferredEventID() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (d *BarnDecoder) IsBarnOwnershipTransferredEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.BarnOwnershipTransferredEventID()
}

func (d *BarnDecoder) IsBarnOwnershipTransferredEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.BarnOwnershipTransferredEventID().String()
}

func (d *BarnDecoder) BarnOwnershipTransferredEventW3(w3l web3types.Log) (BarnOwnershipTransferredEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return BarnOwnershipTransferredEvent{}, err
	}

	return d.BarnOwnershipTransferredEvent(l)
}

func (d *BarnDecoder) BarnOwnershipTransferredEvent(l types.Log) (BarnOwnershipTransferredEvent, error) {
	var out BarnOwnershipTransferredEvent
	if !d.IsBarnOwnershipTransferredEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "OwnershipTransferred", l)
	out.Raw = l
	return out, err
}

type BarnDiamondCutEvent struct {
	DiamondCut []IDiamondCutFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log
}

func (d *BarnDecoder) BarnDiamondCutEventID() common.Hash {
	return common.HexToHash("0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673")
}

func (d *BarnDecoder) IsBarnDiamondCutEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.BarnDiamondCutEventID()
}

func (d *BarnDecoder) IsBarnDiamondCutEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.BarnDiamondCutEventID().String()
}

func (d *BarnDecoder) BarnDiamondCutEventW3(w3l web3types.Log) (BarnDiamondCutEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return BarnDiamondCutEvent{}, err
	}

	return d.BarnDiamondCutEvent(l)
}

func (d *BarnDecoder) BarnDiamondCutEvent(l types.Log) (BarnDiamondCutEvent, error) {
	var out BarnDiamondCutEvent
	if !d.IsBarnDiamondCutEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "DiamondCut", l)
	out.Raw = l
	return out, err
}

type BarnDepositEvent struct {
	User       common.Address
	Amount     *big.Int
	NewBalance *big.Int
	Raw        types.Log
}

func (d *BarnDecoder) BarnDepositEventID() common.Hash {
	return common.HexToHash("0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15")
}

func (d *BarnDecoder) IsBarnDepositEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.BarnDepositEventID()
}

func (d *BarnDecoder) IsBarnDepositEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.BarnDepositEventID().String()
}

func (d *BarnDecoder) BarnDepositEventW3(w3l web3types.Log) (BarnDepositEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return BarnDepositEvent{}, err
	}

	return d.BarnDepositEvent(l)
}

func (d *BarnDecoder) BarnDepositEvent(l types.Log) (BarnDepositEvent, error) {
	var out BarnDepositEvent
	if !d.IsBarnDepositEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "Deposit", l)
	out.Raw = l
	return out, err
}

type BarnDelegatedPowerIncreasedEvent struct {
	From                common.Address
	To                  common.Address
	Amount              *big.Int
	ToNewDelegatedPower *big.Int
	Raw                 types.Log
}

func (d *BarnDecoder) BarnDelegatedPowerIncreasedEventID() common.Hash {
	return common.HexToHash("0x9306546ca617a204223f7da51d942104c887cf8e53f8fd454af55a529aaa689a")
}

func (d *BarnDecoder) IsBarnDelegatedPowerIncreasedEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.BarnDelegatedPowerIncreasedEventID()
}

func (d *BarnDecoder) IsBarnDelegatedPowerIncreasedEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.BarnDelegatedPowerIncreasedEventID().String()
}

func (d *BarnDecoder) BarnDelegatedPowerIncreasedEventW3(w3l web3types.Log) (BarnDelegatedPowerIncreasedEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return BarnDelegatedPowerIncreasedEvent{}, err
	}

	return d.BarnDelegatedPowerIncreasedEvent(l)
}

func (d *BarnDecoder) BarnDelegatedPowerIncreasedEvent(l types.Log) (BarnDelegatedPowerIncreasedEvent, error) {
	var out BarnDelegatedPowerIncreasedEvent
	if !d.IsBarnDelegatedPowerIncreasedEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "DelegatedPowerIncreased", l)
	out.Raw = l
	return out, err
}

type BarnDelegateEvent struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (d *BarnDecoder) BarnDelegateEventID() common.Hash {
	return common.HexToHash("0xab7d75eccd27c9989942a3a6e4137e415df0ad90ec428751b16361f16fe8780f")
}

func (d *BarnDecoder) IsBarnDelegateEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.BarnDelegateEventID()
}

func (d *BarnDecoder) IsBarnDelegateEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.BarnDelegateEventID().String()
}

func (d *BarnDecoder) BarnDelegateEventW3(w3l web3types.Log) (BarnDelegateEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return BarnDelegateEvent{}, err
	}

	return d.BarnDelegateEvent(l)
}

func (d *BarnDecoder) BarnDelegateEvent(l types.Log) (BarnDelegateEvent, error) {
	var out BarnDelegateEvent
	if !d.IsBarnDelegateEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "Delegate", l)
	out.Raw = l
	return out, err
}

type BarnWithdrawEvent struct {
	User           common.Address
	AmountWithdrew *big.Int
	AmountLeft     *big.Int
	Raw            types.Log
}

func (d *BarnDecoder) BarnWithdrawEventID() common.Hash {
	return common.HexToHash("0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568")
}

func (d *BarnDecoder) IsBarnWithdrawEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.BarnWithdrawEventID()
}

func (d *BarnDecoder) IsBarnWithdrawEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.BarnWithdrawEventID().String()
}

func (d *BarnDecoder) BarnWithdrawEventW3(w3l web3types.Log) (BarnWithdrawEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return BarnWithdrawEvent{}, err
	}

	return d.BarnWithdrawEvent(l)
}

func (d *BarnDecoder) BarnWithdrawEvent(l types.Log) (BarnWithdrawEvent, error) {
	var out BarnWithdrawEvent
	if !d.IsBarnWithdrawEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "Withdraw", l)
	out.Raw = l
	return out, err
}

type BarnDelegatedPowerDecreasedEvent struct {
	From                common.Address
	To                  common.Address
	Amount              *big.Int
	ToNewDelegatedPower *big.Int
	Raw                 types.Log
}

func (d *BarnDecoder) BarnDelegatedPowerDecreasedEventID() common.Hash {
	return common.HexToHash("0xfb73cd22fb01f433ef312f758a708c1c7d1442ec871b9dd2546b3ec85a8b4e76")
}

func (d *BarnDecoder) IsBarnDelegatedPowerDecreasedEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.BarnDelegatedPowerDecreasedEventID()
}

func (d *BarnDecoder) IsBarnDelegatedPowerDecreasedEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.BarnDelegatedPowerDecreasedEventID().String()
}

func (d *BarnDecoder) BarnDelegatedPowerDecreasedEventW3(w3l web3types.Log) (BarnDelegatedPowerDecreasedEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return BarnDelegatedPowerDecreasedEvent{}, err
	}

	return d.BarnDelegatedPowerDecreasedEvent(l)
}

func (d *BarnDecoder) BarnDelegatedPowerDecreasedEvent(l types.Log) (BarnDelegatedPowerDecreasedEvent, error) {
	var out BarnDelegatedPowerDecreasedEvent
	if !d.IsBarnDelegatedPowerDecreasedEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "DelegatedPowerDecreased", l)
	out.Raw = l
	return out, err
}
