package leetcode807

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	numX := len(grid[0])
	numY := len(grid)
	eastSkyline := make([]int, numY)
	northSkyline := make([]int, numX)

	for y := range grid {
		for x, height := range grid[y] {
			eastSkyline[y] = Max(eastSkyline[y], height)
			northSkyline[x] = Max(northSkyline[x], height)
		}
	}

	addedHeight := 0
	for y := range grid {
		for x, height := range grid[y] {
			newHeight := Min(eastSkyline[y], northSkyline[x])
			if newHeight > height {
				addedHeight += newHeight - height
			}
		}
	}
	return addedHeight
}
