package leetcode53

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	assert.Equal(t, 6, MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
