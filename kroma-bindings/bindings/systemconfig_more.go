// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const SystemConfigStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1003,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1005,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"overhead\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"scalar\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_bytes32\"},{\"astId\":1008,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"gasLimit\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint64\"},{\"astId\":1009,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_resourceConfig\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_struct(ResourceConfig)1011_storage\"},{\"astId\":1010,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"validatorRewardScalar\",\"offset\":0,\"slot\":\"106\",\"type\":\"t_uint256\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_struct(ResourceConfig)1011_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ResourceMetering.ResourceConfig\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SystemConfigStorageLayout = new(solc.StorageLayout)

var SystemConfigDeployedBin = "0x608060405234801561001057600080fd5b50600436106101775760003560e01c8063c71973f6116100d8578063ed579ad31161008c578063f68016b711610066578063f68016b71461046d578063fc7ffea714610481578063ffa1ad741461049457600080fd5b8063ed579ad314610448578063f2fde38b14610451578063f45e65d81461046457600080fd5b8063cc731b02116100bd578063cc731b02146102f8578063e81b2c6d1461042c578063ecdd939d1461043557600080fd5b8063c71973f6146102d2578063c9b26f61146102e557600080fd5b806354fd4d501161012f5780638da5cb5b116101145780638da5cb5b1461028e578063935f029e146102ac578063b40a817c146102bf57600080fd5b806354fd4d501461023d578063715018a61461028657600080fd5b80631fd19ee1116101605780631fd19ee1146101ad5780634add321d146101f55780634f16540b1461021657600080fd5b80630c18c1621461017c57806318d1391814610198575b600080fd5b61018560655481565b6040519081526020015b60405180910390f35b6101ab6101a6366004611139565b61049c565b005b7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c08545b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161018f565b6101fd610560565b60405167ffffffffffffffff909116815260200161018f565b6101857f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0881565b6102796040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161018f91906111c6565b6101ab61058b565b60335473ffffffffffffffffffffffffffffffffffffffff166101d0565b6101ab6102ba3660046111d9565b61059f565b6101ab6102cd366004611213565b610638565b6101ab6102e036600461133c565b610709565b6101ab6102f3366004611358565b61071d565b6103bc6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c08101825260695463ffffffff8082168352640100000000820460ff9081166020850152650100000000008304169383019390935266010000000000008104831660608301526a0100000000000000000000810490921660808201526e0100000000000000000000000000009091046fffffffffffffffffffffffffffffffff1660a082015290565b60405161018f9190600060c08201905063ffffffff80845116835260ff602085015116602084015260ff6040850151166040840152806060850151166060840152806080850151166080840152506fffffffffffffffffffffffffffffffff60a08401511660a083015292915050565b61018560675481565b6101ab610443366004611358565b61074d565b610185606a5481565b6101ab61045f366004611139565b61081b565b61018560665481565b6068546101fd9067ffffffffffffffff1681565b6101ab61048f366004611371565b6108b5565b610185600081565b6104a4610b3a565b6104cc817f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0855565b6040805173ffffffffffffffffffffffffffffffffffffffff8316602082015260009101604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052905060035b60007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be8360405161055491906111c6565b60405180910390a35050565b6069546000906105869063ffffffff6a010000000000000000000082048116911661141d565b905090565b610593610b3a565b61059d6000610ba1565b565b6105a7610b3a565b606582905560668190556040805160208101849052908101829052600090606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529050600160007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be8360405161062b91906111c6565b60405180910390a3505050565b610640610b3a565b610648610560565b67ffffffffffffffff168167ffffffffffffffff1610156106b05760405162461bcd60e51b815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064015b60405180910390fd5b606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff83169081179091556040805160208082019390935281518082039093018352810190526002610523565b610711610b3a565b61071a81610c18565b50565b610725610b3a565b6067819055604080516020808201849052825180830390910181529082019091526000610523565b610755610b3a565b6127108111156107f35760405162461bcd60e51b815260206004820152604860248201527f53797374656d436f6e6669673a20746865206d61782076616c7565206f66207660448201527f616c696461746f7220726577617264207363616c617220686173206265656e2060648201527f6578636565646564000000000000000000000000000000000000000000000000608482015260a4016106a7565b606a819055604080516020808201849052825180830390910181529082019091526004610523565b610823610b3a565b73ffffffffffffffffffffffffffffffffffffffff81166108ac5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016106a7565b61071a81610ba1565b600054610100900460ff16158080156108d55750600054600160ff909116105b806108ef5750303b1580156108ef575060005460ff166001145b6109615760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016106a7565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156109bf57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6109c761100a565b6109d08961081b565b606588905560668790556067869055606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff87161790557f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c08849055610a4083610c18565b610a48610560565b67ffffffffffffffff168567ffffffffffffffff161015610aab5760405162461bcd60e51b815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016106a7565b606a8290558015610b1357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b60335473ffffffffffffffffffffffffffffffffffffffff16331461059d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016106a7565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b8060a001516fffffffffffffffffffffffffffffffff16816060015163ffffffff161115610cae5760405162461bcd60e51b815260206004820152603560248201527f53797374656d436f6e6669673a206d696e206261736520666565206d7573742060448201527f6265206c657373207468616e206d61782062617365000000000000000000000060648201526084016106a7565b6001816040015160ff1611610d2b5760405162461bcd60e51b815260206004820152602f60248201527f53797374656d436f6e6669673a2064656e6f6d696e61746f72206d757374206260448201527f65206c6172676572207468616e2031000000000000000000000000000000000060648201526084016106a7565b6068546080820151825167ffffffffffffffff90921691610d4c9190611449565b63ffffffff161115610da05760405162461bcd60e51b815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016106a7565b6000816020015160ff1611610e1d5760405162461bcd60e51b815260206004820152602f60248201527f53797374656d436f6e6669673a20656c6173746963697479206d756c7469706c60448201527f6965722063616e6e6f742062652030000000000000000000000000000000000060648201526084016106a7565b8051602082015163ffffffff82169160ff90911690610e3d908290611468565b610e4791906114b2565b63ffffffff1614610ec05760405162461bcd60e51b815260206004820152603760248201527f53797374656d436f6e6669673a20707265636973696f6e206c6f73732077697460448201527f6820746172676574207265736f75726365206c696d697400000000000000000060648201526084016106a7565b805160698054602084015160408501516060860151608087015160a09097015163ffffffff9687167fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000009095169490941764010000000060ff94851602177fffffffffffffffffffffffffffffffffffffffffffff0000000000ffffffffff166501000000000093909216929092027fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff1617660100000000000091851691909102177fffff0000000000000000000000000000000000000000ffffffffffffffffffff166a010000000000000000000093909416929092027fffff00000000000000000000000000000000ffffffffffffffffffffffffffff16929092176e0100000000000000000000000000006fffffffffffffffffffffffffffffffff90921691909102179055565b600054610100900460ff166110875760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106a7565b61059d600054610100900460ff166111075760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106a7565b61059d33610ba1565b803573ffffffffffffffffffffffffffffffffffffffff8116811461113457600080fd5b919050565b60006020828403121561114b57600080fd5b61115482611110565b9392505050565b6000815180845260005b8181101561118157602081850181015186830182015201611165565b81811115611193576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611154602083018461115b565b600080604083850312156111ec57600080fd5b50508035926020909101359150565b803567ffffffffffffffff8116811461113457600080fd5b60006020828403121561122557600080fd5b611154826111fb565b803563ffffffff8116811461113457600080fd5b803560ff8116811461113457600080fd5b80356fffffffffffffffffffffffffffffffff8116811461113457600080fd5b600060c0828403121561128557600080fd5b60405160c0810181811067ffffffffffffffff821117156112cf577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040529050806112de8361122e565b81526112ec60208401611242565b60208201526112fd60408401611242565b604082015261130e6060840161122e565b606082015261131f6080840161122e565b608082015261133060a08401611253565b60a08201525092915050565b600060c0828403121561134e57600080fd5b6111548383611273565b60006020828403121561136a57600080fd5b5035919050565b6000806000806000806000806101a0898b03121561138e57600080fd5b61139789611110565b97506020890135965060408901359550606089013594506113ba60808a016111fb565b93506113c860a08a01611110565b92506113d78a60c08b01611273565b915061018089013590509295985092959890939650565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff808316818516808303821115611440576114406113ee565b01949350505050565b600063ffffffff808316818516808303821115611440576114406113ee565b600063ffffffff808416806114a6577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600063ffffffff808316818516818304811182151516156114d5576114d56113ee565b0294935050505056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(SystemConfigStorageLayoutJSON), SystemConfigStorageLayout); err != nil {
		panic(err)
	}

	layouts["SystemConfig"] = SystemConfigStorageLayout
	deployedBytecodes["SystemConfig"] = SystemConfigDeployedBin
	immutableReferences["SystemConfig"] = false
}
