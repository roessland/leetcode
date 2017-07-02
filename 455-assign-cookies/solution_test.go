package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindContentChildren(t *testing.T) {
	assert.Equal(t, 1, findContentChildren([]int{1,2,3}, []int{1,1}))
	assert.Equal(t, 2, findContentChildren([]int{1,2}, []int{1,2,3}))
	assert.Equal(t, 0, findContentChildren([]int{}, []int{1,2,3}))
	assert.Equal(t, 0, findContentChildren([]int{1,2}, []int{}))
	assert.Equal(t, 0, findContentChildren([]int{}, []int{}))
}