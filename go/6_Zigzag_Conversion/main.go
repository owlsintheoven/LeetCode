package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	w := wide(len(s), numRows)
	zigzag := make([][]int32, numRows)
	for i := 0; i < numRows; i++ {
		zigzag[i] = make([]int32, w)
	}
	goingDown := true
	x := 0
	y := 0
	zigzag[0][0] = int32(s[0])
	for _, c := range s[1:] {
		if goingDown {
			x += 1
			zigzag[x][y] = c
			if x == numRows-1 {
				goingDown = false
			}
		} else {
			x -= 1
			y += 1
			zigzag[x][y] = c
			if x == 0 {
				goingDown = true
			}
		}
	}
	var res string

	for i := 0; i < numRows; i++ {
		for j := 0; j < w; j++ {
			if zigzag[i][j] != 0 {
				res += string(zigzag[i][j])
			}
		}
	}
	return res
}

func wide(l, numRows int) int {
	res := 1
	y := 0
	goingDown := true
	l -= 1
	for l > 0 {
		l -= 1
		if goingDown {
			y += 1
			if y == numRows-1 {
				goingDown = false
			}
		} else {
			res += 1
			y -= 1
			if y == 0 {
				goingDown = true
			}
		}
	}
	return res
}

func main() {
	//s := "PAYPALISHIRING"
	//numRows := 3
	s := "AB"
	numRows := 1
	fmt.Println(convert(s, numRows))
}

//14, 3, 7
//14, 4, 7
//14, 5, 6
//1+4+0+1=6
//5+4+4+1=14
//
//abcdefghijklmn
//a   i
//b  hj
//c g k
//df  ln
//e   m
//a     m
//b    ln
//c   k
//d  j
//e i
//fh
//g
