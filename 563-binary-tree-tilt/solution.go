package leetcode563

// TreeNode represents a tree. Thank you go-linter, for making me type this useless comment.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func findTilt(root *TreeNode) int {
	tilt := 0
	var traverse func(*TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum, rightSum := traverse(node.Left), traverse(node.Right)
		tilt += abs(leftSum - rightSum)
		return leftSum + rightSum + node.Val
	}
	traverse(root)
	return tilt
}
