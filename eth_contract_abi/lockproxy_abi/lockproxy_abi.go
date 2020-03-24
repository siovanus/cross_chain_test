// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lockproxy_abi

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ERC20InterfaceABI is the input ABI used to generate the binding from.
const ERC20InterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20InterfaceFuncSigs maps the 4-byte function signature to its string representation.
var ERC20InterfaceFuncSigs = map[string]string{
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20Interface is an auto generated Go binding around an Ethereum contract.
type ERC20Interface struct {
	ERC20InterfaceCaller     // Read-only binding to the contract
	ERC20InterfaceTransactor // Write-only binding to the contract
	ERC20InterfaceFilterer   // Log filterer for contract events
}

// ERC20InterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20InterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20InterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20InterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20InterfaceSession struct {
	Contract     *ERC20Interface   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20InterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20InterfaceCallerSession struct {
	Contract *ERC20InterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC20InterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20InterfaceTransactorSession struct {
	Contract     *ERC20InterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC20InterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20InterfaceRaw struct {
	Contract *ERC20Interface // Generic contract binding to access the raw methods on
}

// ERC20InterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20InterfaceCallerRaw struct {
	Contract *ERC20InterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20InterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20InterfaceTransactorRaw struct {
	Contract *ERC20InterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Interface creates a new instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20Interface(address common.Address, backend bind.ContractBackend) (*ERC20Interface, error) {
	contract, err := bindERC20Interface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Interface{ERC20InterfaceCaller: ERC20InterfaceCaller{contract: contract}, ERC20InterfaceTransactor: ERC20InterfaceTransactor{contract: contract}, ERC20InterfaceFilterer: ERC20InterfaceFilterer{contract: contract}}, nil
}

// NewERC20InterfaceCaller creates a new read-only instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20InterfaceCaller(address common.Address, caller bind.ContractCaller) (*ERC20InterfaceCaller, error) {
	contract, err := bindERC20Interface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceCaller{contract: contract}, nil
}

// NewERC20InterfaceTransactor creates a new write-only instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20InterfaceTransactor, error) {
	contract, err := bindERC20Interface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceTransactor{contract: contract}, nil
}

// NewERC20InterfaceFilterer creates a new log filterer instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20InterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20InterfaceFilterer, error) {
	contract, err := bindERC20Interface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceFilterer{contract: contract}, nil
}

// bindERC20Interface binds a generic wrapper to an already deployed contract.
func bindERC20Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Interface *ERC20InterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Interface.Contract.ERC20InterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Interface *ERC20InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Interface.Contract.ERC20InterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Interface *ERC20InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Interface.Contract.ERC20InterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Interface *ERC20InterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Interface *ERC20InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Interface *ERC20InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Interface.Contract.contract.Transact(opts, method, params...)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20Interface *ERC20InterfaceTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20Interface *ERC20InterfaceSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.Transfer(&_ERC20Interface.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20Interface *ERC20InterfaceTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.Transfer(&_ERC20Interface.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_ERC20Interface *ERC20InterfaceTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_ERC20Interface *ERC20InterfaceSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.TransferFrom(&_ERC20Interface.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_ERC20Interface *ERC20InterfaceTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.TransferFrom(&_ERC20Interface.TransactOpts, _from, _to, _value)
}

// EthCrossChainABI is the input ABI used to generate the binding from.
const EthCrossChainABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_rawHeader\",\"type\":\"bytes\"},{\"name\":\"_sigList\",\"type\":\"bytes\"}],\"name\":\"SyncBlockHeader\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rawHeader\",\"type\":\"bytes\"},{\"name\":\"_pubKeyList\",\"type\":\"bytes\"},{\"name\":\"_sigList\",\"type\":\"bytes\"}],\"name\":\"ChangeBookKeeper\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Transactions\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MCKeeperPubKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rawHeader\",\"type\":\"bytes\"},{\"name\":\"sigList\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"name\":\"position\",\"type\":\"uint8[]\"},{\"name\":\"toMerkleValueBs\",\"type\":\"bytes\"},{\"name\":\"blockHeight\",\"type\":\"uint64\"}],\"name\":\"SyncAndVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"MCBlockHeaders\",\"outputs\":[{\"name\":\"version\",\"type\":\"uint32\"},{\"name\":\"chainId\",\"type\":\"uint64\"},{\"name\":\"prevBlockHash\",\"type\":\"bytes32\"},{\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"crossStatesRoot\",\"type\":\"bytes32\"},{\"name\":\"blockRoot\",\"type\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint32\"},{\"name\":\"height\",\"type\":\"uint32\"},{\"name\":\"consensusData\",\"type\":\"uint64\"},{\"name\":\"consensusPayload\",\"type\":\"bytes\"},{\"name\":\"nextBookkeeper\",\"type\":\"bytes20\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rawHeader\",\"type\":\"bytes\"},{\"name\":\"_pubKeyList\",\"type\":\"bytes\"}],\"name\":\"InitGenesisBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_toChainId\",\"type\":\"uint64\"},{\"name\":\"_toContract\",\"type\":\"bytes\"},{\"name\":\"_method\",\"type\":\"bytes\"},{\"name\":\"_txData\",\"type\":\"bytes\"}],\"name\":\"crossChain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TransactionId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MCKeeperHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proof\",\"type\":\"bytes32[]\"},{\"name\":\"_position\",\"type\":\"uint8[]\"},{\"name\":\"_toMerkleValueBs\",\"type\":\"bytes\"},{\"name\":\"_blockHeight\",\"type\":\"uint64\"}],\"name\":\"verifyAndExecuteTx\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LatestBookKeeperHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LatestHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"torawHeadersken\",\"type\":\"bytes\"}],\"name\":\"InitGenesisBlockEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rawHeaders\",\"type\":\"bytes\"}],\"name\":\"ChangeBookKeeperEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rawHeaders\",\"type\":\"bytes\"}],\"name\":\"SyncBlockHeaderEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"txId\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"toContract\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"rawdata\",\"type\":\"bytes\"}],\"name\":\"CrossChainEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fromChainID\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"toContract\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"crossChainTxHash\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"fromChainTxHash\",\"type\":\"bytes\"}],\"name\":\"VerifyAndExecuteTxEvent\",\"type\":\"event\"}]"

// EthCrossChainFuncSigs maps the 4-byte function signature to its string representation.
var EthCrossChainFuncSigs = map[string]string{
	"21f966aa": "ChangeBookKeeper(bytes,bytes,bytes)",
	"adbccaf5": "InitGenesisBlock(bytes,bytes)",
	"e244e4b1": "LatestBookKeeperHeight()",
	"e6aefed8": "LatestHeight()",
	"9a3337bc": "MCBlockHeaders(uint64)",
	"c6521e0f": "MCKeeperHeight(uint256)",
	"464f81a2": "MCKeeperPubKeys(uint64,uint256)",
	"48c2f675": "SyncAndVerify(bytes,bytes,bytes32[],uint8[],bytes,uint64)",
	"1e568e18": "SyncBlockHeader(bytes,bytes)",
	"c267dfcb": "TransactionId()",
	"34008b16": "Transactions(uint256)",
	"bd5cf625": "crossChain(uint64,bytes,bytes,bytes)",
	"d5ea92e3": "verifyAndExecuteTx(bytes32[],uint8[],bytes,uint64)",
}

// EthCrossChainBin is the compiled bytecode used for deploying new contracts.
var EthCrossChainBin = "0x608060405234801561001057600080fd5b50614d74806100206000396000f3fe6080604052600436106100c45763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631e568e1881146100c957806321f966aa1461021757806334008b16146103d8578063464f81a21461041457806348c2f6751461046a5780639a3337bc14610741578063adbccaf514610867578063bd5cf625146109a1578063c267dfcb14610b73578063c6521e0f14610b88578063d5ea92e314610bcf578063e244e4b114610d96578063e6aefed814610dab575b600080fd5b3480156100d557600080fd5b50610203600480360360408110156100ec57600080fd5b81019060208101813564010000000081111561010757600080fd5b82018360208201111561011957600080fd5b8035906020019184600183028401116401000000008311171561013b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561018e57600080fd5b8201836020820111156101a057600080fd5b803590602001918460018302840111640100000000831117156101c257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610dc0945050505050565b604080519115158252519081900360200190f35b34801561022357600080fd5b506102036004803603606081101561023a57600080fd5b81019060208101813564010000000081111561025557600080fd5b82018360208201111561026757600080fd5b8035906020019184600183028401116401000000008311171561028957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156102dc57600080fd5b8201836020820111156102ee57600080fd5b8035906020019184600183028401116401000000008311171561031057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561036357600080fd5b82018360208201111561037557600080fd5b8035906020019184600183028401116401000000008311171561039757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506114e3945050505050565b3480156103e457600080fd5b50610402600480360360208110156103fb57600080fd5b5035611a75565b60408051918252519081900360200190f35b34801561042057600080fd5b5061044e6004803603604081101561043757600080fd5b5067ffffffffffffffff8135169060200135611a87565b60408051600160a060020a039092168252519081900360200190f35b34801561047657600080fd5b5061073f600480360360c081101561048d57600080fd5b8101906020810181356401000000008111156104a857600080fd5b8201836020820111156104ba57600080fd5b803590602001918460018302840111640100000000831117156104dc57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561052f57600080fd5b82018360208201111561054157600080fd5b8035906020019184600183028401116401000000008311171561056357600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156105b657600080fd5b8201836020820111156105c857600080fd5b803590602001918460208302840111640100000000831117156105ea57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929594936020810193503591505064010000000081111561063a57600080fd5b82018360208201111561064c57600080fd5b8035906020019184602083028401116401000000008311171561066e57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092959493602081019350359150506401000000008111156106be57600080fd5b8201836020820111156106d057600080fd5b803590602001918460018302840111640100000000831117156106f257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903567ffffffffffffffff169150611abe9050565b005b34801561074d57600080fd5b506107756004803603602081101561076457600080fd5b503567ffffffffffffffff16611b88565b6040805163ffffffff808e16825267ffffffffffffffff808e166020808501919091529383018d9052606083018c9052608083018b905260a083018a905288821660c084015290871660e083015285166101008201526bffffffffffffffffffffffff198316610140820152610160610120820181815285519183019190915284519192909161018084019186019080838360005b8381101561082257818101518382015260200161080a565b50505050905090810190601f16801561084f5780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060405180910390f35b34801561087357600080fd5b5061073f6004803603604081101561088a57600080fd5b8101906020810181356401000000008111156108a557600080fd5b8201836020820111156108b757600080fd5b803590602001918460018302840111640100000000831117156108d957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561092c57600080fd5b82018360208201111561093e57600080fd5b8035906020019184600183028401116401000000008311171561096057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611c99945050505050565b3480156109ad57600080fd5b50610203600480360360808110156109c457600080fd5b67ffffffffffffffff82351691908101906040810160208201356401000000008111156109f057600080fd5b820183602082011115610a0257600080fd5b80359060200191846001830284011164010000000083111715610a2457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050640100000000811115610a7757600080fd5b820183602082011115610a8957600080fd5b80359060200191846001830284011164010000000083111715610aab57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050640100000000811115610afe57600080fd5b820183602082011115610b1057600080fd5b80359060200191846001830284011164010000000083111715610b3257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506120eb945050505050565b348015610b7f57600080fd5b506104026125a5565b348015610b9457600080fd5b50610bb260048036036020811015610bab57600080fd5b50356125ab565b6040805167ffffffffffffffff9092168252519081900360200190f35b348015610bdb57600080fd5b5061020360048036036080811015610bf257600080fd5b810190602081018135640100000000811115610c0d57600080fd5b820183602082011115610c1f57600080fd5b80359060200191846020830284011164010000000083111715610c4157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050640100000000811115610c9157600080fd5b820183602082011115610ca357600080fd5b80359060200191846020830284011164010000000083111715610cc557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050640100000000811115610d1557600080fd5b820183602082011115610d2757600080fd5b80359060200191846001830284011164010000000083111715610d4957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903567ffffffffffffffff1691506125e79050565b348015610da257600080fd5b50610bb2612aaf565b348015610db757600080fd5b50610bb2612abf565b6000610dca614b25565b610dd384612acf565b6101408101519091506bffffffffffffffffffffffff191615610e66576040805160e560020a62461bcd02815260206004820152602960248201527f4865616465722e6e6578746e657874426f6f6b6b65657065722073686f756c6460448201527f20626520656d7074790000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008160e0015163ffffffff16111515610eca576040805160e560020a62461bcd02815260206004820152601760248201527f626c6f636b206865696768742073686f756c64203e2030000000000000000000604482015290519081900360640190fd5b610ed2614b25565b60e08281015163ffffffff9081166000908152600260208181526040928390208351610160810185528154808716825267ffffffffffffffff640100000000918290048116838601526001808501548489015284870154606085015260038501546080850152600485015460a08501526005850154808a1660c086015292830490981698830198909852680100000000000000009004909616610100808801919091526006820180548651978116159092026000190190911693909304601f810183900483028601830190945283855293610120860193909291830182828015610ffd5780601f10610fd257610100808354040283529160200191610ffd565b820191906000526020600020905b815481529060010190602001808311610fe057829003601f168201915b5050509183525050600791909101546c01000000000000000000000000026bffffffffffffffffffffffff191660209091015260c081015190915063ffffffff161561104e576001925050506114dd565b60065460e083015160609167ffffffffffffffff1663ffffffff909116106110e75760065467ffffffffffffffff16600090815260056020908152604091829020805483518184028101840190945280845290918301828280156110db57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116110bd575b5050505050905061122f565b600080611104600480805490508760e0015163ffffffff16612bf2565b91509150801515611185576040805160e560020a62461bcd02815260206004820152603460248201527f63616e206e6f742066696e6420626c6f636b2068656967687420696e20626f6f60448201527f6b206b656570657220686569676874206c697374000000000000000000000000606482015290519081900360840190fd5b600060048367ffffffffffffffff168154811015156111a057fe5b6000918252602080832060048304015460039092166008026101000a90910467ffffffffffffffff168083526005825260409283902080548451818502810185019095528085529194509183018282801561122457602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311611206575b505050505093505050505b80516003600282026000190104810361124a88888584612e95565b15156112a0576040805160e560020a62461bcd02815260206004820152601f60248201527f56657269667920686561646572207369676e6174757265206661696c65642100604482015290519081900360640190fd5b60e08501805163ffffffff9081166000908152600260208181526040928390208a518154838d015163ffffffff19918216928816929092176bffffffffffffffff00000000191664010000000067ffffffffffffffff9384168102919091178455958d0151600184015560608d01519483019490945560808c0151600383015560a08c0151600483015560c08c015160058301805498516101008f0151999096169188169190911767ffffffff00000000191694909616909402929092176fffffffffffffffff0000000000000000191668010000000000000000959093169490940291909117909155610120870151805188936113a5926006850192910190614b80565b5061014091909101516007909101805473ffffffffffffffffffffffffffffffffffffffff19166c0100000000000000000000000090920491909117905560035460e086015167ffffffffffffffff90911663ffffffff90911611156114285760e08501516003805467ffffffffffffffff191663ffffffff9092169190911790555b7f1645cc3eff4fb98fef2ffd7806e7f39bfe18558f6656e69dfc9c8f26e958d4e78560e0015189604051808363ffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611498578181015183820152602001611480565b50505050905090810190601f1680156114c55780820380516001836020036101000a031916815260200191505b50935050505060405180910390a16001955050505050505b92915050565b60006114ed614b25565b6114f685612acf565b60035460e082015191925067ffffffffffffffff1663ffffffff90911611611568576040805160e560020a62461bcd02815260206004820152601c60248201527f54686520686569676874206f662068656164657220696c6c6567616c00000000604482015290519081900360640190fd5b6101408101516bffffffffffffffffffffffff191615156115f9576040805160e560020a62461bcd02815260206004820152602560248201527f546865206e657874426f6f6b4b6565706572206f66206865616465722069732060448201527f656d707479000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60065467ffffffffffffffff1660009081526005602090815260409182902080548351818402810184019094528084526060939283018282801561166657602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311611648575b50508351939450505060036002830260001901048203905061168a88878584612e95565b15156116e0576040805160e560020a62461bcd02815260206004820152601860248201527f566572696679207369676e6174757265206661696c6564210000000000000000604482015290519081900360640190fd5b600060606116ed89613102565b5061014088015191935091506bffffffffffffffffffffffff19808416911614611761576040805160e560020a62461bcd02815260206004820152601360248201527f4e657874426f6f6b65727320696c6c6567616c00000000000000000000000000604482015290519081900360640190fd5b8560e0015163ffffffff16600360006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555085600260008860e0015163ffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548163ffffffff021916908363ffffffff16021790555060208201518160000160046101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010155606082015181600201556080820151816003015560a0820151816004015560c08201518160050160006101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160050160046101000a81548163ffffffff021916908363ffffffff1602179055506101008201518160050160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101208201518160060190805190602001906118d5929190614b80565b5061014091909101516007909101805473ffffffffffffffffffffffffffffffffffffffff19166c0100000000000000000000000090920491909117905560e0860180516006805463ffffffff92831667ffffffffffffffff1990911681179091556004805460018101825560008281527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b92820492909201805467ffffffffffffffff60039093166008026101000a92830219169190930217909155915116815260056020908152604090912082516119b192840190614bfe565b506003546040805167ffffffffffffffff90921680835260208084018381528e51938501939093528d517fe60d33488cba3977bf65766cd2f8ac9617f64bf3b3198aff6240ce5c7d43b6909492938f93929091606084019185019080838360005b83811015611a2a578181015183820152602001611a12565b50505050905090810190601f168015611a575780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15060019998505050505050505050565b60006020819052908152604090205481565b600560205281600052604060002081815481101515611aa257fe5b600091825260209091200154600160a060020a03169150829050565b611ac88686610dc0565b1515611b1e576040805160e560020a62461bcd02815260206004820152601a60248201527f53796e53796e63426c6f636b486561646572206661696c656421000000000000604482015290519081900360640190fd5b611b2a848484846125e7565b1515611b80576040805160e560020a62461bcd02815260206004820152601a60248201527f766572696679416e64457865637574655478206661696c656421000000000000604482015290519081900360640190fd5b505050505050565b6002602081815260009283526040928390208054600180830154838601546003850154600486015460058701546006880180548c51610100988216159890980260001901169a909a04601f81018a90048a0287018a01909b528a865263ffffffff8088169b6401000000009889900467ffffffffffffffff9081169c979b969a959994988385169895850490931696680100000000000000009094041694939192830182828015611c7a5780601f10611c4f57610100808354040283529160200191611c7a565b820191906000526020600020905b815481529060010190602001808311611c5d57829003601f168201915b505050600790930154919250506c01000000000000000000000000028b565b60065468010000000000000000900460ff1615611d00576040805160e560020a62461bcd02815260206004820152601c60248201527f48617320616c7265616479206265656e20696e6974616c697a65642100000000604482015290519081900360640190fd5b611d08614b25565b611d1183612acf565b82519091506043900615611d95576040805160e560020a62461bcd02815260206004820152602160248201527f4c656e677468206f66207075624b65794c69737420697320696c6c6567616c2160448201527f2000000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b8151604390046003600019820104810360006060611db484848861319b565b61014087015191935091506bffffffffffffffffffffffff19808416911614611e27576040805160e560020a62461bcd02815260206004820152601360248201527f4e657874426f6f6b65727320696c6c6567616c00000000000000000000000000604482015290519081900360640190fd5b60e08501805163ffffffff9081166000908152600260208181526040928390208a518154838d015163ffffffff19918216928816929092176bffffffffffffffff00000000191664010000000067ffffffffffffffff9384168102919091178455958d0151600184015560608d01519483019490945560808c0151600383015560a08c0151600483015560c08c015160058301805498516101008f0151999096169188169190911767ffffffff00000000191694909616909402929092176fffffffffffffffff000000000000000019166801000000000000000095909316949094029190911790915561012087015180518893611f2c926006850192910190614b80565b506101409190910151600790910180546c0100000000000000000000000090920473ffffffffffffffffffffffffffffffffffffffff1990921691909117905560e0850180516003805467ffffffffffffffff1990811663ffffffff93841690811783556006805468ff000000000000000019931682179290921668010000000000000000179091556004805460018101825560008281529181047f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b018054919094166008026101000a92830267ffffffffffffffff939093021916919091179091559151168152600560209081526040909120825161202e92840190614bfe565b506003546040805167ffffffffffffffff90921680835260208084018381528b51938501939093528a517f1645cc3eff4fb98fef2ffd7806e7f39bfe18558f6656e69dfc9c8f26e958d4e79492938c93929091606084019185019080838360005b838110156120a757818101518382015260200161208f565b50505050905090810190601f1680156120d45780820380516001836020036101000a031916815260200191505b50935050505060405180910390a150505050505050565b60006120f5614c6c565b6007805460010190819055612109906135b4565b81526007543390612119906135b4565b602083015261212781613658565b604083015267ffffffffffffffff87166060808401919091526080830187905260a0830186905260c08301859052825161216090613673565b61216d8460200151613673565b61217a8560400151613673565b6121878660600151613739565b6121948760800151613673565b6121a18860a00151613673565b6121ae8960c00151613673565b6040516020018088805190602001908083835b602083106121e05780518252601f1990920191602091820191016121c1565b51815160209384036101000a60001901801990921691161790528a5191909301928a0191508083835b602083106122285780518252601f199092019160209182019101612209565b51815160209384036101000a600019018019909216911617905289519190930192890191508083835b602083106122705780518252601f199092019160209182019101612251565b51815160209384036101000a600019018019909216911617905288519190930192880191508083835b602083106122b85780518252601f199092019160209182019101612299565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106123005780518252601f1990920191602091820191016122e1565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106123485780518252601f199092019160209182019101612329565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106123905780518252601f199092019160209182019101612371565b6001836020036101000a0380198251168184511680821785525050505050509050019750505050505050506040516020818303038152906040529050808051906020012060008060075481526020019081526020016000208190555032600160a060020a03167f6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be884838460000151848b8b86604051808060200186600160a060020a0316600160a060020a031681526020018567ffffffffffffffff1667ffffffffffffffff1681526020018060200180602001848103845289818151815260200191508051906020019080838360005b8381101561249757818101518382015260200161247f565b50505050905090810190601f1680156124c45780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b838110156124f75781810151838201526020016124df565b50505050905090810190601f1680156125245780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b8381101561255757818101518382015260200161253f565b50505050905090810190601f1680156125845780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390a2506001979650505050505050565b60075481565b60048054829081106125b957fe5b9060005260206000209060049182820401919006600802915054906101000a900467ffffffffffffffff1681565b60035460009067ffffffffffffffff9081169083161115612652576040805160e560020a62461bcd02815260206004820152601b60248201527f626c6f636b486569676874203e204c6174657374486569676874210000000000604482015290519081900360640190fd5b67ffffffffffffffff821660009081526002602052604090206003015461267b8686838761377c565b15156126f7576040805160e560020a62461bcd02815260206004820152602660248201527f5665726966792063726f7373636861696e206d65726b6c652070726f6f66206660448201527f61696c6564210000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6126ff614cb4565b6127088561383a565b90506001600061271b8360000151613918565b815260208101919091526040016000205460ff16156127aa576040805160e560020a62461bcd02815260206004820152602260248201527f746865207472616e73616374696f6e20686173206265656e206578656375746560448201527f6421000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60018060006127bc8460000151613918565b815260200190815260200160002060006101000a81548160ff02191690831515021790555060006127f482604001516080015161397c565b60408301516060015190915067ffffffffffffffff16600214612887576040805160e560020a62461bcd02815260206004820152602b60248201527f54686973205478206973206e6f742061696d696e67206174204574686572656560448201527f756d206e6574776f726b21000000000000000000000000000000000000000000606482015290519081900360840190fd5b6128b081836040015160a00151846040015160c001518560400151604001518660200151613a06565b1515612906576040805160e560020a62461bcd02815260206004820152601d60248201527f457865637574652043726f7373436861696e205478206661696c656421000000604482015290519081900360640190fd5b7f5944a0908ae763fa254df8e0da54a682cbc62115501336e18be52a7dd774b80b82602001518360400151608001518460000151856040015160000151604051808567ffffffffffffffff1667ffffffffffffffff168152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b838110156129a257818101518382015260200161298a565b50505050905090810190601f1680156129cf5780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b83811015612a025781810151838201526020016129ea565b50505050905090810190601f168015612a2f5780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b83811015612a62578181015183820152602001612a4a565b50505050905090810190601f168015612a8f5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a1506001979650505050505050565b60065467ffffffffffffffff1681565b60035467ffffffffffffffff1681565b612ad7614b25565b612adf614b25565b6000612aeb8482613f0f565b63ffffffff90911683529050612b018482613fb2565b67ffffffffffffffff90911660208401529050612b1e8482614053565b60408401919091529050612b328482614053565b60608401919091529050612b468482614053565b60808401919091529050612b5a8482614053565b60a08401919091529050612b6e8482613f0f565b63ffffffff90911660c08401529050612b878482613f0f565b63ffffffff90911660e08401529050612ba08482613fb2565b67ffffffffffffffff9091166101008401529050612bbe84826140b2565b6101208401919091529050612bd3848261417f565b506bffffffffffffffffffffffff19166101408301525090505b919050565b60008067ffffffffffffffff84168110612c56576040805160e560020a62461bcd02815260206004820152601d60248201527f626f6f6b206b6565706572206c6973742063616e6e6f7420656d707479000000604482015290519081900360640190fd5b845467ffffffffffffffff851614612cb8576040805160e560020a62461bcd02815260206004820152601660248201527f63616e6e6f74207061727469616c6c7920717565727900000000000000000000604482015290519081900360640190fd5b60008311612d10576040805160e560020a62461bcd02815260206004820152601d60248201527f626c6f636b20686569676874206d75737420626520706f736974697665000000604482015290519081900360640190fd5b6000600019850167ffffffffffffffff861660011415612d3a57506000925060019150612e8d9050565b67ffffffffffffffff80821690831611612e12578654600267ffffffffffffffff84840381169190910484019187918a91908416908110612d7757fe5b6000918252602090912060048204015460039091166008026101000a900467ffffffffffffffff161415612db357935060019250612e8d915050565b85888267ffffffffffffffff16815481101515612dcc57fe5b6000918252602090912060048204015460039091166008026101000a900467ffffffffffffffff161015612e0557806001019250612e0c565b6001810391505b50612d3a565b60018267ffffffffffffffff1610158015612e6e575084876001840367ffffffffffffffff16815481101515612e4457fe5b6000918252602090912060048204015460039091166008026101000a900467ffffffffffffffff16105b15612e83575060001901915060019050612e8d565b5060009250829150505b935093915050565b600080600280876040518082805190602001908083835b60208310612ecb5780518252601f199092019160209182019101612eac565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015612f0a573d6000803e3d6000fd5b5050506040513d6020811015612f1f57600080fd5b50516040805160208181019390935281518082038401815290820191829052805190928291908401908083835b60208310612f6b5780518252601f199092019160209182019101612f4c565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015612faa573d6000803e3d6000fd5b5050506040513d6020811015612fbf57600080fd5b505190506000805b8651604190048110156130f4576000612fed612fe8896041850260206141e2565b613918565b90506000613006612fe88a6041860260200160206141e2565b90506000896041850260400181518110151561301e57fe5b90602001015160f860020a900460f860020a0260f860020a9004601b0190506000600187604051602001808281526020019150506040516020818303038152906040528051906020012083868660405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156130c2573d6000803e3d6000fd5b5050506020604051035190506130d88a82614264565b156130e4576001860195505b505060019092019150612fc79050565b509092111595945050505050565b6000606060006043845181151561311557fe5b061561316b576040805160e560020a62461bcd02815260206004820152601b60248201527f5f7075624b65794c697374206c656e67746820696c6c6567616c210000000000604482015290519081900360640190fd5b835160439004600360001982010481036000606061318a84848a61319b565b909990985092965091945050505050565b6000606080806131aa876142ba565b6040516020018083805190602001908083835b602083106131dc5780518252601f1990920191602091820191016131bd565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106132245780518252601f199092019160209182019101613205565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529050606086604051908082528060200260200182016040528015613287578160200160208202803883390190505b50905060005b878110156133bf57826132b56132b06132ab896043860260436141e2565b6142e2565b613673565b6040516020018083805190602001908083835b602083106132e75780518252601f1990920191602091820191016132c8565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061332f5780518252601f199092019160209182019101613310565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529250600061338361337a886043850260436141e2565b600360406141e2565b8051906020012090508060019004838381518110151561339f57fe5b600160a060020a039092166020928302909101909101525060010161328d565b50816133ca876142ba565b6040516020018083805190602001908083835b602083106133fc5780518252601f1990920191602091820191016133dd565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106134445780518252601f199092019160209182019101613425565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529150600060036002846040518082805190602001908083835b602083106134b15780518252601f199092019160209182019101613492565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa1580156134f0573d6000803e3d6000fd5b5050506040513d602081101561350557600080fd5b50516040805160208181019390935281518082038401815290820191829052805190928291908401908083835b602083106135515780518252601f199092019160209182019101613532565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015613590573d6000803e3d6000fd5b5050604051516c010000000000000000000000000299929850919650505050505050565b6060600082101580156135e757507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8211155b151561363d576040805160e560020a62461bcd02815260206004820152601760248201527f56616c75652065786365656473207468652072616e6765000000000000000000604482015290519081900360640190fd5b60405190506020815281602082015260408101604052919050565b604080516014815260609290921b6020830152818101905290565b80516060906136818161445e565b836040516020018083805190602001908083835b602083106136b45780518252601f199092019160209182019101613695565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106136fc5780518252601f1990920191602091820191016136dd565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052915050919050565b6040516008808252606091906000601f5b8282101561376c5785811a82602086010153600191909101906000190161374a565b5050506040818101905292915050565b600080613788836145df565b90506000805b875181101561382d5787818151811015156137a557fe5b602090810290910101519150600087828151811015156137c157fe5b9060200190602002015160018111156137d657fe5b14156137ed576137e682846146e3565b9250613825565b600187828151811015156137fd57fe5b90602001906020020151600181111561381257fe5b14156138255761382283836146e3565b92505b60010161378e565b5050909214949350505050565b613842614cb4565b61384a614cb4565b600061385684826140b2565b90835290506138658482613fb2565b67ffffffffffffffff90911660208401529050613880614c6c565b61388a85836140b2565b908252915061389985836140b2565b602083019190915291506138ad85836140b2565b604083019190915291506138c18583613fb2565b67ffffffffffffffff909116606083015291506138de85836140b2565b608083019190915291506138f285836140b2565b60a0830191909152915061390685836140b2565b5060c082015260408301525092915050565b8051600090602014613974576040805160e560020a62461bcd02815260206004820152601760248201527f6279746573206c656e677468206973206e6f742033322e000000000000000000604482015290519081900360640190fd5b506020015190565b80516000906014146139fe576040805160e560020a62461bcd02815260206004820152602360248201527f6279746573206c656e67746820646f6573206e6f74206d61746368206164647260448201527f6573730000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b506014015190565b6000613a11866147a7565b1515613a8d576040805160e560020a62461bcd02815260206004820152602860248201527f5468652070617373656420696e2061646472657373206973206e6f742061206360448201527f6f6e747261637421000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6060600087600160a060020a0316876040516020018082805190602001908083835b60208310613ace5780518252601f199092019160209182019101613aaf565b51815160001960209485036101000a019081169019919091161790527f2862797465732c62797465732c75696e743634290000000000000000000000009390910192835260408051600b198186030181526014850190915280519082012067ffffffffffffffff8b1660748501526060603485019081528d5160948601528d519196508d95508c948c945090928392605483019260b4019188019080838360005b83811015613b87578181015183820152602001613b6f565b50505050905090810190601f168015613bb45780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015613be7578181015183820152602001613bcf565b50505050905090810190601f168015613c145780820380516001836020036101000a031916815260200191505b509550505050505060405160208183030381529060405260405160200180837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260040182805190602001908083835b60208310613ca15780518252601f199092019160209182019101613c82565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040526040518082805190602001908083835b60208310613d055780518252601f199092019160209182019101613ce6565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114613d67576040519150601f19603f3d011682016040523d82523d6000602084013e613d6c565b606091505b5092509050600181151514613df1576040805160e560020a62461bcd02815260206004820152602b60248201527f45746843726f7373436861696e2063616c6c20627573696e65737320636f6e7460448201527f72616374206661696c6564000000000000000000000000000000000000000000606482015290519081900360840190fd5b8151600010613e70576040805160e560020a62461bcd02815260206004820152602760248201527f4e6f2072657475726e2076616c75652066726f6d20627573696e65737320636f60448201527f6e74726163742100000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000613e7d83601f6147e3565b509050600181151514613f00576040805160e560020a62461bcd02815260206004820152603760248201527f45746843726f7373436861696e2063616c6c20627573696e65737320636f6e7460448201527f726163742072657475726e206973206e6f742074727565000000000000000000606482015290519081900360840190fd5b50600198975050505050505050565b60008083518360040111151515613f5e576040805160e560020a62461bcd0281526020600482015260166024820152600080516020614d29833981519152604482015290519081900360640190fd5b600060405160046000600182038760208a0101515b83831015613f935780821a83860153600183019250600182039150613f73565b50505080820160405260200390035192505050600482015b9250929050565b60008083518360080111151515614001576040805160e560020a62461bcd0281526020600482015260166024820152600080516020614d29833981519152604482015290519081900360640190fd5b600060405160086000600182038760208a0101515b838310156140365780821a83860153600183019250600182039150614016565b505050808201604052602003900351956008949094019450505050565b600080835183602001111515156140a2576040805160e560020a62461bcd0281526020600482015260166024820152600080516020614d29833981519152604482015290519081900360640190fd5b5050602091810182015192910190565b60606000806140c185856148e2565b86519095509091508185011115614110576040805160e560020a62461bcd0281526020600482015260166024820152600080516020614d29833981519152604482015290519081900360640190fd5b60608115801561412b57604051915060208201604052614175565b6040519150601f8316801560200281840101848101888315602002848c0101015b8183101561416457805183526020928301920161414c565b5050848452601f01601f1916604052505b5095930193505050565b600080835183601401111515156141ce576040805160e560020a62461bcd0281526020600482015260166024820152600080516020614d29833981519152604482015290519081900360640190fd5b505081810160200151601482019250929050565b60608183018451101515156141f657600080fd5b6060821580156142115760405191506020820160405261425b565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561424a578051835260209283019201614232565b5050858452601f01601f1916604052505b50949350505050565b6000805b83518110156142b357838181518110151561427f57fe5b90602001906020020151600160a060020a031683600160a060020a031614156142ab57600191506142b3565b600101614268565b5092915050565b6040516002808252606091906000601f85811a6020850153600191909101906000190161374a565b60606022825110151515614340576040805160e560020a62461bcd02815260206004820152601760248201527f6b6579206c656e67676820697320746f6f2073686f7274000000000000000000604482015290519081900360640190fd5b61434d82600060236141e2565b9050600282604281518110151561436057fe5b90602001015160f860020a900460f860020a0260f860020a900460ff1681151561438657fe5b0660ff16600014156143f85780517f020000000000000000000000000000000000000000000000000000000000000090829060029081106143c357fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350612bed565b80517f0300000000000000000000000000000000000000000000000000000000000000908290600290811061442957fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350919050565b606060fd8267ffffffffffffffff1610156144835761447c826149d8565b9050612bed565b61ffff67ffffffffffffffff83161161457e576144bf7ffd000000000000000000000000000000000000000000000000000000000000006149f3565b6144c8836142ba565b6040516020018083805190602001908083835b602083106144fa5780518252601f1990920191602091820191016144db565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106145425780518252601f199092019160209182019101614523565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529050612bed565b63ffffffff67ffffffffffffffff8316116145c5576145bc7ffe000000000000000000000000000000000000000000000000000000000000006149f3565b6144c883614a04565b6145d6600160f860020a03196149f3565b6144c883613739565b6040516000602080830182815284519293600293859387939260210191908401908083835b602083106146235780518252601f199092019160209182019101614604565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040526040518082805190602001908083835b602083106146875780518252601f199092019160209182019101614668565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa1580156146c6573d6000803e3d6000fd5b5050506040513d60208110156146db57600080fd5b505192915050565b6040805160f860020a60208083019190915260218201859052604180830185905283518084039091018152606190920192839052815160009360029392909182918401908083835b6020831061474a5780518252601f19909201916020918201910161472b565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015614789573d6000803e3d6000fd5b5050506040513d602081101561479e57600080fd5b50519392505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47081158015906147db5750808214155b949350505050565b60008083518360010111151515614844576040805160e560020a62461bcd02815260206004820152601460248201527f4f66667365742065786365656473206c696d6974000000000000000000000000604482015290519081900360640190fd5b83830160200151600060f860020a600160f860020a03198316141561486b575060016148d4565b600160f860020a031982161515614884575060006148d4565b6040805160e560020a62461bcd02815260206004820152601460248201527f4e657874426f6f6c2076616c7565206572726f72000000000000000000000000604482015290519081900360640190fd5b956001949094019450505050565b60008060006148f18585614a2c565b945090507ffd00000000000000000000000000000000000000000000000000000000000000600160f860020a031982161415614942576149318585614aa1565b8161ffff1691509250925050613fab565b7ffe00000000000000000000000000000000000000000000000000000000000000600160f860020a0319821614156149915761497e8585613f0f565b8163ffffffff1691509250925050613fab565b600160f860020a031980821614156149c4576149ad8585613fb2565b8167ffffffffffffffff1691509250925050613fab565b60f860020a900460ff169150829050613fab565b604080516001815260f89290921b6020830152818101905290565b60606114dd60f860020a83046149d8565b6040516004808252606091906000601f85811a6020850153600191909101906000190161374a565b60008083518360010111151515614a8d576040805160e560020a62461bcd02815260206004820152601660248201527f4f66667365742065786365656473206d6178696d756d00000000000000000000604482015290519081900360640190fd5b505081810160200151600182019250929050565b60008083518360020111151515614af0576040805160e560020a62461bcd0281526020600482015260166024820152600080516020614d29833981519152604482015290519081900360640190fd5b6000604051846020870101518060011a82538060001a6001830153506002818101604052601d19909101519694019450505050565b604080516101608101825260008082526020820181905291810182905260608082018390526080820183905260a0820183905260c0820183905260e08201839052610100820183905261012082015261014081019190915290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10614bc157805160ff1916838001178555614bee565b82800160010185558215614bee579182015b82811115614bee578251825591602001919060010190614bd3565b50614bfa929150614cda565b5090565b828054828255906000526020600020908101928215614c60579160200282015b82811115614c60578251825473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909116178255602090920191600190910190614c1e565b50614bfa929150614cf7565b60e060405190810160405280606081526020016060815260200160608152602001600067ffffffffffffffff1681526020016060815260200160608152602001606081525090565b60408051610120810182526060815260006020820152908101614cd5614c6c565b905290565b614cf491905b80821115614bfa5760008155600101614ce0565b90565b614cf491905b80821115614bfa57805473ffffffffffffffffffffffffffffffffffffffff19168155600101614cfd56fe6f66667365742065786365656473206d6178696d756d00000000000000000000a165627a7a72305820f1d17115807a7faf9f5cc48f298d6bcef96d1357775c676b1f8dac838d6242200029"

// DeployEthCrossChain deploys a new Ethereum contract, binding an instance of EthCrossChain to it.
func DeployEthCrossChain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthCrossChain, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCrossChainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthCrossChainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthCrossChain{EthCrossChainCaller: EthCrossChainCaller{contract: contract}, EthCrossChainTransactor: EthCrossChainTransactor{contract: contract}, EthCrossChainFilterer: EthCrossChainFilterer{contract: contract}}, nil
}

// EthCrossChain is an auto generated Go binding around an Ethereum contract.
type EthCrossChain struct {
	EthCrossChainCaller     // Read-only binding to the contract
	EthCrossChainTransactor // Write-only binding to the contract
	EthCrossChainFilterer   // Log filterer for contract events
}

// EthCrossChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthCrossChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthCrossChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthCrossChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthCrossChainSession struct {
	Contract     *EthCrossChain    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthCrossChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthCrossChainCallerSession struct {
	Contract *EthCrossChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EthCrossChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthCrossChainTransactorSession struct {
	Contract     *EthCrossChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EthCrossChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthCrossChainRaw struct {
	Contract *EthCrossChain // Generic contract binding to access the raw methods on
}

// EthCrossChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthCrossChainCallerRaw struct {
	Contract *EthCrossChainCaller // Generic read-only contract binding to access the raw methods on
}

// EthCrossChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthCrossChainTransactorRaw struct {
	Contract *EthCrossChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthCrossChain creates a new instance of EthCrossChain, bound to a specific deployed contract.
func NewEthCrossChain(address common.Address, backend bind.ContractBackend) (*EthCrossChain, error) {
	contract, err := bindEthCrossChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthCrossChain{EthCrossChainCaller: EthCrossChainCaller{contract: contract}, EthCrossChainTransactor: EthCrossChainTransactor{contract: contract}, EthCrossChainFilterer: EthCrossChainFilterer{contract: contract}}, nil
}

// NewEthCrossChainCaller creates a new read-only instance of EthCrossChain, bound to a specific deployed contract.
func NewEthCrossChainCaller(address common.Address, caller bind.ContractCaller) (*EthCrossChainCaller, error) {
	contract, err := bindEthCrossChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainCaller{contract: contract}, nil
}

// NewEthCrossChainTransactor creates a new write-only instance of EthCrossChain, bound to a specific deployed contract.
func NewEthCrossChainTransactor(address common.Address, transactor bind.ContractTransactor) (*EthCrossChainTransactor, error) {
	contract, err := bindEthCrossChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainTransactor{contract: contract}, nil
}

// NewEthCrossChainFilterer creates a new log filterer instance of EthCrossChain, bound to a specific deployed contract.
func NewEthCrossChainFilterer(address common.Address, filterer bind.ContractFilterer) (*EthCrossChainFilterer, error) {
	contract, err := bindEthCrossChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainFilterer{contract: contract}, nil
}

// bindEthCrossChain binds a generic wrapper to an already deployed contract.
func bindEthCrossChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCrossChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthCrossChain *EthCrossChainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthCrossChain.Contract.EthCrossChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthCrossChain *EthCrossChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChain.Contract.EthCrossChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthCrossChain *EthCrossChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthCrossChain.Contract.EthCrossChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthCrossChain *EthCrossChainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthCrossChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthCrossChain *EthCrossChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthCrossChain *EthCrossChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthCrossChain.Contract.contract.Transact(opts, method, params...)
}

// LatestBookKeeperHeight is a free data retrieval call binding the contract method 0xe244e4b1.
//
// Solidity: function LatestBookKeeperHeight() constant returns(uint64)
func (_EthCrossChain *EthCrossChainCaller) LatestBookKeeperHeight(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _EthCrossChain.contract.Call(opts, out, "LatestBookKeeperHeight")
	return *ret0, err
}

// LatestBookKeeperHeight is a free data retrieval call binding the contract method 0xe244e4b1.
//
// Solidity: function LatestBookKeeperHeight() constant returns(uint64)
func (_EthCrossChain *EthCrossChainSession) LatestBookKeeperHeight() (uint64, error) {
	return _EthCrossChain.Contract.LatestBookKeeperHeight(&_EthCrossChain.CallOpts)
}

// LatestBookKeeperHeight is a free data retrieval call binding the contract method 0xe244e4b1.
//
// Solidity: function LatestBookKeeperHeight() constant returns(uint64)
func (_EthCrossChain *EthCrossChainCallerSession) LatestBookKeeperHeight() (uint64, error) {
	return _EthCrossChain.Contract.LatestBookKeeperHeight(&_EthCrossChain.CallOpts)
}

// LatestHeight is a free data retrieval call binding the contract method 0xe6aefed8.
//
// Solidity: function LatestHeight() constant returns(uint64)
func (_EthCrossChain *EthCrossChainCaller) LatestHeight(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _EthCrossChain.contract.Call(opts, out, "LatestHeight")
	return *ret0, err
}

// LatestHeight is a free data retrieval call binding the contract method 0xe6aefed8.
//
// Solidity: function LatestHeight() constant returns(uint64)
func (_EthCrossChain *EthCrossChainSession) LatestHeight() (uint64, error) {
	return _EthCrossChain.Contract.LatestHeight(&_EthCrossChain.CallOpts)
}

// LatestHeight is a free data retrieval call binding the contract method 0xe6aefed8.
//
// Solidity: function LatestHeight() constant returns(uint64)
func (_EthCrossChain *EthCrossChainCallerSession) LatestHeight() (uint64, error) {
	return _EthCrossChain.Contract.LatestHeight(&_EthCrossChain.CallOpts)
}

// MCBlockHeaders is a free data retrieval call binding the contract method 0x9a3337bc.
//
// Solidity: function MCBlockHeaders(uint64 ) constant returns(uint32 version, uint64 chainId, bytes32 prevBlockHash, bytes32 transactionsRoot, bytes32 crossStatesRoot, bytes32 blockRoot, uint32 timestamp, uint32 height, uint64 consensusData, bytes consensusPayload, bytes20 nextBookkeeper)
func (_EthCrossChain *EthCrossChainCaller) MCBlockHeaders(opts *bind.CallOpts, arg0 uint64) (struct {
	Version          uint32
	ChainId          uint64
	PrevBlockHash    [32]byte
	TransactionsRoot [32]byte
	CrossStatesRoot  [32]byte
	BlockRoot        [32]byte
	Timestamp        uint32
	Height           uint32
	ConsensusData    uint64
	ConsensusPayload []byte
	NextBookkeeper   [20]byte
}, error) {
	ret := new(struct {
		Version          uint32
		ChainId          uint64
		PrevBlockHash    [32]byte
		TransactionsRoot [32]byte
		CrossStatesRoot  [32]byte
		BlockRoot        [32]byte
		Timestamp        uint32
		Height           uint32
		ConsensusData    uint64
		ConsensusPayload []byte
		NextBookkeeper   [20]byte
	})
	out := ret
	err := _EthCrossChain.contract.Call(opts, out, "MCBlockHeaders", arg0)
	return *ret, err
}

// MCBlockHeaders is a free data retrieval call binding the contract method 0x9a3337bc.
//
// Solidity: function MCBlockHeaders(uint64 ) constant returns(uint32 version, uint64 chainId, bytes32 prevBlockHash, bytes32 transactionsRoot, bytes32 crossStatesRoot, bytes32 blockRoot, uint32 timestamp, uint32 height, uint64 consensusData, bytes consensusPayload, bytes20 nextBookkeeper)
func (_EthCrossChain *EthCrossChainSession) MCBlockHeaders(arg0 uint64) (struct {
	Version          uint32
	ChainId          uint64
	PrevBlockHash    [32]byte
	TransactionsRoot [32]byte
	CrossStatesRoot  [32]byte
	BlockRoot        [32]byte
	Timestamp        uint32
	Height           uint32
	ConsensusData    uint64
	ConsensusPayload []byte
	NextBookkeeper   [20]byte
}, error) {
	return _EthCrossChain.Contract.MCBlockHeaders(&_EthCrossChain.CallOpts, arg0)
}

// MCBlockHeaders is a free data retrieval call binding the contract method 0x9a3337bc.
//
// Solidity: function MCBlockHeaders(uint64 ) constant returns(uint32 version, uint64 chainId, bytes32 prevBlockHash, bytes32 transactionsRoot, bytes32 crossStatesRoot, bytes32 blockRoot, uint32 timestamp, uint32 height, uint64 consensusData, bytes consensusPayload, bytes20 nextBookkeeper)
func (_EthCrossChain *EthCrossChainCallerSession) MCBlockHeaders(arg0 uint64) (struct {
	Version          uint32
	ChainId          uint64
	PrevBlockHash    [32]byte
	TransactionsRoot [32]byte
	CrossStatesRoot  [32]byte
	BlockRoot        [32]byte
	Timestamp        uint32
	Height           uint32
	ConsensusData    uint64
	ConsensusPayload []byte
	NextBookkeeper   [20]byte
}, error) {
	return _EthCrossChain.Contract.MCBlockHeaders(&_EthCrossChain.CallOpts, arg0)
}

// MCKeeperHeight is a free data retrieval call binding the contract method 0xc6521e0f.
//
// Solidity: function MCKeeperHeight(uint256 ) constant returns(uint64)
func (_EthCrossChain *EthCrossChainCaller) MCKeeperHeight(opts *bind.CallOpts, arg0 *big.Int) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _EthCrossChain.contract.Call(opts, out, "MCKeeperHeight", arg0)
	return *ret0, err
}

// MCKeeperHeight is a free data retrieval call binding the contract method 0xc6521e0f.
//
// Solidity: function MCKeeperHeight(uint256 ) constant returns(uint64)
func (_EthCrossChain *EthCrossChainSession) MCKeeperHeight(arg0 *big.Int) (uint64, error) {
	return _EthCrossChain.Contract.MCKeeperHeight(&_EthCrossChain.CallOpts, arg0)
}

// MCKeeperHeight is a free data retrieval call binding the contract method 0xc6521e0f.
//
// Solidity: function MCKeeperHeight(uint256 ) constant returns(uint64)
func (_EthCrossChain *EthCrossChainCallerSession) MCKeeperHeight(arg0 *big.Int) (uint64, error) {
	return _EthCrossChain.Contract.MCKeeperHeight(&_EthCrossChain.CallOpts, arg0)
}

// MCKeeperPubKeys is a free data retrieval call binding the contract method 0x464f81a2.
//
// Solidity: function MCKeeperPubKeys(uint64 , uint256 ) constant returns(address)
func (_EthCrossChain *EthCrossChainCaller) MCKeeperPubKeys(opts *bind.CallOpts, arg0 uint64, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthCrossChain.contract.Call(opts, out, "MCKeeperPubKeys", arg0, arg1)
	return *ret0, err
}

// MCKeeperPubKeys is a free data retrieval call binding the contract method 0x464f81a2.
//
// Solidity: function MCKeeperPubKeys(uint64 , uint256 ) constant returns(address)
func (_EthCrossChain *EthCrossChainSession) MCKeeperPubKeys(arg0 uint64, arg1 *big.Int) (common.Address, error) {
	return _EthCrossChain.Contract.MCKeeperPubKeys(&_EthCrossChain.CallOpts, arg0, arg1)
}

// MCKeeperPubKeys is a free data retrieval call binding the contract method 0x464f81a2.
//
// Solidity: function MCKeeperPubKeys(uint64 , uint256 ) constant returns(address)
func (_EthCrossChain *EthCrossChainCallerSession) MCKeeperPubKeys(arg0 uint64, arg1 *big.Int) (common.Address, error) {
	return _EthCrossChain.Contract.MCKeeperPubKeys(&_EthCrossChain.CallOpts, arg0, arg1)
}

// TransactionId is a free data retrieval call binding the contract method 0xc267dfcb.
//
// Solidity: function TransactionId() constant returns(uint256)
func (_EthCrossChain *EthCrossChainCaller) TransactionId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthCrossChain.contract.Call(opts, out, "TransactionId")
	return *ret0, err
}

