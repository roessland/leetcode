package leetcode807

import "testing"
import "github.com/stretchr/testify/assert"

func TestMaxIncreaseKeepingSkyline(t *testing.T) {
	grid := [][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}
	assert.Equal(t, 35, maxIncreaseKeepingSkyline(grid))
}

func TestEmpty(t *testing.T) {
	assert.Equal(t, 0, maxIncreaseKeepingSkyline([][]int{}))
	assert.Equal(t, 0, maxIncreaseKeepingSkyline([][]int{{}, {}}))
}
