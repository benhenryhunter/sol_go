// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bootoken

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
)

// BootokenMetaData contains all meta data concerning the Bootoken contract.
var BootokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethereals\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASE_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHEREALS\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PER_TX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAllOwned\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getClaimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getOwnedByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ownedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoveryTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"name\":\"setAllowed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"claimTokens\",\"type\":\"bool\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620021e9380380620021e98339810160408190526200003491620001dc565b604051806040016040528060088152602001672137b7aa37b5b2b760c11b81525060405180604001604052806003815260200162424f4f60e81b81525081600390805190602001906200008992919062000136565b5080516200009f90600490602084019062000136565b505050620000bc620000b6620000e060201b60201c565b620000e4565b6005805460ff60a01b1916905560016006556001600160a01b03166080526200024b565b3390565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b82805462000144906200020e565b90600052602060002090601f016020900481019282620001685760008555620001b3565b82601f106200018357805160ff1916838001178555620001b3565b82800160010185558215620001b3579182015b82811115620001b357825182559160200191906001019062000196565b50620001c1929150620001c5565b5090565b5b80821115620001c15760008155600101620001c6565b600060208284031215620001ef57600080fd5b81516001600160a01b03811681146200020757600080fd5b9392505050565b600181811c908216806200022357607f821691505b602082108114156200024557634e487b7160e01b600052602260045260246000fd5b50919050565b608051611f5f6200028a600039600081816103480152818161062d015281816107650152818161091c01528181610da80152610fcd0152611f5f6000f3fe608060405234801561001057600080fd5b50600436106102065760003560e01c806370a082311161011a578063a5b39cfb116100ad578063dd62ed3e1161007c578063dd62ed3e146104c0578063e149f036146104f9578063f2fde38b14610524578063f43a22dc14610537578063f8a14f461461053f57600080fd5b8063a5b39cfb1461044a578063a9059cbb1461046a578063cb03fb1e1461047d578063d63a8e111461049d57600080fd5b806395d89b41116100e957806395d89b41146104095780639dc29fac14610411578063a457c2d714610424578063a583024b1461043757600080fd5b806370a08231146103a7578063715018a6146103d05780637717228a146103d85780638da5cb5b146103f857600080fd5b80632da74ccc1161019d57806341910f901161016c57806341910f90146103215780634697f05d14610330578063552f39401461034357806357f576eb146103825780635c975abb1461039557600080fd5b80632da74ccc146102d9578063313ce567146102ec5780633228337a146102fb578063395093511461030e57600080fd5b806318160ddd116101d957806318160ddd1461028b5780631d0504a8146102935780631e83409a146102b357806323b872dd146102c657600080fd5b80630583e9f81461020b57806306fdde031461023e578063095ea7b3146102535780630fbf0a9314610276575b600080fd5b61022b610219366004611abe565b600c6020526000908152604090205481565b6040519081526020015b60405180910390f35b610246610568565b6040516102359190611ad7565b610266610261366004611b41565b6105fa565b6040519015158152602001610235565b610289610284366004611bb9565b610612565b005b60025461022b565b61022b6102a1366004611bfb565b60076020526000908152604090205481565b6102896102c1366004611bfb565b6109dd565b6102666102d4366004611c1f565b610a46565b61022b6102e7366004611b41565b610a6a565b60405160128152602001610235565b610289610309366004611c6e565b610aee565b61026661031c366004611b41565b610e6f565b61022b678ac7230489e8000081565b61028961033e366004611cc5565b610eae565b61036a7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610235565b610289610390366004611b41565b610f03565b600554600160a01b900460ff16610266565b61022b6103b5366004611bfb565b6001600160a01b031660009081526020819052604090205490565b61028961102d565b6103eb6103e6366004611bfb565b611063565b6040516102359190611cfe565b6005546001600160a01b031661036a565b61024661112c565b61028961041f366004611b41565b61113b565b610266610432366004611b41565b61119d565b61022b610445366004611bfb565b61122f565b61022b610458366004611bfb565b600b6020526000908152604090205481565b610266610478366004611b41565b611263565b61022b61048b366004611bfb565b60086020526000908152604090205481565b6102666104ab366004611bfb565b600d6020526000908152604090205460ff1681565b61022b6104ce366004611d42565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b61022b610507366004611b41565b600a60209081526000928352604080842090915290825290205481565b610289610532366004611bfb565b611271565b61022b603281565b61036a61054d366004611abe565b6009602052600090815260409020546001600160a01b031681565b60606003805461057790611d70565b80601f01602080910402602001604051908101604052809291908181526020018280546105a390611d70565b80156105f05780601f106105c5576101008083540402835291602001916105f0565b820191906000526020600020905b8154815290600101906020018083116105d357829003601f168201915b5050505050905090565b60003361060881858561130c565b5060019392505050565b60405163e985e9c560e01b81523360048201523060248201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063e985e9c590604401602060405180830381865afa15801561067c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a09190611dab565b6106e95760405162461bcd60e51b815260206004820152601560248201527410dbdb9d1c9858dd081b9bdd08185c1c1c9bdd9959605a1b60448201526064015b60405180910390fd5b600554600160a01b900460ff16156107365760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016106e0565b60328111156107575760405162461bcd60e51b81526004016106e090611dc8565b60005b8181101561086e57337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636352211e8585858181106107a4576107a4611e0a565b905060200201356040518263ffffffff1660e01b81526004016107c991815260200190565b602060405180830381865afa1580156107e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061080a9190611e20565b6001600160a01b03161461085c5760405162461bcd60e51b815260206004820152601960248201527821b0b63632b91034b9903737ba103a37b5b2b71037bbb732b960391b60448201526064016106e0565b8061086681611e53565b91505061075a565b5061087833611430565b60005b818110156109d857600083838381811061089757610897611e0a565b60209081029290920135600081815260098452604080822080546001600160a01b03191633908117909155808352600b8087528284208054868652600c8952848620819055838652600a89528486209086528852928420859055908352909452835491945060019392509061090d908490611e6e565b90915550506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166323b872dd333087878781811061095557610955611e0a565b6040516001600160e01b031960e088901b1681526001600160a01b03958616600482015294909316602485015250602090910201356044820152606401600060405180830381600087803b1580156109ac57600080fd5b505af11580156109c0573d6000803e3d6000fd5b505050505080806109d090611e53565b91505061087b565b505050565b60026006541415610a305760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016106e0565b6002600655610a3e81611482565b506001600655565b600033610a54858285611577565b610a5f858585611609565b506001949350505050565b6001600160a01b0382166000908152600b60205260408120548210610ac55760405162461bcd60e51b81526020600482015260116024820152702737b732bc34b9ba32b73a103a37b5b2b760791b60448201526064016106e0565b506001600160a01b03919091166000908152600a60209081526040808320938352929052205490565b60026006541415610b415760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016106e0565b60026006556032821115610b675760405162461bcd60e51b81526004016106e090611dc8565b60005b82811015610c07573360096000868685818110610b8957610b89611e0a565b60209081029290920135835250810191909152604001600020546001600160a01b031614610bf55760405162461bcd60e51b815260206004820152601960248201527821b0b63632b91034b9903737ba103a37b5b2b71037bbb732b960391b60448201526064016106e0565b80610bff81611e53565b915050610b6a565b508015610c1c57610c1733611482565b610c25565b610c2533611430565b60005b82811015610e6457336000908152600a60209081526040808320600b9092528220548290610c5890600190611e86565b8152602001908152602001600020549050600060096000878786818110610c8157610c81611e0a565b90506020020135815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b03160217905550600c6000868685818110610cd057610cd0611e0a565b90506020020135815260200190815260200160002054600c60008381526020019081526020016000208190555080600a6000336001600160a01b03166001600160a01b031681526020019081526020016000206000600c6000898988818110610d3b57610d3b611e0a565b905060200201358152602001908152602001600020548152602001908152602001600020819055506001600b6000336001600160a01b03166001600160a01b031681526020019081526020016000206000828254610d999190611e86565b90915550506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166323b872dd3033888887818110610de157610de1611e0a565b6040516001600160e01b031960e088901b1681526001600160a01b03958616600482015294909316602485015250602090910201356044820152606401600060405180830381600087803b158015610e3857600080fd5b505af1158015610e4c573d6000803e3d6000fd5b50505050508080610e5c90611e53565b915050610c28565b505060016006555050565b3360008181526001602090815260408083206001600160a01b03871684529091528120549091906106089082908690610ea9908790611e6e565b61130c565b6005546001600160a01b03163314610ed85760405162461bcd60e51b81526004016106e090611e9d565b6001600160a01b03919091166000908152600d60205260409020805460ff1916911515919091179055565b6005546001600160a01b03163314610f2d5760405162461bcd60e51b81526004016106e090611e9d565b6000818152600960205260409020546001600160a01b031615610fa15760405162461bcd60e51b815260206004820152602660248201527f546f6b656e20776173206e6f74207472616e73666572726564206163636964656044820152656e74616c6c7960d01b60648201526084016106e0565b6040516323b872dd60e01b81523060048201526001600160a01b038381166024830152604482018390527f000000000000000000000000000000000000000000000000000000000000000016906323b872dd90606401600060405180830381600087803b15801561101157600080fd5b505af1158015611025573d6000803e3d6000fd5b505050505050565b6005546001600160a01b031633146110575760405162461bcd60e51b81526004016106e090611e9d565b61106160006117d7565b565b6001600160a01b0381166000908152600b60205260408120546060919067ffffffffffffffff81111561109857611098611ed2565b6040519080825280602002602001820160405280156110c1578160200160208202803683370190505b50905060005b8151811015611125576001600160a01b0384166000908152600a60209081526040808320848452909152902054825183908390811061110857611108611e0a565b60209081029190910101528061111d81611e53565b9150506110c7565b5092915050565b60606004805461057790611d70565b336000908152600d602052604090205460ff1661118f5760405162461bcd60e51b815260206004820152601260248201527110d85b1b195c881b9bdd08185b1b1bddd95960721b60448201526064016106e0565b6111998282611829565b5050565b3360008181526001602090815260408083206001600160a01b0387168452909152812054909190838110156112225760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016106e0565b610a5f828686840361130c565b600061123a82611977565b6001600160a01b03831660009081526007602052604090205461125d9190611e6e565b92915050565b600033610608818585611609565b6005546001600160a01b0316331461129b5760405162461bcd60e51b81526004016106e090611e9d565b6001600160a01b0381166113005760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016106e0565b611309816117d7565b50565b6001600160a01b03831661136e5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016106e0565b6001600160a01b0382166113cf5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016106e0565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b61143981611977565b6001600160a01b03821660009081526007602052604081208054909190611461908490611e6e565b90915550506001600160a01b03166000908152600860205260409020429055565b600554600160a01b900460ff16156114cf5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016106e0565b336001600160a01b03821614806114f55750336000908152600d602052604090205460ff165b6115365760405162461bcd60e51b815260206004820152601260248201527110d85b1b195c881b9bdd08185b1b1bddd95960721b60448201526064016106e0565b60006115418261122f565b905061154d82826119df565b506001600160a01b0316600090815260076020908152604080832083905560089091529020429055565b6001600160a01b03838116600090815260016020908152604080832093861683529290522054600019811461160357818110156115f65760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016106e0565b611603848484840361130c565b50505050565b6001600160a01b03831661166d5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016106e0565b6001600160a01b0382166116cf5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016106e0565b6001600160a01b038316600090815260208190526040902054818110156117475760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016106e0565b6001600160a01b0380851660009081526020819052604080822085850390559185168152908120805484929061177e908490611e6e565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516117ca91815260200190565b60405180910390a3611603565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b0382166118895760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b60648201526084016106e0565b6001600160a01b038216600090815260208190526040902054818110156118fd5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b60648201526084016106e0565b6001600160a01b038316600090815260208190526040812083830390556002805484929061192c908490611e86565b90915550506040518281526000906001600160a01b038516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a3505050565b6001600160a01b038116600090815260086020526040812054620151809061199f9042611e86565b6001600160a01b0384166000908152600b60205260409020546119cb90678ac7230489e8000090611ee8565b6119d59190611ee8565b61125d9190611f07565b6001600160a01b038216611a355760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016106e0565b8060026000828254611a479190611e6e565b90915550506001600160a01b03821660009081526020819052604081208054839290611a74908490611e6e565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b600060208284031215611ad057600080fd5b5035919050565b600060208083528351808285015260005b81811015611b0457858101830151858201604001528201611ae8565b81811115611b16576000604083870101525b50601f01601f1916929092016040019392505050565b6001600160a01b038116811461130957600080fd5b60008060408385031215611b5457600080fd5b8235611b5f81611b2c565b946020939093013593505050565b60008083601f840112611b7f57600080fd5b50813567ffffffffffffffff811115611b9757600080fd5b6020830191508360208260051b8501011115611bb257600080fd5b9250929050565b60008060208385031215611bcc57600080fd5b823567ffffffffffffffff811115611be357600080fd5b611bef85828601611b6d565b90969095509350505050565b600060208284031215611c0d57600080fd5b8135611c1881611b2c565b9392505050565b600080600060608486031215611c3457600080fd5b8335611c3f81611b2c565b92506020840135611c4f81611b2c565b929592945050506040919091013590565b801515811461130957600080fd5b600080600060408486031215611c8357600080fd5b833567ffffffffffffffff811115611c9a57600080fd5b611ca686828701611b6d565b9094509250506020840135611cba81611c60565b809150509250925092565b60008060408385031215611cd857600080fd5b8235611ce381611b2c565b91506020830135611cf381611c60565b809150509250929050565b6020808252825182820181905260009190848201906040850190845b81811015611d3657835183529284019291840191600101611d1a565b50909695505050505050565b60008060408385031215611d5557600080fd5b8235611d6081611b2c565b91506020830135611cf381611b2c565b600181811c90821680611d8457607f821691505b60208210811415611da557634e487b7160e01b600052602260045260246000fd5b50919050565b600060208284031215611dbd57600080fd5b8151611c1881611c60565b60208082526022908201527f45786365656473206d617820746f6b656e7320706572207472616e736163746960408201526137b760f11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b600060208284031215611e3257600080fd5b8151611c1881611b2c565b634e487b7160e01b600052601160045260246000fd5b6000600019821415611e6757611e67611e3d565b5060010190565b60008219821115611e8157611e81611e3d565b500190565b600082821015611e9857611e98611e3d565b500390565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052604160045260246000fd5b6000816000190483118215151615611f0257611f02611e3d565b500290565b600082611f2457634e487b7160e01b600052601260045260246000fd5b50049056fea2646970667358221220ee4450fcbecfdf26579a3b6a9b7091b833e0ff5c75918bb6c452c8ac72112c2964736f6c634300080b0033",
}

