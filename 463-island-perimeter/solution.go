package leetcode463

func isWater(grid [][]int, i, j int) bool {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return true
	}
	return grid[i][j] == 0
}

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if isWater(grid, i, j) {
				continue
			}
			if isWater(grid, i-1, j) {
				perimeter++
			}
			if isWater(grid, i+1, j) {
				perimeter++
			}
			if isWater(grid, i, j-1) {
				perimeter++
			}
			if isWater(grid, i, j+1) {
				perimeter++
			}
		}
	}
	return perimeter
}
