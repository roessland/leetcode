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
	if len(grid) == 0 {
		return 0
	}
	eastSkyline := make([]int, len(grid))
	northSkyline := make([]int, len(grid[0]))

	for y := range grid {
		for x, height := range grid[y] {
			eastSkyline[y] = Max(eastSkyline[y], height)
			northSkyline[x] = Max(northSkyline[x], height)
		}
	}

	addedHeight := 0
	for y := range grid {
		for x, height := range grid[y] {
			addedHeight += Min(eastSkyline[y], northSkyline[x]) - height
		}
	}
	return addedHeight
}