// BootokenABI is the input ABI used to generate the binding from.
// Deprecated: Use BootokenMetaData.ABI instead.
var BootokenABI = BootokenMetaData.ABI

// BootokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BootokenMetaData.Bin instead.
var BootokenBin = BootokenMetaData.Bin

// DeployBootoken deploys a new Ethereum contract, binding an instance of Bootoken to it.
func DeployBootoken(auth *bind.TransactOpts, backend bind.ContractBackend, ethereals common.Address) (common.Address, *types.Transaction, *Bootoken, error) {
	parsed, err := BootokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BootokenBin), backend, ethereals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bootoken{BootokenCaller: BootokenCaller{contract: contract}, BootokenTransactor: BootokenTransactor{contract: contract}, BootokenFilterer: BootokenFilterer{contract: contract}}, nil
}

// Bootoken is an auto generated Go binding around an Ethereum contract.
type Bootoken struct {
	BootokenCaller     // Read-only binding to the contract
	BootokenTransactor // Write-only binding to the contract
	BootokenFilterer   // Log filterer for contract events
}

// BootokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BootokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BootokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BootokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BootokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BootokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BootokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BootokenSession struct {
	Contract     *Bootoken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BootokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BootokenCallerSession struct {
	Contract *BootokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BootokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BootokenTransactorSession struct {
	Contract     *BootokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BootokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BootokenRaw struct {
	Contract *Bootoken // Generic contract binding to access the raw methods on
}

// BootokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BootokenCallerRaw struct {
	Contract *BootokenCaller // Generic read-only contract binding to access the raw methods on
}

// BootokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BootokenTransactorRaw struct {
	Contract *BootokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBootoken creates a new instance of Bootoken, bound to a specific deployed contract.
func NewBootoken(address common.Address, backend bind.ContractBackend) (*Bootoken, error) {
	contract, err := bindBootoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bootoken{BootokenCaller: BootokenCaller{contract: contract}, BootokenTransactor: BootokenTransactor{contract: contract}, BootokenFilterer: BootokenFilterer{contract: contract}}, nil
}

// NewBootokenCaller creates a new read-only instance of Bootoken, bound to a specific deployed contract.
func NewBootokenCaller(address common.Address, caller bind.ContractCaller) (*BootokenCaller, error) {
	contract, err := bindBootoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BootokenCaller{contract: contract}, nil
}

// NewBootokenTransactor creates a new write-only instance of Bootoken, bound to a specific deployed contract.
func NewBootokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BootokenTransactor, error) {
	contract, err := bindBootoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BootokenTransactor{contract: contract}, nil
}

