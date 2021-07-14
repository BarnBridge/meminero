// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GovernanceProposalParameters is an auto generated low-level Go binding around an user-defined struct.
type GovernanceProposalParameters struct {
	WarmUpDuration      *big.Int
	ActiveDuration      *big.Int
	QueueDuration       *big.Int
	GracePeriodDuration *big.Int
	AcceptanceThreshold *big.Int
	MinQuorum           *big.Int
}

// GovernanceReceipt is an auto generated low-level Go binding around an user-defined struct.
type GovernanceReceipt struct {
	HasVoted bool
	Votes    *big.Int
	Support  bool
}

// GovernanceABI is the input ABI used to generate the binding from.
const GovernanceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AbrogationProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AbrogationProposalStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"name\":\"AbrogationProposalVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"AbrogationProposalVoteCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"VoteCanceled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"abrogateProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"abrogationProposal_cancelVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"}],\"name\":\"abrogationProposal_castVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"abrogationProposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptanceThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"cancelProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"cancelVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"}],\"name\":\"castVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getAbrogationProposalReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"}],\"internalType\":\"structGovernance.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getActions\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalParameters\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"warmUpDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"queueDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gracePeriodDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptanceThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structGovernance.ProposalParameters\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalQuorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"}],\"internalType\":\"structGovernance.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gracePeriodDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"barnAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"latestProposalIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minQuorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canceled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"warmUpDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"queueDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gracePeriodDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptanceThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structGovernance.ProposalParameters\",\"name\":\"parameters\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"queue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"queuedTransactions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"setAcceptanceThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setActiveDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setGracePeriodDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quorum\",\"type\":\"uint256\"}],\"name\":\"setMinQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setQueueDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setWarmUpDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"startAbrogationProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumGovernance.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"warmUpDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Governance is an auto generated Go binding around an Ethereum contract.
type Governance struct {
	GovernanceCaller     // Read-only binding to the contract
	GovernanceTransactor // Write-only binding to the contract
	GovernanceFilterer   // Log filterer for contract events
}

// GovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceSession struct {
	Contract     *Governance       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceCallerSession struct {
	Contract *GovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceTransactorSession struct {
	Contract     *GovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceRaw struct {
	Contract *Governance // Generic contract binding to access the raw methods on
}

// GovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceCallerRaw struct {
	Contract *GovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceTransactorRaw struct {
	Contract *GovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernance creates a new instance of Governance, bound to a specific deployed contract.
func NewGovernance(address common.Address, backend bind.ContractBackend) (*Governance, error) {
	contract, err := bindGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// NewGovernanceCaller creates a new read-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceCaller(address common.Address, caller bind.ContractCaller) (*GovernanceCaller, error) {
	contract, err := bindGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceCaller{contract: contract}, nil
}

// NewGovernanceTransactor creates a new write-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceTransactor, error) {
	contract, err := bindGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTransactor{contract: contract}, nil
}

// NewGovernanceFilterer creates a new log filterer instance of Governance, bound to a specific deployed contract.
func NewGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceFilterer, error) {
	contract, err := bindGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceFilterer{contract: contract}, nil
}

// bindGovernance binds a generic wrapper to an already deployed contract.
func bindGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.GovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transact(opts, method, params...)
}

// AbrogationProposals is a free data retrieval call binding the contract method 0x8b787e30.
//
// Solidity: function abrogationProposals(uint256 ) view returns(address creator, uint256 createTime, string description, uint256 forVotes, uint256 againstVotes)
func (_Governance *GovernanceCaller) AbrogationProposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Creator      common.Address
	CreateTime   *big.Int
	Description  string
	ForVotes     *big.Int
	AgainstVotes *big.Int
}, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "abrogationProposals", arg0)

	outstruct := new(struct {
		Creator      common.Address
		CreateTime   *big.Int
		Description  string
		ForVotes     *big.Int
		AgainstVotes *big.Int
	})

	outstruct.Creator = out[0].(common.Address)
	outstruct.CreateTime = out[1].(*big.Int)
	outstruct.Description = out[2].(string)
	outstruct.ForVotes = out[3].(*big.Int)
	outstruct.AgainstVotes = out[4].(*big.Int)

	return *outstruct, err

}

// AbrogationProposals is a free data retrieval call binding the contract method 0x8b787e30.
//
// Solidity: function abrogationProposals(uint256 ) view returns(address creator, uint256 createTime, string description, uint256 forVotes, uint256 againstVotes)
func (_Governance *GovernanceSession) AbrogationProposals(arg0 *big.Int) (struct {
	Creator      common.Address
	CreateTime   *big.Int
	Description  string
	ForVotes     *big.Int
	AgainstVotes *big.Int
}, error) {
	return _Governance.Contract.AbrogationProposals(&_Governance.CallOpts, arg0)
}

// AbrogationProposals is a free data retrieval call binding the contract method 0x8b787e30.
//
// Solidity: function abrogationProposals(uint256 ) view returns(address creator, uint256 createTime, string description, uint256 forVotes, uint256 againstVotes)
func (_Governance *GovernanceCallerSession) AbrogationProposals(arg0 *big.Int) (struct {
	Creator      common.Address
	CreateTime   *big.Int
	Description  string
	ForVotes     *big.Int
	AgainstVotes *big.Int
}, error) {
	return _Governance.Contract.AbrogationProposals(&_Governance.CallOpts, arg0)
}

// AcceptanceThreshold is a free data retrieval call binding the contract method 0xb0edbb9b.
//
// Solidity: function acceptanceThreshold() view returns(uint256)
func (_Governance *GovernanceCaller) AcceptanceThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "acceptanceThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AcceptanceThreshold is a free data retrieval call binding the contract method 0xb0edbb9b.
//
// Solidity: function acceptanceThreshold() view returns(uint256)
func (_Governance *GovernanceSession) AcceptanceThreshold() (*big.Int, error) {
	return _Governance.Contract.AcceptanceThreshold(&_Governance.CallOpts)
}

// AcceptanceThreshold is a free data retrieval call binding the contract method 0xb0edbb9b.
//
// Solidity: function acceptanceThreshold() view returns(uint256)
func (_Governance *GovernanceCallerSession) AcceptanceThreshold() (*big.Int, error) {
	return _Governance.Contract.AcceptanceThreshold(&_Governance.CallOpts)
}

// ActiveDuration is a free data retrieval call binding the contract method 0x3d05f009.
//
// Solidity: function activeDuration() view returns(uint256)
func (_Governance *GovernanceCaller) ActiveDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "activeDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveDuration is a free data retrieval call binding the contract method 0x3d05f009.
//
// Solidity: function activeDuration() view returns(uint256)
func (_Governance *GovernanceSession) ActiveDuration() (*big.Int, error) {
	return _Governance.Contract.ActiveDuration(&_Governance.CallOpts)
}

// ActiveDuration is a free data retrieval call binding the contract method 0x3d05f009.
//
// Solidity: function activeDuration() view returns(uint256)
func (_Governance *GovernanceCallerSession) ActiveDuration() (*big.Int, error) {
	return _Governance.Contract.ActiveDuration(&_Governance.CallOpts)
}

