package driver

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GetContractBin is NodeDriver contract genesis implementation bin code
// Has to be compiled with flag bin-runtime
// Built from forest-sfc c1d33c81f74abf82c0e22807f16e609578e10ad8, solc 0.5.17+commit.d19bba13.Emscripten.clang, optimize-runs 10000
func GetContractBin() []byte {
	return hexutil.MustDecode("0x608060405234801561001057600080fd5b506004361061011b5760003560e01c80634feb92f3116100b2578063d6a0c7af11610081578063e08d7e6611610066578063e08d7e66146104f5578063e30443bc14610565578063ebdf104c1461059e5761011b565b8063d6a0c7af14610487578063da7fc24f146104c25761011b565b80634feb92f31461031057806379bead38146103bb578063a4066fbe146103f4578063b9cc6b1c146104175761011b565b8063242a6e3f116100ee578063242a6e3f14610202578063267ab4461461027957806339e503ab14610296578063485cc955146102d55761011b565b806307690b2a146101205780630aeeca001461015d57806318f628d41461017a5780631e702f83146101df575b600080fd5b61015b6004803603604081101561013657600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020013516610704565b005b61015b6004803603602081101561017357600080fd5b5035610808565b61015b600480360361012081101561019157600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060208101359060408101359060608101359060808101359060a08101359060c08101359060e08101359061010001356108aa565b61015b600480360360408110156101f557600080fd5b50803590602001356109ce565b61015b6004803603604081101561021857600080fd5b8135919081019060408101602082013564010000000081111561023a57600080fd5b82018360208201111561024c57600080fd5b8035906020019184600183028401116401000000008311171561026e57600080fd5b509092509050610a9b565b61015b6004803603602081101561028f57600080fd5b5035610b8b565b61015b600480360360608110156102ac57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060208101359060400135610c2d565b61015b600480360360408110156102eb57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020013516610d38565b61015b600480360361010081101561032757600080fd5b73ffffffffffffffffffffffffffffffffffffffff8235169160208101359181019060608101604082013564010000000081111561036457600080fd5b82018360208201111561037657600080fd5b8035906020019184600183028401116401000000008311171561039857600080fd5b919350915080359060208101359060408101359060608101359060800135610ee0565b61015b600480360360408110156103d157600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135611038565b61015b6004803603604081101561040a57600080fd5b508035906020013561111f565b61015b6004803603602081101561042d57600080fd5b81019060208101813564010000000081111561044857600080fd5b82018360208201111561045a57600080fd5b8035906020019184600183028401116401000000008311171561047c57600080fd5b5090925090506111c5565b61015b6004803603604081101561049d57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160200135166112b3565b61015b600480360360208110156104d857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661139b565b61015b6004803603602081101561050b57600080fd5b81019060208101813564010000000081111561052657600080fd5b82018360208201111561053857600080fd5b8035906020019184602083028401116401000000008311171561055a57600080fd5b50909250905061148f565b61015b6004803603604081101561057b57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135611585565b61015b600480360360808110156105b457600080fd5b8101906020810181356401000000008111156105cf57600080fd5b8201836020820111156105e157600080fd5b8035906020019184602083028401116401000000008311171561060357600080fd5b91939092909160208101903564010000000081111561062157600080fd5b82018360208201111561063357600080fd5b8035906020019184602083028401116401000000008311171561065557600080fd5b91939092909160208101903564010000000081111561067357600080fd5b82018360208201111561068557600080fd5b803590602001918460208302840111640100000000831117156106a757600080fd5b9193909290916020810190356401000000008111156106c557600080fd5b8201836020820111156106d757600080fd5b803590602001918460208302840111640100000000831117156106f957600080fd5b50909250905061166c565b60345473ffffffffffffffffffffffffffffffffffffffff163314610770576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b603554604080517f07690b2a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301528481166024830152915191909216916307690b2a91604480830192600092919082900301818387803b1580156107ec57600080fd5b505af1158015610800573d6000803e3d6000fd5b505050505050565b60345473ffffffffffffffffffffffffffffffffffffffff163314610874576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b6040805182815290517f0151256d62457b809bbc891b1f81c6dd0b9987552c70ce915b519750cd434dd19181900360200190a150565b33156108fd576040805162461bcd60e51b815260206004820152600c60248201527f6e6f742063616c6c61626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b603454604080517f18f628d400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8c81166004830152602482018c9052604482018b9052606482018a90526084820189905260a4820188905260c4820187905260e482018690526101048201859052915191909216916318f628d49161012480830192600092919082900301818387803b1580156109ab57600080fd5b505af11580156109bf573d6000803e3d6000fd5b50505050505050505050505050565b3315610a21576040805162461bcd60e51b815260206004820152600c60248201527f6e6f742063616c6c61626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b603454604080517f1e702f830000000000000000000000000000000000000000000000000000000081526004810185905260248101849052905173ffffffffffffffffffffffffffffffffffffffff90921691631e702f839160448082019260009290919082900301818387803b1580156107ec57600080fd5b60345473ffffffffffffffffffffffffffffffffffffffff163314610b07576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b827f0f0ef1ab97439def0a9d2c6d9dc166207f1b13b99e62b442b2993d6153c63a6e838360405180806020018281038252848482818152602001925080828437600083820152604051601f9091017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169092018290039550909350505050a2505050565b60345473ffffffffffffffffffffffffffffffffffffffff163314610bf7576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b6040805182815290517f2ccdfd47cf0c1f1069d949f1789bb79b2f12821f021634fc835af1de66ea2feb9181900360200190a150565b60345473ffffffffffffffffffffffffffffffffffffffff163314610c99576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b603554604080517f39e503ab00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301526024820186905260448201859052915191909216916339e503ab91606480830192600092919082900301818387803b158015610d1b57600080fd5b505af1158015610d2f573d6000803e3d6000fd5b50505050505050565b600054610100900460ff1680610d515750610d5161186f565b80610d5f575060005460ff16155b610d9a5760405162461bcd60e51b815260040180806020018281038252602e815260200180611876602e913960400191505060405180910390fd5b600054610100900460ff16158015610e0057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b603480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091556040517f64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef24503583069390600090a2603580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790558015610edb57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b505050565b3315610f33576040805162461bcd60e51b815260206004820152600c60248201527f6e6f742063616c6c61626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634feb92f38a8a8a8a8a8a8a8a8a6040518a63ffffffff1660e01b8152600401808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001898152602001806020018781526020018681526020018581526020018481526020018381526020018281038252898982818152602001925080828437600081840152601f19601f8201169050808301925050509a5050505050505050505050600060405180830381600087803b1580156109ab57600080fd5b60345473ffffffffffffffffffffffffffffffffffffffff1633146110a4576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b603554604080517f79bead3800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015260248201859052915191909216916379bead3891604480830192600092919082900301818387803b1580156107ec57600080fd5b60345473ffffffffffffffffffffffffffffffffffffffff16331461118b576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b60408051828152905183917fb975807576e3b1461be7de07ebf7d20e4790ed802d7a0c4fdd0a1a13df72a935919081900360200190a25050565b60345473ffffffffffffffffffffffffffffffffffffffff163314611231576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b7f47d10eed096a44e3d0abc586c7e3a5d6cb5358cc90e7d437cd0627f7e765fb99828260405180806020018281038252848482818152602001925080828437600083820152604051601f9091017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169092018290039550909350505050a15050565b60345473ffffffffffffffffffffffffffffffffffffffff16331461131f576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b603554604080517fd6a0c7af00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015284811660248301529151919092169163d6a0c7af91604480830192600092919082900301818387803b1580156107ec57600080fd5b60345473ffffffffffffffffffffffffffffffffffffffff163314611407576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b60405173ffffffffffffffffffffffffffffffffffffffff8216907f64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef24503583069390600090a2603480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b33156114e2576040805162461bcd60e51b815260206004820152600c60248201527f6e6f742063616c6c61626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b6034546040517fe08d7e660000000000000000000000000000000000000000000000000000000081526020600482018181526024830185905273ffffffffffffffffffffffffffffffffffffffff9093169263e08d7e6692869286929182916044909101908590850280828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b1580156107ec57600080fd5b60345473ffffffffffffffffffffffffffffffffffffffff1633146115f1576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b603554604080517fe30443bc00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590529151919092169163e30443bc91604480830192600092919082900301818387803b1580156107ec57600080fd5b33156116bf576040805162461bcd60e51b815260206004820152600c60248201527f6e6f742063616c6c61626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ebdf104c89898989898989896040518963ffffffff1660e01b8152600401808060200180602001806020018060200185810385528d8d82818152602001925060200280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01690910186810385528b8152602090810191508c908c0280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169091018681038452898152602090810191508a908a0280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169091018681038352878152602090810191508890880280828437600081840152601f19601f8201169050808301925050509c50505050505050505050505050600060405180830381600087803b15801561184d57600080fd5b505af1158015611861573d6000803e3d6000fd5b505050505050505050505050565b303b159056fe436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564a265627a7a72315820b2d01428f431c4a18d5c62cd2dd3d6b7c2a7b6283e76ffa2526a2819b16a56aa64736f6c63430005110032")
}

// ContractAddress is the NodeDriver contract address
var ContractAddress = common.HexToAddress("0xd100a01e00000000000000000000000000000000")
