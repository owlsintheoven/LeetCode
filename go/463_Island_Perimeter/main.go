package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	h := 0
	v := 0
	// going horizontally
	for i := 0; i < len(grid); i++ {
		j := 0
		for j < len(grid[0]) {
			if grid[i][j] == 1 {
				wide := 1
				for j+wide < len(grid[0]) && grid[i][j+wide] == 1 {
					wide += 1
				}
				j += wide
				h += 2
			} else {
				j += 1
			}
		}
	}
	// going vertically
	for j := 0; j < len(grid[0]); j++ {
		i := 0
		for i < len(grid) {
			if grid[i][j] == 1 {
				wide := 1
				for i+wide < len(grid) && grid[i+wide][j] == 1 {
					wide += 1
				}
				i += wide
				v += 2
			} else {
				i += 1
			}
		}
	}
	return h + v
}

func main() {
	grid := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	fmt.Println(islandPerimeter(grid))
}
