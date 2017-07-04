package leetcode63

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	assert.Equal(t, 2, uniquePathsWithObstacles([][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}))

	assert.Equal(t, 6, uniquePathsWithObstacles([][]int{
		[]int{0, 0, 0},
		[]int{0, 0, 0},
		[]int{0, 0, 0},
	}))

	assert.Equal(t, 0, uniquePathsWithObstacles([][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 1},
		[]int{0, 1, 0},
	}))

	assert.Equal(t, 3, uniquePathsWithObstacles([][]int{
		[]int{0, 1, 0},
		[]int{0, 0, 0},
		[]int{0, 0, 0},
	}))

	assert.Equal(t, 1, uniquePathsWithObstacles([][]int{
		[]int{0},
		[]int{0},
		[]int{0},
	}))

	assert.Equal(t, 1, uniquePathsWithObstacles([][]int{
		[]int{0, 0, 0},
	}))

	assert.Equal(t, 4, uniquePathsWithObstacles([][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
		[]int{0, 0, 0},
	}))
}
