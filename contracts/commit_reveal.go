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
	Bin: "0x608060405260006005553480156200001657600080fd5b50604051620021fd380380620021fd83398181016040528101906200003c919062000292565b600083101562000083576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200007a90620003b3565b60405180910390fd5b824262000091919062000404565b6004819055508160009081620000a8919062000680565b508060019081620000ba919062000680565b5050505062000767565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b620000ed81620000d8565b8114620000f957600080fd5b50565b6000815190506200010d81620000e2565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b62000168826200011d565b810181811067ffffffffffffffff821117156200018a57620001896200012e565b5b80604052505050565b60006200019f620000c4565b9050620001ad82826200015d565b919050565b600067ffffffffffffffff821115620001d057620001cf6200012e565b5b620001db826200011d565b9050602081019050919050565b60005b8381101562000208578082015181840152602081019050620001eb565b60008484015250505050565b60006200022b6200022584620001b2565b62000193565b9050828152602081018484840111156200024a576200024962000118565b5b62000257848285620001e8565b509392505050565b600082601f83011262000277576200027662000113565b5b81516200028984826020860162000214565b91505092915050565b600080600060608486031215620002ae57620002ad620000ce565b5b6000620002be86828701620000fc565b935050602084015167ffffffffffffffff811115620002e257620002e1620000d3565b5b620002f0868287016200025f565b925050604084015167ffffffffffffffff811115620003145762000313620000d3565b5b62000322868287016200025f565b9150509250925092565b600082825260208201905092915050565b7f436f6d6d69742070686173652063616e6e6f74206265206c657373207468616e60008201527f203230207365636f6e64732e0000000000000000000000000000000000000000602082015250565b60006200039b602c836200032c565b9150620003a8826200033d565b604082019050919050565b60006020820190508181036000830152620003ce816200038c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006200041182620000d8565b91506200041e83620000d8565b9250828201905080821115620004395762000438620003d5565b5b92915050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200049257607f821691505b602082108103620004a857620004a76200044a565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620005127fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620004d3565b6200051e8683620004d3565b95508019841693508086168417925050509392505050565b6000819050919050565b6000620005616200055b6200055584620000d8565b62000536565b620000d8565b9050919050565b6000819050919050565b6200057d8362000540565b620005956200058c8262000568565b848454620004e0565b825550505050565b600090565b620005ac6200059d565b620005b981848462000572565b505050565b5b81811015620005e157620005d5600082620005a2565b600181019050620005bf565b5050565b601f8211156200063057620005fa81620004ae565b6200060584620004c3565b8101602085101562000615578190505b6200062d6200062485620004c3565b830182620005be565b50505b505050565b600082821c905092915050565b6000620006556000198460080262000635565b1980831691505092915050565b600062000670838362000642565b9150826002028217905092915050565b6200068b826200043f565b67ffffffffffffffff811115620006a757620006a66200012e565b5b620006b3825462000479565b620006c0828285620005e5565b600060209050601f831160018114620006f85760008415620006e3578287015190505b620006ef858262000662565b8655506200075f565b601f1984166200070886620004ae565b60005b8281101562000732578489015182556001820191506020850194506020810190506200070b565b868310156200075257848901516200074e601f89168262000642565b8355505b6001600288020188555050505b505050505050565b611a8680620007776000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806387cfa5bf1161007157806387cfa5bf146101695780638e7ea5b214610187578063a6193504146101a5578063baa13bb5146101c3578063bf9d0698146101e1578063fd5bd01d14610211576100b4565b8063031b6dbf146100b95780630bbfc206146100d757806322ce9bac146100f35780633a2a2cc5146101115780633e8589231461012f57806349e811c51461014b575b600080fd5b6100c1610241565b6040516100ce9190610cc9565b60405180910390f35b6100f160048036038101906100ec9190610e74565b610247565b005b6100fb610648565b6040516101089190610cc9565b60405180910390f35b61011961064e565b6040516101269190610f8e565b60405180910390f35b61014960048036038101906101449190610fb0565b6106a6565b005b6101536108a0565b604051610160919061105c565b60405180910390f35b61017161092e565b60405161017e919061105c565b60405180910390f35b61018f6109bc565b60405161019c919061105c565b60405180910390f35b6101ad610be0565b6040516101ba9190610cc9565b60405180910390f35b6101cb610be6565b6040516101d89190610cc9565b60405180910390f35b6101fb60048036038101906101f691906110aa565b610bec565b60405161020891906110e6565b60405180910390f35b61022b60048036038101906102269190610fb0565b610c10565b604051610238919061105c565b60405180910390f35b60035481565b600060076000838152602001908152602001600020805461026790611130565b80601f016020809104026020016040519081016040528092919081815260200182805461029390611130565b80156102e05780601f106102b5576101008083540402835291602001916102e0565b820191906000526020600020905b8154815290600101906020018083116102c357829003601f168201915b50505050509050600081510361032b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610322906111d3565b60405180910390fd5b7f4300000000000000000000000000000000000000000000000000000000000000816000815181106103605761035f6111f3565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146103cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103c49061126e565b60405180910390fd5b826040516020016103de91906112ca565b604051602081830303815290604052805190602001208214610435576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042c90611353565b60405180910390fd5b60008390507f31000000000000000000000000000000000000000000000000000000000000008160008151811061046f5761046e6111f3565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916036104f55760016002546104b091906113a2565b6002819055507f69ab1b0836124ec1de36500cc76e0bcd684f01f02d992b7b26f6ad222a2ddf218360006040516104e892919061146f565b60405180910390a16105ec565b7f32000000000000000000000000000000000000000000000000000000000000008160008151811061052a576105296111f3565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916036105b057600160035461056b91906113a2565b6003819055507f69ab1b0836124ec1de36500cc76e0bcd684f01f02d992b7b26f6ad222a2ddf218360016040516105a392919061146f565b60405180910390a16105eb565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e290611537565b60405180910390fd5b5b6040518060400160405280600881526020017f52657665616c656400000000000000000000000000000000000000000000000081525060076000858152602001908152602001600020908161064191906116ee565b5050505050565b60055481565b6060600680548060200260200160405190810160405280929190818152602001828054801561069c57602002820191906000526020600020905b815481526020019060010190808311610688575b5050505050905090565b60045442106106ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e190611832565b60405180910390fd5b600060076000838152602001908152602001600020805461070a90611130565b80601f016020809104026020016040519081016040528092919081815260200182805461073690611130565b80156107835780601f1061075857610100808354040283529160200191610783565b820191906000526020600020905b81548152906001019060200180831161076657829003601f168201915b5050505050905060008151146107ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c5906118c4565b60405180910390fd5b60068290806001815401808255809150506001900390600052602060002001600090919091909150556040518060400160405280600981526020017f436f6d6d6974746564000000000000000000000000000000000000000000000081525060076000848152602001908152602001600020908161084c91906116ee565b5060056000815480929190610860906118e4565b91905055507f336dd6fa514fe196ac01f0e19a495fc73b7885d9357f993cf52e4919f70db0998260405161089491906110e6565b60405180910390a15050565b600080546108ad90611130565b80601f01602080910402602001604051908101604052809291908181526020018280546108d990611130565b80156109265780601f106108fb57610100808354040283529160200191610926565b820191906000526020600020905b81548152906001019060200180831161090957829003601f168201915b505050505081565b6001805461093b90611130565b80601f016020809104026020016040519081016040528092919081815260200182805461096790611130565b80156109b45780601f10610989576101008083540402835291602001916109b4565b820191906000526020600020905b81548152906001019060200180831161099757829003601f168201915b505050505081565b60606004544211610a02576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109f99061199e565b60405180910390fd5b600680549050600354600254610a1891906113a2565b14610a58576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4f90611a30565b60405180910390fd5b6003546002541115610af65760008054610a7190611130565b80601f0160208091040260200160405190810160405280929190818152602001828054610a9d90611130565b8015610aea5780601f10610abf57610100808354040283529160200191610aea565b820191906000526020600020905b815481529060010190602001808311610acd57829003601f168201915b50505050509050610bdd565b6002546003541115610b945760018054610b0f90611130565b80601f0160208091040260200160405190810160405280929190818152602001828054610b3b90611130565b8015610b885780601f10610b5d57610100808354040283529160200191610b88565b820191906000526020600020905b815481529060010190602001808311610b6b57829003601f168201915b50505050509050610bdd565b60035460025403610bdc576040518060400160405280600d81526020017f49742077617320612074696521000000000000000000000000000000000000008152509050610bdd565b5b90565b60045481565b60025481565b60068181548110610bfc57600080fd5b906000526020600020016000915090505481565b60076020528060005260406000206000915090508054610c2f90611130565b80601f0160208091040260200160405190810160405280929190818152602001828054610c5b90611130565b8015610ca85780601f10610c7d57610100808354040283529160200191610ca8565b820191906000526020600020905b815481529060010190602001808311610c8b57829003601f168201915b505050505081565b6000819050919050565b610cc381610cb0565b82525050565b6000602082019050610cde6000830184610cba565b92915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610d4b82610d02565b810181811067ffffffffffffffff82111715610d6a57610d69610d13565b5b80604052505050565b6000610d7d610ce4565b9050610d898282610d42565b919050565b600067ffffffffffffffff821115610da957610da8610d13565b5b610db282610d02565b9050602081019050919050565b82818337600083830152505050565b6000610de1610ddc84610d8e565b610d73565b905082815260208101848484011115610dfd57610dfc610cfd565b5b610e08848285610dbf565b509392505050565b600082601f830112610e2557610e24610cf8565b5b8135610e35848260208601610dce565b91505092915050565b6000819050919050565b610e5181610e3e565b8114610e5c57600080fd5b50565b600081359050610e6e81610e48565b92915050565b60008060408385031215610e8b57610e8a610cee565b5b600083013567ffffffffffffffff811115610ea957610ea8610cf3565b5b610eb585828601610e10565b9250506020610ec685828601610e5f565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610f0581610e3e565b82525050565b6000610f178383610efc565b60208301905092915050565b6000602082019050919050565b6000610f3b82610ed0565b610f458185610edb565b9350610f5083610eec565b8060005b83811015610f81578151610f688882610f0b565b9750610f7383610f23565b925050600181019050610f54565b5085935050505092915050565b60006020820190508181036000830152610fa88184610f30565b905092915050565b600060208284031215610fc657610fc5610cee565b5b6000610fd484828501610e5f565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015611017578082015181840152602081019050610ffc565b60008484015250505050565b600061102e82610fdd565b6110388185610fe8565b9350611048818560208601610ff9565b61105181610d02565b840191505092915050565b600060208201905081810360008301526110768184611023565b905092915050565b61108781610cb0565b811461109257600080fd5b50565b6000813590506110a48161107e565b92915050565b6000602082840312156110c0576110bf610cee565b5b60006110ce84828501611095565b91505092915050565b6110e081610e3e565b82525050565b60006020820190506110fb60008301846110d7565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061114857607f821691505b60208210810361115b5761115a611101565b5b50919050565b7f4120766f74652077697468207468697320766f7465436f6d6d6974207761732060008201527f6e6f7420636173742e0000000000000000000000000000000000000000000000602082015250565b60006111bd602983610fe8565b91506111c882611161565b604082019050919050565b600060208201905081810360008301526111ec816111b0565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f5468697320766f74652077617320616c72656164792072657665616c65642e00600082015250565b6000611258601f83610fe8565b915061126382611222565b602082019050919050565b600060208201905081810360008301526112878161124b565b9050919050565b600081905092915050565b60006112a482610fdd565b6112ae818561128e565b93506112be818560208601610ff9565b80840191505092915050565b60006112d68284611299565b915081905092915050565b7f566f7465206861736820646f6573206e6f74206d6174636820766f746520636f60008201527f6d6d69742e000000000000000000000000000000000000000000000000000000602082015250565b600061133d602583610fe8565b9150611348826112e1565b604082019050919050565b6000602082019050818103600083015261136c81611330565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006113ad82610cb0565b91506113b883610cb0565b92508282019050808211156113d0576113cf611373565b5b92915050565b60008190508160005260206000209050919050565b600081546113f881611130565b6114028186610fe8565b9450600182166000811461141d576001811461143357611466565b60ff198316865281151560200286019350611466565b61143c856113d6565b60005b8381101561145e5781548189015260018201915060208101905061143f565b808801955050505b50505092915050565b600060408201905061148460008301856110d7565b818103602083015261149681846113eb565b90509392505050565b7f566f746520636f756c64206e6f7420626520726561642120566f746573206d7560008201527f737420737461727420776974682074686520415343494920636861726163746560208201527f7220603160206f72206032600000000000000000000000000000000000000000604082015250565b6000611521604c83610fe8565b915061152c8261149f565b606082019050919050565b6000602082019050818103600083015261155081611514565b9050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026115a47fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611567565b6115ae8683611567565b95508019841693508086168417925050509392505050565b6000819050919050565b60006115eb6115e66115e184610cb0565b6115c6565b610cb0565b9050919050565b6000819050919050565b611605836115d0565b611619611611826115f2565b848454611574565b825550505050565b600090565b61162e611621565b6116398184846115fc565b505050565b5b8181101561165d57611652600082611626565b60018101905061163f565b5050565b601f8211156116a257611673816113d6565b61167c84611557565b8101602085101561168b578190505b61169f61169785611557565b83018261163e565b50505b505050565b600082821c905092915050565b60006116c5600019846008026116a7565b1980831691505092915050565b60006116de83836116b4565b9150826002028217905092915050565b6116f782610fdd565b67ffffffffffffffff8111156117105761170f610d13565b5b61171a8254611130565b611725828285611661565b600060209050601f8311600181146117585760008415611746578287015190505b61175085826116d2565b8655506117b8565b601f198416611766866113d6565b60005b8281101561178e57848901518255600182019150602085019450602081019050611769565b868310156117ab57848901516117a7601f8916826116b4565b8355505b6001600288020188555050505b505050505050565b7f4f6e6c7920616c6c6f77656420746f20636f6d6d697420647572696e6720636f60008201527f6d6d697474696e6720706572696f642e00000000000000000000000000000000602082015250565b600061181c603083610fe8565b9150611827826117c0565b604082019050919050565b6000602082019050818103600083015261184b8161180f565b9050919050565b7f5468697320636f6d6d69742068617320616c7265616479206265656e2075736560008201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b60006118ae602183610fe8565b91506118b982611852565b604082019050919050565b600060208201905081810360008301526118dd816118a1565b9050919050565b60006118ef82610cb0565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361192157611920611373565b5b600182019050919050565b7f4f6e6c7920616c6c6f77656420746f206765742077696e6e657220616674657260008201527f20636f6d6d6974207068617365206973206f7665722e00000000000000000000602082015250565b6000611988603683610fe8565b91506119938261192c565b604082019050919050565b600060208201905081810360008301526119b78161197b565b9050919050565b7f43616e206f6e6c79206765742077696e6e6572207768656e20616c6c20766f7460008201527f6573206172652072657665616c65642e00000000000000000000000000000000602082015250565b6000611a1a603083610fe8565b9150611a25826119be565b604082019050919050565b60006020820190508181036000830152611a4981611a0d565b905091905056fea2646970667358221220b67391801ce4a7ca43e53c58f4055474b5038d6af0b05006f80479db2accb1dd64736f6c63430008170033",
}

// CommitRevealABI is the input ABI used to generate the binding from.
// Deprecated: Use CommitRevealMetaData.ABI instead.
var CommitRevealABI = CommitRevealMetaData.ABI

// CommitRevealBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CommitRevealMetaData.Bin instead.
var CommitRevealBin = CommitRevealMetaData.Bin

// DeployCommitReveal deploys a new Ethereum contract, binding an instance of CommitReveal to it.
func DeployCommitReveal(auth *bind.TransactOpts, backend bind.ContractBackend, _phaseLengthInSeconds *big.Int, _choice1 string, _choice2 string) (common.Address, *types.Transaction, *CommitReveal, error) {
	parsed, err := CommitRevealMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CommitRevealBin), backend, _phaseLengthInSeconds, _choice1, _choice2)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CommitReveal{CommitRevealCaller: CommitRevealCaller{contract: contract}, CommitRevealTransactor: CommitRevealTransactor{contract: contract}, CommitRevealFilterer: CommitRevealFilterer{contract: contract}}, nil
}

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
