// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CommitRevealMetaData contains all meta data concerning the CommitReveal contract.
var CommitRevealMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_phaseLengthInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_choice1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_choice2\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"commit\",\"type\":\"bytes32\"}],\"name\":\"NewVoteCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"commit\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"choice\",\"type\":\"string\"}],\"name\":\"NewVoteReveal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"choice1\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"choice2\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commitPhaseEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"voteCommit\",\"type\":\"bytes32\"}],\"name\":\"commitVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVoteCommitsArray\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWinner\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"winner\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfVotesCast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"vote\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"voteCommit\",\"type\":\"bytes32\"}],\"name\":\"revealVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voteCommits\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"voteStatuses\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votesForChoice1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votesForChoice2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CommitRevealABI is the input ABI used to generate the binding from.
// Deprecated: Use CommitRevealMetaData.ABI instead.
var CommitRevealABI = CommitRevealMetaData.ABI

// CommitReveal is an auto generated Go binding around an Ethereum contract.
type CommitReveal struct {
	CommitRevealCaller     // Read-only binding to the contract
	CommitRevealTransactor // Write-only binding to the contract
	CommitRevealFilterer   // Log filterer for contract events
}

// CommitRevealCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommitRevealCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitRevealTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommitRevealTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitRevealFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommitRevealFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitRevealSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommitRevealSession struct {
	Contract     *CommitReveal     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommitRevealCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommitRevealCallerSession struct {
	Contract *CommitRevealCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CommitRevealTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommitRevealTransactorSession struct {
	Contract     *CommitRevealTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CommitRevealRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommitRevealRaw struct {
	Contract *CommitReveal // Generic contract binding to access the raw methods on
}

// CommitRevealCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommitRevealCallerRaw struct {
	Contract *CommitRevealCaller // Generic read-only contract binding to access the raw methods on
}

// CommitRevealTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommitRevealTransactorRaw struct {
	Contract *CommitRevealTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommitReveal creates a new instance of CommitReveal, bound to a specific deployed contract.
func NewCommitReveal(address common.Address, backend bind.ContractBackend) (*CommitReveal, error) {
	contract, err := bindCommitReveal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommitReveal{CommitRevealCaller: CommitRevealCaller{contract: contract}, CommitRevealTransactor: CommitRevealTransactor{contract: contract}, CommitRevealFilterer: CommitRevealFilterer{contract: contract}}, nil
}

// NewCommitRevealCaller creates a new read-only instance of CommitReveal, bound to a specific deployed contract.
func NewCommitRevealCaller(address common.Address, caller bind.ContractCaller) (*CommitRevealCaller, error) {
	contract, err := bindCommitReveal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitRevealCaller{contract: contract}, nil
}

// NewCommitRevealTransactor creates a new write-only instance of CommitReveal, bound to a specific deployed contract.
func NewCommitRevealTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitRevealTransactor, error) {
	contract, err := bindCommitReveal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitRevealTransactor{contract: contract}, nil
}

// NewCommitRevealFilterer creates a new log filterer instance of CommitReveal, bound to a specific deployed contract.
func NewCommitRevealFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitRevealFilterer, error) {
	contract, err := bindCommitReveal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitRevealFilterer{contract: contract}, nil
}

// bindCommitReveal binds a generic wrapper to an already deployed contract.
func bindCommitReveal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommitRevealMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommitReveal *CommitRevealRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitReveal.Contract.CommitRevealCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommitReveal *CommitRevealRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitReveal.Contract.CommitRevealTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommitReveal *CommitRevealRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitReveal.Contract.CommitRevealTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommitReveal *CommitRevealCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitReveal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommitReveal *CommitRevealTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitReveal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommitReveal *CommitRevealTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitReveal.Contract.contract.Transact(opts, method, params...)
}

