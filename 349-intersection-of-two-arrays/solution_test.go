package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	assert.Equal(t, []int{2}, intersection([]int{1, 2, 2, 3}, []int{2, 5, 6}))
	assert.Equal(t, []int{2}, intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	assert.Equal(t, []int{}, intersection([]int{1, 2, 2, 1}, []int{}))
	assert.Equal(t, []int{}, intersection([]int{}, []int{2, 2}))
}
