package utils

import "math/big"

// ToTree number of TREE to Wei
func ToTree(tree uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(tree), big.NewInt(1e18))
}