// Choice1 is a free data retrieval call binding the contract method 0x49e811c5.
//
// Solidity: function choice1() view returns(string)
func (_CommitReveal *CommitRevealCaller) Choice1(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CommitReveal.contract.Call(opts, &out, "choice1")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Choice1 is a free data retrieval call binding the contract method 0x49e811c5.
//
// Solidity: function choice1() view returns(string)
func (_CommitReveal *CommitRevealSession) Choice1() (string, error) {
	return _CommitReveal.Contract.Choice1(&_CommitReveal.CallOpts)
}

// Choice1 is a free data retrieval call binding the contract method 0x49e811c5.
//
// Solidity: function choice1() view returns(string)
func (_CommitReveal *CommitRevealCallerSession) Choice1() (string, error) {
	return _CommitReveal.Contract.Choice1(&_CommitReveal.CallOpts)
}

// Choice2 is a free data retrieval call binding the contract method 0x87cfa5bf.
//
// Solidity: function choice2() view returns(string)
func (_CommitReveal *CommitRevealCaller) Choice2(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CommitReveal.contract.Call(opts, &out, "choice2")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Choice2 is a free data retrieval call binding the contract method 0x87cfa5bf.
//
// Solidity: function choice2() view returns(string)
func (_CommitReveal *CommitRevealSession) Choice2() (string, error) {
	return _CommitReveal.Contract.Choice2(&_CommitReveal.CallOpts)
}

// Choice2 is a free data retrieval call binding the contract method 0x87cfa5bf.
//
// Solidity: function choice2() view returns(string)
func (_CommitReveal *CommitRevealCallerSession) Choice2() (string, error) {
	return _CommitReveal.Contract.Choice2(&_CommitReveal.CallOpts)
}

// CommitPhaseEndTime is a free data retrieval call binding the contract method 0xa6193504.
//
// Solidity: function commitPhaseEndTime() view returns(uint256)
func (_CommitReveal *CommitRevealCaller) CommitPhaseEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CommitReveal.contract.Call(opts, &out, "commitPhaseEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommitPhaseEndTime is a free data retrieval call binding the contract method 0xa6193504.
//
// Solidity: function commitPhaseEndTime() view returns(uint256)
func (_CommitReveal *CommitRevealSession) CommitPhaseEndTime() (*big.Int, error) {
	return _CommitReveal.Contract.CommitPhaseEndTime(&_CommitReveal.CallOpts)
}

// CommitPhaseEndTime is a free data retrieval call binding the contract method 0xa6193504.
//
// Solidity: function commitPhaseEndTime() view returns(uint256)
func (_CommitReveal *CommitRevealCallerSession) CommitPhaseEndTime() (*big.Int, error) {
	return _CommitReveal.Contract.CommitPhaseEndTime(&_CommitReveal.CallOpts)
}

// GetVoteCommitsArray is a free data retrieval call binding the contract method 0x3a2a2cc5.
//
// Solidity: function getVoteCommitsArray() view returns(bytes32[])
func (_CommitReveal *CommitRevealCaller) GetVoteCommitsArray(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _CommitReveal.contract.Call(opts, &out, "getVoteCommitsArray")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetVoteCommitsArray is a free data retrieval call binding the contract method 0x3a2a2cc5.
//
// Solidity: function getVoteCommitsArray() view returns(bytes32[])
func (_CommitReveal *CommitRevealSession) GetVoteCommitsArray() ([][32]byte, error) {
	return _CommitReveal.Contract.GetVoteCommitsArray(&_CommitReveal.CallOpts)
}

// GetVoteCommitsArray is a free data retrieval call binding the contract method 0x3a2a2cc5.
//
// Solidity: function getVoteCommitsArray() view returns(bytes32[])
func (_CommitReveal *CommitRevealCallerSession) GetVoteCommitsArray() ([][32]byte, error) {
	return _CommitReveal.Contract.GetVoteCommitsArray(&_CommitReveal.CallOpts)
}

// GetWinner is a free data retrieval call binding the contract method 0x8e7ea5b2.
//
// Solidity: function getWinner() view returns(string winner)
func (_CommitReveal *CommitRevealCaller) GetWinner(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CommitReveal.contract.Call(opts, &out, "getWinner")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetWinner is a free data retrieval call binding the contract method 0x8e7ea5b2.
//
// Solidity: function getWinner() view returns(string winner)
func (_CommitReveal *CommitRevealSession) GetWinner() (string, error) {
	return _CommitReveal.Contract.GetWinner(&_CommitReveal.CallOpts)
}

// GetWinner is a free data retrieval call binding the contract method 0x8e7ea5b2.
//
// Solidity: function getWinner() view returns(string winner)
func (_CommitReveal *CommitRevealCallerSession) GetWinner() (string, error) {
	return _CommitReveal.Contract.GetWinner(&_CommitReveal.CallOpts)
}

// NumberOfVotesCast is a free data retrieval call binding the contract method 0x22ce9bac.
//
// Solidity: function numberOfVotesCast() view returns(uint256)
func (_CommitReveal *CommitRevealCaller) NumberOfVotesCast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CommitReveal.contract.Call(opts, &out, "numberOfVotesCast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfVotesCast is a free data retrieval call binding the contract method 0x22ce9bac.
//
// Solidity: function numberOfVotesCast() view returns(uint256)
func (_CommitReveal *CommitRevealSession) NumberOfVotesCast() (*big.Int, error) {
	return _CommitReveal.Contract.NumberOfVotesCast(&_CommitReveal.CallOpts)
}

// NumberOfVotesCast is a free data retrieval call binding the contract method 0x22ce9bac.
//
// Solidity: function numberOfVotesCast() view returns(uint256)
func (_CommitReveal *CommitRevealCallerSession) NumberOfVotesCast() (*big.Int, error) {
	return _CommitReveal.Contract.NumberOfVotesCast(&_CommitReveal.CallOpts)
}

// VoteCommits is a free data retrieval call binding the contract method 0xbf9d0698.
//
// Solidity: function voteCommits(uint256 ) view returns(bytes32)
func (_CommitReveal *CommitRevealCaller) VoteCommits(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _CommitReveal.contract.Call(opts, &out, "voteCommits", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VoteCommits is a free data retrieval call binding the contract method 0xbf9d0698.
//
// Solidity: function voteCommits(uint256 ) view returns(bytes32)
func (_CommitReveal *CommitRevealSession) VoteCommits(arg0 *big.Int) ([32]byte, error) {
	return _CommitReveal.Contract.VoteCommits(&_CommitReveal.CallOpts, arg0)
}

// VoteCommits is a free data retrieval call binding the contract method 0xbf9d0698.
//
// Solidity: function voteCommits(uint256 ) view returns(bytes32)
func (_CommitReveal *CommitRevealCallerSession) VoteCommits(arg0 *big.Int) ([32]byte, error) {
	return _CommitReveal.Contract.VoteCommits(&_CommitReveal.CallOpts, arg0)
}

// VoteStatuses is a free data retrieval call binding the contract method 0xfd5bd01d.
//
// Solidity: function voteStatuses(bytes32 ) view returns(string)
func (_CommitReveal *CommitRevealCaller) VoteStatuses(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var out []interface{}
	err := _CommitReveal.contract.Call(opts, &out, "voteStatuses", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VoteStatuses is a free data retrieval call binding the contract method 0xfd5bd01d.
//
// Solidity: function voteStatuses(bytes32 ) view returns(string)
func (_CommitReveal *CommitRevealSession) VoteStatuses(arg0 [32]byte) (string, error) {
	return _CommitReveal.Contract.VoteStatuses(&_CommitReveal.CallOpts, arg0)
}

// VoteStatuses is a free data retrieval call binding the contract method 0xfd5bd01d.
//
// Solidity: function voteStatuses(bytes32 ) view returns(string)
func (_CommitReveal *CommitRevealCallerSession) VoteStatuses(arg0 [32]byte) (string, error) {
	return _CommitReveal.Contract.VoteStatuses(&_CommitReveal.CallOpts, arg0)
}

// VotesForChoice1 is a free data retrieval call binding the contract method 0xbaa13bb5.
//
// Solidity: function votesForChoice1() view returns(uint256)
func (_CommitReveal *CommitRevealCaller) VotesForChoice1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CommitReveal.contract.Call(opts, &out, "votesForChoice1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotesForChoice1 is a free data retrieval call binding the contract method 0xbaa13bb5.
//
// Solidity: function votesForChoice1() view returns(uint256)
func (_CommitReveal *CommitRevealSession) VotesForChoice1() (*big.Int, error) {
	return _CommitReveal.Contract.VotesForChoice1(&_CommitReveal.CallOpts)
}

// VotesForChoice1 is a free data retrieval call binding the contract method 0xbaa13bb5.
//
// Solidity: function votesForChoice1() view returns(uint256)
func (_CommitReveal *CommitRevealCallerSession) VotesForChoice1() (*big.Int, error) {
	return _CommitReveal.Contract.VotesForChoice1(&_CommitReveal.CallOpts)
}

// VotesForChoice2 is a free data retrieval call binding the contract method 0x031b6dbf.
//
// Solidity: function votesForChoice2() view returns(uint256)
func (_CommitReveal *CommitRevealCaller) VotesForChoice2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CommitReveal.contract.Call(opts, &out, "votesForChoice2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotesForChoice2 is a free data retrieval call binding the contract method 0x031b6dbf.
//
// Solidity: function votesForChoice2() view returns(uint256)
func (_CommitReveal *CommitRevealSession) VotesForChoice2() (*big.Int, error) {
	return _CommitReveal.Contract.VotesForChoice2(&_CommitReveal.CallOpts)
}

// VotesForChoice2 is a free data retrieval call binding the contract method 0x031b6dbf.
//
// Solidity: function votesForChoice2() view returns(uint256)
func (_CommitReveal *CommitRevealCallerSession) VotesForChoice2() (*big.Int, error) {
	return _CommitReveal.Contract.VotesForChoice2(&_CommitReveal.CallOpts)
}

// CommitVote is a paid mutator transaction binding the contract method 0x3e858923.
//
// Solidity: function commitVote(bytes32 voteCommit) returns()
func (_CommitReveal *CommitRevealTransactor) CommitVote(opts *bind.TransactOpts, voteCommit [32]byte) (*types.Transaction, error) {
	return _CommitReveal.contract.Transact(opts, "commitVote", voteCommit)
}

// CommitVote is a paid mutator transaction binding the contract method 0x3e858923.
//
// Solidity: function commitVote(bytes32 voteCommit) returns()
func (_CommitReveal *CommitRevealSession) CommitVote(voteCommit [32]byte) (*types.Transaction, error) {
	return _CommitReveal.Contract.CommitVote(&_CommitReveal.TransactOpts, voteCommit)
}

// CommitVote is a paid mutator transaction binding the contract method 0x3e858923.
//
// Solidity: function commitVote(bytes32 voteCommit) returns()
func (_CommitReveal *CommitRevealTransactorSession) CommitVote(voteCommit [32]byte) (*types.Transaction, error) {
	return _CommitReveal.Contract.CommitVote(&_CommitReveal.TransactOpts, voteCommit)
}

// RevealVote is a paid mutator transaction binding the contract method 0x0bbfc206.
//
// Solidity: function revealVote(string vote, bytes32 voteCommit) returns()
func (_CommitReveal *CommitRevealTransactor) RevealVote(opts *bind.TransactOpts, vote string, voteCommit [32]byte) (*types.Transaction, error) {
	return _CommitReveal.contract.Transact(opts, "revealVote", vote, voteCommit)
}

// RevealVote is a paid mutator transaction binding the contract method 0x0bbfc206.
//
// Solidity: function revealVote(string vote, bytes32 voteCommit) returns()
func (_CommitReveal *CommitRevealSession) RevealVote(vote string, voteCommit [32]byte) (*types.Transaction, error) {
	return _CommitReveal.Contract.RevealVote(&_CommitReveal.TransactOpts, vote, voteCommit)
}

// RevealVote is a paid mutator transaction binding the contract method 0x0bbfc206.
//
// Solidity: function revealVote(string vote, bytes32 voteCommit) returns()
func (_CommitReveal *CommitRevealTransactorSession) RevealVote(vote string, voteCommit [32]byte) (*types.Transaction, error) {
	return _CommitReveal.Contract.RevealVote(&_CommitReveal.TransactOpts, vote, voteCommit)
}

// CommitRevealNewVoteCommitIterator is returned from FilterNewVoteCommit and is used to iterate over the raw logs and unpacked data for NewVoteCommit events raised by the CommitReveal contract.
type CommitRevealNewVoteCommitIterator struct {
	Event *CommitRevealNewVoteCommit // Event containing the contract specifics and raw log

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
func (it *CommitRevealNewVoteCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitRevealNewVoteCommit)
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
		it.Event = new(CommitRevealNewVoteCommit)
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
func (it *CommitRevealNewVoteCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitRevealNewVoteCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitRevealNewVoteCommit represents a NewVoteCommit event raised by the CommitReveal contract.
type CommitRevealNewVoteCommit struct {
	Commit [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewVoteCommit is a free log retrieval operation binding the contract event 0x336dd6fa514fe196ac01f0e19a495fc73b7885d9357f993cf52e4919f70db099.
//
// Solidity: event NewVoteCommit(bytes32 commit)
func (_CommitReveal *CommitRevealFilterer) FilterNewVoteCommit(opts *bind.FilterOpts) (*CommitRevealNewVoteCommitIterator, error) {

	logs, sub, err := _CommitReveal.contract.FilterLogs(opts, "NewVoteCommit")
	if err != nil {
		return nil, err
	}
	return &CommitRevealNewVoteCommitIterator{contract: _CommitReveal.contract, event: "NewVoteCommit", logs: logs, sub: sub}, nil
}

// WatchNewVoteCommit is a free log subscription operation binding the contract event 0x336dd6fa514fe196ac01f0e19a495fc73b7885d9357f993cf52e4919f70db099.
//
// Solidity: event NewVoteCommit(bytes32 commit)
func (_CommitReveal *CommitRevealFilterer) WatchNewVoteCommit(opts *bind.WatchOpts, sink chan<- *CommitRevealNewVoteCommit) (event.Subscription, error) {

	logs, sub, err := _CommitReveal.contract.WatchLogs(opts, "NewVoteCommit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitRevealNewVoteCommit)
				if err := _CommitReveal.contract.UnpackLog(event, "NewVoteCommit", log); err != nil {
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

// ParseNewVoteCommit is a log parse operation binding the contract event 0x336dd6fa514fe196ac01f0e19a495fc73b7885d9357f993cf52e4919f70db099.
//
// Solidity: event NewVoteCommit(bytes32 commit)
func (_CommitReveal *CommitRevealFilterer) ParseNewVoteCommit(log types.Log) (*CommitRevealNewVoteCommit, error) {
	event := new(CommitRevealNewVoteCommit)
	if err := _CommitReveal.contract.UnpackLog(event, "NewVoteCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommitRevealNewVoteRevealIterator is returned from FilterNewVoteReveal and is used to iterate over the raw logs and unpacked data for NewVoteReveal events raised by the CommitReveal contract.
type CommitRevealNewVoteRevealIterator struct {
	Event *CommitRevealNewVoteReveal // Event containing the contract specifics and raw log

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
func (it *CommitRevealNewVoteRevealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitRevealNewVoteReveal)
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
		it.Event = new(CommitRevealNewVoteReveal)
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
func (it *CommitRevealNewVoteRevealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitRevealNewVoteRevealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitRevealNewVoteReveal represents a NewVoteReveal event raised by the CommitReveal contract.
type CommitRevealNewVoteReveal struct {
	Commit [32]byte
	Choice string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewVoteReveal is a free log retrieval operation binding the contract event 0x69ab1b0836124ec1de36500cc76e0bcd684f01f02d992b7b26f6ad222a2ddf21.
//
// Solidity: event NewVoteReveal(bytes32 commit, string choice)
func (_CommitReveal *CommitRevealFilterer) FilterNewVoteReveal(opts *bind.FilterOpts) (*CommitRevealNewVoteRevealIterator, error) {

	logs, sub, err := _CommitReveal.contract.FilterLogs(opts, "NewVoteReveal")
	if err != nil {
		return nil, err
	}
	return &CommitRevealNewVoteRevealIterator{contract: _CommitReveal.contract, event: "NewVoteReveal", logs: logs, sub: sub}, nil
}

// WatchNewVoteReveal is a free log subscription operation binding the contract event 0x69ab1b0836124ec1de36500cc76e0bcd684f01f02d992b7b26f6ad222a2ddf21.
//
// Solidity: event NewVoteReveal(bytes32 commit, string choice)
func (_CommitReveal *CommitRevealFilterer) WatchNewVoteReveal(opts *bind.WatchOpts, sink chan<- *CommitRevealNewVoteReveal) (event.Subscription, error) {

	logs, sub, err := _CommitReveal.contract.WatchLogs(opts, "NewVoteReveal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitRevealNewVoteReveal)
				if err := _CommitReveal.contract.UnpackLog(event, "NewVoteReveal", log); err != nil {
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

// ParseNewVoteReveal is a log parse operation binding the contract event 0x69ab1b0836124ec1de36500cc76e0bcd684f01f02d992b7b26f6ad222a2ddf21.
//
// Solidity: event NewVoteReveal(bytes32 commit, string choice)
func (_CommitReveal *CommitRevealFilterer) ParseNewVoteReveal(log types.Log) (*CommitRevealNewVoteReveal, error) {
	event := new(CommitRevealNewVoteReveal)
	if err := _CommitReveal.contract.UnpackLog(event, "NewVoteReveal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
