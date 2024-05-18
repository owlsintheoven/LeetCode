package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// dp[i][j] represents the number of unique paths from obstacleGrid[0][0] to obstacleGrid[i][j]
	// dp[i][j] = dp[i][j-1] + dp[i-1][j] (from the top and from the left)
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	reachedObstacle := false
	for i := 0; i < m; i++ {
		if !reachedObstacle {
			if obstacleGrid[i][0] == 0 {
				dp[i][0] = 1
			} else {
				dp[i][0] = 0
				reachedObstacle = true
			}
		} else {
			dp[i][0] = 0
		}
	}
	reachedObstacle = false
	for j := 0; j < n; j++ {
		if !reachedObstacle {
			if obstacleGrid[0][j] == 0 {
				dp[0][j] = 1
			} else {
				dp[0][j] = 0
				reachedObstacle = true
			}
		} else {
			dp[0][j] = 0
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else {
				dp[i][j] = 0
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	//obstacleGrid := [][]int{
	//	{0, 0, 0},
	//	{0, 1, 0},
	//	{0, 0, 0},
	//}
	obstacleGrid := [][]int{
		{0, 1},
		{0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
