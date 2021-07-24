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

const GovernanceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AbrogationProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AbrogationProposalStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"name\":\"AbrogationProposalVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"AbrogationProposalVoteCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"VoteCanceled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"abrogateProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"abrogationProposal_cancelVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"}],\"name\":\"abrogationProposal_castVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"abrogationProposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptanceThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"cancelProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"cancelVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"}],\"name\":\"castVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getAbrogationProposalReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"}],\"internalType\":\"structGovernance.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getActions\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalParameters\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"warmUpDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"queueDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gracePeriodDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptanceThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structGovernance.ProposalParameters\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalQuorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"}],\"internalType\":\"structGovernance.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gracePeriodDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"barnAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"latestProposalIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minQuorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canceled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"warmUpDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"queueDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gracePeriodDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptanceThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structGovernance.ProposalParameters\",\"name\":\"parameters\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"queue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"queuedTransactions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"setAcceptanceThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setActiveDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setGracePeriodDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quorum\",\"type\":\"uint256\"}],\"name\":\"setMinQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setQueueDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setWarmUpDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"startAbrogationProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumGovernance.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"warmUpDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

var Governance = NewGovernanceDecoder()

type GovernanceDecoder struct {
	*ethgen.Decoder
}

func NewGovernanceDecoder() *GovernanceDecoder {
	dec := ethgen.NewDecoder(GovernanceABI)
	return &GovernanceDecoder{
		dec,
	}
}

type GovernanceVoteCanceledEvent struct {
	ProposalId *big.Int
	User       common.Address
	Raw        types.Log
}

func (d *GovernanceDecoder) GovernanceVoteCanceledEventID() common.Hash {
	return common.HexToHash("0x12beef84830227673717dd5522ee1228a8004e88dc2678d8740f582264efb2b6")
}

func (d *GovernanceDecoder) IsGovernanceVoteCanceledEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceVoteCanceledEventID()
}

func (d *GovernanceDecoder) IsGovernanceVoteCanceledEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceVoteCanceledEventID().String()
}

func (d *GovernanceDecoder) GovernanceVoteCanceledEventW3(w3l web3types.Log) (GovernanceVoteCanceledEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return GovernanceVoteCanceledEvent{}, err
	}

	return d.GovernanceVoteCanceledEvent(l)
}

func (d *GovernanceDecoder) GovernanceVoteCanceledEvent(l types.Log) (GovernanceVoteCanceledEvent, error) {
	var out GovernanceVoteCanceledEvent
	if !d.IsGovernanceVoteCanceledEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "VoteCanceled", l)
	out.Raw = l
	return out, err
}

type GovernanceProposalCanceledEvent struct {
	ProposalId *big.Int
	Caller     common.Address
	Raw        types.Log
}

func (d *GovernanceDecoder) GovernanceProposalCanceledEventID() common.Hash {
	return common.HexToHash("0x253042c67143aeb6d431bb762d75e5905f18fa7850b7b9edb31fedb7c362d7e8")
}

func (d *GovernanceDecoder) IsGovernanceProposalCanceledEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceProposalCanceledEventID()
}

func (d *GovernanceDecoder) IsGovernanceProposalCanceledEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceProposalCanceledEventID().String()
}

func (d *GovernanceDecoder) GovernanceProposalCanceledEventW3(w3l web3types.Log) (GovernanceProposalCanceledEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return GovernanceProposalCanceledEvent{}, err
	}

	return d.GovernanceProposalCanceledEvent(l)
}

func (d *GovernanceDecoder) GovernanceProposalCanceledEvent(l types.Log) (GovernanceProposalCanceledEvent, error) {
	var out GovernanceProposalCanceledEvent
	if !d.IsGovernanceProposalCanceledEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "ProposalCanceled", l)
	out.Raw = l
	return out, err
}

type GovernanceAbrogationProposalStartedEvent struct {
	ProposalId *big.Int
	Caller     common.Address
	Raw        types.Log
}

func (d *GovernanceDecoder) GovernanceAbrogationProposalStartedEventID() common.Hash {
	return common.HexToHash("0x27eba018e1c52b84f732fe4d806fd9750c60752f1d37e7f70bcb4cbec65b1c6a")
}

func (d *GovernanceDecoder) IsGovernanceAbrogationProposalStartedEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceAbrogationProposalStartedEventID()
}

func (d *GovernanceDecoder) IsGovernanceAbrogationProposalStartedEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceAbrogationProposalStartedEventID().String()
}

func (d *GovernanceDecoder) GovernanceAbrogationProposalStartedEventW3(w3l web3types.Log) (GovernanceAbrogationProposalStartedEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return GovernanceAbrogationProposalStartedEvent{}, err
	}

	return d.GovernanceAbrogationProposalStartedEvent(l)
}

