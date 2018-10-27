package leetcode617

import "testing"
import "github.com/stretchr/testify/assert"

func TestMergeTrees(t *testing.T) {
	t1 := &TreeNode{Val: 1}
	t1.Left = &TreeNode{Val: 3}
	t1.Right = &TreeNode{Val: 2}
	t1.Left.Left = &TreeNode{Val: 5}

	t2 := &TreeNode{Val: 2}
	t2.Left = &TreeNode{Val: 1}
	t2.Right = &TreeNode{Val: 3}
	t2.Left.Right = &TreeNode{Val: 4}
	t2.Right.Right = &TreeNode{Val: 7}

	t3 := mergeTrees(t1, t2)

	assert.Equal(t, 3, t3.Val)
	assert.Equal(t, 4, t3.Left.Val)
	assert.Equal(t, 4, t3.Left.Right.Val)
	assert.Equal(t, 7, t3.Right.Right.Val)
}