// TransactionId is a free data retrieval call binding the contract method 0xc267dfcb.
//
// Solidity: function TransactionId() constant returns(uint256)
func (_EthCrossChain *EthCrossChainSession) TransactionId() (*big.Int, error) {
	return _EthCrossChain.Contract.TransactionId(&_EthCrossChain.CallOpts)
}

// TransactionId is a free data retrieval call binding the contract method 0xc267dfcb.
//
// Solidity: function TransactionId() constant returns(uint256)
func (_EthCrossChain *EthCrossChainCallerSession) TransactionId() (*big.Int, error) {
	return _EthCrossChain.Contract.TransactionId(&_EthCrossChain.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x34008b16.
//
// Solidity: function Transactions(uint256 ) constant returns(bytes32)
func (_EthCrossChain *EthCrossChainCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _EthCrossChain.contract.Call(opts, out, "Transactions", arg0)
	return *ret0, err
}

// Transactions is a free data retrieval call binding the contract method 0x34008b16.
//
// Solidity: function Transactions(uint256 ) constant returns(bytes32)
func (_EthCrossChain *EthCrossChainSession) Transactions(arg0 *big.Int) ([32]byte, error) {
	return _EthCrossChain.Contract.Transactions(&_EthCrossChain.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x34008b16.
//
// Solidity: function Transactions(uint256 ) constant returns(bytes32)
func (_EthCrossChain *EthCrossChainCallerSession) Transactions(arg0 *big.Int) ([32]byte, error) {
	return _EthCrossChain.Contract.Transactions(&_EthCrossChain.CallOpts, arg0)
}

// ChangeBookKeeper is a paid mutator transaction binding the contract method 0x21f966aa.
//
// Solidity: function ChangeBookKeeper(bytes _rawHeader, bytes _pubKeyList, bytes _sigList) returns(bool)
func (_EthCrossChain *EthCrossChainTransactor) ChangeBookKeeper(opts *bind.TransactOpts, _rawHeader []byte, _pubKeyList []byte, _sigList []byte) (*types.Transaction, error) {
	return _EthCrossChain.contract.Transact(opts, "ChangeBookKeeper", _rawHeader, _pubKeyList, _sigList)
}

// ChangeBookKeeper is a paid mutator transaction binding the contract method 0x21f966aa.
//
// Solidity: function ChangeBookKeeper(bytes _rawHeader, bytes _pubKeyList, bytes _sigList) returns(bool)
func (_EthCrossChain *EthCrossChainSession) ChangeBookKeeper(_rawHeader []byte, _pubKeyList []byte, _sigList []byte) (*types.Transaction, error) {
	return _EthCrossChain.Contract.ChangeBookKeeper(&_EthCrossChain.TransactOpts, _rawHeader, _pubKeyList, _sigList)
}

// ChangeBookKeeper is a paid mutator transaction binding the contract method 0x21f966aa.
//
// Solidity: function ChangeBookKeeper(bytes _rawHeader, bytes _pubKeyList, bytes _sigList) returns(bool)
func (_EthCrossChain *EthCrossChainTransactorSession) ChangeBookKeeper(_rawHeader []byte, _pubKeyList []byte, _sigList []byte) (*types.Transaction, error) {
	return _EthCrossChain.Contract.ChangeBookKeeper(&_EthCrossChain.TransactOpts, _rawHeader, _pubKeyList, _sigList)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0xadbccaf5.
//
// Solidity: function InitGenesisBlock(bytes _rawHeader, bytes _pubKeyList) returns()
func (_EthCrossChain *EthCrossChainTransactor) InitGenesisBlock(opts *bind.TransactOpts, _rawHeader []byte, _pubKeyList []byte) (*types.Transaction, error) {
	return _EthCrossChain.contract.Transact(opts, "InitGenesisBlock", _rawHeader, _pubKeyList)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0xadbccaf5.
//
// Solidity: function InitGenesisBlock(bytes _rawHeader, bytes _pubKeyList) returns()
func (_EthCrossChain *EthCrossChainSession) InitGenesisBlock(_rawHeader []byte, _pubKeyList []byte) (*types.Transaction, error) {
	return _EthCrossChain.Contract.InitGenesisBlock(&_EthCrossChain.TransactOpts, _rawHeader, _pubKeyList)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0xadbccaf5.
//
// Solidity: function InitGenesisBlock(bytes _rawHeader, bytes _pubKeyList) returns()
func (_EthCrossChain *EthCrossChainTransactorSession) InitGenesisBlock(_rawHeader []byte, _pubKeyList []byte) (*types.Transaction, error) {
	return _EthCrossChain.Contract.InitGenesisBlock(&_EthCrossChain.TransactOpts, _rawHeader, _pubKeyList)
}

// SyncAndVerify is a paid mutator transaction binding the contract method 0x48c2f675.
//
// Solidity: function SyncAndVerify(bytes rawHeader, bytes sigList, bytes32[] proof, uint8[] position, bytes toMerkleValueBs, uint64 blockHeight) returns()
func (_EthCrossChain *EthCrossChainTransactor) SyncAndVerify(opts *bind.TransactOpts, rawHeader []byte, sigList []byte, proof [][32]byte, position []uint8, toMerkleValueBs []byte, blockHeight uint64) (*types.Transaction, error) {
	return _EthCrossChain.contract.Transact(opts, "SyncAndVerify", rawHeader, sigList, proof, position, toMerkleValueBs, blockHeight)
}

// SyncAndVerify is a paid mutator transaction binding the contract method 0x48c2f675.
//
// Solidity: function SyncAndVerify(bytes rawHeader, bytes sigList, bytes32[] proof, uint8[] position, bytes toMerkleValueBs, uint64 blockHeight) returns()
func (_EthCrossChain *EthCrossChainSession) SyncAndVerify(rawHeader []byte, sigList []byte, proof [][32]byte, position []uint8, toMerkleValueBs []byte, blockHeight uint64) (*types.Transaction, error) {
	return _EthCrossChain.Contract.SyncAndVerify(&_EthCrossChain.TransactOpts, rawHeader, sigList, proof, position, toMerkleValueBs, blockHeight)
}

// SyncAndVerify is a paid mutator transaction binding the contract method 0x48c2f675.
//
// Solidity: function SyncAndVerify(bytes rawHeader, bytes sigList, bytes32[] proof, uint8[] position, bytes toMerkleValueBs, uint64 blockHeight) returns()
func (_EthCrossChain *EthCrossChainTransactorSession) SyncAndVerify(rawHeader []byte, sigList []byte, proof [][32]byte, position []uint8, toMerkleValueBs []byte, blockHeight uint64) (*types.Transaction, error) {
	return _EthCrossChain.Contract.SyncAndVerify(&_EthCrossChain.TransactOpts, rawHeader, sigList, proof, position, toMerkleValueBs, blockHeight)
}

// SyncBlockHeader is a paid mutator transaction binding the contract method 0x1e568e18.
//
// Solidity: function SyncBlockHeader(bytes _rawHeader, bytes _sigList) returns(bool)
func (_EthCrossChain *EthCrossChainTransactor) SyncBlockHeader(opts *bind.TransactOpts, _rawHeader []byte, _sigList []byte) (*types.Transaction, error) {
	return _EthCrossChain.contract.Transact(opts, "SyncBlockHeader", _rawHeader, _sigList)
}

// SyncBlockHeader is a paid mutator transaction binding the contract method 0x1e568e18.
//
// Solidity: function SyncBlockHeader(bytes _rawHeader, bytes _sigList) returns(bool)
func (_EthCrossChain *EthCrossChainSession) SyncBlockHeader(_rawHeader []byte, _sigList []byte) (*types.Transaction, error) {
	return _EthCrossChain.Contract.SyncBlockHeader(&_EthCrossChain.TransactOpts, _rawHeader, _sigList)
}

// SyncBlockHeader is a paid mutator transaction binding the contract method 0x1e568e18.
//
// Solidity: function SyncBlockHeader(bytes _rawHeader, bytes _sigList) returns(bool)
func (_EthCrossChain *EthCrossChainTransactorSession) SyncBlockHeader(_rawHeader []byte, _sigList []byte) (*types.Transaction, error) {
	return _EthCrossChain.Contract.SyncBlockHeader(&_EthCrossChain.TransactOpts, _rawHeader, _sigList)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_EthCrossChain *EthCrossChainTransactor) CrossChain(opts *bind.TransactOpts, _toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _EthCrossChain.contract.Transact(opts, "crossChain", _toChainId, _toContract, _method, _txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_EthCrossChain *EthCrossChainSession) CrossChain(_toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _EthCrossChain.Contract.CrossChain(&_EthCrossChain.TransactOpts, _toChainId, _toContract, _method, _txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_EthCrossChain *EthCrossChainTransactorSession) CrossChain(_toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _EthCrossChain.Contract.CrossChain(&_EthCrossChain.TransactOpts, _toChainId, _toContract, _method, _txData)
}

// VerifyAndExecuteTx is a paid mutator transaction binding the contract method 0xd5ea92e3.
//
// Solidity: function verifyAndExecuteTx(bytes32[] _proof, uint8[] _position, bytes _toMerkleValueBs, uint64 _blockHeight) returns(bool)
func (_EthCrossChain *EthCrossChainTransactor) VerifyAndExecuteTx(opts *bind.TransactOpts, _proof [][32]byte, _position []uint8, _toMerkleValueBs []byte, _blockHeight uint64) (*types.Transaction, error) {
	return _EthCrossChain.contract.Transact(opts, "verifyAndExecuteTx", _proof, _position, _toMerkleValueBs, _blockHeight)
}

// VerifyAndExecuteTx is a paid mutator transaction binding the contract method 0xd5ea92e3.
//
// Solidity: function verifyAndExecuteTx(bytes32[] _proof, uint8[] _position, bytes _toMerkleValueBs, uint64 _blockHeight) returns(bool)
func (_EthCrossChain *EthCrossChainSession) VerifyAndExecuteTx(_proof [][32]byte, _position []uint8, _toMerkleValueBs []byte, _blockHeight uint64) (*types.Transaction, error) {
	return _EthCrossChain.Contract.VerifyAndExecuteTx(&_EthCrossChain.TransactOpts, _proof, _position, _toMerkleValueBs, _blockHeight)
}

// VerifyAndExecuteTx is a paid mutator transaction binding the contract method 0xd5ea92e3.
//
// Solidity: function verifyAndExecuteTx(bytes32[] _proof, uint8[] _position, bytes _toMerkleValueBs, uint64 _blockHeight) returns(bool)
func (_EthCrossChain *EthCrossChainTransactorSession) VerifyAndExecuteTx(_proof [][32]byte, _position []uint8, _toMerkleValueBs []byte, _blockHeight uint64) (*types.Transaction, error) {
	return _EthCrossChain.Contract.VerifyAndExecuteTx(&_EthCrossChain.TransactOpts, _proof, _position, _toMerkleValueBs, _blockHeight)
}

// EthCrossChainChangeBookKeeperEventIterator is returned from FilterChangeBookKeeperEvent and is used to iterate over the raw logs and unpacked data for ChangeBookKeeperEvent events raised by the EthCrossChain contract.
type EthCrossChainChangeBookKeeperEventIterator struct {
	Event *EthCrossChainChangeBookKeeperEvent // Event containing the contract specifics and raw log

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
func (it *EthCrossChainChangeBookKeeperEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainChangeBookKeeperEvent)
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
		it.Event = new(EthCrossChainChangeBookKeeperEvent)
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
func (it *EthCrossChainChangeBookKeeperEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainChangeBookKeeperEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainChangeBookKeeperEvent represents a ChangeBookKeeperEvent event raised by the EthCrossChain contract.
type EthCrossChainChangeBookKeeperEvent struct {
	Height     *big.Int
	RawHeaders []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChangeBookKeeperEvent is a free log retrieval operation binding the contract event 0xe60d33488cba3977bf65766cd2f8ac9617f64bf3b3198aff6240ce5c7d43b690.
//
// Solidity: event ChangeBookKeeperEvent(uint256 height, bytes rawHeaders)
func (_EthCrossChain *EthCrossChainFilterer) FilterChangeBookKeeperEvent(opts *bind.FilterOpts) (*EthCrossChainChangeBookKeeperEventIterator, error) {

	logs, sub, err := _EthCrossChain.contract.FilterLogs(opts, "ChangeBookKeeperEvent")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainChangeBookKeeperEventIterator{contract: _EthCrossChain.contract, event: "ChangeBookKeeperEvent", logs: logs, sub: sub}, nil
}

// WatchChangeBookKeeperEvent is a free log subscription operation binding the contract event 0xe60d33488cba3977bf65766cd2f8ac9617f64bf3b3198aff6240ce5c7d43b690.
//
// Solidity: event ChangeBookKeeperEvent(uint256 height, bytes rawHeaders)
func (_EthCrossChain *EthCrossChainFilterer) WatchChangeBookKeeperEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainChangeBookKeeperEvent) (event.Subscription, error) {

	logs, sub, err := _EthCrossChain.contract.WatchLogs(opts, "ChangeBookKeeperEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainChangeBookKeeperEvent)
				if err := _EthCrossChain.contract.UnpackLog(event, "ChangeBookKeeperEvent", log); err != nil {
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

// ParseChangeBookKeeperEvent is a log parse operation binding the contract event 0xe60d33488cba3977bf65766cd2f8ac9617f64bf3b3198aff6240ce5c7d43b690.
//
// Solidity: event ChangeBookKeeperEvent(uint256 height, bytes rawHeaders)
func (_EthCrossChain *EthCrossChainFilterer) ParseChangeBookKeeperEvent(log types.Log) (*EthCrossChainChangeBookKeeperEvent, error) {
	event := new(EthCrossChainChangeBookKeeperEvent)
	if err := _EthCrossChain.contract.UnpackLog(event, "ChangeBookKeeperEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainCrossChainEventIterator is returned from FilterCrossChainEvent and is used to iterate over the raw logs and unpacked data for CrossChainEvent events raised by the EthCrossChain contract.
type EthCrossChainCrossChainEventIterator struct {
	Event *EthCrossChainCrossChainEvent // Event containing the contract specifics and raw log

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
func (it *EthCrossChainCrossChainEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainCrossChainEvent)
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
		it.Event = new(EthCrossChainCrossChainEvent)
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
func (it *EthCrossChainCrossChainEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainCrossChainEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainCrossChainEvent represents a CrossChainEvent event raised by the EthCrossChain contract.
type EthCrossChainCrossChainEvent struct {
	Sender     common.Address
	TxId       []byte
	Token      common.Address
	ToChainId  uint64
	ToContract []byte
	Rawdata    []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCrossChainEvent is a free log retrieval operation binding the contract event 0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483.
//
// Solidity: event CrossChainEvent(address indexed sender, bytes txId, address token, uint64 toChainId, bytes toContract, bytes rawdata)
func (_EthCrossChain *EthCrossChainFilterer) FilterCrossChainEvent(opts *bind.FilterOpts, sender []common.Address) (*EthCrossChainCrossChainEventIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EthCrossChain.contract.FilterLogs(opts, "CrossChainEvent", senderRule)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainCrossChainEventIterator{contract: _EthCrossChain.contract, event: "CrossChainEvent", logs: logs, sub: sub}, nil
}

// WatchCrossChainEvent is a free log subscription operation binding the contract event 0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483.
//
// Solidity: event CrossChainEvent(address indexed sender, bytes txId, address token, uint64 toChainId, bytes toContract, bytes rawdata)
func (_EthCrossChain *EthCrossChainFilterer) WatchCrossChainEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainCrossChainEvent, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EthCrossChain.contract.WatchLogs(opts, "CrossChainEvent", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainCrossChainEvent)
				if err := _EthCrossChain.contract.UnpackLog(event, "CrossChainEvent", log); err != nil {
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

// ParseCrossChainEvent is a log parse operation binding the contract event 0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483.
//
// Solidity: event CrossChainEvent(address indexed sender, bytes txId, address token, uint64 toChainId, bytes toContract, bytes rawdata)
func (_EthCrossChain *EthCrossChainFilterer) ParseCrossChainEvent(log types.Log) (*EthCrossChainCrossChainEvent, error) {
	event := new(EthCrossChainCrossChainEvent)
	if err := _EthCrossChain.contract.UnpackLog(event, "CrossChainEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainInitGenesisBlockEventIterator is returned from FilterInitGenesisBlockEvent and is used to iterate over the raw logs and unpacked data for InitGenesisBlockEvent events raised by the EthCrossChain contract.
type EthCrossChainInitGenesisBlockEventIterator struct {
	Event *EthCrossChainInitGenesisBlockEvent // Event containing the contract specifics and raw log

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
func (it *EthCrossChainInitGenesisBlockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainInitGenesisBlockEvent)
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
		it.Event = new(EthCrossChainInitGenesisBlockEvent)
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
func (it *EthCrossChainInitGenesisBlockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainInitGenesisBlockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainInitGenesisBlockEvent represents a InitGenesisBlockEvent event raised by the EthCrossChain contract.
type EthCrossChainInitGenesisBlockEvent struct {
	Height          *big.Int
	TorawHeadersken []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterInitGenesisBlockEvent is a free log retrieval operation binding the contract event 0xf01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde.
//
// Solidity: event InitGenesisBlockEvent(uint256 height, bytes torawHeadersken)
func (_EthCrossChain *EthCrossChainFilterer) FilterInitGenesisBlockEvent(opts *bind.FilterOpts) (*EthCrossChainInitGenesisBlockEventIterator, error) {

	logs, sub, err := _EthCrossChain.contract.FilterLogs(opts, "InitGenesisBlockEvent")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainInitGenesisBlockEventIterator{contract: _EthCrossChain.contract, event: "InitGenesisBlockEvent", logs: logs, sub: sub}, nil
}

// WatchInitGenesisBlockEvent is a free log subscription operation binding the contract event 0xf01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde.
//
// Solidity: event InitGenesisBlockEvent(uint256 height, bytes torawHeadersken)
func (_EthCrossChain *EthCrossChainFilterer) WatchInitGenesisBlockEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainInitGenesisBlockEvent) (event.Subscription, error) {

	logs, sub, err := _EthCrossChain.contract.WatchLogs(opts, "InitGenesisBlockEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainInitGenesisBlockEvent)
				if err := _EthCrossChain.contract.UnpackLog(event, "InitGenesisBlockEvent", log); err != nil {
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

// ParseInitGenesisBlockEvent is a log parse operation binding the contract event 0xf01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde.
//
// Solidity: event InitGenesisBlockEvent(uint256 height, bytes torawHeadersken)
func (_EthCrossChain *EthCrossChainFilterer) ParseInitGenesisBlockEvent(log types.Log) (*EthCrossChainInitGenesisBlockEvent, error) {
	event := new(EthCrossChainInitGenesisBlockEvent)
	if err := _EthCrossChain.contract.UnpackLog(event, "InitGenesisBlockEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainSyncBlockHeaderEventIterator is returned from FilterSyncBlockHeaderEvent and is used to iterate over the raw logs and unpacked data for SyncBlockHeaderEvent events raised by the EthCrossChain contract.
type EthCrossChainSyncBlockHeaderEventIterator struct {
	Event *EthCrossChainSyncBlockHeaderEvent // Event containing the contract specifics and raw log

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
func (it *EthCrossChainSyncBlockHeaderEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainSyncBlockHeaderEvent)
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
		it.Event = new(EthCrossChainSyncBlockHeaderEvent)
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
func (it *EthCrossChainSyncBlockHeaderEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainSyncBlockHeaderEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainSyncBlockHeaderEvent represents a SyncBlockHeaderEvent event raised by the EthCrossChain contract.
type EthCrossChainSyncBlockHeaderEvent struct {
	Height     *big.Int
	RawHeaders []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSyncBlockHeaderEvent is a free log retrieval operation binding the contract event 0x1645cc3eff4fb98fef2ffd7806e7f39bfe18558f6656e69dfc9c8f26e958d4e7.
//
// Solidity: event SyncBlockHeaderEvent(uint256 height, bytes rawHeaders)
func (_EthCrossChain *EthCrossChainFilterer) FilterSyncBlockHeaderEvent(opts *bind.FilterOpts) (*EthCrossChainSyncBlockHeaderEventIterator, error) {

	logs, sub, err := _EthCrossChain.contract.FilterLogs(opts, "SyncBlockHeaderEvent")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainSyncBlockHeaderEventIterator{contract: _EthCrossChain.contract, event: "SyncBlockHeaderEvent", logs: logs, sub: sub}, nil
}

// WatchSyncBlockHeaderEvent is a free log subscription operation binding the contract event 0x1645cc3eff4fb98fef2ffd7806e7f39bfe18558f6656e69dfc9c8f26e958d4e7.
//
// Solidity: event SyncBlockHeaderEvent(uint256 height, bytes rawHeaders)
func (_EthCrossChain *EthCrossChainFilterer) WatchSyncBlockHeaderEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainSyncBlockHeaderEvent) (event.Subscription, error) {

	logs, sub, err := _EthCrossChain.contract.WatchLogs(opts, "SyncBlockHeaderEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainSyncBlockHeaderEvent)
				if err := _EthCrossChain.contract.UnpackLog(event, "SyncBlockHeaderEvent", log); err != nil {
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

// ParseSyncBlockHeaderEvent is a log parse operation binding the contract event 0x1645cc3eff4fb98fef2ffd7806e7f39bfe18558f6656e69dfc9c8f26e958d4e7.
//
// Solidity: event SyncBlockHeaderEvent(uint256 height, bytes rawHeaders)
func (_EthCrossChain *EthCrossChainFilterer) ParseSyncBlockHeaderEvent(log types.Log) (*EthCrossChainSyncBlockHeaderEvent, error) {
	event := new(EthCrossChainSyncBlockHeaderEvent)
	if err := _EthCrossChain.contract.UnpackLog(event, "SyncBlockHeaderEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainVerifyAndExecuteTxEventIterator is returned from FilterVerifyAndExecuteTxEvent and is used to iterate over the raw logs and unpacked data for VerifyAndExecuteTxEvent events raised by the EthCrossChain contract.
type EthCrossChainVerifyAndExecuteTxEventIterator struct {
	Event *EthCrossChainVerifyAndExecuteTxEvent // Event containing the contract specifics and raw log

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
func (it *EthCrossChainVerifyAndExecuteTxEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainVerifyAndExecuteTxEvent)
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
		it.Event = new(EthCrossChainVerifyAndExecuteTxEvent)
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
func (it *EthCrossChainVerifyAndExecuteTxEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainVerifyAndExecuteTxEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainVerifyAndExecuteTxEvent represents a VerifyAndExecuteTxEvent event raised by the EthCrossChain contract.
type EthCrossChainVerifyAndExecuteTxEvent struct {
	FromChainID      uint64
	ToContract       []byte
	CrossChainTxHash []byte
	FromChainTxHash  []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterVerifyAndExecuteTxEvent is a free log retrieval operation binding the contract event 0x5944a0908ae763fa254df8e0da54a682cbc62115501336e18be52a7dd774b80b.
//
// Solidity: event VerifyAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash)
func (_EthCrossChain *EthCrossChainFilterer) FilterVerifyAndExecuteTxEvent(opts *bind.FilterOpts) (*EthCrossChainVerifyAndExecuteTxEventIterator, error) {

	logs, sub, err := _EthCrossChain.contract.FilterLogs(opts, "VerifyAndExecuteTxEvent")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainVerifyAndExecuteTxEventIterator{contract: _EthCrossChain.contract, event: "VerifyAndExecuteTxEvent", logs: logs, sub: sub}, nil
}

// WatchVerifyAndExecuteTxEvent is a free log subscription operation binding the contract event 0x5944a0908ae763fa254df8e0da54a682cbc62115501336e18be52a7dd774b80b.
//
// Solidity: event VerifyAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash)
func (_EthCrossChain *EthCrossChainFilterer) WatchVerifyAndExecuteTxEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainVerifyAndExecuteTxEvent) (event.Subscription, error) {

	logs, sub, err := _EthCrossChain.contract.WatchLogs(opts, "VerifyAndExecuteTxEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainVerifyAndExecuteTxEvent)
				if err := _EthCrossChain.contract.UnpackLog(event, "VerifyAndExecuteTxEvent", log); err != nil {
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

// ParseVerifyAndExecuteTxEvent is a log parse operation binding the contract event 0x5944a0908ae763fa254df8e0da54a682cbc62115501336e18be52a7dd774b80b.
//
// Solidity: event VerifyAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash)
func (_EthCrossChain *EthCrossChainFilterer) ParseVerifyAndExecuteTxEvent(log types.Log) (*EthCrossChainVerifyAndExecuteTxEvent, error) {
	event := new(EthCrossChainVerifyAndExecuteTxEvent)
	if err := _EthCrossChain.contract.UnpackLog(event, "VerifyAndExecuteTxEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockProxyABI is the input ABI used to generate the binding from.
const LockProxyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"argsBs\",\"type\":\"bytes\"},{\"name\":\"fromContractAddr\",\"type\":\"bytes\"},{\"name\":\"fromChainId\",\"type\":\"uint64\"}],\"name\":\"unlock\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"assetTCCSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"managerContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"toChainId\",\"type\":\"uint64\"},{\"name\":\"targetProxyHash\",\"type\":\"bytes\"}],\"name\":\"bindProxyHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"assetTCCLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"assetHashMap\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txData\",\"type\":\"bytes\"}],\"name\":\"test_deserializeTxArgs\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sourceAssetHash\",\"type\":\"address\"},{\"name\":\"toChainId\",\"type\":\"uint64\"},{\"name\":\"toAddress\",\"type\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"proxyHashMap\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"targetAssetHash\",\"type\":\"bytes\"},{\"name\":\"toAddress\",\"type\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"test_serializeTxArgs\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ethCrossChainContractAddr\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sourceAssetHash\",\"type\":\"address\"},{\"name\":\"toChainId\",\"type\":\"uint64\"},{\"name\":\"targetAssetHash\",\"type\":\"bytes\"},{\"name\":\"assetLimit\",\"type\":\"uint256\"},{\"name\":\"isTargetChainAsset\",\"type\":\"bool\"}],\"name\":\"bindAssetHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"SetManagerEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"targetProxyHash\",\"type\":\"bytes\"}],\"name\":\"BindProxyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sourceAssetHash\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"targetProxyHash\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"assetLimit\",\"type\":\"uint256\"}],\"name\":\"BindAssetEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fromContractAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"toAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnlockEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"thisContract\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"toContract\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"txArgs\",\"type\":\"bytes\"}],\"name\":\"LockEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DebugBytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DebugUint256\",\"type\":\"event\"}]"

// LockProxyFuncSigs maps the 4-byte function signature to its string representation.
var LockProxyFuncSigs = map[string]string{
	"4f7d9808": "assetHashMap(address,uint64)",
	"3eb91723": "assetTCCLimit(address,uint64)",
	"170066f8": "assetTCCSupply(address,uint64)",
	"f525e595": "bindAssetHash(address,uint64,bytes,uint256,bool)",
	"379b98f6": "bindProxyHash(uint64,bytes)",
	"84a6d055": "lock(address,uint64,bytes,uint256)",
	"2710ac80": "managerContract()",
	"570ca735": "operator()",
	"9e5767aa": "proxyHashMap(uint64)",
	"d0ebdbe7": "setManager(address)",
	"6c34ab55": "test_deserializeTxArgs(bytes)",
	"b06ca5c8": "test_serializeTxArgs(bytes,bytes,uint256)",
	"06af4b9f": "unlock(bytes,bytes,uint64)",
}

// LockProxyBin is the compiled bytecode used for deploying new contracts.
var LockProxyBin = "0x608060405234801561001057600080fd5b50610022640100000000610047810204565b60008054600160a060020a031916600160a060020a039290921691909117905561004b565b3390565b612c41806200005b6000396000f3fe6080604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306af4b9f81146100c9578063170066f8146102235780632710ac8014610278578063379b98f6146102a95780633eb917231461036d5780634f7d9808146103b0578063570ca735146104685780636c34ab551461047d57806384a6d055146105305780639e5767aa146105f7578063b06ca5c81461062b578063d0ebdbe714610767578063f525e5951461079c575b600080fd5b3480156100d557600080fd5b5061020f600480360360608110156100ec57600080fd5b81019060208101813564010000000081111561010757600080fd5b82018360208201111561011957600080fd5b8035906020019184600183028401116401000000008311171561013b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561018e57600080fd5b8201836020820111156101a057600080fd5b803590602001918460018302840111640100000000831117156101c257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903567ffffffffffffffff1691506108759050565b604080519115158252519081900360200190f35b34801561022f57600080fd5b506102666004803603604081101561024657600080fd5b508035600160a060020a0316906020013567ffffffffffffffff16610bb0565b60408051918252519081900360200190f35b34801561028457600080fd5b5061028d610bcd565b60408051600160a060020a039092168252519081900360200190f35b3480156102b557600080fd5b5061020f600480360360408110156102cc57600080fd5b67ffffffffffffffff82351691908101906040810160208201356401000000008111156102f857600080fd5b82018360208201111561030a57600080fd5b8035906020019184600183028401116401000000008311171561032c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610bdc945050505050565b34801561037957600080fd5b506102666004803603604081101561039057600080fd5b508035600160a060020a0316906020013567ffffffffffffffff16610ced565b3480156103bc57600080fd5b506103f3600480360360408110156103d357600080fd5b508035600160a060020a0316906020013567ffffffffffffffff16610d0a565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561042d578181015183820152602001610415565b50505050905090810190601f16801561045a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561047457600080fd5b5061028d610dae565b34801561048957600080fd5b5061020f600480360360208110156104a057600080fd5b8101906020810181356401000000008111156104bb57600080fd5b8201836020820111156104cd57600080fd5b803590602001918460018302840111640100000000831117156104ef57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610dbd945050505050565b61020f6004803603608081101561054657600080fd5b600160a060020a038235169167ffffffffffffffff6020820135169181019060608101604082013564010000000081111561058057600080fd5b82018360208201111561059257600080fd5b803590602001918460018302840111640100000000831117156105b457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610f45915050565b34801561060357600080fd5b506103f36004803603602081101561061a57600080fd5b503567ffffffffffffffff166115b6565b34801561063757600080fd5b5061020f6004803603606081101561064e57600080fd5b81019060208101813564010000000081111561066957600080fd5b82018360208201111561067b57600080fd5b8035906020019184600183028401116401000000008311171561069d57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156106f057600080fd5b82018360208201111561070257600080fd5b8035906020019184600183028401116401000000008311171561072457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061161c915050565b34801561077357600080fd5b5061079a6004803603602081101561078a57600080fd5b5035600160a060020a03166116f6565b005b3480156107a857600080fd5b5061020f600480360360a08110156107bf57600080fd5b600160a060020a038235169167ffffffffffffffff602082013516918101906060810160408201356401000000008111156107f957600080fd5b82018360208201111561080b57600080fd5b8035906020019184600183028401116401000000008311171561082d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135151561177e565b600154600090600160a060020a031661088c611a0a565b600160a060020a03161461089f57600080fd5b6108a7612b3b565b6108b085611a0f565b67ffffffffffffffff841660009081526002602052604090209091506108d69085611a5d565b1515610952576040805160e560020a62461bcd02815260206004820152602260248201527f46726f6d2050726f787920636f6e74726163742061646472657373206572726f60448201527f7221000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60006109618260000151611b13565b905060006109728360200151611b13565b604080850151600160a060020a03851660009081526005602090815283822067ffffffffffffffff8b168352905291909120549192506109b8919063ffffffff611b9d16565b600160a060020a038316600090815260056020908152604080832067ffffffffffffffff8a16845290915290819020919091558301516109fb9083908390611be6565b1515610a77576040805160e560020a62461bcd02815260206004820152603c60248201527f7472616e736665722061737365742066726f6d206c6f636b5f70726f7879206360448201527f6f6e747261637420746f20746f41646472657373206661696c65642100000000606482015290519081900360840190fd5b7f31c3212616f0a6c018b96d403900949984f6cf1ac90e443ea80236327746935286868560200151866040015160405180806020018567ffffffffffffffff1667ffffffffffffffff16815260200180602001848152602001838103835287818151815260200191508051906020019080838360005b83811015610b05578181015183820152602001610aed565b50505050905090810190601f168015610b325780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015610b65578181015183820152602001610b4d565b50505050905090810190601f168015610b925780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a15060019695505050505050565b600560209081526000928352604080842090915290825290205481565b600154600160a060020a031681565b60008054600160a060020a0316610bf1611a0a565b600160a060020a031614610c0457600080fd5b67ffffffffffffffff831660009081526002602090815260409091208351610c2e92850190612b5d565b507fdacd7d303272a3b58aec6620d6d1fb588f4996a5b46858ed437f1c34348f2d0f8383604051808367ffffffffffffffff1667ffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610ca9578181015183820152602001610c91565b50505050905090810190601f168015610cd65780820380516001836020036101000a031916815260200191505b50935050505060405180910390a150600192915050565b600460209081526000928352604080842090915290825290205481565b60036020908152600092835260408084208252918352918190208054825160026001831615610100026000190190921691909104601f810185900485028201850190935282815292909190830182828015610da65780601f10610d7b57610100808354040283529160200191610da6565b820191906000526020600020905b815481529060010190602001808311610d8957829003601f168201915b505050505081565b600054600160a060020a031681565b6000610dc7612b3b565b610dd083611a0f565b805160408051602080825283518183015283519495507faf5a5af90a78ece430d7df503b54fc4070844db69884a9a4afb00710a4816e5394919283929083019185019080838360005b83811015610e31578181015183820152602001610e19565b50505050905090810190601f168015610e5e5780820380516001836020036101000a031916815260200191505b509250505060405180910390a17faf5a5af90a78ece430d7df503b54fc4070844db69884a9a4afb00710a4816e5381602001516040518080602001828103825283818151815260200191508051906020019080838360005b83811015610ece578181015183820152602001610eb6565b50505050905090810190601f168015610efb5780820380516001836020036101000a031916815260200191505b509250505060405180910390a1604080820151815190815290517f43d4b4706539f9e22baf8767ebea21ad24f723f14b6981664ac4d0af596dddbe9181900360200190a150919050565b600080821015610f9f576040805160e560020a62461bcd02815260206004820152601960248201527f616d6f756e74206973206c657373207468616e207a65726f2100000000000000604482015290519081900360640190fd5b600160a060020a038516600090815260056020908152604080832067ffffffffffffffff88168452909152902054610fdd908363ffffffff611cc516565b600160a060020a038616600081815260056020908152604080832067ffffffffffffffff8a16808552818452828520879055948452600483528184209484529382529091205491905210156110a2576040805160e560020a62461bcd02815260206004820152603060248201527f617373657420696e2074617267657420636861696e2077696c6c20657863656560448201527f64206c696d697420636f6e74726f6c2100000000000000000000000000000000606482015290519081900360840190fd5b6110ac8583611d22565b1515611128576040805160e560020a62461bcd02815260206004820152603f60248201527f7472616e736665722061737365742066726f6d2066726f6d416464726573732060448201527f746f206c6f636b5f70726f787920636f6e747261637420206661696c65642100606482015290519081900360840190fd5b600160a060020a038516600090815260036020908152604080832067ffffffffffffffff8816845282529182902080548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452606093928301828280156111d95780601f106111ae576101008083540402835291602001916111d9565b820191906000526020600020905b8154815290600101906020018083116111bc57829003601f168201915b505050505090506111e8612b3b565b606060405190810160405280838152602001868152602001858152509050606061121182611e5a565b6001805467ffffffffffffffff8a1660008181526002602081905260409182902091517fbd5cf6250000000000000000000000000000000000000000000000000000000081526004810193845260806024820190815283546000199781161561010002979097019096169190910460848201819052959650600160a060020a0390931694859463bd5cf625948e9489939092916044810191606482019160a40190879080156113015780601f106112d657610100808354040283529160200191611301565b820191906000526020600020905b8154815290600101906020018083116112e457829003601f168201915b5050848103835260068152602001807f756e6c6f636b0000000000000000000000000000000000000000000000000000815250602001848103825285818151815260200191508051906020019080838360005b8381101561136c578181015183820152602001611354565b50505050905090810190601f1680156113995780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b1580156113bc57600080fd5b505af11580156113d0573d6000803e3d6000fd5b505050506040513d60208110156113e657600080fd5b50511515611464576040805160e560020a62461bcd02815260206004820152602860248201527f45746843726f7373436861696e2063726f7373436861696e206578656375746560448201527f64206572726f7221000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b67ffffffffffffffff881660008181526002602081815260409283902083513080825292810195909552608093850184815281546000196101006001831615020116939093049385018490527f28094cc2d1bbe0fe894550907e1f9c6b8c9bb18cd72c4830534347b5974645c09491938d93919288929091606083019060a0840190869080156115355780601f1061150a57610100808354040283529160200191611535565b820191906000526020600020905b81548152906001019060200180831161151857829003601f168201915b5050838103825284518152845160209182019186019080838360005b83811015611569578181015183820152602001611551565b50505050905090810190601f1680156115965780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a150600198975050505050505050565b600260208181526000928352604092839020805484516001821615610100026000190190911693909304601f8101839004830284018301909452838352919290830182828015610da65780601f10610d7b57610100808354040283529160200191610da6565b6000611626612b3b565b606060405190810160405280868152602001858152602001848152509050606061164f82611e5a565b90507faf5a5af90a78ece430d7df503b54fc4070844db69884a9a4afb00710a4816e53816040518080602001828103825283818151815260200191508051906020019080838360005b838110156116b0578181015183820152602001611698565b50505050905090810190601f1680156116dd5780820380516001836020036101000a031916815260200191505b509250505060405180910390a150600195945050505050565b600054600160a060020a031661170a611a0a565b600160a060020a03161461171d57600080fd5b60018054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051918252517f45c21a4c9bc0568ee70fe6ae892fae1ce77b914d6b972041296072954ab21f8f9181900360200190a150565b60008054600160a060020a0316611793611a0a565b600160a060020a0316146117a657600080fd5b600160a060020a038616600090815260036020908152604080832067ffffffffffffffff89168452825290912085516117e192870190612b5d565b50811561191857600160a060020a038616600090815260046020908152604080832067ffffffffffffffff8916845290915290205480841015611894576040805160e560020a62461bcd02815260206004820152602c60248201527f61737365744c696d69742063616e206f6e6c792062652075706461746564206960448201527f6e6372656173696e676c79210000000000000000000000000000000000000000606482015290519081900360840190fd5b60006118a6858363ffffffff611b9d16565b600160a060020a038916600090815260056020908152604080832067ffffffffffffffff8c1684529091529020549091506118e7908263ffffffff611cc516565b600160a060020a038916600090815260056020908152604080832067ffffffffffffffff8c16845290915290205550505b600160a060020a038616600081815260046020908152604080832067ffffffffffffffff8a1680855290835281842088905581519485528483015260608401879052608090840181815288519185019190915287517f1628c8374c1bdfeb2275fd9f4c90562fd3fae974783dc522c8234e36abcfc58e948b948b948b948b94909260a0850192870191908190849084905b838110156119c15781810151838201526020016119a9565b50505050905090810190601f1680156119ee5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a150600195945050505050565b335b90565b611a17612b3b565b611a1f612b3b565b6000611a2b8482611f87565b9083529050611a3a8482611f87565b60208401919091529050611a4e8482612059565b5060408301525090505b919050565b60008060019050835460026001808316156101000203821604845180821460018114611a8c5760009450611b07565b8215611b07576020831060018114611aea57600189600052602060002060208a018581015b600284828410011415611ae1578151835414611ad05760009950600093505b600183019250602082019150611ab1565b50505050611b05565b610100808604029450602088015185141515611b0557600095505b505b50929695505050505050565b8051600090601414611b95576040805160e560020a62461bcd02815260206004820152602360248201527f6279746573206c656e67746820646f6573206e6f74206d61746368206164647260448201527f6573730000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b506014015190565b6000611bdf83836040805190810160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061217c565b9392505050565b6000600160a060020a0384161515611c3457604051600160a060020a0384169083156108fc029084906000818181858888f19350505050158015611c2e573d6000803e3d6000fd5b50611cbb565b611c3f848484612216565b1515611cbb576040805160e560020a62461bcd02815260206004820152603360248201527f7472616e7366657220657263323020617373657420746f206c6f636b5f70726f60448201527f787920636f6e7472616374206661696c65642100000000000000000000000000606482015290519081900360840190fd5b5060019392505050565b600082820183811015611bdf576040805160e560020a62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6000600160a060020a038316158015611d3b5750600034115b15611dc257348214611dbd576040805160e560020a62461bcd02815260206004820152602960248201527f7472616e73666572726564206574686572206973206e6f7420657175616c207460448201527f6f20616d6f756e74210000000000000000000000000000000000000000000000606482015290519081900360840190fd5b611e51565b611dd583611dce611a0a565b3085612313565b1515611e51576040805160e560020a62461bcd02815260206004820152603360248201527f7472616e7366657220657263323020617373657420746f206c6f636b5f70726f60448201527f787920636f6e7472616374206661696c65642100000000000000000000000000606482015290519081900360840190fd5b50600192915050565b606080611e6a8360000151612419565b611e778460200151612419565b611e8485604001516124df565b6040516020018084805190602001908083835b60208310611eb65780518252601f199092019160209182019101611e97565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b60208310611efe5780518252601f199092019160209182019101611edf565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310611f465780518252601f199092019160209182019101611f27565b6001836020036101000a0380198251168184511680821785525050505050509050019350505050604051602081830303815290604052905080915050919050565b6060600080611f9685856125a7565b86519095509091508185011115611fe5576040805160e560020a62461bcd0281526020600482015260166024820152600080516020612bf6833981519152604482015290519081900360640190fd5b6060811580156120005760405191506020820160405261204a565b6040519150601f8316801560200281840101848101888315602002848c0101015b81831015612039578051835260209283019201612021565b5050848452601f01601f1916604052505b509250830190505b9250929050565b600080835183602001111515156120a8576040805160e560020a62461bcd0281526020600482015260166024820152600080516020612bf6833981519152604482015290519081900360640190fd5b600060405160206000600182038760208a0101515b838310156120dd5780821a838601536001830192506001820391506120bd565b50505081016040525190506000811080159061211957507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8111155b151561216f576040805160e560020a62461bcd02815260206004820152601760248201527f56616c75652065786365656473207468652072616e6765000000000000000000604482015290519081900360640190fd5b9460209390930193505050565b6000818484111561220e5760405160e560020a62461bcd0281526004018080602001828103825283818151815260200191508051906020019080838360005b838110156121d35781810151838201526020016121bb565b50505050905090810190601f1680156122005780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015260248201849052915160009286929083169163a9059cbb9160448082019260209290919082900301818887803b15801561228657600080fd5b505af115801561229a573d6000803e3d6000fd5b505050506040513d60208110156122b057600080fd5b50511515612308576040805160e560020a62461bcd02815260206004820152601c60248201527f747261736e66657220455243323020546f6b656e206661696c65642100000000604482015290519081900360640190fd5b506001949350505050565b604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015284811660248301526044820184905291516000928792908316916323b872dd9160648082019260209290919082900301818887803b15801561238b57600080fd5b505af115801561239f573d6000803e3d6000fd5b505050506040513d60208110156123b557600080fd5b5051151561240d576040805160e560020a62461bcd02815260206004820152601c60248201527f747261736e66657220455243323020546f6b656e206661696c65642100000000604482015290519081900360640190fd5b50600195945050505050565b8051606090612427816126b9565b836040516020018083805190602001908083835b6020831061245a5780518252601f19909201916020918201910161243b565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106124a25780518252601f199092019160209182019101612483565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052915050919050565b60606000821015801561251257507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8211155b1515612568576040805160e560020a62461bcd02815260206004820152601760248201527f56616c75652065786365656473207468652072616e6765000000000000000000604482015290519081900360640190fd5b60405160208082526000601f5b828210156125975785811a826020860101536001919091019060001901612575565b5050506040818101905292915050565b60008060006125b6858561283a565b945090507ffd00000000000000000000000000000000000000000000000000000000000000600160f860020a031982161415612607576125f685856128af565b8161ffff1691509250925050612052565b7ffe00000000000000000000000000000000000000000000000000000000000000600160f860020a031982161415612656576126438585612933565b8163ffffffff1691509250925050612052565b600160f860020a031980821614156126895761267285856129d4565b8167ffffffffffffffff1691509250925050612052565b7f0100000000000000000000000000000000000000000000000000000000000000900460ff169150829050612052565b606060fd8267ffffffffffffffff1610156126de576126d782612a75565b9050611a58565b61ffff67ffffffffffffffff8316116127d95761271a7ffd00000000000000000000000000000000000000000000000000000000000000612a90565b61272383612ac3565b6040516020018083805190602001908083835b602083106127555780518252601f199092019160209182019101612736565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061279d5780518252601f19909201916020918201910161277e565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529050611a58565b63ffffffff67ffffffffffffffff831611612820576128177ffe00000000000000000000000000000000000000000000000000000000000000612a90565b61272383612aeb565b612831600160f860020a0319612a90565b61272383612b13565b6000808351836001011115151561289b576040805160e560020a62461bcd02815260206004820152601660248201527f4f66667365742065786365656473206d6178696d756d00000000000000000000604482015290519081900360640190fd5b505081810160200151600182019250929050565b600080835183600201111515156128fe576040805160e560020a62461bcd0281526020600482015260166024820152600080516020612bf6833981519152604482015290519081900360640190fd5b6000604051846020870101518060011a82538060001a6001830153506002818101604052601d19909101519694019450505050565b60008083518360040111151515612982576040805160e560020a62461bcd0281526020600482015260166024820152600080516020612bf6833981519152604482015290519081900360640190fd5b600060405160046000600182038760208a0101515b838310156129b75780821a83860153600183019250600182039150612997565b505050808201604052602003900351956004949094019450505050565b60008083518360080111151515612a23576040805160e560020a62461bcd0281526020600482015260166024820152600080516020612bf6833981519152604482015290519081900360640190fd5b600060405160086000600182038760208a0101515b83831015612a585780821a83860153600183019250600182039150612a38565b505050808201604052602003900351956008949094019450505050565b604080516001815260f89290921b6020830152818101905290565b6060612abd7f01000000000000000000000000000000000000000000000000000000000000008304612a75565b92915050565b6040516002808252606091906000601f85811a60208501536001919091019060001901612575565b6040516004808252606091906000601f85811a60208501536001919091019060001901612575565b6040516008808252606091906000601f85811a60208501536001919091019060001901612575565b6060604051908101604052806060815260200160608152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10612b9e57805160ff1916838001178555612bcb565b82800160010185558215612bcb579182015b82811115612bcb578251825591602001919060010190612bb0565b50612bd7929150612bdb565b5090565b611a0c91905b80821115612bd75760008155600101612be156fe6f66667365742065786365656473206d6178696d756d00000000000000000000a165627a7a72305820f41dad409631fe5bda4861210f1447547c95a146cdd99f04ad2de9d8f9c888250029"

// DeployLockProxy deploys a new Ethereum contract, binding an instance of LockProxy to it.
func DeployLockProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LockProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(LockProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LockProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LockProxy{LockProxyCaller: LockProxyCaller{contract: contract}, LockProxyTransactor: LockProxyTransactor{contract: contract}, LockProxyFilterer: LockProxyFilterer{contract: contract}}, nil
}

// LockProxy is an auto generated Go binding around an Ethereum contract.
type LockProxy struct {
	LockProxyCaller     // Read-only binding to the contract
	LockProxyTransactor // Write-only binding to the contract
	LockProxyFilterer   // Log filterer for contract events
}

// LockProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockProxySession struct {
	Contract     *LockProxy        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockProxyCallerSession struct {
	Contract *LockProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// LockProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockProxyTransactorSession struct {
	Contract     *LockProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LockProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockProxyRaw struct {
	Contract *LockProxy // Generic contract binding to access the raw methods on
}

// LockProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockProxyCallerRaw struct {
	Contract *LockProxyCaller // Generic read-only contract binding to access the raw methods on
}

// LockProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockProxyTransactorRaw struct {
	Contract *LockProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLockProxy creates a new instance of LockProxy, bound to a specific deployed contract.
func NewLockProxy(address common.Address, backend bind.ContractBackend) (*LockProxy, error) {
	contract, err := bindLockProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LockProxy{LockProxyCaller: LockProxyCaller{contract: contract}, LockProxyTransactor: LockProxyTransactor{contract: contract}, LockProxyFilterer: LockProxyFilterer{contract: contract}}, nil
}

// NewLockProxyCaller creates a new read-only instance of LockProxy, bound to a specific deployed contract.
func NewLockProxyCaller(address common.Address, caller bind.ContractCaller) (*LockProxyCaller, error) {
	contract, err := bindLockProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockProxyCaller{contract: contract}, nil
}

// NewLockProxyTransactor creates a new write-only instance of LockProxy, bound to a specific deployed contract.
func NewLockProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*LockProxyTransactor, error) {
	contract, err := bindLockProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockProxyTransactor{contract: contract}, nil
}

// NewLockProxyFilterer creates a new log filterer instance of LockProxy, bound to a specific deployed contract.
func NewLockProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*LockProxyFilterer, error) {
	contract, err := bindLockProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockProxyFilterer{contract: contract}, nil
}

// bindLockProxy binds a generic wrapper to an already deployed contract.
func bindLockProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockProxy *LockProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LockProxy.Contract.LockProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockProxy *LockProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockProxy.Contract.LockProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockProxy *LockProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockProxy.Contract.LockProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockProxy *LockProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LockProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockProxy *LockProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockProxy *LockProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockProxy.Contract.contract.Transact(opts, method, params...)
}

// AssetHashMap is a free data retrieval call binding the contract method 0x4f7d9808.
//
// Solidity: function assetHashMap(address , uint64 ) constant returns(bytes)
func (_LockProxy *LockProxyCaller) AssetHashMap(opts *bind.CallOpts, arg0 common.Address, arg1 uint64) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _LockProxy.contract.Call(opts, out, "assetHashMap", arg0, arg1)
	return *ret0, err
}

// AssetHashMap is a free data retrieval call binding the contract method 0x4f7d9808.
//
// Solidity: function assetHashMap(address , uint64 ) constant returns(bytes)
func (_LockProxy *LockProxySession) AssetHashMap(arg0 common.Address, arg1 uint64) ([]byte, error) {
	return _LockProxy.Contract.AssetHashMap(&_LockProxy.CallOpts, arg0, arg1)
}

// AssetHashMap is a free data retrieval call binding the contract method 0x4f7d9808.
//
// Solidity: function assetHashMap(address , uint64 ) constant returns(bytes)
func (_LockProxy *LockProxyCallerSession) AssetHashMap(arg0 common.Address, arg1 uint64) ([]byte, error) {
	return _LockProxy.Contract.AssetHashMap(&_LockProxy.CallOpts, arg0, arg1)
}

// AssetTCCLimit is a free data retrieval call binding the contract method 0x3eb91723.
//
// Solidity: function assetTCCLimit(address , uint64 ) constant returns(uint256)
func (_LockProxy *LockProxyCaller) AssetTCCLimit(opts *bind.CallOpts, arg0 common.Address, arg1 uint64) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LockProxy.contract.Call(opts, out, "assetTCCLimit", arg0, arg1)
	return *ret0, err
}

// AssetTCCLimit is a free data retrieval call binding the contract method 0x3eb91723.
//
// Solidity: function assetTCCLimit(address , uint64 ) constant returns(uint256)
func (_LockProxy *LockProxySession) AssetTCCLimit(arg0 common.Address, arg1 uint64) (*big.Int, error) {
	return _LockProxy.Contract.AssetTCCLimit(&_LockProxy.CallOpts, arg0, arg1)
}

// AssetTCCLimit is a free data retrieval call binding the contract method 0x3eb91723.
//
// Solidity: function assetTCCLimit(address , uint64 ) constant returns(uint256)
func (_LockProxy *LockProxyCallerSession) AssetTCCLimit(arg0 common.Address, arg1 uint64) (*big.Int, error) {
	return _LockProxy.Contract.AssetTCCLimit(&_LockProxy.CallOpts, arg0, arg1)
}

// AssetTCCSupply is a free data retrieval call binding the contract method 0x170066f8.
//
// Solidity: function assetTCCSupply(address , uint64 ) constant returns(uint256)
func (_LockProxy *LockProxyCaller) AssetTCCSupply(opts *bind.CallOpts, arg0 common.Address, arg1 uint64) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LockProxy.contract.Call(opts, out, "assetTCCSupply", arg0, arg1)
	return *ret0, err
}

// AssetTCCSupply is a free data retrieval call binding the contract method 0x170066f8.
//
// Solidity: function assetTCCSupply(address , uint64 ) constant returns(uint256)
func (_LockProxy *LockProxySession) AssetTCCSupply(arg0 common.Address, arg1 uint64) (*big.Int, error) {
	return _LockProxy.Contract.AssetTCCSupply(&_LockProxy.CallOpts, arg0, arg1)
}

// AssetTCCSupply is a free data retrieval call binding the contract method 0x170066f8.
//
// Solidity: function assetTCCSupply(address , uint64 ) constant returns(uint256)
func (_LockProxy *LockProxyCallerSession) AssetTCCSupply(arg0 common.Address, arg1 uint64) (*big.Int, error) {
	return _LockProxy.Contract.AssetTCCSupply(&_LockProxy.CallOpts, arg0, arg1)
}

// ManagerContract is a free data retrieval call binding the contract method 0x2710ac80.
//
// Solidity: function managerContract() constant returns(address)
func (_LockProxy *LockProxyCaller) ManagerContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LockProxy.contract.Call(opts, out, "managerContract")
	return *ret0, err
}

// ManagerContract is a free data retrieval call binding the contract method 0x2710ac80.
//
// Solidity: function managerContract() constant returns(address)
func (_LockProxy *LockProxySession) ManagerContract() (common.Address, error) {
	return _LockProxy.Contract.ManagerContract(&_LockProxy.CallOpts)
}

// ManagerContract is a free data retrieval call binding the contract method 0x2710ac80.
//
// Solidity: function managerContract() constant returns(address)
func (_LockProxy *LockProxyCallerSession) ManagerContract() (common.Address, error) {
	return _LockProxy.Contract.ManagerContract(&_LockProxy.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_LockProxy *LockProxyCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LockProxy.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_LockProxy *LockProxySession) Operator() (common.Address, error) {
	return _LockProxy.Contract.Operator(&_LockProxy.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_LockProxy *LockProxyCallerSession) Operator() (common.Address, error) {
	return _LockProxy.Contract.Operator(&_LockProxy.CallOpts)
}

// ProxyHashMap is a free data retrieval call binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) constant returns(bytes)
func (_LockProxy *LockProxyCaller) ProxyHashMap(opts *bind.CallOpts, arg0 uint64) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _LockProxy.contract.Call(opts, out, "proxyHashMap", arg0)
	return *ret0, err
}

// ProxyHashMap is a free data retrieval call binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) constant returns(bytes)
func (_LockProxy *LockProxySession) ProxyHashMap(arg0 uint64) ([]byte, error) {
	return _LockProxy.Contract.ProxyHashMap(&_LockProxy.CallOpts, arg0)
}

// ProxyHashMap is a free data retrieval call binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) constant returns(bytes)
func (_LockProxy *LockProxyCallerSession) ProxyHashMap(arg0 uint64) ([]byte, error) {
	return _LockProxy.Contract.ProxyHashMap(&_LockProxy.CallOpts, arg0)
}

// BindAssetHash is a paid mutator transaction binding the contract method 0xf525e595.
//
// Solidity: function bindAssetHash(address sourceAssetHash, uint64 toChainId, bytes targetAssetHash, uint256 assetLimit, bool isTargetChainAsset) returns(bool)
func (_LockProxy *LockProxyTransactor) BindAssetHash(opts *bind.TransactOpts, sourceAssetHash common.Address, toChainId uint64, targetAssetHash []byte, assetLimit *big.Int, isTargetChainAsset bool) (*types.Transaction, error) {
	return _LockProxy.contract.Transact(opts, "bindAssetHash", sourceAssetHash, toChainId, targetAssetHash, assetLimit, isTargetChainAsset)
}

// BindAssetHash is a paid mutator transaction binding the contract method 0xf525e595.
//
// Solidity: function bindAssetHash(address sourceAssetHash, uint64 toChainId, bytes targetAssetHash, uint256 assetLimit, bool isTargetChainAsset) returns(bool)
func (_LockProxy *LockProxySession) BindAssetHash(sourceAssetHash common.Address, toChainId uint64, targetAssetHash []byte, assetLimit *big.Int, isTargetChainAsset bool) (*types.Transaction, error) {
	return _LockProxy.Contract.BindAssetHash(&_LockProxy.TransactOpts, sourceAssetHash, toChainId, targetAssetHash, assetLimit, isTargetChainAsset)
}

// BindAssetHash is a paid mutator transaction binding the contract method 0xf525e595.
//
// Solidity: function bindAssetHash(address sourceAssetHash, uint64 toChainId, bytes targetAssetHash, uint256 assetLimit, bool isTargetChainAsset) returns(bool)
func (_LockProxy *LockProxyTransactorSession) BindAssetHash(sourceAssetHash common.Address, toChainId uint64, targetAssetHash []byte, assetLimit *big.Int, isTargetChainAsset bool) (*types.Transaction, error) {
	return _LockProxy.Contract.BindAssetHash(&_LockProxy.TransactOpts, sourceAssetHash, toChainId, targetAssetHash, assetLimit, isTargetChainAsset)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 toChainId, bytes targetProxyHash) returns(bool)
func (_LockProxy *LockProxyTransactor) BindProxyHash(opts *bind.TransactOpts, toChainId uint64, targetProxyHash []byte) (*types.Transaction, error) {
	return _LockProxy.contract.Transact(opts, "bindProxyHash", toChainId, targetProxyHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 toChainId, bytes targetProxyHash) returns(bool)
func (_LockProxy *LockProxySession) BindProxyHash(toChainId uint64, targetProxyHash []byte) (*types.Transaction, error) {
	return _LockProxy.Contract.BindProxyHash(&_LockProxy.TransactOpts, toChainId, targetProxyHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 toChainId, bytes targetProxyHash) returns(bool)
func (_LockProxy *LockProxyTransactorSession) BindProxyHash(toChainId uint64, targetProxyHash []byte) (*types.Transaction, error) {
	return _LockProxy.Contract.BindProxyHash(&_LockProxy.TransactOpts, toChainId, targetProxyHash)
}

// Lock is a paid mutator transaction binding the contract method 0x84a6d055.
//
// Solidity: function lock(address sourceAssetHash, uint64 toChainId, bytes toAddress, uint256 amount) returns(bool)
func (_LockProxy *LockProxyTransactor) Lock(opts *bind.TransactOpts, sourceAssetHash common.Address, toChainId uint64, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _LockProxy.contract.Transact(opts, "lock", sourceAssetHash, toChainId, toAddress, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x84a6d055.
//
// Solidity: function lock(address sourceAssetHash, uint64 toChainId, bytes toAddress, uint256 amount) returns(bool)
func (_LockProxy *LockProxySession) Lock(sourceAssetHash common.Address, toChainId uint64, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _LockProxy.Contract.Lock(&_LockProxy.TransactOpts, sourceAssetHash, toChainId, toAddress, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x84a6d055.
//
// Solidity: function lock(address sourceAssetHash, uint64 toChainId, bytes toAddress, uint256 amount) returns(bool)
func (_LockProxy *LockProxyTransactorSession) Lock(sourceAssetHash common.Address, toChainId uint64, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _LockProxy.Contract.Lock(&_LockProxy.TransactOpts, sourceAssetHash, toChainId, toAddress, amount)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address ethCrossChainContractAddr) returns()
func (_LockProxy *LockProxyTransactor) SetManager(opts *bind.TransactOpts, ethCrossChainContractAddr common.Address) (*types.Transaction, error) {
	return _LockProxy.contract.Transact(opts, "setManager", ethCrossChainContractAddr)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address ethCrossChainContractAddr) returns()
func (_LockProxy *LockProxySession) SetManager(ethCrossChainContractAddr common.Address) (*types.Transaction, error) {
	return _LockProxy.Contract.SetManager(&_LockProxy.TransactOpts, ethCrossChainContractAddr)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address ethCrossChainContractAddr) returns()
func (_LockProxy *LockProxyTransactorSession) SetManager(ethCrossChainContractAddr common.Address) (*types.Transaction, error) {
	return _LockProxy.Contract.SetManager(&_LockProxy.TransactOpts, ethCrossChainContractAddr)
}

// TestDeserializeTxArgs is a paid mutator transaction binding the contract method 0x6c34ab55.
//
// Solidity: function test_deserializeTxArgs(bytes txData) returns(bool)
func (_LockProxy *LockProxyTransactor) TestDeserializeTxArgs(opts *bind.TransactOpts, txData []byte) (*types.Transaction, error) {
	return _LockProxy.contract.Transact(opts, "test_deserializeTxArgs", txData)
}

// TestDeserializeTxArgs is a paid mutator transaction binding the contract method 0x6c34ab55.
//
// Solidity: function test_deserializeTxArgs(bytes txData) returns(bool)
func (_LockProxy *LockProxySession) TestDeserializeTxArgs(txData []byte) (*types.Transaction, error) {
	return _LockProxy.Contract.TestDeserializeTxArgs(&_LockProxy.TransactOpts, txData)
}

// TestDeserializeTxArgs is a paid mutator transaction binding the contract method 0x6c34ab55.
//
// Solidity: function test_deserializeTxArgs(bytes txData) returns(bool)
func (_LockProxy *LockProxyTransactorSession) TestDeserializeTxArgs(txData []byte) (*types.Transaction, error) {
	return _LockProxy.Contract.TestDeserializeTxArgs(&_LockProxy.TransactOpts, txData)
}

// TestSerializeTxArgs is a paid mutator transaction binding the contract method 0xb06ca5c8.
//
// Solidity: function test_serializeTxArgs(bytes targetAssetHash, bytes toAddress, uint256 amount) returns(bool)
func (_LockProxy *LockProxyTransactor) TestSerializeTxArgs(opts *bind.TransactOpts, targetAssetHash []byte, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _LockProxy.contract.Transact(opts, "test_serializeTxArgs", targetAssetHash, toAddress, amount)
}

// TestSerializeTxArgs is a paid mutator transaction binding the contract method 0xb06ca5c8.
//
// Solidity: function test_serializeTxArgs(bytes targetAssetHash, bytes toAddress, uint256 amount) returns(bool)
func (_LockProxy *LockProxySession) TestSerializeTxArgs(targetAssetHash []byte, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _LockProxy.Contract.TestSerializeTxArgs(&_LockProxy.TransactOpts, targetAssetHash, toAddress, amount)
}

// TestSerializeTxArgs is a paid mutator transaction binding the contract method 0xb06ca5c8.
//
// Solidity: function test_serializeTxArgs(bytes targetAssetHash, bytes toAddress, uint256 amount) returns(bool)
func (_LockProxy *LockProxyTransactorSession) TestSerializeTxArgs(targetAssetHash []byte, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _LockProxy.Contract.TestSerializeTxArgs(&_LockProxy.TransactOpts, targetAssetHash, toAddress, amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x06af4b9f.
//
// Solidity: function unlock(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_LockProxy *LockProxyTransactor) Unlock(opts *bind.TransactOpts, argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _LockProxy.contract.Transact(opts, "unlock", argsBs, fromContractAddr, fromChainId)
}

// Unlock is a paid mutator transaction binding the contract method 0x06af4b9f.
//
// Solidity: function unlock(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_LockProxy *LockProxySession) Unlock(argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _LockProxy.Contract.Unlock(&_LockProxy.TransactOpts, argsBs, fromContractAddr, fromChainId)
}

// Unlock is a paid mutator transaction binding the contract method 0x06af4b9f.
//
// Solidity: function unlock(bytes argsBs, bytes fromContractAddr, uint64 fromChainId) returns(bool)
func (_LockProxy *LockProxyTransactorSession) Unlock(argsBs []byte, fromContractAddr []byte, fromChainId uint64) (*types.Transaction, error) {
	return _LockProxy.Contract.Unlock(&_LockProxy.TransactOpts, argsBs, fromContractAddr, fromChainId)
}

// LockProxyBindAssetEventIterator is returned from FilterBindAssetEvent and is used to iterate over the raw logs and unpacked data for BindAssetEvent events raised by the LockProxy contract.
type LockProxyBindAssetEventIterator struct {
	Event *LockProxyBindAssetEvent // Event containing the contract specifics and raw log

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
func (it *LockProxyBindAssetEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockProxyBindAssetEvent)
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
		it.Event = new(LockProxyBindAssetEvent)
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
func (it *LockProxyBindAssetEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockProxyBindAssetEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockProxyBindAssetEvent represents a BindAssetEvent event raised by the LockProxy contract.
type LockProxyBindAssetEvent struct {
	SourceAssetHash common.Address
	ToChainId       uint64
	TargetProxyHash []byte
	AssetLimit      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBindAssetEvent is a free log retrieval operation binding the contract event 0x1628c8374c1bdfeb2275fd9f4c90562fd3fae974783dc522c8234e36abcfc58e.
//
// Solidity: event BindAssetEvent(address sourceAssetHash, uint64 toChainId, bytes targetProxyHash, uint256 assetLimit)
func (_LockProxy *LockProxyFilterer) FilterBindAssetEvent(opts *bind.FilterOpts) (*LockProxyBindAssetEventIterator, error) {

	logs, sub, err := _LockProxy.contract.FilterLogs(opts, "BindAssetEvent")
	if err != nil {
		return nil, err
	}
	return &LockProxyBindAssetEventIterator{contract: _LockProxy.contract, event: "BindAssetEvent", logs: logs, sub: sub}, nil
}

// WatchBindAssetEvent is a free log subscription operation binding the contract event 0x1628c8374c1bdfeb2275fd9f4c90562fd3fae974783dc522c8234e36abcfc58e.
//
// Solidity: event BindAssetEvent(address sourceAssetHash, uint64 toChainId, bytes targetProxyHash, uint256 assetLimit)
func (_LockProxy *LockProxyFilterer) WatchBindAssetEvent(opts *bind.WatchOpts, sink chan<- *LockProxyBindAssetEvent) (event.Subscription, error) {

	logs, sub, err := _LockProxy.contract.WatchLogs(opts, "BindAssetEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockProxyBindAssetEvent)
				if err := _LockProxy.contract.UnpackLog(event, "BindAssetEvent", log); err != nil {
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

// ParseBindAssetEvent is a log parse operation binding the contract event 0x1628c8374c1bdfeb2275fd9f4c90562fd3fae974783dc522c8234e36abcfc58e.
//
// Solidity: event BindAssetEvent(address sourceAssetHash, uint64 toChainId, bytes targetProxyHash, uint256 assetLimit)
func (_LockProxy *LockProxyFilterer) ParseBindAssetEvent(log types.Log) (*LockProxyBindAssetEvent, error) {
	event := new(LockProxyBindAssetEvent)
	if err := _LockProxy.contract.UnpackLog(event, "BindAssetEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockProxyBindProxyEventIterator is returned from FilterBindProxyEvent and is used to iterate over the raw logs and unpacked data for BindProxyEvent events raised by the LockProxy contract.
type LockProxyBindProxyEventIterator struct {
	Event *LockProxyBindProxyEvent // Event containing the contract specifics and raw log

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
func (it *LockProxyBindProxyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockProxyBindProxyEvent)
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
		it.Event = new(LockProxyBindProxyEvent)
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
func (it *LockProxyBindProxyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockProxyBindProxyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockProxyBindProxyEvent represents a BindProxyEvent event raised by the LockProxy contract.
type LockProxyBindProxyEvent struct {
	ToChainId       uint64
	TargetProxyHash []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBindProxyEvent is a free log retrieval operation binding the contract event 0xdacd7d303272a3b58aec6620d6d1fb588f4996a5b46858ed437f1c34348f2d0f.
//
// Solidity: event BindProxyEvent(uint64 toChainId, bytes targetProxyHash)
func (_LockProxy *LockProxyFilterer) FilterBindProxyEvent(opts *bind.FilterOpts) (*LockProxyBindProxyEventIterator, error) {

	logs, sub, err := _LockProxy.contract.FilterLogs(opts, "BindProxyEvent")
	if err != nil {
		return nil, err
	}
	return &LockProxyBindProxyEventIterator{contract: _LockProxy.contract, event: "BindProxyEvent", logs: logs, sub: sub}, nil
}

// WatchBindProxyEvent is a free log subscription operation binding the contract event 0xdacd7d303272a3b58aec6620d6d1fb588f4996a5b46858ed437f1c34348f2d0f.
//
// Solidity: event BindProxyEvent(uint64 toChainId, bytes targetProxyHash)
func (_LockProxy *LockProxyFilterer) WatchBindProxyEvent(opts *bind.WatchOpts, sink chan<- *LockProxyBindProxyEvent) (event.Subscription, error) {

	logs, sub, err := _LockProxy.contract.WatchLogs(opts, "BindProxyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockProxyBindProxyEvent)
				if err := _LockProxy.contract.UnpackLog(event, "BindProxyEvent", log); err != nil {
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

// ParseBindProxyEvent is a log parse operation binding the contract event 0xdacd7d303272a3b58aec6620d6d1fb588f4996a5b46858ed437f1c34348f2d0f.
//
// Solidity: event BindProxyEvent(uint64 toChainId, bytes targetProxyHash)
func (_LockProxy *LockProxyFilterer) ParseBindProxyEvent(log types.Log) (*LockProxyBindProxyEvent, error) {
	event := new(LockProxyBindProxyEvent)
	if err := _LockProxy.contract.UnpackLog(event, "BindProxyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockProxyDebugBytesIterator is returned from FilterDebugBytes and is used to iterate over the raw logs and unpacked data for DebugBytes events raised by the LockProxy contract.
type LockProxyDebugBytesIterator struct {
	Event *LockProxyDebugBytes // Event containing the contract specifics and raw log

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
func (it *LockProxyDebugBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockProxyDebugBytes)
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
		it.Event = new(LockProxyDebugBytes)
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
func (it *LockProxyDebugBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockProxyDebugBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockProxyDebugBytes represents a DebugBytes event raised by the LockProxy contract.
type LockProxyDebugBytes struct {
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDebugBytes is a free log retrieval operation binding the contract event 0xaf5a5af90a78ece430d7df503b54fc4070844db69884a9a4afb00710a4816e53.
//
// Solidity: event DebugBytes(bytes data)
func (_LockProxy *LockProxyFilterer) FilterDebugBytes(opts *bind.FilterOpts) (*LockProxyDebugBytesIterator, error) {

	logs, sub, err := _LockProxy.contract.FilterLogs(opts, "DebugBytes")
	if err != nil {
		return nil, err
	}
	return &LockProxyDebugBytesIterator{contract: _LockProxy.contract, event: "DebugBytes", logs: logs, sub: sub}, nil
}

// WatchDebugBytes is a free log subscription operation binding the contract event 0xaf5a5af90a78ece430d7df503b54fc4070844db69884a9a4afb00710a4816e53.
//
// Solidity: event DebugBytes(bytes data)
func (_LockProxy *LockProxyFilterer) WatchDebugBytes(opts *bind.WatchOpts, sink chan<- *LockProxyDebugBytes) (event.Subscription, error) {

	logs, sub, err := _LockProxy.contract.WatchLogs(opts, "DebugBytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockProxyDebugBytes)
				if err := _LockProxy.contract.UnpackLog(event, "DebugBytes", log); err != nil {
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

// ParseDebugBytes is a log parse operation binding the contract event 0xaf5a5af90a78ece430d7df503b54fc4070844db69884a9a4afb00710a4816e53.
//
// Solidity: event DebugBytes(bytes data)
func (_LockProxy *LockProxyFilterer) ParseDebugBytes(log types.Log) (*LockProxyDebugBytes, error) {
	event := new(LockProxyDebugBytes)
	if err := _LockProxy.contract.UnpackLog(event, "DebugBytes", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockProxyDebugUint256Iterator is returned from FilterDebugUint256 and is used to iterate over the raw logs and unpacked data for DebugUint256 events raised by the LockProxy contract.
type LockProxyDebugUint256Iterator struct {
	Event *LockProxyDebugUint256 // Event containing the contract specifics and raw log

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
func (it *LockProxyDebugUint256Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockProxyDebugUint256)
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
		it.Event = new(LockProxyDebugUint256)
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
func (it *LockProxyDebugUint256Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockProxyDebugUint256Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockProxyDebugUint256 represents a DebugUint256 event raised by the LockProxy contract.
type LockProxyDebugUint256 struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDebugUint256 is a free log retrieval operation binding the contract event 0x43d4b4706539f9e22baf8767ebea21ad24f723f14b6981664ac4d0af596dddbe.
//
// Solidity: event DebugUint256(uint256 amount)
func (_LockProxy *LockProxyFilterer) FilterDebugUint256(opts *bind.FilterOpts) (*LockProxyDebugUint256Iterator, error) {

	logs, sub, err := _LockProxy.contract.FilterLogs(opts, "DebugUint256")
	if err != nil {
		return nil, err
	}
	return &LockProxyDebugUint256Iterator{contract: _LockProxy.contract, event: "DebugUint256", logs: logs, sub: sub}, nil
}

// WatchDebugUint256 is a free log subscription operation binding the contract event 0x43d4b4706539f9e22baf8767ebea21ad24f723f14b6981664ac4d0af596dddbe.
//
// Solidity: event DebugUint256(uint256 amount)
func (_LockProxy *LockProxyFilterer) WatchDebugUint256(opts *bind.WatchOpts, sink chan<- *LockProxyDebugUint256) (event.Subscription, error) {

	logs, sub, err := _LockProxy.contract.WatchLogs(opts, "DebugUint256")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockProxyDebugUint256)
				if err := _LockProxy.contract.UnpackLog(event, "DebugUint256", log); err != nil {
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

// ParseDebugUint256 is a log parse operation binding the contract event 0x43d4b4706539f9e22baf8767ebea21ad24f723f14b6981664ac4d0af596dddbe.
//
// Solidity: event DebugUint256(uint256 amount)
func (_LockProxy *LockProxyFilterer) ParseDebugUint256(log types.Log) (*LockProxyDebugUint256, error) {
	event := new(LockProxyDebugUint256)
	if err := _LockProxy.contract.UnpackLog(event, "DebugUint256", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockProxyLockEventIterator is returned from FilterLockEvent and is used to iterate over the raw logs and unpacked data for LockEvent events raised by the LockProxy contract.
type LockProxyLockEventIterator struct {
	Event *LockProxyLockEvent // Event containing the contract specifics and raw log

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
func (it *LockProxyLockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockProxyLockEvent)
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
		it.Event = new(LockProxyLockEvent)
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
func (it *LockProxyLockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockProxyLockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockProxyLockEvent represents a LockEvent event raised by the LockProxy contract.
type LockProxyLockEvent struct {
	ThisContract common.Address
	ChainId      uint64
	ToContract   []byte
	TxArgs       []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLockEvent is a free log retrieval operation binding the contract event 0x28094cc2d1bbe0fe894550907e1f9c6b8c9bb18cd72c4830534347b5974645c0.
//
// Solidity: event LockEvent(address thisContract, uint64 chainId, bytes toContract, bytes txArgs)
func (_LockProxy *LockProxyFilterer) FilterLockEvent(opts *bind.FilterOpts) (*LockProxyLockEventIterator, error) {

	logs, sub, err := _LockProxy.contract.FilterLogs(opts, "LockEvent")
	if err != nil {
		return nil, err
	}
	return &LockProxyLockEventIterator{contract: _LockProxy.contract, event: "LockEvent", logs: logs, sub: sub}, nil
}

// WatchLockEvent is a free log subscription operation binding the contract event 0x28094cc2d1bbe0fe894550907e1f9c6b8c9bb18cd72c4830534347b5974645c0.
//
// Solidity: event LockEvent(address thisContract, uint64 chainId, bytes toContract, bytes txArgs)
func (_LockProxy *LockProxyFilterer) WatchLockEvent(opts *bind.WatchOpts, sink chan<- *LockProxyLockEvent) (event.Subscription, error) {

	logs, sub, err := _LockProxy.contract.WatchLogs(opts, "LockEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockProxyLockEvent)
				if err := _LockProxy.contract.UnpackLog(event, "LockEvent", log); err != nil {
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

// ParseLockEvent is a log parse operation binding the contract event 0x28094cc2d1bbe0fe894550907e1f9c6b8c9bb18cd72c4830534347b5974645c0.
//
// Solidity: event LockEvent(address thisContract, uint64 chainId, bytes toContract, bytes txArgs)
func (_LockProxy *LockProxyFilterer) ParseLockEvent(log types.Log) (*LockProxyLockEvent, error) {
	event := new(LockProxyLockEvent)
	if err := _LockProxy.contract.UnpackLog(event, "LockEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockProxySetManagerEventIterator is returned from FilterSetManagerEvent and is used to iterate over the raw logs and unpacked data for SetManagerEvent events raised by the LockProxy contract.
type LockProxySetManagerEventIterator struct {
	Event *LockProxySetManagerEvent // Event containing the contract specifics and raw log

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
func (it *LockProxySetManagerEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockProxySetManagerEvent)
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
		it.Event = new(LockProxySetManagerEvent)
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
func (it *LockProxySetManagerEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockProxySetManagerEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockProxySetManagerEvent represents a SetManagerEvent event raised by the LockProxy contract.
type LockProxySetManagerEvent struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetManagerEvent is a free log retrieval operation binding the contract event 0x45c21a4c9bc0568ee70fe6ae892fae1ce77b914d6b972041296072954ab21f8f.
//
// Solidity: event SetManagerEvent(address manager)
func (_LockProxy *LockProxyFilterer) FilterSetManagerEvent(opts *bind.FilterOpts) (*LockProxySetManagerEventIterator, error) {

	logs, sub, err := _LockProxy.contract.FilterLogs(opts, "SetManagerEvent")
	if err != nil {
		return nil, err
	}
	return &LockProxySetManagerEventIterator{contract: _LockProxy.contract, event: "SetManagerEvent", logs: logs, sub: sub}, nil
}

// WatchSetManagerEvent is a free log subscription operation binding the contract event 0x45c21a4c9bc0568ee70fe6ae892fae1ce77b914d6b972041296072954ab21f8f.
//
// Solidity: event SetManagerEvent(address manager)
func (_LockProxy *LockProxyFilterer) WatchSetManagerEvent(opts *bind.WatchOpts, sink chan<- *LockProxySetManagerEvent) (event.Subscription, error) {

	logs, sub, err := _LockProxy.contract.WatchLogs(opts, "SetManagerEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockProxySetManagerEvent)
				if err := _LockProxy.contract.UnpackLog(event, "SetManagerEvent", log); err != nil {
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

// ParseSetManagerEvent is a log parse operation binding the contract event 0x45c21a4c9bc0568ee70fe6ae892fae1ce77b914d6b972041296072954ab21f8f.
//
// Solidity: event SetManagerEvent(address manager)
func (_LockProxy *LockProxyFilterer) ParseSetManagerEvent(log types.Log) (*LockProxySetManagerEvent, error) {
	event := new(LockProxySetManagerEvent)
	if err := _LockProxy.contract.UnpackLog(event, "SetManagerEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockProxyUnlockEventIterator is returned from FilterUnlockEvent and is used to iterate over the raw logs and unpacked data for UnlockEvent events raised by the LockProxy contract.
type LockProxyUnlockEventIterator struct {
	Event *LockProxyUnlockEvent // Event containing the contract specifics and raw log

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
func (it *LockProxyUnlockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockProxyUnlockEvent)
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
		it.Event = new(LockProxyUnlockEvent)
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
func (it *LockProxyUnlockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockProxyUnlockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockProxyUnlockEvent represents a UnlockEvent event raised by the LockProxy contract.
type LockProxyUnlockEvent struct {
	FromContractAddr []byte
	FromChainId      uint64
	ToAddress        []byte
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUnlockEvent is a free log retrieval operation binding the contract event 0x31c3212616f0a6c018b96d403900949984f6cf1ac90e443ea802363277469352.
//
// Solidity: event UnlockEvent(bytes fromContractAddr, uint64 fromChainId, bytes toAddress, uint256 amount)
func (_LockProxy *LockProxyFilterer) FilterUnlockEvent(opts *bind.FilterOpts) (*LockProxyUnlockEventIterator, error) {

	logs, sub, err := _LockProxy.contract.FilterLogs(opts, "UnlockEvent")
	if err != nil {
		return nil, err
	}
	return &LockProxyUnlockEventIterator{contract: _LockProxy.contract, event: "UnlockEvent", logs: logs, sub: sub}, nil
}

// WatchUnlockEvent is a free log subscription operation binding the contract event 0x31c3212616f0a6c018b96d403900949984f6cf1ac90e443ea802363277469352.
//
// Solidity: event UnlockEvent(bytes fromContractAddr, uint64 fromChainId, bytes toAddress, uint256 amount)
func (_LockProxy *LockProxyFilterer) WatchUnlockEvent(opts *bind.WatchOpts, sink chan<- *LockProxyUnlockEvent) (event.Subscription, error) {

	logs, sub, err := _LockProxy.contract.WatchLogs(opts, "UnlockEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockProxyUnlockEvent)
				if err := _LockProxy.contract.UnpackLog(event, "UnlockEvent", log); err != nil {
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

// ParseUnlockEvent is a log parse operation binding the contract event 0x31c3212616f0a6c018b96d403900949984f6cf1ac90e443ea802363277469352.
//
// Solidity: event UnlockEvent(bytes fromContractAddr, uint64 fromChainId, bytes toAddress, uint256 amount)
func (_LockProxy *LockProxyFilterer) ParseUnlockEvent(log types.Log) (*LockProxyUnlockEvent, error) {
	event := new(LockProxyUnlockEvent)
	if err := _LockProxy.contract.UnpackLog(event, "UnlockEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820397f259a2608799bbc8f09ce4165ec2dcbd3b2a407e845d88eb0a306b3deac430029"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// UtilsABI is the input ABI used to generate the binding from.
const UtilsABI = "[]"

// UtilsBin is the compiled bytecode used for deploying new contracts.
var UtilsBin = "0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058207ddc58c7d02c94146520322140e560af58ecf00ffbdc6e168c1e94d4be8a4be90029"

// DeployUtils deploys a new Ethereum contract, binding an instance of Utils to it.
func DeployUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utils, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// Utils is an auto generated Go binding around an Ethereum contract.
type Utils struct {
	UtilsCaller     // Read-only binding to the contract
	UtilsTransactor // Write-only binding to the contract
	UtilsFilterer   // Log filterer for contract events
}

// UtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilsSession struct {
	Contract     *Utils            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilsCallerSession struct {
	Contract *UtilsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilsTransactorSession struct {
	Contract     *UtilsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilsRaw struct {
	Contract *Utils // Generic contract binding to access the raw methods on
}

// UtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilsCallerRaw struct {
	Contract *UtilsCaller // Generic read-only contract binding to access the raw methods on
}

// UtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilsTransactorRaw struct {
	Contract *UtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtils creates a new instance of Utils, bound to a specific deployed contract.
func NewUtils(address common.Address, backend bind.ContractBackend) (*Utils, error) {
	contract, err := bindUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// NewUtilsCaller creates a new read-only instance of Utils, bound to a specific deployed contract.
func NewUtilsCaller(address common.Address, caller bind.ContractCaller) (*UtilsCaller, error) {
	contract, err := bindUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsCaller{contract: contract}, nil
}

// NewUtilsTransactor creates a new write-only instance of Utils, bound to a specific deployed contract.
func NewUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilsTransactor, error) {
	contract, err := bindUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsTransactor{contract: contract}, nil
}

// NewUtilsFilterer creates a new log filterer instance of Utils, bound to a specific deployed contract.
func NewUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilsFilterer, error) {
	contract, err := bindUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilsFilterer{contract: contract}, nil
}

// bindUtils binds a generic wrapper to an already deployed contract.
func bindUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.UtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transact(opts, method, params...)
}

// ZeroCopySinkABI is the input ABI used to generate the binding from.
const ZeroCopySinkABI = "[]"

// ZeroCopySinkBin is the compiled bytecode used for deploying new contracts.
var ZeroCopySinkBin = "0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820d4740ef5014a59f4b4b8f71e9a2d4b039e59a70465c7bfaec7f80e648bb0c8390029"

// DeployZeroCopySink deploys a new Ethereum contract, binding an instance of ZeroCopySink to it.
func DeployZeroCopySink(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZeroCopySink, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySinkABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZeroCopySinkBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZeroCopySink{ZeroCopySinkCaller: ZeroCopySinkCaller{contract: contract}, ZeroCopySinkTransactor: ZeroCopySinkTransactor{contract: contract}, ZeroCopySinkFilterer: ZeroCopySinkFilterer{contract: contract}}, nil
}

// ZeroCopySink is an auto generated Go binding around an Ethereum contract.
type ZeroCopySink struct {
	ZeroCopySinkCaller     // Read-only binding to the contract
	ZeroCopySinkTransactor // Write-only binding to the contract
	ZeroCopySinkFilterer   // Log filterer for contract events
}

// ZeroCopySinkCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZeroCopySinkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySinkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZeroCopySinkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySinkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZeroCopySinkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySinkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZeroCopySinkSession struct {
	Contract     *ZeroCopySink     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroCopySinkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZeroCopySinkCallerSession struct {
	Contract *ZeroCopySinkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ZeroCopySinkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZeroCopySinkTransactorSession struct {
	Contract     *ZeroCopySinkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZeroCopySinkRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZeroCopySinkRaw struct {
	Contract *ZeroCopySink // Generic contract binding to access the raw methods on
}

// ZeroCopySinkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZeroCopySinkCallerRaw struct {
	Contract *ZeroCopySinkCaller // Generic read-only contract binding to access the raw methods on
}

// ZeroCopySinkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZeroCopySinkTransactorRaw struct {
	Contract *ZeroCopySinkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZeroCopySink creates a new instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySink(address common.Address, backend bind.ContractBackend) (*ZeroCopySink, error) {
	contract, err := bindZeroCopySink(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySink{ZeroCopySinkCaller: ZeroCopySinkCaller{contract: contract}, ZeroCopySinkTransactor: ZeroCopySinkTransactor{contract: contract}, ZeroCopySinkFilterer: ZeroCopySinkFilterer{contract: contract}}, nil
}

// NewZeroCopySinkCaller creates a new read-only instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySinkCaller(address common.Address, caller bind.ContractCaller) (*ZeroCopySinkCaller, error) {
	contract, err := bindZeroCopySink(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySinkCaller{contract: contract}, nil
}

// NewZeroCopySinkTransactor creates a new write-only instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySinkTransactor(address common.Address, transactor bind.ContractTransactor) (*ZeroCopySinkTransactor, error) {
	contract, err := bindZeroCopySink(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySinkTransactor{contract: contract}, nil
}

// NewZeroCopySinkFilterer creates a new log filterer instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySinkFilterer(address common.Address, filterer bind.ContractFilterer) (*ZeroCopySinkFilterer, error) {
	contract, err := bindZeroCopySink(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySinkFilterer{contract: contract}, nil
}

// bindZeroCopySink binds a generic wrapper to an already deployed contract.
func bindZeroCopySink(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySinkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySink *ZeroCopySinkRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySink.Contract.ZeroCopySinkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySink *ZeroCopySinkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.ZeroCopySinkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySink *ZeroCopySinkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.ZeroCopySinkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySink *ZeroCopySinkCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySink.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySink *ZeroCopySinkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySink *ZeroCopySinkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.contract.Transact(opts, method, params...)
}

// ZeroCopySourceABI is the input ABI used to generate the binding from.
const ZeroCopySourceABI = "[]"

// ZeroCopySourceBin is the compiled bytecode used for deploying new contracts.
var ZeroCopySourceBin = "0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a7230582075d21fa89ce37c909e46a7abae91e05837e2e613801886e78cf143c1649be7b10029"

// DeployZeroCopySource deploys a new Ethereum contract, binding an instance of ZeroCopySource to it.
func DeployZeroCopySource(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZeroCopySource, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySourceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZeroCopySourceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZeroCopySource{ZeroCopySourceCaller: ZeroCopySourceCaller{contract: contract}, ZeroCopySourceTransactor: ZeroCopySourceTransactor{contract: contract}, ZeroCopySourceFilterer: ZeroCopySourceFilterer{contract: contract}}, nil
}

// ZeroCopySource is an auto generated Go binding around an Ethereum contract.
type ZeroCopySource struct {
	ZeroCopySourceCaller     // Read-only binding to the contract
	ZeroCopySourceTransactor // Write-only binding to the contract
	ZeroCopySourceFilterer   // Log filterer for contract events
}

// ZeroCopySourceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZeroCopySourceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySourceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZeroCopySourceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySourceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZeroCopySourceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySourceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZeroCopySourceSession struct {
	Contract     *ZeroCopySource   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroCopySourceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZeroCopySourceCallerSession struct {
	Contract *ZeroCopySourceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ZeroCopySourceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZeroCopySourceTransactorSession struct {
	Contract     *ZeroCopySourceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ZeroCopySourceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZeroCopySourceRaw struct {
	Contract *ZeroCopySource // Generic contract binding to access the raw methods on
}

// ZeroCopySourceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZeroCopySourceCallerRaw struct {
	Contract *ZeroCopySourceCaller // Generic read-only contract binding to access the raw methods on
}

// ZeroCopySourceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZeroCopySourceTransactorRaw struct {
	Contract *ZeroCopySourceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZeroCopySource creates a new instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySource(address common.Address, backend bind.ContractBackend) (*ZeroCopySource, error) {
	contract, err := bindZeroCopySource(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySource{ZeroCopySourceCaller: ZeroCopySourceCaller{contract: contract}, ZeroCopySourceTransactor: ZeroCopySourceTransactor{contract: contract}, ZeroCopySourceFilterer: ZeroCopySourceFilterer{contract: contract}}, nil
}

// NewZeroCopySourceCaller creates a new read-only instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySourceCaller(address common.Address, caller bind.ContractCaller) (*ZeroCopySourceCaller, error) {
	contract, err := bindZeroCopySource(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySourceCaller{contract: contract}, nil
}

// NewZeroCopySourceTransactor creates a new write-only instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySourceTransactor(address common.Address, transactor bind.ContractTransactor) (*ZeroCopySourceTransactor, error) {
	contract, err := bindZeroCopySource(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySourceTransactor{contract: contract}, nil
}

// NewZeroCopySourceFilterer creates a new log filterer instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySourceFilterer(address common.Address, filterer bind.ContractFilterer) (*ZeroCopySourceFilterer, error) {
	contract, err := bindZeroCopySource(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySourceFilterer{contract: contract}, nil
}

// bindZeroCopySource binds a generic wrapper to an already deployed contract.
func bindZeroCopySource(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySourceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySource *ZeroCopySourceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySource.Contract.ZeroCopySourceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySource *ZeroCopySourceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.ZeroCopySourceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySource *ZeroCopySourceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.ZeroCopySourceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySource *ZeroCopySourceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySource.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySource *ZeroCopySourceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySource *ZeroCopySourceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.contract.Transact(opts, method, params...)
}
