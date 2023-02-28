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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"firstParam\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"secondParam\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"thirdParam\",\"type\":\"bool\"}],\"name\":\"ActorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"ActorDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"firstParam\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"secondParam\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"thirdParam\",\"type\":\"bool\"}],\"name\":\"addActor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"delActor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620027eb380380620027eb833981810160405281019062000037919062000386565b6040518060400160405280600981526020017f506f6a616c6565736800000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f504f4a0000000000000000000000000000000000000000000000000000000000815250620000c3620000b76200010260201b60201c565b6200010a60201b60201c565b8160049081620000d4919062000628565b508060059081620000e6919062000628565b505050620000fb3382620001ce60201b60201c565b506200082a565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160362000240576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002379062000770565b60405180910390fd5b62000254600083836200033c60201b60201c565b8060036000828254620002689190620007c1565b9250508190555080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516200031c91906200080d565b60405180910390a362000338600083836200034160201b60201c565b5050565b505050565b505050565b600080fd5b6000819050919050565b62000360816200034b565b81146200036c57600080fd5b50565b600081519050620003808162000355565b92915050565b6000602082840312156200039f576200039e62000346565b5b6000620003af848285016200036f565b91505092915050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200043a57607f821691505b60208210810362000450576200044f620003f2565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620004ba7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200047b565b620004c686836200047b565b95508019841693508086168417925050509392505050565b6000819050919050565b60006200050962000503620004fd846200034b565b620004de565b6200034b565b9050919050565b6000819050919050565b6200052583620004e8565b6200053d620005348262000510565b84845462000488565b825550505050565b600090565b6200055462000545565b620005618184846200051a565b505050565b5b8181101562000589576200057d6000826200054a565b60018101905062000567565b5050565b601f821115620005d857620005a28162000456565b620005ad846200046b565b81016020851015620005bd578190505b620005d5620005cc856200046b565b83018262000566565b50505b505050565b600082821c905092915050565b6000620005fd60001984600802620005dd565b1980831691505092915050565b6000620006188383620005ea565b9150826002028217905092915050565b6200063382620003b8565b67ffffffffffffffff8111156200064f576200064e620003c3565b5b6200065b825462000421565b620006688282856200058d565b600060209050601f831160018114620006a057600084156200068b578287015190505b6200069785826200060a565b86555062000707565b601f198416620006b08662000456565b60005b82811015620006da57848901518255600182019150602085019450602081019050620006b3565b86831015620006fa5784890151620006f6601f891682620005ea565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b600062000758601f836200070f565b9150620007658262000720565b602082019050919050565b600060208201905081810360008301526200078b8162000749565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000620007ce826200034b565b9150620007db836200034b565b9250828201905080821115620007f657620007f562000792565b5b92915050565b62000807816200034b565b82525050565b6000602082019050620008246000830184620007fc565b92915050565b611fb1806200083a6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c806356af2655116100a257806395d89b411161007157806395d89b41146102a6578063a457c2d7146102c4578063a9059cbb146102f4578063dd62ed3e14610324578063f2fde38b146103545761010b565b806356af26551461023257806370a082311461024e578063715018a61461027e5780638da5cb5b146102885761010b565b8063313ce567116100de578063313ce567146101ac57806339509351146101ca57806340c10f19146101fa578063569775a6146102165761010b565b806306fdde0314610110578063095ea7b31461012e57806318160ddd1461015e57806323b872dd1461017c575b600080fd5b610118610370565b6040516101259190611136565b60405180910390f35b61014860048036038101906101439190611200565b610402565b604051610155919061125b565b60405180910390f35b610166610425565b6040516101739190611285565b60405180910390f35b610196600480360381019061019191906112a0565b61042f565b6040516101a3919061125b565b60405180910390f35b6101b461045e565b6040516101c1919061130f565b60405180910390f35b6101e460048036038101906101df9190611200565b610467565b6040516101f1919061125b565b60405180910390f35b610214600480360381019061020f9190611200565b61049e565b005b610230600480360381019061022b919061145f565b6104b4565b005b61024c6004803603810190610247919061150a565b610543565b005b6102686004803603810190610263919061158d565b610614565b6040516102759190611285565b60405180910390f35b61028661065d565b005b610290610671565b60405161029d91906115c9565b60405180910390f35b6102ae61069a565b6040516102bb9190611136565b60405180910390f35b6102de60048036038101906102d99190611200565b61072c565b6040516102eb919061125b565b60405180910390f35b61030e60048036038101906103099190611200565b6107a3565b60405161031b919061125b565b60405180910390f35b61033e600480360381019061033991906115e4565b6107c6565b60405161034b9190611285565b60405180910390f35b61036e6004803603810190610369919061158d565b61084d565b005b60606004805461037f90611653565b80601f01602080910402602001604051908101604052809291908181526020018280546103ab90611653565b80156103f85780601f106103cd576101008083540402835291602001916103f8565b820191906000526020600020905b8154815290600101906020018083116103db57829003601f168201915b5050505050905090565b60008061040d6108d0565b905061041a8185856108d8565b600191505092915050565b6000600354905090565b60008061043a6108d0565b9050610447858285610aa1565b610452858585610b2d565b60019150509392505050565b60006012905090565b6000806104726108d0565b905061049381858561048485896107c6565b61048e91906116b3565b6108d8565b600191505092915050565b6104a6610da6565b6104b08282610e24565b5050565b6006816040516104c49190611723565b9081526020016040518091039020600080820160006104e39190611049565b600182016000905560028201600090556003820160006101000a81549060ff021916905550507f439cc593a85f8085eb9c8d29c9d38661a8ba12d449d1a89f126048d6a5e013e8816040516105389190611136565b60405180910390a150565b60405180608001604052808581526020018481526020018381526020018215158152506006856040516105769190611723565b9081526020016040518091039020600082015181600001908161059991906118e6565b50602082015181600101556040820151816002015560608201518160030160006101000a81548160ff0219169083151502179055509050507f3f53892a00eab68ee250fc793c38d2be4d77c378a98257e847fa8ba8d4f9dcb08484848460405161060694939291906119c7565b60405180910390a150505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b610665610da6565b61066f6000610f7b565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600580546106a990611653565b80601f01602080910402602001604051908101604052809291908181526020018280546106d590611653565b80156107225780601f106106f757610100808354040283529160200191610722565b820191906000526020600020905b81548152906001019060200180831161070557829003601f168201915b5050505050905090565b6000806107376108d0565b9050600061074582866107c6565b90508381101561078a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078190611a85565b60405180910390fd5b61079782868684036108d8565b60019250505092915050565b6000806107ae6108d0565b90506107bb818585610b2d565b600191505092915050565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b610855610da6565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036108c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108bb90611b17565b60405180910390fd5b6108cd81610f7b565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610947576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093e90611ba9565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036109b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ad90611c3b565b60405180910390fd5b80600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610a949190611285565b60405180910390a3505050565b6000610aad84846107c6565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610b275781811015610b19576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b1090611ca7565b60405180910390fd5b610b2684848484036108d8565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610b9c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b9390611d39565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c0b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c0290611dcb565b60405180910390fd5b610c1683838361103f565b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610c9d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9490611e5d565b60405180910390fd5b818103600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610d8d9190611285565b60405180910390a3610da0848484611044565b50505050565b610dae6108d0565b73ffffffffffffffffffffffffffffffffffffffff16610dcc610671565b73ffffffffffffffffffffffffffffffffffffffff1614610e22576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1990611ec9565b60405180910390fd5b565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610e93576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e8a90611f35565b60405180910390fd5b610e9f6000838361103f565b8060036000828254610eb191906116b3565b9250508190555080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610f639190611285565b60405180910390a3610f7760008383611044565b5050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b505050565b505050565b50805461105590611653565b6000825580601f106110675750611086565b601f0160209004906000526020600020908101906110859190611089565b5b50565b5b808211156110a257600081600090555060010161108a565b5090565b600081519050919050565b600082825260208201905092915050565b60005b838110156110e05780820151818401526020810190506110c5565b60008484015250505050565b6000601f19601f8301169050919050565b6000611108826110a6565b61111281856110b1565b93506111228185602086016110c2565b61112b816110ec565b840191505092915050565b6000602082019050818103600083015261115081846110fd565b905092915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006111978261116c565b9050919050565b6111a78161118c565b81146111b257600080fd5b50565b6000813590506111c48161119e565b92915050565b6000819050919050565b6111dd816111ca565b81146111e857600080fd5b50565b6000813590506111fa816111d4565b92915050565b6000806040838503121561121757611216611162565b5b6000611225858286016111b5565b9250506020611236858286016111eb565b9150509250929050565b60008115159050919050565b61125581611240565b82525050565b6000602082019050611270600083018461124c565b92915050565b61127f816111ca565b82525050565b600060208201905061129a6000830184611276565b92915050565b6000806000606084860312156112b9576112b8611162565b5b60006112c7868287016111b5565b93505060206112d8868287016111b5565b92505060406112e9868287016111eb565b9150509250925092565b600060ff82169050919050565b611309816112f3565b82525050565b60006020820190506113246000830184611300565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61136c826110ec565b810181811067ffffffffffffffff8211171561138b5761138a611334565b5b80604052505050565b600061139e611158565b90506113aa8282611363565b919050565b600067ffffffffffffffff8211156113ca576113c9611334565b5b6113d3826110ec565b9050602081019050919050565b82818337600083830152505050565b60006114026113fd846113af565b611394565b90508281526020810184848401111561141e5761141d61132f565b5b6114298482856113e0565b509392505050565b600082601f8301126114465761144561132a565b5b81356114568482602086016113ef565b91505092915050565b60006020828403121561147557611474611162565b5b600082013567ffffffffffffffff81111561149357611492611167565b5b61149f84828501611431565b91505092915050565b6000819050919050565b6114bb816114a8565b81146114c657600080fd5b50565b6000813590506114d8816114b2565b92915050565b6114e781611240565b81146114f257600080fd5b50565b600081359050611504816114de565b92915050565b6000806000806080858703121561152457611523611162565b5b600085013567ffffffffffffffff81111561154257611541611167565b5b61154e87828801611431565b945050602061155f878288016114c9565b9350506040611570878288016111eb565b9250506060611581878288016114f5565b91505092959194509250565b6000602082840312156115a3576115a2611162565b5b60006115b1848285016111b5565b91505092915050565b6115c38161118c565b82525050565b60006020820190506115de60008301846115ba565b92915050565b600080604083850312156115fb576115fa611162565b5b6000611609858286016111b5565b925050602061161a858286016111b5565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061166b57607f821691505b60208210810361167e5761167d611624565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006116be826111ca565b91506116c9836111ca565b92508282019050808211156116e1576116e0611684565b5b92915050565b600081905092915050565b60006116fd826110a6565b61170781856116e7565b93506117178185602086016110c2565b80840191505092915050565b600061172f82846116f2565b915081905092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261179c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261175f565b6117a6868361175f565b95508019841693508086168417925050509392505050565b6000819050919050565b60006117e36117de6117d9846111ca565b6117be565b6111ca565b9050919050565b6000819050919050565b6117fd836117c8565b611811611809826117ea565b84845461176c565b825550505050565b600090565b611826611819565b6118318184846117f4565b505050565b5b818110156118555761184a60008261181e565b600181019050611837565b5050565b601f82111561189a5761186b8161173a565b6118748461174f565b81016020851015611883578190505b61189761188f8561174f565b830182611836565b50505b505050565b600082821c905092915050565b60006118bd6000198460080261189f565b1980831691505092915050565b60006118d683836118ac565b9150826002028217905092915050565b6118ef826110a6565b67ffffffffffffffff81111561190857611907611334565b5b6119128254611653565b61191d828285611859565b600060209050601f831160018114611950576000841561193e578287015190505b61194885826118ca565b8655506119b0565b601f19841661195e8661173a565b60005b8281101561198657848901518255600182019150602085019450602081019050611961565b868310156119a3578489015161199f601f8916826118ac565b8355505b6001600288020188555050505b505050505050565b6119c1816114a8565b82525050565b600060808201905081810360008301526119e181876110fd565b90506119f060208301866119b8565b6119fd6040830185611276565b611a0a606083018461124c565b95945050505050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b6000611a6f6025836110b1565b9150611a7a82611a13565b604082019050919050565b60006020820190508181036000830152611a9e81611a62565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611b016026836110b1565b9150611b0c82611aa5565b604082019050919050565b60006020820190508181036000830152611b3081611af4565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000611b936024836110b1565b9150611b9e82611b37565b604082019050919050565b60006020820190508181036000830152611bc281611b86565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b6000611c256022836110b1565b9150611c3082611bc9565b604082019050919050565b60006020820190508181036000830152611c5481611c18565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b6000611c91601d836110b1565b9150611c9c82611c5b565b602082019050919050565b60006020820190508181036000830152611cc081611c84565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b6000611d236025836110b1565b9150611d2e82611cc7565b604082019050919050565b60006020820190508181036000830152611d5281611d16565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000611db56023836110b1565b9150611dc082611d59565b604082019050919050565b60006020820190508181036000830152611de481611da8565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b6000611e476026836110b1565b9150611e5282611deb565b604082019050919050565b60006020820190508181036000830152611e7681611e3a565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611eb36020836110b1565b9150611ebe82611e7d565b602082019050919050565b60006020820190508181036000830152611ee281611ea6565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000611f1f601f836110b1565b9150611f2a82611ee9565b602082019050919050565b60006020820190508181036000830152611f4e81611f12565b905091905056fea26469706673582212200320985687c4197e262172f6e1e55276f99e6f72a4246436a011ac4e537a6f9e64736f6c637828302e382e32302d646576656c6f702e323032332e322e32352b636f6d6d69742e39383334303737360059",
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

