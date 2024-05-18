package main

import "fmt"

func numIslands(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				tour([2]int{i, j}, grid, visited)
				res += 1
			}
		}
	}
	return res
}

func tour(pos [2]int, grid [][]byte, visited [][]bool) {
	x := pos[0]
	y := pos[1]
	if x < 0 || x == len(grid) || y < 0 || y == len(grid[0]) || visited[x][y] || grid[x][y] == '0' {
		return
	}
	visited[x][y] = true
	// left, right, up, down
	tour([2]int{x - 1, y}, grid, visited)
	tour([2]int{x + 1, y}, grid, visited)
	tour([2]int{x, y - 1}, grid, visited)
	tour([2]int{x, y + 1}, grid, visited)
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	//grid := [][]byte{
	//	{'1', '1', '0', '0', '0'},
	//	{'1', '1', '0', '0', '0'},
	//	{'0', '0', '1', '0', '0'},
	//	{'0', '0', '0', '1', '1'},
	//}
	fmt.Println(numIslands(grid))
}