// GetAbrogationProposalReceipt is a free data retrieval call binding the contract method 0xe4e2d261.
//
// Solidity: function getAbrogationProposalReceipt(uint256 proposalId, address voter) view returns((bool,uint256,bool))
func (_Governance *GovernanceCaller) GetAbrogationProposalReceipt(opts *bind.CallOpts, proposalId *big.Int, voter common.Address) (GovernanceReceipt, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getAbrogationProposalReceipt", proposalId, voter)

	if err != nil {
		return *new(GovernanceReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(GovernanceReceipt)).(*GovernanceReceipt)

	return out0, err

}

// GetAbrogationProposalReceipt is a free data retrieval call binding the contract method 0xe4e2d261.
//
// Solidity: function getAbrogationProposalReceipt(uint256 proposalId, address voter) view returns((bool,uint256,bool))
func (_Governance *GovernanceSession) GetAbrogationProposalReceipt(proposalId *big.Int, voter common.Address) (GovernanceReceipt, error) {
	return _Governance.Contract.GetAbrogationProposalReceipt(&_Governance.CallOpts, proposalId, voter)
}

// GetAbrogationProposalReceipt is a free data retrieval call binding the contract method 0xe4e2d261.
//
// Solidity: function getAbrogationProposalReceipt(uint256 proposalId, address voter) view returns((bool,uint256,bool))
func (_Governance *GovernanceCallerSession) GetAbrogationProposalReceipt(proposalId *big.Int, voter common.Address) (GovernanceReceipt, error) {
	return _Governance.Contract.GetAbrogationProposalReceipt(&_Governance.CallOpts, proposalId, voter)
}

// GetActions is a free data retrieval call binding the contract method 0x328dd982.
//
// Solidity: function getActions(uint256 proposalId) view returns(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas)
func (_Governance *GovernanceCaller) GetActions(opts *bind.CallOpts, proposalId *big.Int) (struct {
	Targets    []common.Address
	Values     []*big.Int
	Signatures []string
	Calldatas  [][]byte
}, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getActions", proposalId)

	outstruct := new(struct {
		Targets    []common.Address
		Values     []*big.Int
		Signatures []string
		Calldatas  [][]byte
	})

	outstruct.Targets = out[0].([]common.Address)
	outstruct.Values = out[1].([]*big.Int)
	outstruct.Signatures = out[2].([]string)
	outstruct.Calldatas = out[3].([][]byte)

	return *outstruct, err

}

// GetActions is a free data retrieval call binding the contract method 0x328dd982.
//
// Solidity: function getActions(uint256 proposalId) view returns(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas)
func (_Governance *GovernanceSession) GetActions(proposalId *big.Int) (struct {
	Targets    []common.Address
	Values     []*big.Int
	Signatures []string
	Calldatas  [][]byte
}, error) {
	return _Governance.Contract.GetActions(&_Governance.CallOpts, proposalId)
}

// GetActions is a free data retrieval call binding the contract method 0x328dd982.
//
// Solidity: function getActions(uint256 proposalId) view returns(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas)
func (_Governance *GovernanceCallerSession) GetActions(proposalId *big.Int) (struct {
	Targets    []common.Address
	Values     []*big.Int
	Signatures []string
	Calldatas  [][]byte
}, error) {
	return _Governance.Contract.GetActions(&_Governance.CallOpts, proposalId)
}

// GetProposalParameters is a free data retrieval call binding the contract method 0x45892155.
//
// Solidity: function getProposalParameters(uint256 proposalId) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_Governance *GovernanceCaller) GetProposalParameters(opts *bind.CallOpts, proposalId *big.Int) (GovernanceProposalParameters, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getProposalParameters", proposalId)

	if err != nil {
		return *new(GovernanceProposalParameters), err
	}

	out0 := *abi.ConvertType(out[0], new(GovernanceProposalParameters)).(*GovernanceProposalParameters)

	return out0, err

}

// GetProposalParameters is a free data retrieval call binding the contract method 0x45892155.
//
// Solidity: function getProposalParameters(uint256 proposalId) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_Governance *GovernanceSession) GetProposalParameters(proposalId *big.Int) (GovernanceProposalParameters, error) {
	return _Governance.Contract.GetProposalParameters(&_Governance.CallOpts, proposalId)
}

// GetProposalParameters is a free data retrieval call binding the contract method 0x45892155.
//
// Solidity: function getProposalParameters(uint256 proposalId) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_Governance *GovernanceCallerSession) GetProposalParameters(proposalId *big.Int) (GovernanceProposalParameters, error) {
	return _Governance.Contract.GetProposalParameters(&_Governance.CallOpts, proposalId)
}

