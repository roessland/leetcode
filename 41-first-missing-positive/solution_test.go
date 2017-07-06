package leetcode41

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {
	assert.Equal(t, 3, firstMissingPositive([]int{1, 2, 0}))
	assert.Equal(t, 2, firstMissingPositive([]int{3, 4, -1, 1}))
	assert.Equal(t, 3, firstMissingPositive([]int{1, 2, 2, 0}))
	assert.Equal(t, 2, firstMissingPositive([]int{3, 4, 4, -1, 1}))
	assert.Equal(t, 1, firstMissingPositive([]int{}))
	assert.Equal(t, 1, firstMissingPositive([]int{5}))
	assert.Equal(t, 1, firstMissingPositive([]int{-1}))
	assert.Equal(t, 2, firstMissingPositive([]int{1}))
	assert.Equal(t, 3, firstMissingPositive([]int{2, 1}))
}
