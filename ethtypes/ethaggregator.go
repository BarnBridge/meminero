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

const EthaggregatorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_accessController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"confirmAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"phaseAggregators\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"phaseId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"proposeAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposedAggregator\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"proposedGetRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposedLatestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accessController\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

var Ethaggregator = NewEthaggregatorDecoder()

type EthaggregatorDecoder struct {
	*ethgen.Decoder
}

func NewEthaggregatorDecoder() *EthaggregatorDecoder {
	dec := ethgen.NewDecoder(EthaggregatorABI)
	return &EthaggregatorDecoder{
		dec,
	}
}

type EthaggregatorNewRoundEvent struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log
}

func (d *EthaggregatorDecoder) EthaggregatorNewRoundEventID() common.Hash {
	return common.HexToHash("0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271")
}

func (d *EthaggregatorDecoder) IsEthaggregatorNewRoundEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EthaggregatorNewRoundEventID()
}

func (d *EthaggregatorDecoder) EthaggregatorNewRoundEventW3(w3l web3types.Log) (EthaggregatorNewRoundEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EthaggregatorNewRoundEvent{}, err
	}

	return d.EthaggregatorNewRoundEvent(l)
}

func (d *EthaggregatorDecoder) EthaggregatorNewRoundEvent(l types.Log) (EthaggregatorNewRoundEvent, error) {
	var out EthaggregatorNewRoundEvent
	if !d.IsEthaggregatorNewRoundEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "NewRound", l)
	out.Raw = l
	return out, err
}

type EthaggregatorAnswerUpdatedEvent struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log
}

func (d *EthaggregatorDecoder) EthaggregatorAnswerUpdatedEventID() common.Hash {
	return common.HexToHash("0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f")
}

func (d *EthaggregatorDecoder) IsEthaggregatorAnswerUpdatedEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EthaggregatorAnswerUpdatedEventID()
}

func (d *EthaggregatorDecoder) EthaggregatorAnswerUpdatedEventW3(w3l web3types.Log) (EthaggregatorAnswerUpdatedEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EthaggregatorAnswerUpdatedEvent{}, err
	}

	return d.EthaggregatorAnswerUpdatedEvent(l)
}

func (d *EthaggregatorDecoder) EthaggregatorAnswerUpdatedEvent(l types.Log) (EthaggregatorAnswerUpdatedEvent, error) {
	var out EthaggregatorAnswerUpdatedEvent
	if !d.IsEthaggregatorAnswerUpdatedEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "AnswerUpdated", l)
	out.Raw = l
	return out, err
}

type EthaggregatorOwnershipTransferredEvent struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (d *EthaggregatorDecoder) EthaggregatorOwnershipTransferredEventID() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (d *EthaggregatorDecoder) IsEthaggregatorOwnershipTransferredEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EthaggregatorOwnershipTransferredEventID()
}

func (d *EthaggregatorDecoder) EthaggregatorOwnershipTransferredEventW3(w3l web3types.Log) (EthaggregatorOwnershipTransferredEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EthaggregatorOwnershipTransferredEvent{}, err
	}

	return d.EthaggregatorOwnershipTransferredEvent(l)
}

func (d *EthaggregatorDecoder) EthaggregatorOwnershipTransferredEvent(l types.Log) (EthaggregatorOwnershipTransferredEvent, error) {
	var out EthaggregatorOwnershipTransferredEvent
	if !d.IsEthaggregatorOwnershipTransferredEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "OwnershipTransferred", l)
	out.Raw = l
	return out, err
}

type EthaggregatorOwnershipTransferRequestedEvent struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (d *EthaggregatorDecoder) EthaggregatorOwnershipTransferRequestedEventID() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (d *EthaggregatorDecoder) IsEthaggregatorOwnershipTransferRequestedEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.EthaggregatorOwnershipTransferRequestedEventID()
}

func (d *EthaggregatorDecoder) EthaggregatorOwnershipTransferRequestedEventW3(w3l web3types.Log) (EthaggregatorOwnershipTransferRequestedEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return EthaggregatorOwnershipTransferRequestedEvent{}, err
	}

	return d.EthaggregatorOwnershipTransferRequestedEvent(l)
}

func (d *EthaggregatorDecoder) EthaggregatorOwnershipTransferRequestedEvent(l types.Log) (EthaggregatorOwnershipTransferRequestedEvent, error) {
	var out EthaggregatorOwnershipTransferRequestedEvent
	if !d.IsEthaggregatorOwnershipTransferRequestedEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "OwnershipTransferRequested", l)
	out.Raw = l
	return out, err
}