// GetProposalQuorum is a free data retrieval call binding the contract method 0xd0cd595e.
//
// Solidity: function getProposalQuorum(uint256 proposalId) view returns(uint256)
func (_Governance *GovernanceCaller) GetProposalQuorum(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getProposalQuorum", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposalQuorum is a free data retrieval call binding the contract method 0xd0cd595e.
//
// Solidity: function getProposalQuorum(uint256 proposalId) view returns(uint256)
func (_Governance *GovernanceSession) GetProposalQuorum(proposalId *big.Int) (*big.Int, error) {
	return _Governance.Contract.GetProposalQuorum(&_Governance.CallOpts, proposalId)
}

// GetProposalQuorum is a free data retrieval call binding the contract method 0xd0cd595e.
//
// Solidity: function getProposalQuorum(uint256 proposalId) view returns(uint256)
func (_Governance *GovernanceCallerSession) GetProposalQuorum(proposalId *big.Int) (*big.Int, error) {
	return _Governance.Contract.GetProposalQuorum(&_Governance.CallOpts, proposalId)
}

// GetReceipt is a free data retrieval call binding the contract method 0xe23a9a52.
//
// Solidity: function getReceipt(uint256 proposalId, address voter) view returns((bool,uint256,bool))
func (_Governance *GovernanceCaller) GetReceipt(opts *bind.CallOpts, proposalId *big.Int, voter common.Address) (GovernanceReceipt, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getReceipt", proposalId, voter)

	if err != nil {
		return *new(GovernanceReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(GovernanceReceipt)).(*GovernanceReceipt)

	return out0, err

}

// GetReceipt is a free data retrieval call binding the contract method 0xe23a9a52.
//
// Solidity: function getReceipt(uint256 proposalId, address voter) view returns((bool,uint256,bool))
func (_Governance *GovernanceSession) GetReceipt(proposalId *big.Int, voter common.Address) (GovernanceReceipt, error) {
	return _Governance.Contract.GetReceipt(&_Governance.CallOpts, proposalId, voter)
}

// GetReceipt is a free data retrieval call binding the contract method 0xe23a9a52.
//
// Solidity: function getReceipt(uint256 proposalId, address voter) view returns((bool,uint256,bool))
func (_Governance *GovernanceCallerSession) GetReceipt(proposalId *big.Int, voter common.Address) (GovernanceReceipt, error) {
	return _Governance.Contract.GetReceipt(&_Governance.CallOpts, proposalId, voter)
}

// GracePeriodDuration is a free data retrieval call binding the contract method 0xc099f575.
//
// Solidity: function gracePeriodDuration() view returns(uint256)
func (_Governance *GovernanceCaller) GracePeriodDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "gracePeriodDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GracePeriodDuration is a free data retrieval call binding the contract method 0xc099f575.
//
// Solidity: function gracePeriodDuration() view returns(uint256)
func (_Governance *GovernanceSession) GracePeriodDuration() (*big.Int, error) {
	return _Governance.Contract.GracePeriodDuration(&_Governance.CallOpts)
}

// GracePeriodDuration is a free data retrieval call binding the contract method 0xc099f575.
//
// Solidity: function gracePeriodDuration() view returns(uint256)
func (_Governance *GovernanceCallerSession) GracePeriodDuration() (*big.Int, error) {
	return _Governance.Contract.GracePeriodDuration(&_Governance.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_Governance *GovernanceCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "isActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_Governance *GovernanceSession) IsActive() (bool, error) {
	return _Governance.Contract.IsActive(&_Governance.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_Governance *GovernanceCallerSession) IsActive() (bool, error) {
	return _Governance.Contract.IsActive(&_Governance.CallOpts)
}

// LastProposalId is a free data retrieval call binding the contract method 0x74cb3041.
//
// Solidity: function lastProposalId() view returns(uint256)
func (_Governance *GovernanceCaller) LastProposalId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "lastProposalId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastProposalId is a free data retrieval call binding the contract method 0x74cb3041.
//
// Solidity: function lastProposalId() view returns(uint256)
func (_Governance *GovernanceSession) LastProposalId() (*big.Int, error) {
	return _Governance.Contract.LastProposalId(&_Governance.CallOpts)
}

// LastProposalId is a free data retrieval call binding the contract method 0x74cb3041.
//
// Solidity: function lastProposalId() view returns(uint256)
func (_Governance *GovernanceCallerSession) LastProposalId() (*big.Int, error) {
	return _Governance.Contract.LastProposalId(&_Governance.CallOpts)
}

// LatestProposalIds is a free data retrieval call binding the contract method 0x17977c61.
//
// Solidity: function latestProposalIds(address ) view returns(uint256)
func (_Governance *GovernanceCaller) LatestProposalIds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "latestProposalIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestProposalIds is a free data retrieval call binding the contract method 0x17977c61.
//
// Solidity: function latestProposalIds(address ) view returns(uint256)
func (_Governance *GovernanceSession) LatestProposalIds(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.LatestProposalIds(&_Governance.CallOpts, arg0)
}

// LatestProposalIds is a free data retrieval call binding the contract method 0x17977c61.
//
// Solidity: function latestProposalIds(address ) view returns(uint256)
func (_Governance *GovernanceCallerSession) LatestProposalIds(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.LatestProposalIds(&_Governance.CallOpts, arg0)
}

// MinQuorum is a free data retrieval call binding the contract method 0xb5a127e5.
//
// Solidity: function minQuorum() view returns(uint256)
func (_Governance *GovernanceCaller) MinQuorum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "minQuorum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuorum is a free data retrieval call binding the contract method 0xb5a127e5.
//
// Solidity: function minQuorum() view returns(uint256)
func (_Governance *GovernanceSession) MinQuorum() (*big.Int, error) {
	return _Governance.Contract.MinQuorum(&_Governance.CallOpts)
}

// MinQuorum is a free data retrieval call binding the contract method 0xb5a127e5.
//
// Solidity: function minQuorum() view returns(uint256)
func (_Governance *GovernanceCallerSession) MinQuorum() (*big.Int, error) {
	return _Governance.Contract.MinQuorum(&_Governance.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(uint256 id, address proposer, string description, string title, uint256 createTime, uint256 eta, uint256 forVotes, uint256 againstVotes, bool canceled, bool executed, (uint256,uint256,uint256,uint256,uint256,uint256) parameters)
func (_Governance *GovernanceCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id           *big.Int
	Proposer     common.Address
	Description  string
	Title        string
	CreateTime   *big.Int
	Eta          *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	Canceled     bool
	Executed     bool
	Parameters   GovernanceProposalParameters
}, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Id           *big.Int
		Proposer     common.Address
		Description  string
		Title        string
		CreateTime   *big.Int
		Eta          *big.Int
		ForVotes     *big.Int
		AgainstVotes *big.Int
		Canceled     bool
		Executed     bool
		Parameters   GovernanceProposalParameters
	})

	outstruct.Id = out[0].(*big.Int)
	outstruct.Proposer = out[1].(common.Address)
	outstruct.Description = out[2].(string)
	outstruct.Title = out[3].(string)
	outstruct.CreateTime = out[4].(*big.Int)
	outstruct.Eta = out[5].(*big.Int)
	outstruct.ForVotes = out[6].(*big.Int)
	outstruct.AgainstVotes = out[7].(*big.Int)
	outstruct.Canceled = out[8].(bool)
	outstruct.Executed = out[9].(bool)
	outstruct.Parameters = *abi.ConvertType(out[10], new(GovernanceProposalParameters)).(*GovernanceProposalParameters)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(uint256 id, address proposer, string description, string title, uint256 createTime, uint256 eta, uint256 forVotes, uint256 againstVotes, bool canceled, bool executed, (uint256,uint256,uint256,uint256,uint256,uint256) parameters)
func (_Governance *GovernanceSession) Proposals(arg0 *big.Int) (struct {
	Id           *big.Int
	Proposer     common.Address
	Description  string
	Title        string
	CreateTime   *big.Int
	Eta          *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	Canceled     bool
	Executed     bool
	Parameters   GovernanceProposalParameters
}, error) {
	return _Governance.Contract.Proposals(&_Governance.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(uint256 id, address proposer, string description, string title, uint256 createTime, uint256 eta, uint256 forVotes, uint256 againstVotes, bool canceled, bool executed, (uint256,uint256,uint256,uint256,uint256,uint256) parameters)
func (_Governance *GovernanceCallerSession) Proposals(arg0 *big.Int) (struct {
	Id           *big.Int
	Proposer     common.Address
	Description  string
	Title        string
	CreateTime   *big.Int
	Eta          *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	Canceled     bool
	Executed     bool
	Parameters   GovernanceProposalParameters
}, error) {
	return _Governance.Contract.Proposals(&_Governance.CallOpts, arg0)
}

// QueueDuration is a free data retrieval call binding the contract method 0x2e8e34e1.
//
// Solidity: function queueDuration() view returns(uint256)
func (_Governance *GovernanceCaller) QueueDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "queueDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueueDuration is a free data retrieval call binding the contract method 0x2e8e34e1.
//
// Solidity: function queueDuration() view returns(uint256)
func (_Governance *GovernanceSession) QueueDuration() (*big.Int, error) {
	return _Governance.Contract.QueueDuration(&_Governance.CallOpts)
}

// QueueDuration is a free data retrieval call binding the contract method 0x2e8e34e1.
//
// Solidity: function queueDuration() view returns(uint256)
func (_Governance *GovernanceCallerSession) QueueDuration() (*big.Int, error) {
	return _Governance.Contract.QueueDuration(&_Governance.CallOpts)
}

// QueuedTransactions is a free data retrieval call binding the contract method 0xf2b06537.
//
// Solidity: function queuedTransactions(bytes32 ) view returns(bool)
func (_Governance *GovernanceCaller) QueuedTransactions(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "queuedTransactions", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// QueuedTransactions is a free data retrieval call binding the contract method 0xf2b06537.
//
// Solidity: function queuedTransactions(bytes32 ) view returns(bool)
func (_Governance *GovernanceSession) QueuedTransactions(arg0 [32]byte) (bool, error) {
	return _Governance.Contract.QueuedTransactions(&_Governance.CallOpts, arg0)
}

// QueuedTransactions is a free data retrieval call binding the contract method 0xf2b06537.
//
// Solidity: function queuedTransactions(bytes32 ) view returns(bool)
func (_Governance *GovernanceCallerSession) QueuedTransactions(arg0 [32]byte) (bool, error) {
	return _Governance.Contract.QueuedTransactions(&_Governance.CallOpts, arg0)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_Governance *GovernanceCaller) State(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_Governance *GovernanceSession) State(proposalId *big.Int) (uint8, error) {
	return _Governance.Contract.State(&_Governance.CallOpts, proposalId)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_Governance *GovernanceCallerSession) State(proposalId *big.Int) (uint8, error) {
	return _Governance.Contract.State(&_Governance.CallOpts, proposalId)
}

// WarmUpDuration is a free data retrieval call binding the contract method 0x5f2e9f60.
//
// Solidity: function warmUpDuration() view returns(uint256)
func (_Governance *GovernanceCaller) WarmUpDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "warmUpDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WarmUpDuration is a free data retrieval call binding the contract method 0x5f2e9f60.
//
// Solidity: function warmUpDuration() view returns(uint256)
func (_Governance *GovernanceSession) WarmUpDuration() (*big.Int, error) {
	return _Governance.Contract.WarmUpDuration(&_Governance.CallOpts)
}

// WarmUpDuration is a free data retrieval call binding the contract method 0x5f2e9f60.
//
// Solidity: function warmUpDuration() view returns(uint256)
func (_Governance *GovernanceCallerSession) WarmUpDuration() (*big.Int, error) {
	return _Governance.Contract.WarmUpDuration(&_Governance.CallOpts)
}

// AbrogateProposal is a paid mutator transaction binding the contract method 0x24b43787.
//
// Solidity: function abrogateProposal(uint256 proposalId) returns()
func (_Governance *GovernanceTransactor) AbrogateProposal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "abrogateProposal", proposalId)
}

// AbrogateProposal is a paid mutator transaction binding the contract method 0x24b43787.
//
// Solidity: function abrogateProposal(uint256 proposalId) returns()
func (_Governance *GovernanceSession) AbrogateProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.AbrogateProposal(&_Governance.TransactOpts, proposalId)
}

// AbrogateProposal is a paid mutator transaction binding the contract method 0x24b43787.
//
// Solidity: function abrogateProposal(uint256 proposalId) returns()
func (_Governance *GovernanceTransactorSession) AbrogateProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.AbrogateProposal(&_Governance.TransactOpts, proposalId)
}

// AbrogationProposalCancelVote is a paid mutator transaction binding the contract method 0x1e9d1e64.
//
// Solidity: function abrogationProposal_cancelVote(uint256 proposalId) returns()
func (_Governance *GovernanceTransactor) AbrogationProposalCancelVote(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "abrogationProposal_cancelVote", proposalId)
}

// AbrogationProposalCancelVote is a paid mutator transaction binding the contract method 0x1e9d1e64.
//
// Solidity: function abrogationProposal_cancelVote(uint256 proposalId) returns()
func (_Governance *GovernanceSession) AbrogationProposalCancelVote(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.AbrogationProposalCancelVote(&_Governance.TransactOpts, proposalId)
}

// AbrogationProposalCancelVote is a paid mutator transaction binding the contract method 0x1e9d1e64.
//
// Solidity: function abrogationProposal_cancelVote(uint256 proposalId) returns()
func (_Governance *GovernanceTransactorSession) AbrogationProposalCancelVote(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.AbrogationProposalCancelVote(&_Governance.TransactOpts, proposalId)
}

// AbrogationProposalCastVote is a paid mutator transaction binding the contract method 0x9e70a234.
//
// Solidity: function abrogationProposal_castVote(uint256 proposalId, bool support) returns()
func (_Governance *GovernanceTransactor) AbrogationProposalCastVote(opts *bind.TransactOpts, proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "abrogationProposal_castVote", proposalId, support)
}

// AbrogationProposalCastVote is a paid mutator transaction binding the contract method 0x9e70a234.
//
// Solidity: function abrogationProposal_castVote(uint256 proposalId, bool support) returns()
func (_Governance *GovernanceSession) AbrogationProposalCastVote(proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _Governance.Contract.AbrogationProposalCastVote(&_Governance.TransactOpts, proposalId, support)
}

// AbrogationProposalCastVote is a paid mutator transaction binding the contract method 0x9e70a234.
//
// Solidity: function abrogationProposal_castVote(uint256 proposalId, bool support) returns()
func (_Governance *GovernanceTransactorSession) AbrogationProposalCastVote(proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _Governance.Contract.AbrogationProposalCastVote(&_Governance.TransactOpts, proposalId, support)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Governance *GovernanceTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Governance *GovernanceSession) Activate() (*types.Transaction, error) {
	return _Governance.Contract.Activate(&_Governance.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Governance *GovernanceTransactorSession) Activate() (*types.Transaction, error) {
	return _Governance.Contract.Activate(&_Governance.TransactOpts)
}

// CancelProposal is a paid mutator transaction binding the contract method 0xe0a8f6f5.
//
// Solidity: function cancelProposal(uint256 proposalId) returns()
func (_Governance *GovernanceTransactor) CancelProposal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "cancelProposal", proposalId)
}

