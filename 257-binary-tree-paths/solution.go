package leetcode257

import (
	"strconv"
)

// TreeNode represents a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Really ugly code. I was tired.
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	paths := []string{}
	var traverse func(*TreeNode, string)
	traverse = func(node *TreeNode, currPath string) {
		if node == nil {
		} else if node.Left == nil && node.Right == nil {
			paths = append(paths, currPath+"->"+strconv.Itoa(node.Val))
		} else {
			traverse(node.Left, currPath+"->"+strconv.Itoa(node.Val))
			traverse(node.Right, currPath+"->"+strconv.Itoa(node.Val))
		}
	}
	traverse(root.Left, strconv.Itoa(root.Val))
	traverse(root.Right, strconv.Itoa(root.Val))
	return paths
}
