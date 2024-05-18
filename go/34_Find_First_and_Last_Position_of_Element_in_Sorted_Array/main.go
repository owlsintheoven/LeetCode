package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	l := lowerBound(nums, target)
	if l < len(nums) && nums[l] == target {
		return []int{l, higherBound(nums, target)}
	}
	return []int{-1, -1}
}

func lowerBound(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func higherBound(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}

func main() {
	//nums := []int{5, 7, 7, 8, 8, 10}
	//target := 8
	nums := []int{1}
	target := 1
	fmt.Println(searchRange(nums, target))
}
