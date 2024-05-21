package main

import (
	"fmt"
	"sort"
)

// recursive method
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	find(nums, []int{}, &res)
	return res
}

func find(nums []int, curr []int, res *[][]int) {
	*res = append(*res, curr)
	for i, n := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		var newCurr []int
		newCurr = append(newCurr, curr...)
		newCurr = append(newCurr, n)
		var newNums []int
		newNums = append(newNums, nums[i+1:]...)
		find(newNums, newCurr, res)
	}
}

//// method: bitwise / specifically check for duplicates
//func subsetsWithDup(nums []int) [][]int {
//	n := len(nums)
//	sort.Ints(nums)
//	var res [][]int
//	for i := 0; i < 1<<n; i++ {
//		var subset []int
//		for j := 0; j < n; j++ {
//			if i&(1<<j) != 0 {
//				subset = append(subset, nums[j])
//			}
//		}
//		duplicated := false
//		for _, r := range res {
//			if intSliecEqual(subset, r) {
//				duplicated = true
//				break
//			}
//		}
//		if !duplicated {
//			res = append(res, subset)
//		}
//	}
//	return res
//}
//
//func intSliecEqual(a, b []int) bool {
//	if len(a) != len(b) {
//		return false
//	}
//	for i := 0; i < len(a); i++ {
//		if a[i] != b[i] {
//			return false
//		}
//	}
//	return true
//}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
