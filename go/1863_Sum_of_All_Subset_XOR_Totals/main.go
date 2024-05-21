package main

import "fmt"

func subsetXORSum(nums []int) int {
	n := len(nums)
	var res int
	for i := 1; i < 1<<n; i++ {
		var t int
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				t ^= nums[j]
			}
		}
		res += t
	}
	return res
}

func main() {
	nums := []int{1, 3}
	fmt.Println(subsetXORSum(nums))
}