// CancelProposal is a paid mutator transaction binding the contract method 0xe0a8f6f5.
//
// Solidity: function cancelProposal(uint256 proposalId) returns()
func (_Governance *GovernanceSession) CancelProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.CancelProposal(&_Governance.TransactOpts, proposalId)
}

// CancelProposal is a paid mutator transaction binding the contract method 0xe0a8f6f5.
//
// Solidity: function cancelProposal(uint256 proposalId) returns()
func (_Governance *GovernanceTransactorSession) CancelProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.CancelProposal(&_Governance.TransactOpts, proposalId)
}

// CancelVote is a paid mutator transaction binding the contract method 0xbacbe2da.
//
// Solidity: function cancelVote(uint256 proposalId) returns()
func (_Governance *GovernanceTransactor) CancelVote(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "cancelVote", proposalId)
}

// CancelVote is a paid mutator transaction binding the contract method 0xbacbe2da.
//
// Solidity: function cancelVote(uint256 proposalId) returns()
func (_Governance *GovernanceSession) CancelVote(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.CancelVote(&_Governance.TransactOpts, proposalId)
}

// CancelVote is a paid mutator transaction binding the contract method 0xbacbe2da.
//
// Solidity: function cancelVote(uint256 proposalId) returns()
func (_Governance *GovernanceTransactorSession) CancelVote(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.CancelVote(&_Governance.TransactOpts, proposalId)
}

