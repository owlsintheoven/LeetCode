package main

import (
	"fmt"
)

func nextPermutation(nums []int) {
	// find break point where nums[bp] < nums[bp+1]
	bp := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			bp = i
			break
		}
	}
	if bp == -1 {
		reverse(nums, 0, len(nums)-1)
	} else {
		// find the smallest num that's bigger than nums[bp] on the right half
		// and swap it with nums[bp]
		target := nums[bp+1]
		targetIdx := bp + 1
		for i := bp + 2; i < len(nums); i++ {
			if nums[i] > nums[bp] {
				if nums[i] < target {
					target = nums[i]
					targetIdx = i
				} else if nums[i] == target {
					targetIdx = i
				}
			}
		}
		swap(nums, bp, targetIdx)
		// reverse the rest from bp+1 to the right end
		reverse(nums, bp+1, len(nums)-1)
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func reverse(nums []int, l, r int) {
	for i, j := l, r; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	nums := []int{2, 3, 1, 3, 3}
	//nums := []int{1, 2}
	//nums := []int{1}
	nextPermutation(nums)
	fmt.Println(nums)
}

//2,3,1,3,3
//2,3,3,3,1
//2,3,3,1,3