func (d *GovernanceDecoder) GovernanceAbrogationProposalStartedEvent(l types.Log) (GovernanceAbrogationProposalStartedEvent, error) {
	var out GovernanceAbrogationProposalStartedEvent
	if !d.IsGovernanceAbrogationProposalStartedEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "AbrogationProposalStarted", l)
	out.Raw = l
	return out, err
}

type GovernanceAbrogationProposalVoteCancelledEvent struct {
	ProposalId *big.Int
	User       common.Address
	Raw        types.Log
}

func (d *GovernanceDecoder) GovernanceAbrogationProposalVoteCancelledEventID() common.Hash {
	return common.HexToHash("0x5e8ee24f838173ed2ae7989835696f6e11945ac8fbc5259aef01cc4d7f0d4920")
}

func (d *GovernanceDecoder) IsGovernanceAbrogationProposalVoteCancelledEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceAbrogationProposalVoteCancelledEventID()
}

func (d *GovernanceDecoder) IsGovernanceAbrogationProposalVoteCancelledEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceAbrogationProposalVoteCancelledEventID().String()
}

func (d *GovernanceDecoder) GovernanceAbrogationProposalVoteCancelledEventW3(w3l web3types.Log) (GovernanceAbrogationProposalVoteCancelledEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return GovernanceAbrogationProposalVoteCancelledEvent{}, err
	}

	return d.GovernanceAbrogationProposalVoteCancelledEvent(l)
}

func (d *GovernanceDecoder) GovernanceAbrogationProposalVoteCancelledEvent(l types.Log) (GovernanceAbrogationProposalVoteCancelledEvent, error) {
	var out GovernanceAbrogationProposalVoteCancelledEvent
	if !d.IsGovernanceAbrogationProposalVoteCancelledEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "AbrogationProposalVoteCancelled", l)
	out.Raw = l
	return out, err
}

type GovernanceAbrogationProposalExecutedEvent struct {
	ProposalId *big.Int
	Caller     common.Address
	Raw        types.Log
}

func (d *GovernanceDecoder) GovernanceAbrogationProposalExecutedEventID() common.Hash {
	return common.HexToHash("0x6d7acd63bebeaf524f1761a88687951f776fc7c182205f820424c7fb572c7235")
}

func (d *GovernanceDecoder) IsGovernanceAbrogationProposalExecutedEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceAbrogationProposalExecutedEventID()
}

func (d *GovernanceDecoder) IsGovernanceAbrogationProposalExecutedEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceAbrogationProposalExecutedEventID().String()
}

func (d *GovernanceDecoder) GovernanceAbrogationProposalExecutedEventW3(w3l web3types.Log) (GovernanceAbrogationProposalExecutedEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return GovernanceAbrogationProposalExecutedEvent{}, err
	}

	return d.GovernanceAbrogationProposalExecutedEvent(l)
}

func (d *GovernanceDecoder) GovernanceAbrogationProposalExecutedEvent(l types.Log) (GovernanceAbrogationProposalExecutedEvent, error) {
	var out GovernanceAbrogationProposalExecutedEvent
	if !d.IsGovernanceAbrogationProposalExecutedEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "AbrogationProposalExecuted", l)
	out.Raw = l
	return out, err
}

type GovernanceAbrogationProposalVoteEvent struct {
	ProposalId *big.Int
	User       common.Address
	Support    bool
	Power      *big.Int
	Raw        types.Log
}

func (d *GovernanceDecoder) GovernanceAbrogationProposalVoteEventID() common.Hash {
	return common.HexToHash("0x80f2ad7e3e83d197670402663f224adb2f649967b9629c67dcfafa40c94d30f9")
}

func (d *GovernanceDecoder) IsGovernanceAbrogationProposalVoteEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceAbrogationProposalVoteEventID()
}

func (d *GovernanceDecoder) IsGovernanceAbrogationProposalVoteEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceAbrogationProposalVoteEventID().String()
}

func (d *GovernanceDecoder) GovernanceAbrogationProposalVoteEventW3(w3l web3types.Log) (GovernanceAbrogationProposalVoteEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return GovernanceAbrogationProposalVoteEvent{}, err
	}

	return d.GovernanceAbrogationProposalVoteEvent(l)
}

func (d *GovernanceDecoder) GovernanceAbrogationProposalVoteEvent(l types.Log) (GovernanceAbrogationProposalVoteEvent, error) {
	var out GovernanceAbrogationProposalVoteEvent
	if !d.IsGovernanceAbrogationProposalVoteEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "AbrogationProposalVote", l)
	out.Raw = l
	return out, err
}

