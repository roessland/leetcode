package leetcode654

import "testing"

import "github.com/stretchr/testify/assert"

func TestConstructMaximumBinaryTree(t *testing.T) {
	nums := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree(nums)
	assert.Equal(t, 6, root.Val)
	assert.Equal(t, 3, root.Left.Val)
	assert.Equal(t, 2, root.Left.Right.Val)
	assert.Equal(t, 1, root.Left.Right.Right.Val)
	assert.Equal(t, 5, root.Right.Val)
	assert.Equal(t, 0, root.Right.Left.Val)
}
