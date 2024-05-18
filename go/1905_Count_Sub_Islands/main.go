package main

import "fmt"

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	visited := make([][]bool, len(grid2))
	for i := 0; i < len(grid2); i++ {
		visited[i] = make([]bool, len(grid2[0]))
	}
	res := 0
	for i := 0; i < len(grid2); i++ {
		for j := 0; j < len(grid2[0]); j++ {
			if !visited[i][j] && grid2[i][j] == 1 {
				if tour([2]int{i, j}, grid2, grid1, visited) {
					res += 1
				}
			}
		}
	}
	return res
}

func tour(pos [2]int, grid2, grid1 [][]int, visited [][]bool) bool {
	x := pos[0]
	y := pos[1]
	if x < 0 || x == len(grid2) || y < 0 || y == len(grid2[0]) || visited[x][y] || grid2[x][y] == 0 {
		return true
	}
	visited[x][y] = true
	valid := grid1[x][y] == 1
	// left, right, up, down
	l := tour([2]int{x - 1, y}, grid2, grid1, visited)
	r := tour([2]int{x + 1, y}, grid2, grid1, visited)
	u := tour([2]int{x, y - 1}, grid2, grid1, visited)
	d := tour([2]int{x, y + 1}, grid2, grid1, visited)
	return valid && l && r && u && d
}

func main() {
	grid1 := [][]int{
		{1, 1, 1, 0, 0},
		{0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 1, 1},
	}
	grid2 := [][]int{
		{1, 1, 1, 0, 0},
		{0, 0, 1, 1, 1},
		{0, 1, 0, 0, 0},
		{1, 0, 1, 1, 0},
		{0, 1, 0, 1, 0},
	}
	fmt.Println(countSubIslands(grid1, grid2))
}