// NewBootokenFilterer creates a new log filterer instance of Bootoken, bound to a specific deployed contract.
func NewBootokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BootokenFilterer, error) {
	contract, err := bindBootoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BootokenFilterer{contract: contract}, nil
}

// bindBootoken binds a generic wrapper to an already deployed contract.
func bindBootoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BootokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bootoken *BootokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bootoken.Contract.BootokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bootoken *BootokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bootoken.Contract.BootokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bootoken *BootokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bootoken.Contract.BootokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bootoken *BootokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bootoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bootoken *BootokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bootoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bootoken *BootokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bootoken.Contract.contract.Transact(opts, method, params...)
}

// BASERATE is a free data retrieval call binding the contract method 0x41910f90.
//
// Solidity: function BASE_RATE() view returns(uint256)
func (_Bootoken *BootokenCaller) BASERATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "BASE_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASERATE is a free data retrieval call binding the contract method 0x41910f90.
//
// Solidity: function BASE_RATE() view returns(uint256)
func (_Bootoken *BootokenSession) BASERATE() (*big.Int, error) {
	return _Bootoken.Contract.BASERATE(&_Bootoken.CallOpts)
}

// BASERATE is a free data retrieval call binding the contract method 0x41910f90.
//
// Solidity: function BASE_RATE() view returns(uint256)
func (_Bootoken *BootokenCallerSession) BASERATE() (*big.Int, error) {
	return _Bootoken.Contract.BASERATE(&_Bootoken.CallOpts)
}

