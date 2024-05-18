package main

import "fmt"

func generate(numRows int) [][]int {
	var res [][]int
	if numRows == 1 {
		res = append(res, []int{1})
	} else if numRows >= 2 {
		res = append(res, []int{1})
		res = append(res, []int{1, 1})
	}
	pre := res[len(res)-1]
	for i := 2; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = pre[0]
		row[i] = pre[len(pre)-1]
		for j := 1; j < i; j++ {
			row[j] = pre[j-1] + pre[j]
		}
		res = append(res, row)
		pre = row
	}
	return res
}

func main() {
	numRows := 5
	fmt.Println(generate(numRows))
}
