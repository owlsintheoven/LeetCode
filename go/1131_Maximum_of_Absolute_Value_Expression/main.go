package main

import "fmt"

// maxAbsValExpr returns maximum of |arr1[i] - arr1[j]| + |arr2[i] - arr2[j]| + |i - j|
func maxAbsValExpr(arr1 []int, arr2 []int) int {
	best := 0
	for i := 0; i < len(arr1); i++ {
		for j := i + 1; j < len(arr1); j++ {
			val := absIntDiff(arr1[i], arr1[j]) + absIntDiff(arr2[i], arr2[j]) + absIntDiff(i, j)
			if best < val {
				best = val
			}
		}
	}
	return best
}

func absIntDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{-1, 4, 5, 6}
	//arr1 := []int{1, -2, -5, 0, 10}
	//arr2 := []int{0, -2, -1, -7, -4}
	fmt.Println(maxAbsValExpr(arr1, arr2))
}
