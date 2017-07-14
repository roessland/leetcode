package leetcode501

import (
	"sort"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type KeyVal struct {
	Key int
	Val int
}

// FindMode finds the numbers in the tree with the highest frequency
func FindMode(root *TreeNode) []int {
	frequencies := map[int]int{}
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		frequencies[node.Val]++
		traverse(node.Left)
		traverse(node.Right)
	}
	traverse(root)
	return ArgMax(frequencies)
}

// ArgMax finds the key(s) in a map with the highest value
func ArgMax(m map[int]int) []int {
	if len(m) == 0 {
		return nil
	}
	keyvals := []KeyVal{}
	for key, val := range m {
		keyvals = append(keyvals, KeyVal{key, val})
	}
	sort.Slice(keyvals, func(i, j int) bool {
		return keyvals[i].Val > keyvals[j].Val
	})

	maxval := keyvals[0].Val
	maxkeys := []int{}
	for _, keyval := range keyvals {
		if keyval.Val == maxval {
			maxkeys = append(maxkeys, keyval.Key)
		} else {
			break
		}
	}
	return maxkeys
}

var findMode = FindMode