// ETHEREALS is a free data retrieval call binding the contract method 0x552f3940.
//
// Solidity: function ETHEREALS() view returns(address)
func (_Bootoken *BootokenCaller) ETHEREALS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "ETHEREALS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHEREALS is a free data retrieval call binding the contract method 0x552f3940.
//
// Solidity: function ETHEREALS() view returns(address)
func (_Bootoken *BootokenSession) ETHEREALS() (common.Address, error) {
	return _Bootoken.Contract.ETHEREALS(&_Bootoken.CallOpts)
}

// ETHEREALS is a free data retrieval call binding the contract method 0x552f3940.
//
// Solidity: function ETHEREALS() view returns(address)
func (_Bootoken *BootokenCallerSession) ETHEREALS() (common.Address, error) {
	return _Bootoken.Contract.ETHEREALS(&_Bootoken.CallOpts)
}

// MAXPERTX is a free data retrieval call binding the contract method 0xf43a22dc.
//
// Solidity: function MAX_PER_TX() view returns(uint256)
func (_Bootoken *BootokenCaller) MAXPERTX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "MAX_PER_TX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPERTX is a free data retrieval call binding the contract method 0xf43a22dc.
//
// Solidity: function MAX_PER_TX() view returns(uint256)
func (_Bootoken *BootokenSession) MAXPERTX() (*big.Int, error) {
	return _Bootoken.Contract.MAXPERTX(&_Bootoken.CallOpts)
}

