// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"firstParam\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"secondParam\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"thirdParam\",\"type\":\"int256\"}],\"name\":\"ActorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"ActorDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"firstParam\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"secondParam\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"thirdParam\",\"type\":\"int256\"}],\"name\":\"addToActor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"delFromActor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200279b3803806200279b833981810160405281019062000037919062000386565b6040518060400160405280600981526020017f506f6a616c6565736800000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f504f4a0000000000000000000000000000000000000000000000000000000000815250620000c3620000b76200010260201b60201c565b6200010a60201b60201c565b8160049081620000d4919062000628565b508060059081620000e6919062000628565b505050620000fb3382620001ce60201b60201c565b506200082a565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160362000240576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002379062000770565b60405180910390fd5b62000254600083836200033c60201b60201c565b8060036000828254620002689190620007c1565b9250508190555080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516200031c91906200080d565b60405180910390a362000338600083836200034160201b60201c565b5050565b505050565b505050565b600080fd5b6000819050919050565b62000360816200034b565b81146200036c57600080fd5b50565b600081519050620003808162000355565b92915050565b6000602082840312156200039f576200039e62000346565b5b6000620003af848285016200036f565b91505092915050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200043a57607f821691505b60208210810362000450576200044f620003f2565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620004ba7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200047b565b620004c686836200047b565b95508019841693508086168417925050509392505050565b6000819050919050565b60006200050962000503620004fd846200034b565b620004de565b6200034b565b9050919050565b6000819050919050565b6200052583620004e8565b6200053d620005348262000510565b84845462000488565b825550505050565b600090565b6200055462000545565b620005618184846200051a565b505050565b5b8181101562000589576200057d6000826200054a565b60018101905062000567565b5050565b601f821115620005d857620005a28162000456565b620005ad846200046b565b81016020851015620005bd578190505b620005d5620005cc856200046b565b83018262000566565b50505b505050565b600082821c905092915050565b6000620005fd60001984600802620005dd565b1980831691505092915050565b6000620006188383620005ea565b9150826002028217905092915050565b6200063382620003b8565b67ffffffffffffffff8111156200064f576200064e620003c3565b5b6200065b825462000421565b620006688282856200058d565b600060209050601f831160018114620006a057600084156200068b578287015190505b6200069785826200060a565b86555062000707565b601f198416620006b08662000456565b60005b82811015620006da57848901518255600182019150602085019450602081019050620006b3565b86831015620006fa5784890151620006f6601f891682620005ea565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b600062000758601f836200070f565b9150620007658262000720565b602082019050919050565b600060208201905081810360008301526200078b8162000749565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000620007ce826200034b565b9150620007db836200034b565b9250828201905080821115620007f657620007f562000792565b5b92915050565b62000807816200034b565b82525050565b6000602082019050620008246000830184620007fc565b92915050565b611f61806200083a6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c8063715018a6116100a2578063a9059cbb11610071578063a9059cbb146102bc578063a9c7b0db146102ec578063ceadadd514610308578063dd62ed3e14610324578063f2fde38b146103545761010b565b8063715018a6146102465780638da5cb5b1461025057806395d89b411461026e578063a457c2d71461028c5761010b565b8063313ce567116100de578063313ce567146101ac57806339509351146101ca57806340c10f19146101fa57806370a08231146102165761010b565b806306fdde0314610110578063095ea7b31461012e57806318160ddd1461015e57806323b872dd1461017c575b600080fd5b610118610370565b6040516101259190611112565b60405180910390f35b610148600480360381019061014391906111dc565b610402565b6040516101559190611237565b60405180910390f35b610166610425565b6040516101739190611261565b60405180910390f35b6101966004803603810190610191919061127c565b61042f565b6040516101a39190611237565b60405180910390f35b6101b461045e565b6040516101c191906112eb565b60405180910390f35b6101e460048036038101906101df91906111dc565b610467565b6040516101f19190611237565b60405180910390f35b610214600480360381019061020f91906111dc565b61049e565b005b610230600480360381019061022b9190611306565b6104b4565b60405161023d9190611261565b60405180910390f35b61024e6104fd565b005b610258610511565b6040516102659190611342565b60405180910390f35b61027661053a565b6040516102839190611112565b60405180910390f35b6102a660048036038101906102a191906111dc565b6105cc565b6040516102b39190611237565b60405180910390f35b6102d660048036038101906102d191906111dc565b610643565b6040516102e39190611237565b60405180910390f35b61030660048036038101906103019190611492565b610666565b005b610322600480360381019061031d9190611511565b6106e9565b005b61033e60048036038101906103399190611594565b6107a2565b60405161034b9190611261565b60405180910390f35b61036e60048036038101906103699190611306565b610829565b005b60606004805461037f90611603565b80601f01602080910402602001604051908101604052809291908181526020018280546103ab90611603565b80156103f85780601f106103cd576101008083540402835291602001916103f8565b820191906000526020600020905b8154815290600101906020018083116103db57829003601f168201915b5050505050905090565b60008061040d6108ac565b905061041a8185856108b4565b600191505092915050565b6000600354905090565b60008061043a6108ac565b9050610447858285610a7d565b610452858585610b09565b60019150509392505050565b60006012905090565b6000806104726108ac565b905061049381858561048485896107a2565b61048e9190611663565b6108b4565b600191505092915050565b6104a6610d82565b6104b08282610e00565b5050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b610505610d82565b61050f6000610f57565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606005805461054990611603565b80601f016020809104026020016040519081016040528092919081815260200182805461057590611603565b80156105c25780601f10610597576101008083540402835291602001916105c2565b820191906000526020600020905b8154815290600101906020018083116105a557829003601f168201915b5050505050905090565b6000806105d76108ac565b905060006105e582866107a2565b90508381101561062a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062190611709565b60405180910390fd5b61063782868684036108b4565b60019250505092915050565b60008061064e6108ac565b905061065b818585610b09565b600191505092915050565b6006816040516106769190611765565b9081526020016040518091039020600080820160006106959190611025565b60018201600090556002820160009055600382016000905550507f439cc593a85f8085eb9c8d29c9d38661a8ba12d449d1a89f126048d6a5e013e8816040516106de9190611112565b60405180910390a150565b60405180608001604052808581526020018481526020018381526020018281525060068560405161071a9190611765565b9081526020016040518091039020600082015181600001908161073d9190611928565b506020820151816001015560408201518160020155606082015181600301559050507f308a8ca6a4a357459c647019da8dca3ee08dfd385e18976a66d56e2186bad423848484846040516107949493929190611a09565b60405180910390a150505050565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b610831610d82565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036108a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089790611ac7565b60405180910390fd5b6108a981610f57565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610923576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091a90611b59565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610992576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098990611beb565b60405180910390fd5b80600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610a709190611261565b60405180910390a3505050565b6000610a8984846107a2565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610b035781811015610af5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aec90611c57565b60405180910390fd5b610b0284848484036108b4565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610b78576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6f90611ce9565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610be7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bde90611d7b565b60405180910390fd5b610bf283838361101b565b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610c79576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7090611e0d565b60405180910390fd5b818103600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610d699190611261565b60405180910390a3610d7c848484611020565b50505050565b610d8a6108ac565b73ffffffffffffffffffffffffffffffffffffffff16610da8610511565b73ffffffffffffffffffffffffffffffffffffffff1614610dfe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610df590611e79565b60405180910390fd5b565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610e6f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e6690611ee5565b60405180910390fd5b610e7b6000838361101b565b8060036000828254610e8d9190611663565b9250508190555080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610f3f9190611261565b60405180910390a3610f5360008383611020565b5050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b505050565b505050565b50805461103190611603565b6000825580601f106110435750611062565b601f0160209004906000526020600020908101906110619190611065565b5b50565b5b8082111561107e576000816000905550600101611066565b5090565b600081519050919050565b600082825260208201905092915050565b60005b838110156110bc5780820151818401526020810190506110a1565b60008484015250505050565b6000601f19601f8301169050919050565b60006110e482611082565b6110ee818561108d565b93506110fe81856020860161109e565b611107816110c8565b840191505092915050565b6000602082019050818103600083015261112c81846110d9565b905092915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061117382611148565b9050919050565b61118381611168565b811461118e57600080fd5b50565b6000813590506111a08161117a565b92915050565b6000819050919050565b6111b9816111a6565b81146111c457600080fd5b50565b6000813590506111d6816111b0565b92915050565b600080604083850312156111f3576111f261113e565b5b600061120185828601611191565b9250506020611212858286016111c7565b9150509250929050565b60008115159050919050565b6112318161121c565b82525050565b600060208201905061124c6000830184611228565b92915050565b61125b816111a6565b82525050565b60006020820190506112766000830184611252565b92915050565b6000806000606084860312156112955761129461113e565b5b60006112a386828701611191565b93505060206112b486828701611191565b92505060406112c5868287016111c7565b9150509250925092565b600060ff82169050919050565b6112e5816112cf565b82525050565b600060208201905061130060008301846112dc565b92915050565b60006020828403121561131c5761131b61113e565b5b600061132a84828501611191565b91505092915050565b61133c81611168565b82525050565b60006020820190506113576000830184611333565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61139f826110c8565b810181811067ffffffffffffffff821117156113be576113bd611367565b5b80604052505050565b60006113d1611134565b90506113dd8282611396565b919050565b600067ffffffffffffffff8211156113fd576113fc611367565b5b611406826110c8565b9050602081019050919050565b82818337600083830152505050565b6000611435611430846113e2565b6113c7565b90508281526020810184848401111561145157611450611362565b5b61145c848285611413565b509392505050565b600082601f8301126114795761147861135d565b5b8135611489848260208601611422565b91505092915050565b6000602082840312156114a8576114a761113e565b5b600082013567ffffffffffffffff8111156114c6576114c5611143565b5b6114d284828501611464565b91505092915050565b6000819050919050565b6114ee816114db565b81146114f957600080fd5b50565b60008135905061150b816114e5565b92915050565b6000806000806080858703121561152b5761152a61113e565b5b600085013567ffffffffffffffff81111561154957611548611143565b5b61155587828801611464565b9450506020611566878288016114fc565b9350506040611577878288016114fc565b9250506060611588878288016114fc565b91505092959194509250565b600080604083850312156115ab576115aa61113e565b5b60006115b985828601611191565b92505060206115ca85828601611191565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061161b57607f821691505b60208210810361162e5761162d6115d4565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061166e826111a6565b9150611679836111a6565b925082820190508082111561169157611690611634565b5b92915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b60006116f360258361108d565b91506116fe82611697565b604082019050919050565b60006020820190508181036000830152611722816116e6565b9050919050565b600081905092915050565b600061173f82611082565b6117498185611729565b935061175981856020860161109e565b80840191505092915050565b60006117718284611734565b915081905092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026117de7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826117a1565b6117e886836117a1565b95508019841693508086168417925050509392505050565b6000819050919050565b600061182561182061181b846111a6565b611800565b6111a6565b9050919050565b6000819050919050565b61183f8361180a565b61185361184b8261182c565b8484546117ae565b825550505050565b600090565b61186861185b565b611873818484611836565b505050565b5b818110156118975761188c600082611860565b600181019050611879565b5050565b601f8211156118dc576118ad8161177c565b6118b684611791565b810160208510156118c5578190505b6118d96118d185611791565b830182611878565b50505b505050565b600082821c905092915050565b60006118ff600019846008026118e1565b1980831691505092915050565b600061191883836118ee565b9150826002028217905092915050565b61193182611082565b67ffffffffffffffff81111561194a57611949611367565b5b6119548254611603565b61195f82828561189b565b600060209050601f8311600181146119925760008415611980578287015190505b61198a858261190c565b8655506119f2565b601f1984166119a08661177c565b60005b828110156119c8578489015182556001820191506020850194506020810190506119a3565b868310156119e557848901516119e1601f8916826118ee565b8355505b6001600288020188555050505b505050505050565b611a03816114db565b82525050565b60006080820190508181036000830152611a2381876110d9565b9050611a3260208301866119fa565b611a3f60408301856119fa565b611a4c60608301846119fa565b95945050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611ab160268361108d565b9150611abc82611a55565b604082019050919050565b60006020820190508181036000830152611ae081611aa4565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000611b4360248361108d565b9150611b4e82611ae7565b604082019050919050565b60006020820190508181036000830152611b7281611b36565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b6000611bd560228361108d565b9150611be082611b79565b604082019050919050565b60006020820190508181036000830152611c0481611bc8565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b6000611c41601d8361108d565b9150611c4c82611c0b565b602082019050919050565b60006020820190508181036000830152611c7081611c34565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b6000611cd360258361108d565b9150611cde82611c77565b604082019050919050565b60006020820190508181036000830152611d0281611cc6565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000611d6560238361108d565b9150611d7082611d09565b604082019050919050565b60006020820190508181036000830152611d9481611d58565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b6000611df760268361108d565b9150611e0282611d9b565b604082019050919050565b60006020820190508181036000830152611e2681611dea565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611e6360208361108d565b9150611e6e82611e2d565b602082019050919050565b60006020820190508181036000830152611e9281611e56565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000611ecf601f8361108d565b9150611eda82611e99565b602082019050919050565b60006020820190508181036000830152611efe81611ec2565b905091905056fea2646970667358221220a67ec90e5ffb0fd75b2c60825cd3ee3d1d57937392fec0d428bb2bed662742e564736f6c637828302e382e32302d646576656c6f702e323032332e322e32352b636f6d6d69742e39383334303737360059",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// MainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MainMetaData.Bin instead.
var MainBin = MainMetaData.Bin

