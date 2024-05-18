package main

import "fmt"

func removeDuplicates(nums []int) int {
	res := 0
	prev := -101
	for _, num := range nums {
		if num != prev {
			nums[res] = num
			prev = num
			res += 1
		}
	}
	return res
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
