// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/kroma-network/kroma/bindings/solc"
)

const ValidatorPoolStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/ValidatorPool.sol:ValidatorPool\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/ValidatorPool.sol:ValidatorPool\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/ValidatorPool.sol:ValidatorPool\",\"label\":\"_status\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"contracts/L1/ValidatorPool.sol:ValidatorPool\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_array(t_uint256)1011_storage\"},{\"astId\":1004,\"contract\":\"contracts/L1/ValidatorPool.sol:ValidatorPool\",\"label\":\"balances\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1005,\"contract\":\"contracts/L1/ValidatorPool.sol:ValidatorPool\",\"label\":\"bonds\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_mapping(t_uint256,t_struct(Bond)1012_storage)\"},{\"astId\":1006,\"contract\":\"contracts/L1/ValidatorPool.sol:ValidatorPool\",\"label\":\"nextUnbondOutputIndex\",\"offset\":0,\"slot\":\"53\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"contracts/L1/ValidatorPool.sol:ValidatorPool\",\"label\":\"validators\",\"offset\":0,\"slot\":\"54\",\"type\":\"t_array(t_address)dyn_storage\"},{\"astId\":1008,\"contract\":\"contracts/L1/ValidatorPool.sol:ValidatorPool\",\"label\":\"validatorIndexes\",\"offset\":0,\"slot\":\"55\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1009,\"contract\":\"contracts/L1/ValidatorPool.sol:ValidatorPool\",\"label\":\"nextPriorityValidator\",\"offset\":0,\"slot\":\"56\",\"type\":\"t_address\"},{\"astId\":1010,\"contract\":\"contracts/L1/ValidatorPool.sol:ValidatorPool\",\"label\":\"pendingBonds\",\"offset\":0,\"slot\":\"57\",\"type\":\"t_mapping(t_uint256,t_mapping(t_address,t_uint128))\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_address)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"address[]\",\"numberOfBytes\":\"32\"},\"t_array(t_uint256)1011_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_uint128)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint128)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint128\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_mapping(t_address,t_uint128))\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e mapping(address =\u003e uint128))\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_mapping(t_address,t_uint128)\"},\"t_mapping(t_uint256,t_struct(Bond)1012_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e struct Types.Bond)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_struct(Bond)1012_storage\"},\"t_struct(Bond)1012_storage\":{\"encoding\":\"inplace\",\"label\":\"struct Types.Bond\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var ValidatorPoolStorageLayout = new(solc.StorageLayout)

