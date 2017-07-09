package leetcode437

// TreeNode represents a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSumStartingAtNode(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}
	ret := 0
	if sum == node.Val {
		ret = 1
	}
	return ret + pathSumStartingAtNode(node.Left, sum-node.Val) + pathSumStartingAtNode(node.Right, sum-node.Val)
}

// PathSum finds the number of paths in the tree that sum to a given value.
// Paths can start anywhere and go downwards to anywhere.
func PathSum(root *TreeNode, sum int) int {
	var traverse func(*TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return pathSumStartingAtNode(node, sum) + traverse(node.Left) + traverse(node.Right)
	}
	return traverse(root)
}

var pathSum = PathSum
