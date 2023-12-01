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
	Bin: "0x60806040525f60055534801562000014575f80fd5b50604051620021243803806200212483398181016040528101906200003a91906200027d565b5f83101562000080576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000779062000398565b60405180910390fd5b82426200008e9190620003e5565b600481905550815f9081620000a491906200064d565b508060019081620000b691906200064d565b5050505062000731565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b620000e581620000d1565b8114620000f0575f80fd5b50565b5f815190506200010381620000da565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b620001598262000111565b810181811067ffffffffffffffff821117156200017b576200017a62000121565b5b80604052505050565b5f6200018f620000c0565b90506200019d82826200014e565b919050565b5f67ffffffffffffffff821115620001bf57620001be62000121565b5b620001ca8262000111565b9050602081019050919050565b5f5b83811015620001f6578082015181840152602081019050620001d9565b5f8484015250505050565b5f620002176200021184620001a2565b62000184565b9050828152602081018484840111156200023657620002356200010d565b5b62000243848285620001d7565b509392505050565b5f82601f83011262000262576200026162000109565b5b81516200027484826020860162000201565b91505092915050565b5f805f60608486031215620002975762000296620000c9565b5b5f620002a686828701620000f3565b935050602084015167ffffffffffffffff811115620002ca57620002c9620000cd565b5b620002d8868287016200024b565b925050604084015167ffffffffffffffff811115620002fc57620002fb620000cd565b5b6200030a868287016200024b565b9150509250925092565b5f82825260208201905092915050565b7f436f6d6d69742070686173652063616e6e6f74206265206c657373207468616e5f8201527f203230207365636f6e64732e0000000000000000000000000000000000000000602082015250565b5f62000380602c8362000314565b91506200038d8262000324565b604082019050919050565b5f6020820190508181035f830152620003b18162000372565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f620003f182620000d1565b9150620003fe83620000d1565b9250828201905080821115620004195762000418620003b8565b5b92915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200046e57607f821691505b60208210810362000484576200048362000429565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620004e87fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620004ab565b620004f48683620004ab565b95508019841693508086168417925050509392505050565b5f819050919050565b5f620005356200052f6200052984620000d1565b6200050c565b620000d1565b9050919050565b5f819050919050565b620005508362000515565b620005686200055f826200053c565b848454620004b7565b825550505050565b5f90565b6200057e62000570565b6200058b81848462000545565b505050565b5b81811015620005b257620005a65f8262000574565b60018101905062000591565b5050565b601f8211156200060157620005cb816200048a565b620005d6846200049c565b81016020851015620005e6578190505b620005fe620005f5856200049c565b83018262000590565b50505b505050565b5f82821c905092915050565b5f620006235f198460080262000606565b1980831691505092915050565b5f6200063d838362000612565b9150826002028217905092915050565b62000658826200041f565b67ffffffffffffffff81111562000674576200067362000121565b5b62000680825462000456565b6200068d828285620005b6565b5f60209050601f831160018114620006c3575f8415620006ae578287015190505b620006ba858262000630565b86555062000729565b601f198416620006d3866200048a565b5f5b82811015620006fc57848901518255600182019150602085019450602081019050620006d5565b868310156200071c578489015162000718601f89168262000612565b8355505b6001600288020188555050505b505050505050565b6119e5806200073f5f395ff3fe608060405234801561000f575f80fd5b50600436106100b2575f3560e01c806387cfa5bf1161006f57806387cfa5bf146101665780638e7ea5b214610184578063a6193504146101a2578063baa13bb5146101c0578063bf9d0698146101de578063fd5bd01d1461020e576100b2565b8063031b6dbf146100b65780630bbfc206146100d457806322ce9bac146100f05780633a2a2cc51461010e5780633e8589231461012c57806349e811c514610148575b5f80fd5b6100be61023e565b6040516100cb9190610c97565b60405180910390f35b6100ee60048036038101906100e99190610e30565b610244565b005b6100f8610638565b6040516101059190610c97565b60405180910390f35b61011661063e565b6040516101239190610f41565b60405180910390f35b61014660048036038101906101419190610f61565b610694565b005b610150610882565b60405161015d9190611006565b60405180910390f35b61016e61090d565b60405161017b9190611006565b60405180910390f35b61018c610999565b6040516101999190611006565b60405180910390f35b6101aa610bb8565b6040516101b79190610c97565b60405180910390f35b6101c8610bbe565b6040516101d59190610c97565b60405180910390f35b6101f860048036038101906101f39190611050565b610bc4565b604051610205919061108a565b60405180910390f35b61022860048036038101906102239190610f61565b610be4565b6040516102359190611006565b60405180910390f35b60035481565b5f60075f8381526020019081526020015f208054610261906110d0565b80601f016020809104026020016040519081016040528092919081815260200182805461028d906110d0565b80156102d85780601f106102af576101008083540402835291602001916102d8565b820191905f5260205f20905b8154815290600101906020018083116102bb57829003601f168201915b505050505090505f815103610322576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031990611170565b60405180910390fd5b7f4300000000000000000000000000000000000000000000000000000000000000815f815181106103565761035561118e565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146103c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ba90611205565b60405180910390fd5b826040516020016103d4919061125d565b60405160208183030381529060405280519060200120821461042b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610422906112e3565b60405180910390fd5b5f8390507f3100000000000000000000000000000000000000000000000000000000000000815f815181106104635761046261118e565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916036104e85760016002546104a4919061132e565b6002819055507f69ab1b0836124ec1de36500cc76e0bcd684f01f02d992b7b26f6ad222a2ddf21835f6040516104db9291906113f4565b60405180910390a16105de565b7f3200000000000000000000000000000000000000000000000000000000000000815f8151811061051c5761051b61118e565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916036105a257600160035461055d919061132e565b6003819055507f69ab1b0836124ec1de36500cc76e0bcd684f01f02d992b7b26f6ad222a2ddf218360016040516105959291906113f4565b60405180910390a16105dd565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d4906114b8565b60405180910390fd5b5b6040518060400160405280600881526020017f52657665616c656400000000000000000000000000000000000000000000000081525060075f8581526020019081526020015f2090816106319190611661565b5050505050565b60055481565b6060600680548060200260200160405190810160405280929190818152602001828054801561068a57602002820191905f5260205f20905b815481526020019060010190808311610676575b5050505050905090565b60045442106106d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106cf906117a0565b60405180910390fd5b5f60075f8381526020019081526020015f2080546106f5906110d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610721906110d0565b801561076c5780601f106107435761010080835404028352916020019161076c565b820191905f5260205f20905b81548152906001019060200180831161074f57829003601f168201915b505050505090505f8151146107b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ad9061182e565b60405180910390fd5b600682908060018154018082558091505060019003905f5260205f20015f90919091909150556040518060400160405280600981526020017f436f6d6d6974746564000000000000000000000000000000000000000000000081525060075f8481526020019081526020015f20908161082f9190611661565b5060055f8154809291906108429061184c565b91905055507f336dd6fa514fe196ac01f0e19a495fc73b7885d9357f993cf52e4919f70db09982604051610876919061108a565b60405180910390a15050565b5f805461088e906110d0565b80601f01602080910402602001604051908101604052809291908181526020018280546108ba906110d0565b80156109055780601f106108dc57610100808354040283529160200191610905565b820191905f5260205f20905b8154815290600101906020018083116108e857829003601f168201915b505050505081565b6001805461091a906110d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610946906110d0565b80156109915780601f1061096857610100808354040283529160200191610991565b820191905f5260205f20905b81548152906001019060200180831161097457829003601f168201915b505050505081565b606060045442116109df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d690611903565b60405180910390fd5b6006805490506003546002546109f5919061132e565b14610a35576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2c90611991565b60405180910390fd5b6003546002541115610ad0575f8054610a4d906110d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610a79906110d0565b8015610ac45780601f10610a9b57610100808354040283529160200191610ac4565b820191905f5260205f20905b815481529060010190602001808311610aa757829003601f168201915b50505050509050610bb5565b6002546003541115610b6c5760018054610ae9906110d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610b15906110d0565b8015610b605780601f10610b3757610100808354040283529160200191610b60565b820191905f5260205f20905b815481529060010190602001808311610b4357829003601f168201915b50505050509050610bb5565b60035460025403610bb4576040518060400160405280600d81526020017f49742077617320612074696521000000000000000000000000000000000000008152509050610bb5565b5b90565b60045481565b60025481565b60068181548110610bd3575f80fd5b905f5260205f20015f915090505481565b6007602052805f5260405f205f915090508054610c00906110d0565b80601f0160208091040260200160405190810160405280929190818152602001828054610c2c906110d0565b8015610c775780601f10610c4e57610100808354040283529160200191610c77565b820191905f5260205f20905b815481529060010190602001808311610c5a57829003601f168201915b505050505081565b5f819050919050565b610c9181610c7f565b82525050565b5f602082019050610caa5f830184610c88565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610d0f82610cc9565b810181811067ffffffffffffffff82111715610d2e57610d2d610cd9565b5b80604052505050565b5f610d40610cb0565b9050610d4c8282610d06565b919050565b5f67ffffffffffffffff821115610d6b57610d6a610cd9565b5b610d7482610cc9565b9050602081019050919050565b828183375f83830152505050565b5f610da1610d9c84610d51565b610d37565b905082815260208101848484011115610dbd57610dbc610cc5565b5b610dc8848285610d81565b509392505050565b5f82601f830112610de457610de3610cc1565b5b8135610df4848260208601610d8f565b91505092915050565b5f819050919050565b610e0f81610dfd565b8114610e19575f80fd5b50565b5f81359050610e2a81610e06565b92915050565b5f8060408385031215610e4657610e45610cb9565b5b5f83013567ffffffffffffffff811115610e6357610e62610cbd565b5b610e6f85828601610dd0565b9250506020610e8085828601610e1c565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610ebc81610dfd565b82525050565b5f610ecd8383610eb3565b60208301905092915050565b5f602082019050919050565b5f610eef82610e8a565b610ef98185610e94565b9350610f0483610ea4565b805f5b83811015610f34578151610f1b8882610ec2565b9750610f2683610ed9565b925050600181019050610f07565b5085935050505092915050565b5f6020820190508181035f830152610f598184610ee5565b905092915050565b5f60208284031215610f7657610f75610cb9565b5b5f610f8384828501610e1c565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610fc3578082015181840152602081019050610fa8565b5f8484015250505050565b5f610fd882610f8c565b610fe28185610f96565b9350610ff2818560208601610fa6565b610ffb81610cc9565b840191505092915050565b5f6020820190508181035f83015261101e8184610fce565b905092915050565b61102f81610c7f565b8114611039575f80fd5b50565b5f8135905061104a81611026565b92915050565b5f6020828403121561106557611064610cb9565b5b5f6110728482850161103c565b91505092915050565b61108481610dfd565b82525050565b5f60208201905061109d5f83018461107b565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806110e757607f821691505b6020821081036110fa576110f96110a3565b5b50919050565b7f4120766f74652077697468207468697320766f7465436f6d6d697420776173205f8201527f6e6f7420636173742e0000000000000000000000000000000000000000000000602082015250565b5f61115a602983610f96565b915061116582611100565b604082019050919050565b5f6020820190508181035f8301526111878161114e565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f5468697320766f74652077617320616c72656164792072657665616c65642e005f82015250565b5f6111ef601f83610f96565b91506111fa826111bb565b602082019050919050565b5f6020820190508181035f83015261121c816111e3565b9050919050565b5f81905092915050565b5f61123782610f8c565b6112418185611223565b9350611251818560208601610fa6565b80840191505092915050565b5f611268828461122d565b915081905092915050565b7f566f7465206861736820646f6573206e6f74206d6174636820766f746520636f5f8201527f6d6d69742e000000000000000000000000000000000000000000000000000000602082015250565b5f6112cd602583610f96565b91506112d882611273565b604082019050919050565b5f6020820190508181035f8301526112fa816112c1565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61133882610c7f565b915061134383610c7f565b925082820190508082111561135b5761135a611301565b5b92915050565b5f819050815f5260205f209050919050565b5f815461137f816110d0565b6113898186610f96565b9450600182165f81146113a357600181146113b9576113eb565b60ff1983168652811515602002860193506113eb565b6113c285611361565b5f5b838110156113e3578154818901526001820191506020810190506113c4565b808801955050505b50505092915050565b5f6040820190506114075f83018561107b565b81810360208301526114198184611373565b90509392505050565b7f566f746520636f756c64206e6f7420626520726561642120566f746573206d755f8201527f737420737461727420776974682074686520415343494920636861726163746560208201527f7220603160206f72206032600000000000000000000000000000000000000000604082015250565b5f6114a2604c83610f96565b91506114ad82611422565b606082019050919050565b5f6020820190508181035f8301526114cf81611496565b9050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026115207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826114e5565b61152a86836114e5565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61156561156061155b84610c7f565b611542565b610c7f565b9050919050565b5f819050919050565b61157e8361154b565b61159261158a8261156c565b8484546114f1565b825550505050565b5f90565b6115a661159a565b6115b1818484611575565b505050565b5b818110156115d4576115c95f8261159e565b6001810190506115b7565b5050565b601f821115611619576115ea81611361565b6115f3846114d6565b81016020851015611602578190505b61161661160e856114d6565b8301826115b6565b50505b505050565b5f82821c905092915050565b5f6116395f198460080261161e565b1980831691505092915050565b5f611651838361162a565b9150826002028217905092915050565b61166a82610f8c565b67ffffffffffffffff81111561168357611682610cd9565b5b61168d82546110d0565b6116988282856115d8565b5f60209050601f8311600181146116c9575f84156116b7578287015190505b6116c18582611646565b865550611728565b601f1984166116d786611361565b5f5b828110156116fe578489015182556001820191506020850194506020810190506116d9565b8683101561171b5784890151611717601f89168261162a565b8355505b6001600288020188555050505b505050505050565b7f4f6e6c7920616c6c6f77656420746f20636f6d6d697420647572696e6720636f5f8201527f6d6d697474696e6720706572696f642e00000000000000000000000000000000602082015250565b5f61178a603083610f96565b915061179582611730565b604082019050919050565b5f6020820190508181035f8301526117b78161177e565b9050919050565b7f5468697320636f6d6d69742068617320616c7265616479206265656e207573655f8201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b5f611818602183610f96565b9150611823826117be565b604082019050919050565b5f6020820190508181035f8301526118458161180c565b9050919050565b5f61185682610c7f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361188857611887611301565b5b600182019050919050565b7f4f6e6c7920616c6c6f77656420746f206765742077696e6e65722061667465725f8201527f20636f6d6d6974207068617365206973206f7665722e00000000000000000000602082015250565b5f6118ed603683610f96565b91506118f882611893565b604082019050919050565b5f6020820190508181035f83015261191a816118e1565b9050919050565b7f43616e206f6e6c79206765742077696e6e6572207768656e20616c6c20766f745f8201527f6573206172652072657665616c65642e00000000000000000000000000000000602082015250565b5f61197b603083610f96565b915061198682611921565b604082019050919050565b5f6020820190508181035f8301526119a88161196f565b905091905056fea2646970667358221220749dde9a54d04f5c715a75f1178bb5b87785233c6941023ef84286f61446eb0264736f6c63430008170033",
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