// DeployMain deploys a new Ethereum contract, binding an instance of Main to it.
func DeployMain(auth *bind.TransactOpts, backend bind.ContractBackend, amount *big.Int) (common.Address, *types.Transaction, *Main, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MainBin), backend, amount)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Main *MainCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Main *MainSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Main.Contract.Allowance(&_Main.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Main *MainCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Main.Contract.Allowance(&_Main.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Main *MainCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Main *MainSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Main.Contract.BalanceOf(&_Main.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Main *MainCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Main.Contract.BalanceOf(&_Main.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Main *MainCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Main *MainSession) Decimals() (uint8, error) {
	return _Main.Contract.Decimals(&_Main.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Main *MainCallerSession) Decimals() (uint8, error) {
	return _Main.Contract.Decimals(&_Main.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Main *MainCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Main *MainSession) Name() (string, error) {
	return _Main.Contract.Name(&_Main.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Main *MainCallerSession) Name() (string, error) {
	return _Main.Contract.Name(&_Main.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainCallerSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Main *MainCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Main *MainSession) Symbol() (string, error) {
	return _Main.Contract.Symbol(&_Main.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Main *MainCallerSession) Symbol() (string, error) {
	return _Main.Contract.Symbol(&_Main.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Main *MainCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Main *MainSession) TotalSupply() (*big.Int, error) {
	return _Main.Contract.TotalSupply(&_Main.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Main *MainCallerSession) TotalSupply() (*big.Int, error) {
	return _Main.Contract.TotalSupply(&_Main.CallOpts)
}

// AddToActor is a paid mutator transaction binding the contract method 0xceadadd5.
//
// Solidity: function addToActor(string id, int256 firstParam, int256 secondParam, int256 thirdParam) returns()
func (_Main *MainTransactor) AddToActor(opts *bind.TransactOpts, id string, firstParam *big.Int, secondParam *big.Int, thirdParam *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addToActor", id, firstParam, secondParam, thirdParam)
}

// AddToActor is a paid mutator transaction binding the contract method 0xceadadd5.
//
// Solidity: function addToActor(string id, int256 firstParam, int256 secondParam, int256 thirdParam) returns()
func (_Main *MainSession) AddToActor(id string, firstParam *big.Int, secondParam *big.Int, thirdParam *big.Int) (*types.Transaction, error) {
	return _Main.Contract.AddToActor(&_Main.TransactOpts, id, firstParam, secondParam, thirdParam)
}

// AddToActor is a paid mutator transaction binding the contract method 0xceadadd5.
//
// Solidity: function addToActor(string id, int256 firstParam, int256 secondParam, int256 thirdParam) returns()
func (_Main *MainTransactorSession) AddToActor(id string, firstParam *big.Int, secondParam *big.Int, thirdParam *big.Int) (*types.Transaction, error) {
	return _Main.Contract.AddToActor(&_Main.TransactOpts, id, firstParam, secondParam, thirdParam)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Main *MainTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Main *MainSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Approve(&_Main.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Main *MainTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Approve(&_Main.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Main *MainTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Main *MainSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Main.Contract.DecreaseAllowance(&_Main.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Main *MainTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Main.Contract.DecreaseAllowance(&_Main.TransactOpts, spender, subtractedValue)
}

// DelFromActor is a paid mutator transaction binding the contract method 0xa9c7b0db.
//
// Solidity: function delFromActor(string id) returns()
func (_Main *MainTransactor) DelFromActor(opts *bind.TransactOpts, id string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "delFromActor", id)
}

// DelFromActor is a paid mutator transaction binding the contract method 0xa9c7b0db.
//
// Solidity: function delFromActor(string id) returns()
func (_Main *MainSession) DelFromActor(id string) (*types.Transaction, error) {
	return _Main.Contract.DelFromActor(&_Main.TransactOpts, id)
}

// DelFromActor is a paid mutator transaction binding the contract method 0xa9c7b0db.
//
// Solidity: function delFromActor(string id) returns()
func (_Main *MainTransactorSession) DelFromActor(id string) (*types.Transaction, error) {
	return _Main.Contract.DelFromActor(&_Main.TransactOpts, id)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Main *MainTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Main *MainSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Main.Contract.IncreaseAllowance(&_Main.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Main *MainTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Main.Contract.IncreaseAllowance(&_Main.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Main *MainTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Main *MainSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Mint(&_Main.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Main *MainTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Mint(&_Main.TransactOpts, account, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Main *MainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Main *MainSession) RenounceOwnership() (*types.Transaction, error) {
	return _Main.Contract.RenounceOwnership(&_Main.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Main *MainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Main.Contract.RenounceOwnership(&_Main.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Main *MainTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Main *MainSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Transfer(&_Main.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Main *MainTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Transfer(&_Main.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Main *MainTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Main *MainSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TransferFrom(&_Main.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Main *MainTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TransferFrom(&_Main.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Main *MainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Main *MainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Main.Contract.TransferOwnership(&_Main.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Main *MainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Main.Contract.TransferOwnership(&_Main.TransactOpts, newOwner)
}

// MainActorAddedIterator is returned from FilterActorAdded and is used to iterate over the raw logs and unpacked data for ActorAdded events raised by the Main contract.
type MainActorAddedIterator struct {
	Event *MainActorAdded // Event containing the contract specifics and raw log

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
func (it *MainActorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainActorAdded)
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
		it.Event = new(MainActorAdded)
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
func (it *MainActorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainActorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainActorAdded represents a ActorAdded event raised by the Main contract.
type MainActorAdded struct {
	Id          string
	FirstParam  *big.Int
	SecondParam *big.Int
	ThirdParam  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterActorAdded is a free log retrieval operation binding the contract event 0x308a8ca6a4a357459c647019da8dca3ee08dfd385e18976a66d56e2186bad423.
//
// Solidity: event ActorAdded(string id, int256 firstParam, int256 secondParam, int256 thirdParam)
func (_Main *MainFilterer) FilterActorAdded(opts *bind.FilterOpts) (*MainActorAddedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "ActorAdded")
	if err != nil {
		return nil, err
	}
	return &MainActorAddedIterator{contract: _Main.contract, event: "ActorAdded", logs: logs, sub: sub}, nil
}

// WatchActorAdded is a free log subscription operation binding the contract event 0x308a8ca6a4a357459c647019da8dca3ee08dfd385e18976a66d56e2186bad423.
//
// Solidity: event ActorAdded(string id, int256 firstParam, int256 secondParam, int256 thirdParam)
func (_Main *MainFilterer) WatchActorAdded(opts *bind.WatchOpts, sink chan<- *MainActorAdded) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "ActorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainActorAdded)
				if err := _Main.contract.UnpackLog(event, "ActorAdded", log); err != nil {
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

// ParseActorAdded is a log parse operation binding the contract event 0x308a8ca6a4a357459c647019da8dca3ee08dfd385e18976a66d56e2186bad423.
//
// Solidity: event ActorAdded(string id, int256 firstParam, int256 secondParam, int256 thirdParam)
func (_Main *MainFilterer) ParseActorAdded(log types.Log) (*MainActorAdded, error) {
	event := new(MainActorAdded)
	if err := _Main.contract.UnpackLog(event, "ActorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainActorDeletedIterator is returned from FilterActorDeleted and is used to iterate over the raw logs and unpacked data for ActorDeleted events raised by the Main contract.
type MainActorDeletedIterator struct {
	Event *MainActorDeleted // Event containing the contract specifics and raw log

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
func (it *MainActorDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainActorDeleted)
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
		it.Event = new(MainActorDeleted)
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
func (it *MainActorDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainActorDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainActorDeleted represents a ActorDeleted event raised by the Main contract.
type MainActorDeleted struct {
	Id  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterActorDeleted is a free log retrieval operation binding the contract event 0x439cc593a85f8085eb9c8d29c9d38661a8ba12d449d1a89f126048d6a5e013e8.
//
// Solidity: event ActorDeleted(string id)
func (_Main *MainFilterer) FilterActorDeleted(opts *bind.FilterOpts) (*MainActorDeletedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "ActorDeleted")
	if err != nil {
		return nil, err
	}
	return &MainActorDeletedIterator{contract: _Main.contract, event: "ActorDeleted", logs: logs, sub: sub}, nil
}

// WatchActorDeleted is a free log subscription operation binding the contract event 0x439cc593a85f8085eb9c8d29c9d38661a8ba12d449d1a89f126048d6a5e013e8.
//
// Solidity: event ActorDeleted(string id)
func (_Main *MainFilterer) WatchActorDeleted(opts *bind.WatchOpts, sink chan<- *MainActorDeleted) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "ActorDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainActorDeleted)
				if err := _Main.contract.UnpackLog(event, "ActorDeleted", log); err != nil {
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

// ParseActorDeleted is a log parse operation binding the contract event 0x439cc593a85f8085eb9c8d29c9d38661a8ba12d449d1a89f126048d6a5e013e8.
//
// Solidity: event ActorDeleted(string id)
func (_Main *MainFilterer) ParseActorDeleted(log types.Log) (*MainActorDeleted, error) {
	event := new(MainActorDeleted)
	if err := _Main.contract.UnpackLog(event, "ActorDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Main contract.
type MainApprovalIterator struct {
	Event *MainApproval // Event containing the contract specifics and raw log

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
func (it *MainApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainApproval)
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
		it.Event = new(MainApproval)
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
func (it *MainApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainApproval represents a Approval event raised by the Main contract.
type MainApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Main *MainFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MainApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MainApprovalIterator{contract: _Main.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Main *MainFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MainApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainApproval)
				if err := _Main.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Main *MainFilterer) ParseApproval(log types.Log) (*MainApproval, error) {
	event := new(MainApproval)
	if err := _Main.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Main contract.
type MainOwnershipTransferredIterator struct {
	Event *MainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainOwnershipTransferred)
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
		it.Event = new(MainOwnershipTransferred)
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
func (it *MainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainOwnershipTransferred represents a OwnershipTransferred event raised by the Main contract.
type MainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Main *MainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MainOwnershipTransferredIterator{contract: _Main.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Main *MainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainOwnershipTransferred)
				if err := _Main.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Main *MainFilterer) ParseOwnershipTransferred(log types.Log) (*MainOwnershipTransferred, error) {
	event := new(MainOwnershipTransferred)
	if err := _Main.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Main contract.
type MainTransferIterator struct {
	Event *MainTransfer // Event containing the contract specifics and raw log

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
func (it *MainTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainTransfer)
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
		it.Event = new(MainTransfer)
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
func (it *MainTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainTransfer represents a Transfer event raised by the Main contract.
type MainTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Main *MainFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MainTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MainTransferIterator{contract: _Main.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Main *MainFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MainTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainTransfer)
				if err := _Main.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Main *MainFilterer) ParseTransfer(log types.Log) (*MainTransfer, error) {
	event := new(MainTransfer)
	if err := _Main.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
