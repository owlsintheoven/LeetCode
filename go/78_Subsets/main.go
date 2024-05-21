package main

import "fmt"

//// recursive method
//func subsets(nums []int) [][]int {
//	var res [][]int
//	find(nums, []int{}, &res)
//	return res
//}
//
//func find(nums []int, curr []int, res *[][]int) {
//	*res = append(*res, curr)
//	for i, n := range nums {
//		var newCurr []int
//		newCurr = append(newCurr, curr...)
//		newCurr = append(newCurr, n)
//		var newNums []int
//		newNums = append(newNums, nums[i+1:]...)
//		find(newNums, newCurr, res)
//	}
//}

// bitwise method
func subsets(nums []int) [][]int {
	n := len(nums)
	var res [][]int
	for i := 0; i < 1<<n; i++ {
		var subset []int
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				subset = append(subset, nums[j])
			}
		}
		res = append(res, subset)
	}
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
