package main

import "fmt"

func exist(board [][]byte, word string) bool {
	r, c := len(board), len(board[0])
	walked := make([][]bool, r)
	for idx := range walked {
		walked[idx] = make([]bool, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if board[i][j] == word[0] {
				walked[i][j] = true
				if walk(board, word[1:], [2]int{i, j}, walked) {
					return true
				}
				walked[i][j] = false
			}
		}
	}

	return false
}

func walk(board [][]byte, word string, pos [2]int, walked [][]bool) bool {
	if len(word) == 0 {
		return true
	}
	x, y := pos[0], pos[1]
	// up
	if x > 0 && walked[x-1][y] == false && board[x-1][y] == word[0] {
		walked[x-1][y] = true
		if walk(board, word[1:], [2]int{x - 1, y}, walked) {
			return true
		}
		walked[x-1][y] = false
	}
	// down
	if x < len(board)-1 && walked[x+1][y] == false && board[x+1][y] == word[0] {
		walked[x+1][y] = true
		if walk(board, word[1:], [2]int{x + 1, y}, walked) {
			return true
		}
		walked[x+1][y] = false
	}
	// left
	if y > 0 && walked[x][y-1] == false && board[x][y-1] == word[0] {
		walked[x][y-1] = true
		if walk(board, word[1:], [2]int{x, y - 1}, walked) {
			return true
		}
		walked[x][y-1] = false
	}
	// right
	if y < len(board[0])-1 && walked[x][y+1] == false && board[x][y+1] == word[0] {
		walked[x][y+1] = true
		if walk(board, word[1:], [2]int{x, y + 1}, walked) {
			return true
		}
		walked[x][y+1] = false
	}

	return false
}

func main() {
	//board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	//word := "ABCCED"
	//board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	//word := "SEE"
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCB"
	fmt.Println(exist(board, word))
}
