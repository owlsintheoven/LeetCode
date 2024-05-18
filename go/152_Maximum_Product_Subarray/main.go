package main

import (
	"fmt"
)

//func maxProduct(nums []int) int {
//	// dp[i][0] is the biggest possible product of the subarray ending at i
//	// dp[i][1] is the smallest possible product of the subarray ending at i
//	dp := make([][]int, len(nums))
//	for i := range dp {
//		dp[i] = make([]int, 2)
//	}
//	dp[0] = []int{nums[0], nums[0]}
//	maxProd := dp[0][0]
//	for i, num := range nums[1:] {
//		maxN := num
//		minN := num
//		for _, elem := range []int{num * dp[i][0], num * dp[i][1]} {
//			if maxN < elem {
//				maxN = elem
//			}
//			if minN > elem {
//				minN = elem
//			}
//		}
//		dp[i+1][0] = maxN
//		dp[i+1][1] = minN
//		if maxProd < maxN {
//			maxProd = maxN
//		}
//	}
//	return maxProd
//}

func maxProduct(nums []int) int {
	maxProd := nums[0]
	minProd := nums[0]
	bestProd := nums[0]
	for _, num := range nums[1:] {
		maxN := num
		minN := num
		for _, n := range []int{maxProd * num, minProd * num} {
			if maxN < n {
				maxN = n
			}
			if minN > n {
				minN = n
			}
		}
		maxProd = maxN
		minProd = minN
		if bestProd < maxN {
			bestProd = maxN
		}
	}
	return bestProd
}

func main() {
	//nums := []int{2, 3, -2, 4}
	//nums := []int{-2, 0, -1}
	nums := []int{-2, 3, -4}
	fmt.Println(maxProduct(nums))
}
