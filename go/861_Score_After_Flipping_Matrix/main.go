package main

import "fmt"

func matrixScore(grid [][]int) int {
	// goal: ensure most significant values are 1s
	r, c := len(grid), len(grid[0])
	numOneInCol := make([]int, c)
	for i := 0; i < r; i++ {
		var needToFlip bool
		if grid[i][0] == 0 {
			needToFlip = true
		}
		for j := 0; j < c; j++ {
			if needToFlip {
				grid[i][j] = flip(grid[i][j])
			}
			if grid[i][j] == 1 {
				numOneInCol[j]++
			}
		}
	}
	for j := 0; j < c; j++ {
		if numOneInCol[j] <= r/2 {
			for i := 0; i < r; i++ {
				grid[i][j] = flip(grid[i][j])
			}
		}
	}
	res := 0
	x := 1
	for j := c - 1; j >= 0; j-- {
		for i := 0; i < r; i++ {
			res += grid[i][j] * x
		}
		x *= 2
	}
	return res
}

func flip(n int) int {
	if n == 0 {
		return 1
	}
	return 0
}

func main() {
	//grid := [][]int{
	//	{0, 0, 1, 1},
	//	{1, 0, 1, 0},
	//	{1, 1, 0, 0},
	//}
	grid := [][]int{
		{0, 1},
		{0, 1},
		{0, 1},
		{0, 0},
	}
	// 11
	// 11
	// 11
	// 10
	fmt.Println(matrixScore(grid))
}