// CastVote is a paid mutator transaction binding the contract method 0x15373e3d.
//
// Solidity: function castVote(uint256 proposalId, bool support) returns()
func (_Governance *GovernanceTransactor) CastVote(opts *bind.TransactOpts, proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "castVote", proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x15373e3d.
//
// Solidity: function castVote(uint256 proposalId, bool support) returns()
func (_Governance *GovernanceSession) CastVote(proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _Governance.Contract.CastVote(&_Governance.TransactOpts, proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x15373e3d.
//
// Solidity: function castVote(uint256 proposalId, bool support) returns()
func (_Governance *GovernanceTransactorSession) CastVote(proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _Governance.Contract.CastVote(&_Governance.TransactOpts, proposalId, support)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns()
func (_Governance *GovernanceTransactor) Execute(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "execute", proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns()
func (_Governance *GovernanceSession) Execute(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Execute(&_Governance.TransactOpts, proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns()
func (_Governance *GovernanceTransactorSession) Execute(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Execute(&_Governance.TransactOpts, proposalId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address barnAddr) returns()
func (_Governance *GovernanceTransactor) Initialize(opts *bind.TransactOpts, barnAddr common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "initialize", barnAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address barnAddr) returns()
func (_Governance *GovernanceSession) Initialize(barnAddr common.Address) (*types.Transaction, error) {
	return _Governance.Contract.Initialize(&_Governance.TransactOpts, barnAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address barnAddr) returns()
func (_Governance *GovernanceTransactorSession) Initialize(barnAddr common.Address) (*types.Transaction, error) {
	return _Governance.Contract.Initialize(&_Governance.TransactOpts, barnAddr)
}

// Propose is a paid mutator transaction binding the contract method 0x490145c8.
//
// Solidity: function propose(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, string description, string title) returns(uint256)
func (_Governance *GovernanceTransactor) Propose(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, signatures []string, calldatas [][]byte, description string, title string) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "propose", targets, values, signatures, calldatas, description, title)
}

// Propose is a paid mutator transaction binding the contract method 0x490145c8.
//
// Solidity: function propose(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, string description, string title) returns(uint256)
func (_Governance *GovernanceSession) Propose(targets []common.Address, values []*big.Int, signatures []string, calldatas [][]byte, description string, title string) (*types.Transaction, error) {
	return _Governance.Contract.Propose(&_Governance.TransactOpts, targets, values, signatures, calldatas, description, title)
}

// Propose is a paid mutator transaction binding the contract method 0x490145c8.
//
// Solidity: function propose(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, string description, string title) returns(uint256)
func (_Governance *GovernanceTransactorSession) Propose(targets []common.Address, values []*big.Int, signatures []string, calldatas [][]byte, description string, title string) (*types.Transaction, error) {
	return _Governance.Contract.Propose(&_Governance.TransactOpts, targets, values, signatures, calldatas, description, title)
}

// Queue is a paid mutator transaction binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 proposalId) returns()
func (_Governance *GovernanceTransactor) Queue(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "queue", proposalId)
}

// Queue is a paid mutator transaction binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 proposalId) returns()
func (_Governance *GovernanceSession) Queue(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Queue(&_Governance.TransactOpts, proposalId)
}

// Queue is a paid mutator transaction binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 proposalId) returns()
func (_Governance *GovernanceTransactorSession) Queue(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Queue(&_Governance.TransactOpts, proposalId)
}

// SetAcceptanceThreshold is a paid mutator transaction binding the contract method 0xd1291f19.
//
// Solidity: function setAcceptanceThreshold(uint256 threshold) returns()
func (_Governance *GovernanceTransactor) SetAcceptanceThreshold(opts *bind.TransactOpts, threshold *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setAcceptanceThreshold", threshold)
}

// SetAcceptanceThreshold is a paid mutator transaction binding the contract method 0xd1291f19.
//
// Solidity: function setAcceptanceThreshold(uint256 threshold) returns()
func (_Governance *GovernanceSession) SetAcceptanceThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetAcceptanceThreshold(&_Governance.TransactOpts, threshold)
}

// SetAcceptanceThreshold is a paid mutator transaction binding the contract method 0xd1291f19.
//
// Solidity: function setAcceptanceThreshold(uint256 threshold) returns()
func (_Governance *GovernanceTransactorSession) SetAcceptanceThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetAcceptanceThreshold(&_Governance.TransactOpts, threshold)
}

// SetActiveDuration is a paid mutator transaction binding the contract method 0x24cd62d3.
//
// Solidity: function setActiveDuration(uint256 period) returns()
func (_Governance *GovernanceTransactor) SetActiveDuration(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setActiveDuration", period)
}

// SetActiveDuration is a paid mutator transaction binding the contract method 0x24cd62d3.
//
// Solidity: function setActiveDuration(uint256 period) returns()
func (_Governance *GovernanceSession) SetActiveDuration(period *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetActiveDuration(&_Governance.TransactOpts, period)
}

// SetActiveDuration is a paid mutator transaction binding the contract method 0x24cd62d3.
//
// Solidity: function setActiveDuration(uint256 period) returns()
func (_Governance *GovernanceTransactorSession) SetActiveDuration(period *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetActiveDuration(&_Governance.TransactOpts, period)
}

// SetGracePeriodDuration is a paid mutator transaction binding the contract method 0x342d067a.
//
// Solidity: function setGracePeriodDuration(uint256 period) returns()
func (_Governance *GovernanceTransactor) SetGracePeriodDuration(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setGracePeriodDuration", period)
}

// SetGracePeriodDuration is a paid mutator transaction binding the contract method 0x342d067a.
//
// Solidity: function setGracePeriodDuration(uint256 period) returns()
func (_Governance *GovernanceSession) SetGracePeriodDuration(period *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetGracePeriodDuration(&_Governance.TransactOpts, period)
}

// SetGracePeriodDuration is a paid mutator transaction binding the contract method 0x342d067a.
//
// Solidity: function setGracePeriodDuration(uint256 period) returns()
func (_Governance *GovernanceTransactorSession) SetGracePeriodDuration(period *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetGracePeriodDuration(&_Governance.TransactOpts, period)
}

// SetMinQuorum is a paid mutator transaction binding the contract method 0x563909de.
//
// Solidity: function setMinQuorum(uint256 quorum) returns()
func (_Governance *GovernanceTransactor) SetMinQuorum(opts *bind.TransactOpts, quorum *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setMinQuorum", quorum)
}

// SetMinQuorum is a paid mutator transaction binding the contract method 0x563909de.
//
// Solidity: function setMinQuorum(uint256 quorum) returns()
func (_Governance *GovernanceSession) SetMinQuorum(quorum *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetMinQuorum(&_Governance.TransactOpts, quorum)
}

// SetMinQuorum is a paid mutator transaction binding the contract method 0x563909de.
//
// Solidity: function setMinQuorum(uint256 quorum) returns()
func (_Governance *GovernanceTransactorSession) SetMinQuorum(quorum *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetMinQuorum(&_Governance.TransactOpts, quorum)
}

// SetQueueDuration is a paid mutator transaction binding the contract method 0x53e5056a.
//
// Solidity: function setQueueDuration(uint256 period) returns()
func (_Governance *GovernanceTransactor) SetQueueDuration(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setQueueDuration", period)
}

// SetQueueDuration is a paid mutator transaction binding the contract method 0x53e5056a.
//
// Solidity: function setQueueDuration(uint256 period) returns()
func (_Governance *GovernanceSession) SetQueueDuration(period *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetQueueDuration(&_Governance.TransactOpts, period)
}

// SetQueueDuration is a paid mutator transaction binding the contract method 0x53e5056a.
//
// Solidity: function setQueueDuration(uint256 period) returns()
func (_Governance *GovernanceTransactorSession) SetQueueDuration(period *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetQueueDuration(&_Governance.TransactOpts, period)
}

// SetWarmUpDuration is a paid mutator transaction binding the contract method 0x984690db.
//
// Solidity: function setWarmUpDuration(uint256 period) returns()
func (_Governance *GovernanceTransactor) SetWarmUpDuration(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setWarmUpDuration", period)
}

// SetWarmUpDuration is a paid mutator transaction binding the contract method 0x984690db.
//
// Solidity: function setWarmUpDuration(uint256 period) returns()
func (_Governance *GovernanceSession) SetWarmUpDuration(period *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetWarmUpDuration(&_Governance.TransactOpts, period)
}

// SetWarmUpDuration is a paid mutator transaction binding the contract method 0x984690db.
//
// Solidity: function setWarmUpDuration(uint256 period) returns()
func (_Governance *GovernanceTransactorSession) SetWarmUpDuration(period *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetWarmUpDuration(&_Governance.TransactOpts, period)
}

// StartAbrogationProposal is a paid mutator transaction binding the contract method 0x39e778ee.
//
// Solidity: function startAbrogationProposal(uint256 proposalId, string description) returns()
func (_Governance *GovernanceTransactor) StartAbrogationProposal(opts *bind.TransactOpts, proposalId *big.Int, description string) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "startAbrogationProposal", proposalId, description)
}

// StartAbrogationProposal is a paid mutator transaction binding the contract method 0x39e778ee.
//
// Solidity: function startAbrogationProposal(uint256 proposalId, string description) returns()
func (_Governance *GovernanceSession) StartAbrogationProposal(proposalId *big.Int, description string) (*types.Transaction, error) {
	return _Governance.Contract.StartAbrogationProposal(&_Governance.TransactOpts, proposalId, description)
}

// StartAbrogationProposal is a paid mutator transaction binding the contract method 0x39e778ee.
//
// Solidity: function startAbrogationProposal(uint256 proposalId, string description) returns()
func (_Governance *GovernanceTransactorSession) StartAbrogationProposal(proposalId *big.Int, description string) (*types.Transaction, error) {
	return _Governance.Contract.StartAbrogationProposal(&_Governance.TransactOpts, proposalId, description)
}

// GovernanceAbrogationProposalExecutedIterator is returned from FilterAbrogationProposalExecuted and is used to iterate over the raw logs and unpacked data for AbrogationProposalExecuted events raised by the Governance contract.
type GovernanceAbrogationProposalExecutedIterator struct {
	Event *GovernanceAbrogationProposalExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceAbrogationProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceAbrogationProposalExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceAbrogationProposalExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceAbrogationProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceAbrogationProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceAbrogationProposalExecuted represents a AbrogationProposalExecuted event raised by the Governance contract.
type GovernanceAbrogationProposalExecuted struct {
	ProposalId *big.Int
	Caller     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAbrogationProposalExecuted is a free log retrieval operation binding the contract event 0x6d7acd63bebeaf524f1761a88687951f776fc7c182205f820424c7fb572c7235.
//
// Solidity: event AbrogationProposalExecuted(uint256 indexed proposalId, address caller)
func (_Governance *GovernanceFilterer) FilterAbrogationProposalExecuted(opts *bind.FilterOpts, proposalId []*big.Int) (*GovernanceAbrogationProposalExecutedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "AbrogationProposalExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceAbrogationProposalExecutedIterator{contract: _Governance.contract, event: "AbrogationProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchAbrogationProposalExecuted is a free log subscription operation binding the contract event 0x6d7acd63bebeaf524f1761a88687951f776fc7c182205f820424c7fb572c7235.
//
// Solidity: event AbrogationProposalExecuted(uint256 indexed proposalId, address caller)
func (_Governance *GovernanceFilterer) WatchAbrogationProposalExecuted(opts *bind.WatchOpts, sink chan<- *GovernanceAbrogationProposalExecuted, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "AbrogationProposalExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceAbrogationProposalExecuted)
				if err := _Governance.contract.UnpackLog(event, "AbrogationProposalExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAbrogationProposalExecuted is a log parse operation binding the contract event 0x6d7acd63bebeaf524f1761a88687951f776fc7c182205f820424c7fb572c7235.
//
// Solidity: event AbrogationProposalExecuted(uint256 indexed proposalId, address caller)
func (_Governance *GovernanceFilterer) ParseAbrogationProposalExecuted(log types.Log) (*GovernanceAbrogationProposalExecuted, error) {
	event := new(GovernanceAbrogationProposalExecuted)
	if err := _Governance.contract.UnpackLog(event, "AbrogationProposalExecuted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceAbrogationProposalStartedIterator is returned from FilterAbrogationProposalStarted and is used to iterate over the raw logs and unpacked data for AbrogationProposalStarted events raised by the Governance contract.
type GovernanceAbrogationProposalStartedIterator struct {
	Event *GovernanceAbrogationProposalStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceAbrogationProposalStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceAbrogationProposalStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceAbrogationProposalStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceAbrogationProposalStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceAbrogationProposalStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceAbrogationProposalStarted represents a AbrogationProposalStarted event raised by the Governance contract.
type GovernanceAbrogationProposalStarted struct {
	ProposalId *big.Int
	Caller     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAbrogationProposalStarted is a free log retrieval operation binding the contract event 0x27eba018e1c52b84f732fe4d806fd9750c60752f1d37e7f70bcb4cbec65b1c6a.
//
// Solidity: event AbrogationProposalStarted(uint256 indexed proposalId, address caller)
func (_Governance *GovernanceFilterer) FilterAbrogationProposalStarted(opts *bind.FilterOpts, proposalId []*big.Int) (*GovernanceAbrogationProposalStartedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "AbrogationProposalStarted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceAbrogationProposalStartedIterator{contract: _Governance.contract, event: "AbrogationProposalStarted", logs: logs, sub: sub}, nil
}

// WatchAbrogationProposalStarted is a free log subscription operation binding the contract event 0x27eba018e1c52b84f732fe4d806fd9750c60752f1d37e7f70bcb4cbec65b1c6a.
//
// Solidity: event AbrogationProposalStarted(uint256 indexed proposalId, address caller)
func (_Governance *GovernanceFilterer) WatchAbrogationProposalStarted(opts *bind.WatchOpts, sink chan<- *GovernanceAbrogationProposalStarted, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "AbrogationProposalStarted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceAbrogationProposalStarted)
				if err := _Governance.contract.UnpackLog(event, "AbrogationProposalStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAbrogationProposalStarted is a log parse operation binding the contract event 0x27eba018e1c52b84f732fe4d806fd9750c60752f1d37e7f70bcb4cbec65b1c6a.
//
// Solidity: event AbrogationProposalStarted(uint256 indexed proposalId, address caller)
func (_Governance *GovernanceFilterer) ParseAbrogationProposalStarted(log types.Log) (*GovernanceAbrogationProposalStarted, error) {
	event := new(GovernanceAbrogationProposalStarted)
	if err := _Governance.contract.UnpackLog(event, "AbrogationProposalStarted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceAbrogationProposalVoteIterator is returned from FilterAbrogationProposalVote and is used to iterate over the raw logs and unpacked data for AbrogationProposalVote events raised by the Governance contract.
type GovernanceAbrogationProposalVoteIterator struct {
	Event *GovernanceAbrogationProposalVote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceAbrogationProposalVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceAbrogationProposalVote)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceAbrogationProposalVote)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceAbrogationProposalVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceAbrogationProposalVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceAbrogationProposalVote represents a AbrogationProposalVote event raised by the Governance contract.
type GovernanceAbrogationProposalVote struct {
	ProposalId *big.Int
	User       common.Address
	Support    bool
	Power      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAbrogationProposalVote is a free log retrieval operation binding the contract event 0x80f2ad7e3e83d197670402663f224adb2f649967b9629c67dcfafa40c94d30f9.
//
// Solidity: event AbrogationProposalVote(uint256 indexed proposalId, address indexed user, bool support, uint256 power)
func (_Governance *GovernanceFilterer) FilterAbrogationProposalVote(opts *bind.FilterOpts, proposalId []*big.Int, user []common.Address) (*GovernanceAbrogationProposalVoteIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "AbrogationProposalVote", proposalIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceAbrogationProposalVoteIterator{contract: _Governance.contract, event: "AbrogationProposalVote", logs: logs, sub: sub}, nil
}

// WatchAbrogationProposalVote is a free log subscription operation binding the contract event 0x80f2ad7e3e83d197670402663f224adb2f649967b9629c67dcfafa40c94d30f9.
//
// Solidity: event AbrogationProposalVote(uint256 indexed proposalId, address indexed user, bool support, uint256 power)
func (_Governance *GovernanceFilterer) WatchAbrogationProposalVote(opts *bind.WatchOpts, sink chan<- *GovernanceAbrogationProposalVote, proposalId []*big.Int, user []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "AbrogationProposalVote", proposalIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceAbrogationProposalVote)
				if err := _Governance.contract.UnpackLog(event, "AbrogationProposalVote", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAbrogationProposalVote is a log parse operation binding the contract event 0x80f2ad7e3e83d197670402663f224adb2f649967b9629c67dcfafa40c94d30f9.
//
// Solidity: event AbrogationProposalVote(uint256 indexed proposalId, address indexed user, bool support, uint256 power)
func (_Governance *GovernanceFilterer) ParseAbrogationProposalVote(log types.Log) (*GovernanceAbrogationProposalVote, error) {
	event := new(GovernanceAbrogationProposalVote)
	if err := _Governance.contract.UnpackLog(event, "AbrogationProposalVote", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceAbrogationProposalVoteCancelledIterator is returned from FilterAbrogationProposalVoteCancelled and is used to iterate over the raw logs and unpacked data for AbrogationProposalVoteCancelled events raised by the Governance contract.
type GovernanceAbrogationProposalVoteCancelledIterator struct {
	Event *GovernanceAbrogationProposalVoteCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceAbrogationProposalVoteCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceAbrogationProposalVoteCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceAbrogationProposalVoteCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceAbrogationProposalVoteCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceAbrogationProposalVoteCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceAbrogationProposalVoteCancelled represents a AbrogationProposalVoteCancelled event raised by the Governance contract.
type GovernanceAbrogationProposalVoteCancelled struct {
	ProposalId *big.Int
	User       common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAbrogationProposalVoteCancelled is a free log retrieval operation binding the contract event 0x5e8ee24f838173ed2ae7989835696f6e11945ac8fbc5259aef01cc4d7f0d4920.
//
// Solidity: event AbrogationProposalVoteCancelled(uint256 indexed proposalId, address indexed user)
func (_Governance *GovernanceFilterer) FilterAbrogationProposalVoteCancelled(opts *bind.FilterOpts, proposalId []*big.Int, user []common.Address) (*GovernanceAbrogationProposalVoteCancelledIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "AbrogationProposalVoteCancelled", proposalIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceAbrogationProposalVoteCancelledIterator{contract: _Governance.contract, event: "AbrogationProposalVoteCancelled", logs: logs, sub: sub}, nil
}

// WatchAbrogationProposalVoteCancelled is a free log subscription operation binding the contract event 0x5e8ee24f838173ed2ae7989835696f6e11945ac8fbc5259aef01cc4d7f0d4920.
//
// Solidity: event AbrogationProposalVoteCancelled(uint256 indexed proposalId, address indexed user)
func (_Governance *GovernanceFilterer) WatchAbrogationProposalVoteCancelled(opts *bind.WatchOpts, sink chan<- *GovernanceAbrogationProposalVoteCancelled, proposalId []*big.Int, user []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "AbrogationProposalVoteCancelled", proposalIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceAbrogationProposalVoteCancelled)
				if err := _Governance.contract.UnpackLog(event, "AbrogationProposalVoteCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAbrogationProposalVoteCancelled is a log parse operation binding the contract event 0x5e8ee24f838173ed2ae7989835696f6e11945ac8fbc5259aef01cc4d7f0d4920.
//
// Solidity: event AbrogationProposalVoteCancelled(uint256 indexed proposalId, address indexed user)
func (_Governance *GovernanceFilterer) ParseAbrogationProposalVoteCancelled(log types.Log) (*GovernanceAbrogationProposalVoteCancelled, error) {
	event := new(GovernanceAbrogationProposalVoteCancelled)
	if err := _Governance.contract.UnpackLog(event, "AbrogationProposalVoteCancelled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceProposalCanceledIterator is returned from FilterProposalCanceled and is used to iterate over the raw logs and unpacked data for ProposalCanceled events raised by the Governance contract.
type GovernanceProposalCanceledIterator struct {
	Event *GovernanceProposalCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceProposalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceProposalCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceProposalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalCanceled represents a ProposalCanceled event raised by the Governance contract.
type GovernanceProposalCanceled struct {
	ProposalId *big.Int
	Caller     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCanceled is a free log retrieval operation binding the contract event 0x253042c67143aeb6d431bb762d75e5905f18fa7850b7b9edb31fedb7c362d7e8.
//
// Solidity: event ProposalCanceled(uint256 indexed proposalId, address caller)
func (_Governance *GovernanceFilterer) FilterProposalCanceled(opts *bind.FilterOpts, proposalId []*big.Int) (*GovernanceProposalCanceledIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalCanceled", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalCanceledIterator{contract: _Governance.contract, event: "ProposalCanceled", logs: logs, sub: sub}, nil
}

// WatchProposalCanceled is a free log subscription operation binding the contract event 0x253042c67143aeb6d431bb762d75e5905f18fa7850b7b9edb31fedb7c362d7e8.
//
// Solidity: event ProposalCanceled(uint256 indexed proposalId, address caller)
func (_Governance *GovernanceFilterer) WatchProposalCanceled(opts *bind.WatchOpts, sink chan<- *GovernanceProposalCanceled, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalCanceled", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalCanceled)
				if err := _Governance.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalCanceled is a log parse operation binding the contract event 0x253042c67143aeb6d431bb762d75e5905f18fa7850b7b9edb31fedb7c362d7e8.
//
// Solidity: event ProposalCanceled(uint256 indexed proposalId, address caller)
func (_Governance *GovernanceFilterer) ParseProposalCanceled(log types.Log) (*GovernanceProposalCanceled, error) {
	event := new(GovernanceProposalCanceled)
	if err := _Governance.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the Governance contract.
type GovernanceProposalCreatedIterator struct {
	Event *GovernanceProposalCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceProposalCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalCreated represents a ProposalCreated event raised by the Governance contract.
type GovernanceProposalCreated struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) FilterProposalCreated(opts *bind.FilterOpts, proposalId []*big.Int) (*GovernanceProposalCreatedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalCreated", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalCreatedIterator{contract: _Governance.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *GovernanceProposalCreated, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalCreated", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalCreated)
				if err := _Governance.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalCreated is a log parse operation binding the contract event 0xc2c021f5d73c63c481d336fbbafec58f694fc45095f00b02d2deb8cca59afe07.
//
// Solidity: event ProposalCreated(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) ParseProposalCreated(log types.Log) (*GovernanceProposalCreated, error) {
	event := new(GovernanceProposalCreated)
	if err := _Governance.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the Governance contract.
type GovernanceProposalExecutedIterator struct {
	Event *GovernanceProposalExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceProposalExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalExecuted represents a ProposalExecuted event raised by the Governance contract.
type GovernanceProposalExecuted struct {
	ProposalId *big.Int
	Caller     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x9c85b616f29fca57a17eafe71cf9ff82ffef41766e2cf01ea7f8f7878dd3ec24.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalId, address caller)
func (_Governance *GovernanceFilterer) FilterProposalExecuted(opts *bind.FilterOpts, proposalId []*big.Int) (*GovernanceProposalExecutedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalExecutedIterator{contract: _Governance.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x9c85b616f29fca57a17eafe71cf9ff82ffef41766e2cf01ea7f8f7878dd3ec24.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalId, address caller)
func (_Governance *GovernanceFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *GovernanceProposalExecuted, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalExecuted)
				if err := _Governance.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalExecuted is a log parse operation binding the contract event 0x9c85b616f29fca57a17eafe71cf9ff82ffef41766e2cf01ea7f8f7878dd3ec24.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalId, address caller)
func (_Governance *GovernanceFilterer) ParseProposalExecuted(log types.Log) (*GovernanceProposalExecuted, error) {
	event := new(GovernanceProposalExecuted)
	if err := _Governance.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceProposalQueuedIterator is returned from FilterProposalQueued and is used to iterate over the raw logs and unpacked data for ProposalQueued events raised by the Governance contract.
type GovernanceProposalQueuedIterator struct {
	Event *GovernanceProposalQueued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceProposalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalQueued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceProposalQueued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceProposalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalQueued represents a ProposalQueued event raised by the Governance contract.
type GovernanceProposalQueued struct {
	ProposalId *big.Int
	Caller     common.Address
	Eta        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalQueued is a free log retrieval operation binding the contract event 0xf7230a453b4c21e4f2d0ef1ad055635b08cb2c884eaf24a5ddc7147c79fd8c22.
//
// Solidity: event ProposalQueued(uint256 indexed proposalId, address caller, uint256 eta)
func (_Governance *GovernanceFilterer) FilterProposalQueued(opts *bind.FilterOpts, proposalId []*big.Int) (*GovernanceProposalQueuedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalQueued", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalQueuedIterator{contract: _Governance.contract, event: "ProposalQueued", logs: logs, sub: sub}, nil
}

// WatchProposalQueued is a free log subscription operation binding the contract event 0xf7230a453b4c21e4f2d0ef1ad055635b08cb2c884eaf24a5ddc7147c79fd8c22.
//
// Solidity: event ProposalQueued(uint256 indexed proposalId, address caller, uint256 eta)
func (_Governance *GovernanceFilterer) WatchProposalQueued(opts *bind.WatchOpts, sink chan<- *GovernanceProposalQueued, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalQueued", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalQueued)
				if err := _Governance.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalQueued is a log parse operation binding the contract event 0xf7230a453b4c21e4f2d0ef1ad055635b08cb2c884eaf24a5ddc7147c79fd8c22.
//
// Solidity: event ProposalQueued(uint256 indexed proposalId, address caller, uint256 eta)
func (_Governance *GovernanceFilterer) ParseProposalQueued(log types.Log) (*GovernanceProposalQueued, error) {
	event := new(GovernanceProposalQueued)
	if err := _Governance.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the Governance contract.
type GovernanceVoteIterator struct {
	Event *GovernanceVote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceVote)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceVote)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceVote represents a Vote event raised by the Governance contract.
type GovernanceVote struct {
	ProposalId *big.Int
	User       common.Address
	Support    bool
	Power      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0x88d35328232823f54954b6627e9f732371656f6daa40cb1b01b27dc7875a7b47.
//
// Solidity: event Vote(uint256 indexed proposalId, address indexed user, bool support, uint256 power)
func (_Governance *GovernanceFilterer) FilterVote(opts *bind.FilterOpts, proposalId []*big.Int, user []common.Address) (*GovernanceVoteIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Vote", proposalIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceVoteIterator{contract: _Governance.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0x88d35328232823f54954b6627e9f732371656f6daa40cb1b01b27dc7875a7b47.
//
// Solidity: event Vote(uint256 indexed proposalId, address indexed user, bool support, uint256 power)
func (_Governance *GovernanceFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *GovernanceVote, proposalId []*big.Int, user []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Vote", proposalIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceVote)
				if err := _Governance.contract.UnpackLog(event, "Vote", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVote is a log parse operation binding the contract event 0x88d35328232823f54954b6627e9f732371656f6daa40cb1b01b27dc7875a7b47.
//
// Solidity: event Vote(uint256 indexed proposalId, address indexed user, bool support, uint256 power)
func (_Governance *GovernanceFilterer) ParseVote(log types.Log) (*GovernanceVote, error) {
	event := new(GovernanceVote)
	if err := _Governance.contract.UnpackLog(event, "Vote", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceVoteCanceledIterator is returned from FilterVoteCanceled and is used to iterate over the raw logs and unpacked data for VoteCanceled events raised by the Governance contract.
type GovernanceVoteCanceledIterator struct {
	Event *GovernanceVoteCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceVoteCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceVoteCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceVoteCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceVoteCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceVoteCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceVoteCanceled represents a VoteCanceled event raised by the Governance contract.
type GovernanceVoteCanceled struct {
	ProposalId *big.Int
	User       common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCanceled is a free log retrieval operation binding the contract event 0x12beef84830227673717dd5522ee1228a8004e88dc2678d8740f582264efb2b6.
//
// Solidity: event VoteCanceled(uint256 indexed proposalId, address indexed user)
func (_Governance *GovernanceFilterer) FilterVoteCanceled(opts *bind.FilterOpts, proposalId []*big.Int, user []common.Address) (*GovernanceVoteCanceledIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "VoteCanceled", proposalIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceVoteCanceledIterator{contract: _Governance.contract, event: "VoteCanceled", logs: logs, sub: sub}, nil
}

// WatchVoteCanceled is a free log subscription operation binding the contract event 0x12beef84830227673717dd5522ee1228a8004e88dc2678d8740f582264efb2b6.
//
// Solidity: event VoteCanceled(uint256 indexed proposalId, address indexed user)
func (_Governance *GovernanceFilterer) WatchVoteCanceled(opts *bind.WatchOpts, sink chan<- *GovernanceVoteCanceled, proposalId []*big.Int, user []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "VoteCanceled", proposalIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceVoteCanceled)
				if err := _Governance.contract.UnpackLog(event, "VoteCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteCanceled is a log parse operation binding the contract event 0x12beef84830227673717dd5522ee1228a8004e88dc2678d8740f582264efb2b6.
//
// Solidity: event VoteCanceled(uint256 indexed proposalId, address indexed user)
func (_Governance *GovernanceFilterer) ParseVoteCanceled(log types.Log) (*GovernanceVoteCanceled, error) {
	event := new(GovernanceVoteCanceled)
	if err := _Governance.contract.UnpackLog(event, "VoteCanceled", log); err != nil {
		return nil, err
	}
	return event, nil
}
