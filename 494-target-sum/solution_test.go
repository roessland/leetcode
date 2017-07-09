package leetcode494

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTargetSumWays(t *testing.T) {
	assert.Equal(t, 5, FindTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	assert.Equal(t, 1, FindTargetSumWays([]int{3}, 3))
	assert.Equal(t, 1, FindTargetSumWays([]int{3}, -3))
	assert.Equal(t, 1, FindTargetSumWays([]int{-3}, 3))
	assert.Equal(t, 1, FindTargetSumWays([]int{5, 5}, 10))
	assert.Equal(t, 2, FindTargetSumWays([]int{5, 5}, 0))
	assert.Equal(t, 5656, FindTargetSumWays([]int{33, 36, 38, 40, 25, 49, 1, 8, 50, 13, 41, 50, 29, 27, 18, 25, 37, 8, 0, 48}, 0))
	assert.Equal(t, 2, FindTargetSumWays([]int{5, 0}, 5))
}

func BenchmarkSlowExponentialFindTargetSumWays(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = SlowExponentialFindTargetSumWays([]int{33, 36, 38, 40, 25, 49, 1, 8, 50, 13, 41, 50, 29, 27, 18, 25, 37, 8, 0, 48}, 0)
	}
}

func BenchmarkFindTargetSumWays(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = FindTargetSumWays([]int{33, 36, 38, 40, 25, 49, 1, 8, 50, 13, 41, 50, 29, 27, 18, 25, 37, 8, 0, 48}, 0)
	}
}
