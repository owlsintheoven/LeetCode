package main

import "fmt"

func removeElement(nums []int, val int) int {
	res := 0
	for _, num := range nums {
		if num != val {
			nums[res] = num
			res += 1
		}
	}
	return res
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)
}
