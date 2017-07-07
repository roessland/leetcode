package leetcode66

import "testing"
import "github.com/stretchr/testify/assert"

func TestPlusOne(t *testing.T) {
	assert.Equal(t, []int{1}, PlusOne([]int{0}))
	assert.Equal(t, []int{1, 1}, PlusOne([]int{1, 0}))
	assert.Equal(t, []int{1, 0}, PlusOne([]int{9}))
	assert.Equal(t, []int{1, 0, 0}, PlusOne([]int{9, 9}))
	assert.Equal(t, []int{1, 5, 0}, PlusOne([]int{1, 4, 9}))
}
