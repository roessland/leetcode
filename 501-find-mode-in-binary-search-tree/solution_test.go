package leetcode501

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMode(t *testing.T) {
	var tree *TreeNode

	tree = &TreeNode{5, &TreeNode{5, nil, nil}, nil}
	assert.Equal(t, []int{5}, FindMode(tree))

	tree = &TreeNode{5, &TreeNode{5, nil, nil}, &TreeNode{4, nil, &TreeNode{4, nil, nil}}}
	assert.Equal(t, []int{5, 4}, FindMode(tree))

	tree = nil
	assert.Equal(t, []int(nil), FindMode(tree))
}
