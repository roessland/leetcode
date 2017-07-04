package leetcode63

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	numRows := len(obstacleGrid)
	numCols := len(obstacleGrid[0])
	pathCounts := make([][]int, numRows)
	for row := 0; row < numRows; row++ {
		pathCounts[row] = make([]int, numCols)
	}
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if obstacleGrid[row][col] == 1 {
				continue
			}
			if row == 0 && col == 0 {
				pathCounts[0][0] = 1
			}
			if row != 0 {
				pathCounts[row][col] += pathCounts[row-1][col]
			}
			if col != 0 {
				pathCounts[row][col] += pathCounts[row][col-1]
			}
		}
	}
	return pathCounts[numRows-1][numCols-1]
}
