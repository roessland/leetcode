package leetcode62

import "math/big"

func uniquePaths(m int, n int) int {
	k := big.NewInt(0)
	k.Binomial(int64(m+n-2), int64(m-1))
	return int(k.Int64())
}