var ValidatorPoolDeployedBin = "0x60806040526004361061017f5760003560e01c80636641ea08116100d65780639fbc4a5f1161007f578063d8fe764211610059578063d8fe7642146104aa578063dd215c5d146104fa578063facd743b1461051a57600080fd5b80639fbc4a5f1461043e578063ab91f19014610472578063d0e30db0146104a257600080fd5b80638f09ade4116100b05780638f09ade4146103a9578063946765fd146103ea57806396946f751461041e57600080fd5b80636641ea081461032a57806370a082311461035e5780638129fc1c1461039457600080fd5b80633a549046116101385780635a544742116101125780635a544742146102d55780635b86f599146102f55780635df6a6bc1461031557600080fd5b80633a5490461461026a5780633ee4d4a31461027f57806354fd4d50146102b357600080fd5b80630ff754ea116101695780630ff754ea146101f45780632e1a7d4d1461022857806336b834691461024a57600080fd5b80621c2ff6146101845780630f43a677146101d5575b600080fd5b34801561019057600080fd5b506101b87f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156101e157600080fd5b506036545b6040519081526020016101cc565b34801561020057600080fd5b506101b87f000000000000000000000000000000000000000000000000000000000000000081565b34801561023457600080fd5b506102486102433660046122ce565b61054a565b005b34801561025657600080fd5b506102486102653660046122fc565b61064a565b34801561027657600080fd5b506101b86108bb565b34801561028b57600080fd5b506101b87f000000000000000000000000000000000000000000000000000000000000000081565b3480156102bf57600080fd5b506102c8610a79565b6040516101cc91906123b8565b3480156102e157600080fd5b506102486102f03660046123cb565b610b1c565b34801561030157600080fd5b506102486103103660046123fb565b610d89565b34801561032157600080fd5b50610248610e9f565b34801561033657600080fd5b506101e67f000000000000000000000000000000000000000000000000000000000000000081565b34801561036a57600080fd5b506101e6610379366004612427565b6001600160a01b031660009081526033602052604090205490565b3480156103a057600080fd5b50610248610f21565b3480156103b557600080fd5b506103c96103c43660046123cb565b611098565b6040516fffffffffffffffffffffffffffffffff90911681526020016101cc565b3480156103f657600080fd5b506101e67f000000000000000000000000000000000000000000000000000000000000000081565b34801561042a57600080fd5b50610248610439366004612462565b611147565b34801561044a57600080fd5b506101e67f000000000000000000000000000000000000000000000000000000000000000081565b34801561047e57600080fd5b50610489620186a081565b60405167ffffffffffffffff90911681526020016101cc565b610248611480565b3480156104b657600080fd5b506104ca6104c53660046122ce565b61148c565b6040805182516fffffffffffffffffffffffffffffffff90811682526020938401511692810192909252016101cc565b34801561050657600080fd5b506102486105153660046123cb565b6115ab565b34801561052657600080fd5b5061053a610535366004612427565b6118a6565b60405190151581526020016101cc565b6002600154036105a15760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064015b60405180910390fd5b60026001556105b0338261192c565b60006105cd335a8460405180602001604052806000815250611b6c565b9050806106425760405162461bcd60e51b815260206004820152602260248201527f56616c696461746f72506f6f6c3a20455448207472616e73666572206661696c60448201527f65640000000000000000000000000000000000000000000000000000000000006064820152608401610598565b505060018055565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639e45e8f46040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106cc9190612499565b6001600160a01b0316336001600160a01b0316146107525760405162461bcd60e51b815260206004820152602660248201527f56616c696461746f72506f6f6c3a2073656e646572206973206e6f7420436f6c60448201527f6f737365756d00000000000000000000000000000000000000000000000000006064820152608401610598565b60008381526039602090815260408083206001600160a01b03861684529091529020546fffffffffffffffffffffffffffffffff16806107fa5760405162461bcd60e51b815260206004820152602e60248201527f56616c696461746f72506f6f6c3a207468652070656e64696e6720626f6e642060448201527f646f6573206e6f742065786973740000000000000000000000000000000000006064820152608401610598565b60008481526039602090815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffff0000000000000000000000000000000016905561085e826fffffffffffffffffffffffffffffffff8316611b86565b6040516fffffffffffffffffffffffffffffffff821681526001600160a01b03808416919085169086907f8c95336a279406edcc768d685e8eb6667368a77d840a188144b8e3719423198f9060200160405180910390a450505050565b6038546000906001600160a01b031615610a545760007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663dcec33486040518163ffffffff1660e01b8152600401602060405180830381865afa15801561092f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095391906124b6565b905060006001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663d1de856c6109918460016124fe565b6040518263ffffffff1660e01b81526004016109af91815260200190565b602060405180830381865afa1580156109cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109f091906124b6565b9050804210610a41576000610a058242612516565b90507f0000000000000000000000000000000000000000000000000000000000000000811115610a3f576001600160a01b03935050505090565b505b50506038546001600160a01b0316919050565b507f000000000000000000000000000000000000000000000000000000000000000090565b6060610aa47f0000000000000000000000000000000000000000000000000000000000000000611c6f565b610acd7f0000000000000000000000000000000000000000000000000000000000000000611c6f565b610af67f0000000000000000000000000000000000000000000000000000000000000000611c6f565b604051602001610b089392919061252d565b604051602081830303815290604052905090565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639e45e8f46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b7a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b9e9190612499565b6001600160a01b0316336001600160a01b031614610c245760405162461bcd60e51b815260206004820152602660248201527f56616c696461746f72506f6f6c3a2073656e646572206973206e6f7420436f6c60448201527f6f737365756d00000000000000000000000000000000000000000000000000006064820152608401610598565b60008281526034602052604090208054427001000000000000000000000000000000009091046fffffffffffffffffffffffffffffffff161015610cd05760405162461bcd60e51b815260206004820152602e60248201527f56616c696461746f72506f6f6c3a20746865206f757470757420697320616c7260448201527f656164792066696e616c697a65640000000000000000000000000000000000006064820152608401610598565b80546fffffffffffffffffffffffffffffffff16610cee838261192c565b60008481526039602090815260408083206001600160a01b0387168085529083529281902080547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff8616908117909155905190815286917f2904258f32adf74dd8f23ad6f17ff50209896039c8ee3d4728ff55bd05c4cf2a91015b60405180910390a350505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639e45e8f46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610de7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e0b9190612499565b6001600160a01b0316336001600160a01b031614610e915760405162461bcd60e51b815260206004820152602660248201527f56616c696461746f72506f6f6c3a2073656e646572206973206e6f7420436f6c60448201527f6f737365756d00000000000000000000000000000000000000000000000000006064820152608401610598565b610e9b8282611b86565b5050565b6000610ea9611dac565b905080610f1e5760405162461bcd60e51b815260206004820152602960248201527f56616c696461746f72506f6f6c3a206e6f20626f6e6420746861742063616e2060448201527f626520756e626f6e6400000000000000000000000000000000000000000000006064820152608401610598565b50565b600054610100900460ff1615808015610f415750600054600160ff909116105b80610f5b5750303b158015610f5b575060005460ff166001145b610fcd5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610598565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561102b57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b611033611fd7565b8015610f1e57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b60008281526039602090815260408083206001600160a01b03851684529091528120546fffffffffffffffffffffffffffffffff16806111405760405162461bcd60e51b815260206004820152602e60248201527f56616c696461746f72506f6f6c3a207468652070656e64696e6720626f6e642060448201527f646f6573206e6f742065786973740000000000000000000000000000000000006064820152608401610598565b9392505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146111e55760405162461bcd60e51b815260206004820152602b60248201527f56616c696461746f72506f6f6c3a2073656e646572206973206e6f74204c324f60448201527f75747075744f7261636c650000000000000000000000000000000000000000006064820152608401610598565b7f0000000000000000000000000000000000000000000000000000000000000000826fffffffffffffffffffffffffffffffff16101561128d5760405162461bcd60e51b815260206004820152602b60248201527f56616c696461746f72506f6f6c3a2074686520626f6e6420616d6f756e74206960448201527f7320746f6f20736d616c6c0000000000000000000000000000000000000000006064820152608401610598565b6000838152603460205260409020805470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16156113365760405162461bcd60e51b815260206004820152603c60248201527f56616c696461746f72506f6f6c3a20626f6e64206f662074686520676976656e60448201527f206f757470757420696e64657820616c726561647920657869737473000000006064820152608401610598565b61133e611dac565b506040517fb0ea09a8000000000000000000000000000000000000000000000000000000008152600481018590526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063b0ea09a890602401602060405180830381865afa1580156113c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113e49190612499565b905061140281856fffffffffffffffffffffffffffffffff1661192c565b6fffffffffffffffffffffffffffffffff8481167001000000000000000000000000000000009185169182028117845560408051918252602082019290925286916001600160a01b038416917f5ca130257b8f76f72ad2965efcbe166f3918d820e4a49956e70081ea311f97c4910160405180910390a35050505050565b61148a3334611b86565b565b6040805180820190915260008082526020820152600082815260346020526040902080546fffffffffffffffffffffffffffffffff16158015906114f65750805470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff1615155b6115685760405162461bcd60e51b815260206004820152602660248201527f56616c696461746f72506f6f6c3a2074686520626f6e6420646f6573206e6f7460448201527f20657869737400000000000000000000000000000000000000000000000000006064820152608401610598565b6040805180820190915290546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000090910416602082015292915050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639e45e8f46040518163ffffffff1660e01b8152600401602060405180830381865afa158015611609573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061162d9190612499565b6001600160a01b0316336001600160a01b0316146116b35760405162461bcd60e51b815260206004820152602660248201527f56616c696461746f72506f6f6c3a2073656e646572206973206e6f7420436f6c60448201527f6f737365756d00000000000000000000000000000000000000000000000000006064820152608401610598565b60008281526034602052604090208054427001000000000000000000000000000000009091046fffffffffffffffffffffffffffffffff16101561175f5760405162461bcd60e51b815260206004820152602e60248201527f56616c696461746f72506f6f6c3a20746865206f757470757420697320616c7260448201527f656164792066696e616c697a65640000000000000000000000000000000000006064820152608401610598565b60008381526039602090815260408083206001600160a01b03861684529091529020546fffffffffffffffffffffffffffffffff16806118075760405162461bcd60e51b815260206004820152602e60248201527f56616c696461746f72506f6f6c3a207468652070656e64696e6720626f6e642060448201527f646f6573206e6f742065786973740000000000000000000000000000000000006064820152608401610598565b60008481526039602090815260408083206001600160a01b0387168085529083529281902080547fffffffffffffffffffffffffffffffff0000000000000000000000000000000090811690915585549081166fffffffffffffffffffffffffffffffff918216860182161786559051908416815286917f383f9b8b5a1fc2ec555726eb895621a312042e18b764135fa12ef1a520ad30db9101610d7b565b60365460009081036118ba57506000919050565b6001600160a01b0382166118d057506000919050565b6001600160a01b0382166000818152603760205260409020546036805491929183908110611900576119006125a3565b6000918252602090912001546001600160a01b0316149392505050565b6001600160a01b03163b151590565b6001600160a01b038216600090815260336020526040902054818110156119ba5760405162461bcd60e51b8152602060048201526024808201527f56616c696461746f72506f6f6c3a20696e73756666696369656e742062616c6160448201527f6e636573000000000000000000000000000000000000000000000000000000006064820152608401610598565b6119c48282612516565b90507f0000000000000000000000000000000000000000000000000000000000000000811080156119f957506119f9836118a6565b15611b4c57603654600090611a1090600190612516565b90508015611ac8576001600160a01b0384166000908152603760205260408120546036805491929184908110611a4857611a486125a3565b600091825260209091200154603680546001600160a01b039092169250829184908110611a7757611a776125a3565b600091825260208083209190910180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03948516179055929091168152603790915260409020555b6001600160a01b0384166000908152603760205260408120556036805480611af257611af26125d2565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055505b6001600160a01b0390921660009081526033602052604090209190915550565b600080600080845160208601878a8af19695505050505050565b6001600160a01b038216600090815260336020526040812054611baa9083906124fe565b90507f00000000000000000000000000000000000000000000000000000000000000008110158015611be25750611be0836118a6565b155b15611b4c57603680546001600160a01b03949094166000818152603760209081526040808320889055600188019094557f4a11f94e20a93c79f6ec743a1954ec4fc2c08429ae2122118bf234b2185c81b890960180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690921790915560339094529092209190915550565b606081600003611cb257505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115611cdc5780611cc681612601565b9150611cd59050600a83612668565b9150611cb6565b60008167ffffffffffffffff811115611cf757611cf761267c565b6040519080825280601f01601f191660200182016040528015611d21576020820181803683370190505b5090505b8415611da457611d36600183612516565b9150611d43600a866126ab565b611d4e9060306124fe565b60f81b818381518110611d6357611d636125a3565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611d9d600a86612668565b9450611d25565b949350505050565b60355460408051608081018252600080825260208201819052918101829052606081018290529091908290819060005b7f0000000000000000000000000000000000000000000000000000000000000000811015611fa757600085815260346020526040902080546fffffffffffffffffffffffffffffffff80821696509194507001000000000000000000000000000000009004164210801590611e6357506000846fffffffffffffffffffffffffffffffff16115b15611fa75760008581526034602052604080822091909155517fa25ae557000000000000000000000000000000000000000000000000000000008152600481018690527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063a25ae55790602401608060405180830381865afa158015611ef7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f1b91906126bf565b9150611f3d8260000151856fffffffffffffffffffffffffffffffff16611b86565b81516040516fffffffffffffffffffffffffffffffff861681526001600160a01b039091169086907f7047a0fb8bfae78c0ebbd4117437945bb85240453235ac4fd2e55712eb5bf0c39060200160405180910390a3611f9b8261205a565b60019485019401611ddc565b8015611fcb57611fba82602001516121cb565b505050603591909155506001919050565b60009550505050505090565b600054610100900460ff166120545760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610598565b60018055565b805160608201516040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169263c30af3889273420000000000000000000000000000000000000892620186a0927f21670f220000000000000000000000000000000000000000000000000000000092612104926024016001600160a01b039290921682526fffffffffffffffffffffffffffffffff16602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e086901b9092168252612196939291600401612762565b600060405180830381600087803b1580156121b057600080fd5b505af11580156121c4573d6000803e3d6000fd5b5050505050565b60365480156122a25760008183434160405160200161222293929190928352602083019190915260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016604082015260540190565b6040516020818303038152906040528051906020012060001c61224591906126ab565b90506036818154811061225a5761225a6125a3565b600091825260209091200154603880547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03909216919091179055505050565b603880547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555050565b6000602082840312156122e057600080fd5b5035919050565b6001600160a01b0381168114610f1e57600080fd5b60008060006060848603121561231157600080fd5b833592506020840135612323816122e7565b91506040840135612333816122e7565b809150509250925092565b60005b83811015612359578181015183820152602001612341565b83811115612368576000848401525b50505050565b6000815180845261238681602086016020860161233e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611140602083018461236e565b600080604083850312156123de57600080fd5b8235915060208301356123f0816122e7565b809150509250929050565b6000806040838503121561240e57600080fd5b8235612419816122e7565b946020939093013593505050565b60006020828403121561243957600080fd5b8135611140816122e7565b6fffffffffffffffffffffffffffffffff81168114610f1e57600080fd5b60008060006060848603121561247757600080fd5b83359250602084013561248981612444565b9150604084013561233381612444565b6000602082840312156124ab57600080fd5b8151611140816122e7565b6000602082840312156124c857600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115612511576125116124cf565b500190565b600082821015612528576125286124cf565b500390565b6000845161253f81846020890161233e565b80830190507f2e00000000000000000000000000000000000000000000000000000000000000808252855161257b816001850160208a0161233e565b6001920191820152835161259681600284016020880161233e565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612632576126326124cf565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261267757612677612639565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000826126ba576126ba612639565b500690565b6000608082840312156126d157600080fd5b6040516080810181811067ffffffffffffffff8211171561271b577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040528251612729816122e7565b815260208381015190820152604083015161274381612444565b6040820152606083015161275681612444565b60608201529392505050565b6001600160a01b038416815267ffffffffffffffff83166020820152606060408201526000612794606083018461236e565b9594505050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(ValidatorPoolStorageLayoutJSON), ValidatorPoolStorageLayout); err != nil {
		panic(err)
	}

	layouts["ValidatorPool"] = ValidatorPoolStorageLayout
	deployedBytecodes["ValidatorPool"] = ValidatorPoolDeployedBin
}
