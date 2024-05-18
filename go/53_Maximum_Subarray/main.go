package main

import "fmt"

func maxSubArray(nums []int) int {
	currentSum := nums[0]
	maxSum := currentSum
	for _, num := range nums[1:] {
		if currentSum > 0 {
			currentSum += num
		} else {
			currentSum = num
		}
		if maxSum < currentSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func main() {
	//nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//nums := []int{1}
	nums := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(nums))
}