type GovernanceVoteEvent struct {
	ProposalId *big.Int
	User       common.Address
	Support    bool
	Power      *big.Int
	Raw        types.Log
}

func (d *GovernanceDecoder) GovernanceVoteEventID() common.Hash {
	return common.HexToHash("0x88d35328232823f54954b6627e9f732371656f6daa40cb1b01b27dc7875a7b47")
}

func (d *GovernanceDecoder) IsGovernanceVoteEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceVoteEventID()
}

func (d *GovernanceDecoder) IsGovernanceVoteEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceVoteEventID().String()
}

func (d *GovernanceDecoder) GovernanceVoteEventW3(w3l web3types.Log) (GovernanceVoteEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return GovernanceVoteEvent{}, err
	}

	return d.GovernanceVoteEvent(l)
}

func (d *GovernanceDecoder) GovernanceVoteEvent(l types.Log) (GovernanceVoteEvent, error) {
	var out GovernanceVoteEvent
	if !d.IsGovernanceVoteEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "Vote", l)
	out.Raw = l
	return out, err
}

type GovernanceProposalExecutedEvent struct {
	ProposalId *big.Int
	Caller     common.Address
	Raw        types.Log
}

func (d *GovernanceDecoder) GovernanceProposalExecutedEventID() common.Hash {
	return common.HexToHash("0x9c85b616f29fca57a17eafe71cf9ff82ffef41766e2cf01ea7f8f7878dd3ec24")
}

func (d *GovernanceDecoder) IsGovernanceProposalExecutedEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceProposalExecutedEventID()
}

func (d *GovernanceDecoder) IsGovernanceProposalExecutedEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceProposalExecutedEventID().String()
}

func (d *GovernanceDecoder) GovernanceProposalExecutedEventW3(w3l web3types.Log) (GovernanceProposalExecutedEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return GovernanceProposalExecutedEvent{}, err
	}

	return d.GovernanceProposalExecutedEvent(l)
}

func (d *GovernanceDecoder) GovernanceProposalExecutedEvent(l types.Log) (GovernanceProposalExecutedEvent, error) {
	var out GovernanceProposalExecutedEvent
	if !d.IsGovernanceProposalExecutedEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "ProposalExecuted", l)
	out.Raw = l
	return out, err
}

type GovernanceProposalCreatedEvent struct {
	ProposalId *big.Int
	Raw        types.Log
}

func (d *GovernanceDecoder) GovernanceProposalCreatedEventID() common.Hash {
	return common.HexToHash("0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07")
}

func (d *GovernanceDecoder) IsGovernanceProposalCreatedEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceProposalCreatedEventID()
}

func (d *GovernanceDecoder) IsGovernanceProposalCreatedEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceProposalCreatedEventID().String()
}

func (d *GovernanceDecoder) GovernanceProposalCreatedEventW3(w3l web3types.Log) (GovernanceProposalCreatedEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return GovernanceProposalCreatedEvent{}, err
	}

	return d.GovernanceProposalCreatedEvent(l)
}

func (d *GovernanceDecoder) GovernanceProposalCreatedEvent(l types.Log) (GovernanceProposalCreatedEvent, error) {
	var out GovernanceProposalCreatedEvent
	if !d.IsGovernanceProposalCreatedEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "ProposalCreated", l)
	out.Raw = l
	return out, err
}

type GovernanceProposalQueuedEvent struct {
	ProposalId *big.Int
	Caller     common.Address
	Eta        *big.Int
	Raw        types.Log
}

func (d *GovernanceDecoder) GovernanceProposalQueuedEventID() common.Hash {
	return common.HexToHash("0xf7230a453b4c21e4f2d0ef1ad055635b08cb2c884eaf24a5ddc7147c79fd8c22")
}

func (d *GovernanceDecoder) IsGovernanceProposalQueuedEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceProposalQueuedEventID()
}

func (d *GovernanceDecoder) IsGovernanceProposalQueuedEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.GovernanceProposalQueuedEventID().String()
}

func (d *GovernanceDecoder) GovernanceProposalQueuedEventW3(w3l web3types.Log) (GovernanceProposalQueuedEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return GovernanceProposalQueuedEvent{}, err
	}

	return d.GovernanceProposalQueuedEvent(l)
}

func (d *GovernanceDecoder) GovernanceProposalQueuedEvent(l types.Log) (GovernanceProposalQueuedEvent, error) {
	var out GovernanceProposalQueuedEvent
	if !d.IsGovernanceProposalQueuedEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "ProposalQueued", l)
	out.Raw = l
	return out, err
}