// AddActor is a paid mutator transaction binding the contract method 0x56af2655.
//
// Solidity: function addActor(string id, int256 firstParam, uint256 secondParam, bool thirdParam) returns()
func (_Main *MainTransactor) AddActor(opts *bind.TransactOpts, id string, firstParam *big.Int, secondParam *big.Int, thirdParam bool) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addActor", id, firstParam, secondParam, thirdParam)
}

// AddActor is a paid mutator transaction binding the contract method 0x56af2655.
//
// Solidity: function addActor(string id, int256 firstParam, uint256 secondParam, bool thirdParam) returns()
func (_Main *MainSession) AddActor(id string, firstParam *big.Int, secondParam *big.Int, thirdParam bool) (*types.Transaction, error) {
	return _Main.Contract.AddActor(&_Main.TransactOpts, id, firstParam, secondParam, thirdParam)
}

// AddActor is a paid mutator transaction binding the contract method 0x56af2655.
//
// Solidity: function addActor(string id, int256 firstParam, uint256 secondParam, bool thirdParam) returns()
func (_Main *MainTransactorSession) AddActor(id string, firstParam *big.Int, secondParam *big.Int, thirdParam bool) (*types.Transaction, error) {
	return _Main.Contract.AddActor(&_Main.TransactOpts, id, firstParam, secondParam, thirdParam)
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

// DelActor is a paid mutator transaction binding the contract method 0x569775a6.
//
// Solidity: function delActor(string id) returns()
func (_Main *MainTransactor) DelActor(opts *bind.TransactOpts, id string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "delActor", id)
}

