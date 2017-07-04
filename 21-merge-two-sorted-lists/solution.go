package leetcode21

import (
	"math"
)

// ListNode represents a node in a singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists rearranges two singly linked sorted lists into a single sorted lists
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	H := &ListNode{math.MinInt32, nil}
	T := H
	for {
		if l1 == nil {
			T.Next = l2
			break
		}
		if l2 == nil {
			T.Next = l1
			break
		}
		if l1.Val < l2.Val {
			T.Next, T, l1 = l1, l1, l1.Next
		} else {
			T.Next, T, l2 = l2, l2, l2.Next
		}
	}
	return H.Next
}
