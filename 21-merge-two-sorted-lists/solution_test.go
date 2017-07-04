package leetcode21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newLinkedList(nums []int) *ListNode {
	var next, head *ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		head = &ListNode{nums[i], next}
		next = head
	}
	return head
}

func toSlice(l *ListNode) []int {
	s := []int{}
	for l != nil {
		s = append(s, l.Val)
		l = l.Next
	}
	return s
}

func TestNewLinkedList(t *testing.T) {
	l := newLinkedList([]int{1, 2, 3})
	assert.Equal(t, 1, l.Val)
	assert.Equal(t, 2, l.Next.Val)
	assert.Equal(t, 3, l.Next.Next.Val)
	assert.Nil(t, l.Next.Next.Next)
}

func TestMergeTwoLists(t *testing.T) {
	assert.Equal(t, []int{1}, toSlice(mergeTwoLists(newLinkedList([]int{1}), newLinkedList([]int{}))))
	assert.Equal(t, []int{1, 1}, toSlice(mergeTwoLists(newLinkedList([]int{1}), newLinkedList([]int{1}))))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, toSlice(mergeTwoLists(newLinkedList([]int{1, 2, 4}), newLinkedList([]int{3, 5}))))
}