// DelActor is a paid mutator transaction binding the contract method 0x569775a6.
//
// Solidity: function delActor(string id) returns()
func (_Main *MainSession) DelActor(id string) (*types.Transaction, error) {
	return _Main.Contract.DelActor(&_Main.TransactOpts, id)
}

// DelActor is a paid mutator transaction binding the contract method 0x569775a6.
//
// Solidity: function delActor(string id) returns()
func (_Main *MainTransactorSession) DelActor(id string) (*types.Transaction, error) {
	return _Main.Contract.DelActor(&_Main.TransactOpts, id)
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
	ThirdParam  bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterActorAdded is a free log retrieval operation binding the contract event 0x3f53892a00eab68ee250fc793c38d2be4d77c378a98257e847fa8ba8d4f9dcb0.
//
// Solidity: event ActorAdded(string id, int256 firstParam, uint256 secondParam, bool thirdParam)
func (_Main *MainFilterer) FilterActorAdded(opts *bind.FilterOpts) (*MainActorAddedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "ActorAdded")
	if err != nil {
		return nil, err
	}
	return &MainActorAddedIterator{contract: _Main.contract, event: "ActorAdded", logs: logs, sub: sub}, nil
}

// WatchActorAdded is a free log subscription operation binding the contract event 0x3f53892a00eab68ee250fc793c38d2be4d77c378a98257e847fa8ba8d4f9dcb0.
//
// Solidity: event ActorAdded(string id, int256 firstParam, uint256 secondParam, bool thirdParam)
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

// ParseActorAdded is a log parse operation binding the contract event 0x3f53892a00eab68ee250fc793c38d2be4d77c378a98257e847fa8ba8d4f9dcb0.
//
// Solidity: event ActorAdded(string id, int256 firstParam, uint256 secondParam, bool thirdParam)
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
