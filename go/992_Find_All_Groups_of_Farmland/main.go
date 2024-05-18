package main

import "fmt"

func findFarmland(land [][]int) [][]int {
	visited := make([][]bool, len(land))
	for i := 0; i < len(land); i++ {
		visited[i] = make([]bool, len(land[0]))
	}
	var res [][]int
	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[0]); j++ {
			if !visited[i][j] && land[i][j] == 1 {
				cord := tour([2]int{i, j}, land, visited)
				res = append(res, cord)
			}
		}
	}
	return res
}

func tour(pos [2]int, land [][]int, visited [][]bool) []int {
	x := pos[0]
	y := pos[1]
	// find bottom border
	bb := x
	for bb+1 < len(land) && land[bb+1][y] == 1 {
		bb++
	}
	// find right border
	rb := y
	for rb+1 < len(land[0]) && land[x][rb+1] == 1 {
		rb++
	}
	for i := x; i <= bb; i++ {
		for j := y; j <= rb; j++ {
			visited[i][j] = true
		}
	}
	return []int{x, y, bb, rb}
}

func main() {
	land := [][]int{
		{1, 0, 0},
		{0, 1, 1},
		{0, 1, 1},
	}
	fmt.Println(findFarmland(land))
}