// MAXPERTX is a free data retrieval call binding the contract method 0xf43a22dc.
//
// Solidity: function MAX_PER_TX() view returns(uint256)
func (_Bootoken *BootokenCallerSession) MAXPERTX() (*big.Int, error) {
	return _Bootoken.Contract.MAXPERTX(&_Bootoken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bootoken *BootokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bootoken *BootokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bootoken.Contract.Allowance(&_Bootoken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bootoken *BootokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bootoken.Contract.Allowance(&_Bootoken.CallOpts, owner, spender)
}

// Allowed is a free data retrieval call binding the contract method 0xd63a8e11.
//
// Solidity: function allowed(address ) view returns(bool)
func (_Bootoken *BootokenCaller) Allowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "allowed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Allowed is a free data retrieval call binding the contract method 0xd63a8e11.
//
// Solidity: function allowed(address ) view returns(bool)
func (_Bootoken *BootokenSession) Allowed(arg0 common.Address) (bool, error) {
	return _Bootoken.Contract.Allowed(&_Bootoken.CallOpts, arg0)
}

// Allowed is a free data retrieval call binding the contract method 0xd63a8e11.
//
// Solidity: function allowed(address ) view returns(bool)
func (_Bootoken *BootokenCallerSession) Allowed(arg0 common.Address) (bool, error) {
	return _Bootoken.Contract.Allowed(&_Bootoken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bootoken *BootokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bootoken *BootokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Bootoken.Contract.BalanceOf(&_Bootoken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bootoken *BootokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Bootoken.Contract.BalanceOf(&_Bootoken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bootoken *BootokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bootoken *BootokenSession) Decimals() (uint8, error) {
	return _Bootoken.Contract.Decimals(&_Bootoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bootoken *BootokenCallerSession) Decimals() (uint8, error) {
	return _Bootoken.Contract.Decimals(&_Bootoken.CallOpts)
}

// GetAllOwned is a free data retrieval call binding the contract method 0x7717228a.
//
// Solidity: function getAllOwned(address account) view returns(uint256[])
func (_Bootoken *BootokenCaller) GetAllOwned(opts *bind.CallOpts, account common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "getAllOwned", account)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllOwned is a free data retrieval call binding the contract method 0x7717228a.
//
// Solidity: function getAllOwned(address account) view returns(uint256[])
func (_Bootoken *BootokenSession) GetAllOwned(account common.Address) ([]*big.Int, error) {
	return _Bootoken.Contract.GetAllOwned(&_Bootoken.CallOpts, account)
}

// GetAllOwned is a free data retrieval call binding the contract method 0x7717228a.
//
// Solidity: function getAllOwned(address account) view returns(uint256[])
func (_Bootoken *BootokenCallerSession) GetAllOwned(account common.Address) ([]*big.Int, error) {
	return _Bootoken.Contract.GetAllOwned(&_Bootoken.CallOpts, account)
}

// GetClaimable is a free data retrieval call binding the contract method 0xa583024b.
//
// Solidity: function getClaimable(address account) view returns(uint256)
func (_Bootoken *BootokenCaller) GetClaimable(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "getClaimable", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimable is a free data retrieval call binding the contract method 0xa583024b.
//
// Solidity: function getClaimable(address account) view returns(uint256)
func (_Bootoken *BootokenSession) GetClaimable(account common.Address) (*big.Int, error) {
	return _Bootoken.Contract.GetClaimable(&_Bootoken.CallOpts, account)
}

// GetClaimable is a free data retrieval call binding the contract method 0xa583024b.
//
// Solidity: function getClaimable(address account) view returns(uint256)
func (_Bootoken *BootokenCallerSession) GetClaimable(account common.Address) (*big.Int, error) {
	return _Bootoken.Contract.GetClaimable(&_Bootoken.CallOpts, account)
}

// GetOwnedByIndex is a free data retrieval call binding the contract method 0x2da74ccc.
//
// Solidity: function getOwnedByIndex(address account, uint256 index) view returns(uint256)
func (_Bootoken *BootokenCaller) GetOwnedByIndex(opts *bind.CallOpts, account common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "getOwnedByIndex", account, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOwnedByIndex is a free data retrieval call binding the contract method 0x2da74ccc.
//
// Solidity: function getOwnedByIndex(address account, uint256 index) view returns(uint256)
func (_Bootoken *BootokenSession) GetOwnedByIndex(account common.Address, index *big.Int) (*big.Int, error) {
	return _Bootoken.Contract.GetOwnedByIndex(&_Bootoken.CallOpts, account, index)
}

// GetOwnedByIndex is a free data retrieval call binding the contract method 0x2da74ccc.
//
// Solidity: function getOwnedByIndex(address account, uint256 index) view returns(uint256)
func (_Bootoken *BootokenCallerSession) GetOwnedByIndex(account common.Address, index *big.Int) (*big.Int, error) {
	return _Bootoken.Contract.GetOwnedByIndex(&_Bootoken.CallOpts, account, index)
}

// LastUpdate is a free data retrieval call binding the contract method 0xcb03fb1e.
//
// Solidity: function lastUpdate(address ) view returns(uint256)
func (_Bootoken *BootokenCaller) LastUpdate(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "lastUpdate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdate is a free data retrieval call binding the contract method 0xcb03fb1e.
//
// Solidity: function lastUpdate(address ) view returns(uint256)
func (_Bootoken *BootokenSession) LastUpdate(arg0 common.Address) (*big.Int, error) {
	return _Bootoken.Contract.LastUpdate(&_Bootoken.CallOpts, arg0)
}

// LastUpdate is a free data retrieval call binding the contract method 0xcb03fb1e.
//
// Solidity: function lastUpdate(address ) view returns(uint256)
func (_Bootoken *BootokenCallerSession) LastUpdate(arg0 common.Address) (*big.Int, error) {
	return _Bootoken.Contract.LastUpdate(&_Bootoken.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bootoken *BootokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bootoken *BootokenSession) Name() (string, error) {
	return _Bootoken.Contract.Name(&_Bootoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bootoken *BootokenCallerSession) Name() (string, error) {
	return _Bootoken.Contract.Name(&_Bootoken.CallOpts)
}

// OwnedTokens is a free data retrieval call binding the contract method 0xe149f036.
//
// Solidity: function ownedTokens(address , uint256 ) view returns(uint256)
func (_Bootoken *BootokenCaller) OwnedTokens(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "ownedTokens", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnedTokens is a free data retrieval call binding the contract method 0xe149f036.
//
// Solidity: function ownedTokens(address , uint256 ) view returns(uint256)
func (_Bootoken *BootokenSession) OwnedTokens(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bootoken.Contract.OwnedTokens(&_Bootoken.CallOpts, arg0, arg1)
}

// OwnedTokens is a free data retrieval call binding the contract method 0xe149f036.
//
// Solidity: function ownedTokens(address , uint256 ) view returns(uint256)
func (_Bootoken *BootokenCallerSession) OwnedTokens(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bootoken.Contract.OwnedTokens(&_Bootoken.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bootoken *BootokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bootoken *BootokenSession) Owner() (common.Address, error) {
	return _Bootoken.Contract.Owner(&_Bootoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bootoken *BootokenCallerSession) Owner() (common.Address, error) {
	return _Bootoken.Contract.Owner(&_Bootoken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bootoken *BootokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bootoken *BootokenSession) Paused() (bool, error) {
	return _Bootoken.Contract.Paused(&_Bootoken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bootoken *BootokenCallerSession) Paused() (bool, error) {
	return _Bootoken.Contract.Paused(&_Bootoken.CallOpts)
}

// StakedTokens is a free data retrieval call binding the contract method 0xa5b39cfb.
//
// Solidity: function stakedTokens(address ) view returns(uint256)
func (_Bootoken *BootokenCaller) StakedTokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "stakedTokens", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedTokens is a free data retrieval call binding the contract method 0xa5b39cfb.
//
// Solidity: function stakedTokens(address ) view returns(uint256)
func (_Bootoken *BootokenSession) StakedTokens(arg0 common.Address) (*big.Int, error) {
	return _Bootoken.Contract.StakedTokens(&_Bootoken.CallOpts, arg0)
}

// StakedTokens is a free data retrieval call binding the contract method 0xa5b39cfb.
//
// Solidity: function stakedTokens(address ) view returns(uint256)
func (_Bootoken *BootokenCallerSession) StakedTokens(arg0 common.Address) (*big.Int, error) {
	return _Bootoken.Contract.StakedTokens(&_Bootoken.CallOpts, arg0)
}

// Stash is a free data retrieval call binding the contract method 0x1d0504a8.
//
// Solidity: function stash(address ) view returns(uint256)
func (_Bootoken *BootokenCaller) Stash(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "stash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Stash is a free data retrieval call binding the contract method 0x1d0504a8.
//
// Solidity: function stash(address ) view returns(uint256)
func (_Bootoken *BootokenSession) Stash(arg0 common.Address) (*big.Int, error) {
	return _Bootoken.Contract.Stash(&_Bootoken.CallOpts, arg0)
}

// Stash is a free data retrieval call binding the contract method 0x1d0504a8.
//
// Solidity: function stash(address ) view returns(uint256)
func (_Bootoken *BootokenCallerSession) Stash(arg0 common.Address) (*big.Int, error) {
	return _Bootoken.Contract.Stash(&_Bootoken.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bootoken *BootokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bootoken *BootokenSession) Symbol() (string, error) {
	return _Bootoken.Contract.Symbol(&_Bootoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bootoken *BootokenCallerSession) Symbol() (string, error) {
	return _Bootoken.Contract.Symbol(&_Bootoken.CallOpts)
}

// TokenIndex is a free data retrieval call binding the contract method 0x0583e9f8.
//
// Solidity: function tokenIndex(uint256 ) view returns(uint256)
func (_Bootoken *BootokenCaller) TokenIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "tokenIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIndex is a free data retrieval call binding the contract method 0x0583e9f8.
//
// Solidity: function tokenIndex(uint256 ) view returns(uint256)
func (_Bootoken *BootokenSession) TokenIndex(arg0 *big.Int) (*big.Int, error) {
	return _Bootoken.Contract.TokenIndex(&_Bootoken.CallOpts, arg0)
}

// TokenIndex is a free data retrieval call binding the contract method 0x0583e9f8.
//
// Solidity: function tokenIndex(uint256 ) view returns(uint256)
func (_Bootoken *BootokenCallerSession) TokenIndex(arg0 *big.Int) (*big.Int, error) {
	return _Bootoken.Contract.TokenIndex(&_Bootoken.CallOpts, arg0)
}

// TokenOwners is a free data retrieval call binding the contract method 0xf8a14f46.
//
// Solidity: function tokenOwners(uint256 ) view returns(address)
func (_Bootoken *BootokenCaller) TokenOwners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "tokenOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenOwners is a free data retrieval call binding the contract method 0xf8a14f46.
//
// Solidity: function tokenOwners(uint256 ) view returns(address)
func (_Bootoken *BootokenSession) TokenOwners(arg0 *big.Int) (common.Address, error) {
	return _Bootoken.Contract.TokenOwners(&_Bootoken.CallOpts, arg0)
}

// TokenOwners is a free data retrieval call binding the contract method 0xf8a14f46.
//
// Solidity: function tokenOwners(uint256 ) view returns(address)
func (_Bootoken *BootokenCallerSession) TokenOwners(arg0 *big.Int) (common.Address, error) {
	return _Bootoken.Contract.TokenOwners(&_Bootoken.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bootoken *BootokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bootoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bootoken *BootokenSession) TotalSupply() (*big.Int, error) {
	return _Bootoken.Contract.TotalSupply(&_Bootoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bootoken *BootokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Bootoken.Contract.TotalSupply(&_Bootoken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bootoken *BootokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootoken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bootoken *BootokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootoken.Contract.Approve(&_Bootoken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bootoken *BootokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootoken.Contract.Approve(&_Bootoken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_Bootoken *BootokenTransactor) Burn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootoken.contract.Transact(opts, "burn", from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_Bootoken *BootokenSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootoken.Contract.Burn(&_Bootoken.TransactOpts, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_Bootoken *BootokenTransactorSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootoken.Contract.Burn(&_Bootoken.TransactOpts, from, amount)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address account) returns()
func (_Bootoken *BootokenTransactor) Claim(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Bootoken.contract.Transact(opts, "claim", account)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address account) returns()
func (_Bootoken *BootokenSession) Claim(account common.Address) (*types.Transaction, error) {
	return _Bootoken.Contract.Claim(&_Bootoken.TransactOpts, account)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address account) returns()
func (_Bootoken *BootokenTransactorSession) Claim(account common.Address) (*types.Transaction, error) {
	return _Bootoken.Contract.Claim(&_Bootoken.TransactOpts, account)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bootoken *BootokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bootoken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bootoken *BootokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bootoken.Contract.DecreaseAllowance(&_Bootoken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bootoken *BootokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bootoken.Contract.DecreaseAllowance(&_Bootoken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bootoken *BootokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bootoken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bootoken *BootokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bootoken.Contract.IncreaseAllowance(&_Bootoken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bootoken *BootokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bootoken.Contract.IncreaseAllowance(&_Bootoken.TransactOpts, spender, addedValue)
}

// RecoveryTransfer is a paid mutator transaction binding the contract method 0x57f576eb.
//
// Solidity: function recoveryTransfer(address to, uint256 tokenId) returns()
func (_Bootoken *BootokenTransactor) RecoveryTransfer(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bootoken.contract.Transact(opts, "recoveryTransfer", to, tokenId)
}

// RecoveryTransfer is a paid mutator transaction binding the contract method 0x57f576eb.
//
// Solidity: function recoveryTransfer(address to, uint256 tokenId) returns()
func (_Bootoken *BootokenSession) RecoveryTransfer(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bootoken.Contract.RecoveryTransfer(&_Bootoken.TransactOpts, to, tokenId)
}

// RecoveryTransfer is a paid mutator transaction binding the contract method 0x57f576eb.
//
// Solidity: function recoveryTransfer(address to, uint256 tokenId) returns()
func (_Bootoken *BootokenTransactorSession) RecoveryTransfer(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bootoken.Contract.RecoveryTransfer(&_Bootoken.TransactOpts, to, tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bootoken *BootokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bootoken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bootoken *BootokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bootoken.Contract.RenounceOwnership(&_Bootoken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bootoken *BootokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bootoken.Contract.RenounceOwnership(&_Bootoken.TransactOpts)
}

// SetAllowed is a paid mutator transaction binding the contract method 0x4697f05d.
//
// Solidity: function setAllowed(address account, bool isAllowed) returns()
func (_Bootoken *BootokenTransactor) SetAllowed(opts *bind.TransactOpts, account common.Address, isAllowed bool) (*types.Transaction, error) {
	return _Bootoken.contract.Transact(opts, "setAllowed", account, isAllowed)
}

// SetAllowed is a paid mutator transaction binding the contract method 0x4697f05d.
//
// Solidity: function setAllowed(address account, bool isAllowed) returns()
func (_Bootoken *BootokenSession) SetAllowed(account common.Address, isAllowed bool) (*types.Transaction, error) {
	return _Bootoken.Contract.SetAllowed(&_Bootoken.TransactOpts, account, isAllowed)
}

// SetAllowed is a paid mutator transaction binding the contract method 0x4697f05d.
//
// Solidity: function setAllowed(address account, bool isAllowed) returns()
func (_Bootoken *BootokenTransactorSession) SetAllowed(account common.Address, isAllowed bool) (*types.Transaction, error) {
	return _Bootoken.Contract.SetAllowed(&_Bootoken.TransactOpts, account, isAllowed)
}

// Stake is a paid mutator transaction binding the contract method 0x0fbf0a93.
//
// Solidity: function stake(uint256[] tokenIds) returns()
func (_Bootoken *BootokenTransactor) Stake(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Bootoken.contract.Transact(opts, "stake", tokenIds)
}

// Stake is a paid mutator transaction binding the contract method 0x0fbf0a93.
//
// Solidity: function stake(uint256[] tokenIds) returns()
func (_Bootoken *BootokenSession) Stake(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Bootoken.Contract.Stake(&_Bootoken.TransactOpts, tokenIds)
}

// Stake is a paid mutator transaction binding the contract method 0x0fbf0a93.
//
// Solidity: function stake(uint256[] tokenIds) returns()
func (_Bootoken *BootokenTransactorSession) Stake(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Bootoken.Contract.Stake(&_Bootoken.TransactOpts, tokenIds)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Bootoken *BootokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootoken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Bootoken *BootokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootoken.Contract.Transfer(&_Bootoken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Bootoken *BootokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootoken.Contract.Transfer(&_Bootoken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Bootoken *BootokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootoken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Bootoken *BootokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootoken.Contract.TransferFrom(&_Bootoken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Bootoken *BootokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bootoken.Contract.TransferFrom(&_Bootoken.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bootoken *BootokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bootoken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bootoken *BootokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bootoken.Contract.TransferOwnership(&_Bootoken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bootoken *BootokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bootoken.Contract.TransferOwnership(&_Bootoken.TransactOpts, newOwner)
}

// Unstake is a paid mutator transaction binding the contract method 0x3228337a.
//
// Solidity: function unstake(uint256[] tokenIds, bool claimTokens) returns()
func (_Bootoken *BootokenTransactor) Unstake(opts *bind.TransactOpts, tokenIds []*big.Int, claimTokens bool) (*types.Transaction, error) {
	return _Bootoken.contract.Transact(opts, "unstake", tokenIds, claimTokens)
}

// Unstake is a paid mutator transaction binding the contract method 0x3228337a.
//
// Solidity: function unstake(uint256[] tokenIds, bool claimTokens) returns()
func (_Bootoken *BootokenSession) Unstake(tokenIds []*big.Int, claimTokens bool) (*types.Transaction, error) {
	return _Bootoken.Contract.Unstake(&_Bootoken.TransactOpts, tokenIds, claimTokens)
}

// Unstake is a paid mutator transaction binding the contract method 0x3228337a.
//
// Solidity: function unstake(uint256[] tokenIds, bool claimTokens) returns()
func (_Bootoken *BootokenTransactorSession) Unstake(tokenIds []*big.Int, claimTokens bool) (*types.Transaction, error) {
	return _Bootoken.Contract.Unstake(&_Bootoken.TransactOpts, tokenIds, claimTokens)
}

// BootokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Bootoken contract.
type BootokenApprovalIterator struct {
	Event *BootokenApproval // Event containing the contract specifics and raw log

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
func (it *BootokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootokenApproval)
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
		it.Event = new(BootokenApproval)
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
func (it *BootokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootokenApproval represents a Approval event raised by the Bootoken contract.
type BootokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bootoken *BootokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BootokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bootoken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BootokenApprovalIterator{contract: _Bootoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bootoken *BootokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BootokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bootoken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootokenApproval)
				if err := _Bootoken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bootoken *BootokenFilterer) ParseApproval(log types.Log) (*BootokenApproval, error) {
	event := new(BootokenApproval)
	if err := _Bootoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bootoken contract.
type BootokenOwnershipTransferredIterator struct {
	Event *BootokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BootokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootokenOwnershipTransferred)
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
		it.Event = new(BootokenOwnershipTransferred)
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
func (it *BootokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootokenOwnershipTransferred represents a OwnershipTransferred event raised by the Bootoken contract.
type BootokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bootoken *BootokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BootokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bootoken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BootokenOwnershipTransferredIterator{contract: _Bootoken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bootoken *BootokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BootokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bootoken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootokenOwnershipTransferred)
				if err := _Bootoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bootoken *BootokenFilterer) ParseOwnershipTransferred(log types.Log) (*BootokenOwnershipTransferred, error) {
	event := new(BootokenOwnershipTransferred)
	if err := _Bootoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootokenPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bootoken contract.
type BootokenPausedIterator struct {
	Event *BootokenPaused // Event containing the contract specifics and raw log

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
func (it *BootokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootokenPaused)
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
		it.Event = new(BootokenPaused)
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
func (it *BootokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootokenPaused represents a Paused event raised by the Bootoken contract.
type BootokenPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bootoken *BootokenFilterer) FilterPaused(opts *bind.FilterOpts) (*BootokenPausedIterator, error) {

	logs, sub, err := _Bootoken.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BootokenPausedIterator{contract: _Bootoken.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bootoken *BootokenFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BootokenPaused) (event.Subscription, error) {

	logs, sub, err := _Bootoken.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootokenPaused)
				if err := _Bootoken.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bootoken *BootokenFilterer) ParsePaused(log types.Log) (*BootokenPaused, error) {
	event := new(BootokenPaused)
	if err := _Bootoken.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bootoken contract.
type BootokenTransferIterator struct {
	Event *BootokenTransfer // Event containing the contract specifics and raw log

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
func (it *BootokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootokenTransfer)
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
		it.Event = new(BootokenTransfer)
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
func (it *BootokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootokenTransfer represents a Transfer event raised by the Bootoken contract.
type BootokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bootoken *BootokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BootokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bootoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BootokenTransferIterator{contract: _Bootoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bootoken *BootokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BootokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bootoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootokenTransfer)
				if err := _Bootoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bootoken *BootokenFilterer) ParseTransfer(log types.Log) (*BootokenTransfer, error) {
	event := new(BootokenTransfer)
	if err := _Bootoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootokenUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bootoken contract.
type BootokenUnpausedIterator struct {
	Event *BootokenUnpaused // Event containing the contract specifics and raw log

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
func (it *BootokenUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootokenUnpaused)
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
		it.Event = new(BootokenUnpaused)
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
func (it *BootokenUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootokenUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootokenUnpaused represents a Unpaused event raised by the Bootoken contract.
type BootokenUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bootoken *BootokenFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BootokenUnpausedIterator, error) {

	logs, sub, err := _Bootoken.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BootokenUnpausedIterator{contract: _Bootoken.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bootoken *BootokenFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BootokenUnpaused) (event.Subscription, error) {

	logs, sub, err := _Bootoken.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootokenUnpaused)
				if err := _Bootoken.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bootoken *BootokenFilterer) ParseUnpaused(log types.Log) (*BootokenUnpaused, error) {
	event := new(BootokenUnpaused)
	if err := _Bootoken.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
