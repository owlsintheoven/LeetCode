package main

import (
	"fmt"
	"sort"
)

/*
*
There are n workers. You are given two integer arrays quality and wage where quality[i] is the quality of the ith worker and wage[i] is the minimum wage expectation for the ith worker.

We want to hire exactly k workers to form a paid group. To hire a group of k workers, we must pay them according to the following rules:

Every worker in the paid group must be paid at least their minimum wage expectation.
In the group, each worker's pay must be directly proportional to their quality. This means if a worker’s quality is double that of another worker in the group, then they must be paid twice as much as the other worker.
Given the integer k, return the least amount of money needed to form a paid group satisfying the above conditions. Answers within 10^-5 of the actual answer will be accepted.

Example 1:

Input: quality = [10,20,5], wage = [70,50,30], k = 2
Output: 105.00000
Explanation: We pay 70 to 0th worker and 35 to 2nd worker.
Example 2:

Input: quality = [3,1,10,10,1], wage = [4,8,2,2,7], k = 3
Output: 30.66667
Explanation: We pay 4 to 0th worker, 13.33333 to 2nd and 3rd workers separately.
*/

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	// 2.5 6 7 / 20 5 10
	// sort quality in asc order of wage/quality
	// init first k workers as candidates, and use a stack to record worker with the highest quality, also with the total of quality and current quotient
	// 20 5 / 25 / 6
	// whenever the stack exceeds total count of k, pop the first elem, and update candidates and quality total
	// when iterating the rest workers, find the min of current cost and the one if the newly added worker and rebased wage/quality
	// update candidates, stack, quality total and cost if a new min is found
	quality, qws := sortToQW(quality, wage)
	var candidates []int
	var totalQuality int
	var qw float64
	var cost float64
	for i := 0; i < k; i++ {
		totalQuality += quality[i]
		candidates = pushPriorityQ(candidates, quality[i])
	}
	qw = qws[k-1]
	cost = float64(totalQuality) * qw
	for i := k; i < len(quality); i++ {
		qw = qws[i]
		candidates = pushPriorityQ(candidates, quality[i])
		var prev int
		candidates, prev = popPriorityQ(candidates)
		totalQuality = totalQuality - prev + quality[i]
		if cost > float64(totalQuality)*qw {
			cost = float64(totalQuality) * qw
		}
	}
	return cost
}

type qqw struct {
	quality  int
	quotient float64
}

type sortQQW []qqw

func (s sortQQW) Less(i, j int) bool {
	return s[i].quotient < s[j].quotient
}

func (s sortQQW) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortQQW) Len() int {
	return len(s)
}

func sortToQW(quality, wage []int) ([]int, []float64) {
	qqws := make([]qqw, len(quality))
	for i := 0; i < len(quality); i++ {
		qqws[i] = qqw{
			quality:  quality[i],
			quotient: float64(wage[i]) / float64(quality[i]),
		}
	}
	sort.Sort(sortQQW(qqws))
	qws := make([]float64, len(quality))
	for i := 0; i < len(quality); i++ {
		quality[i] = qqws[i].quality
		qws[i] = qqws[i].quotient
	}
	return quality, qws
}

func pushPriorityQ(nums []int, n int) []int {
	res := make([]int, len(nums)+1)
	pushed := false
	pos := 0
	for pos < len(nums) {
		if n > nums[pos] {
			res[pos] = n
			pushed = true
			break
		} else {
			res[pos] = nums[pos]
		}
		pos++
	}
	if !pushed {
		res[pos] = n
	} else {
		for i := pos; i < len(nums); i++ {
			res[i+1] = nums[i]
		}
	}
	return res
}

func popPriorityQ(nums []int) ([]int, int) {
	var res []int
	res = append(res, nums[1:]...)
	return res, nums[0]
}

//// naive method
//func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
//	res := math.Pow(10, 8)
//	// find who receives minimum wage first
//	// see if the rest of workers can receive proportional wage which is over their minimum wage
//	// there must be at least (k - 1) workers to be qualified
//	// or the first worker will be considered invalid as the one to receive minimum wage
//	// find all possible solutions
//	// find which solution cost the least
//	for i := 0; i < len(quality); i++ {
//		p := float64(wage[i]) / float64(quality[i])
//		var qualified []float64
//		for j := 0; j < len(quality); j++ {
//			if i != j {
//				w := float64(quality[j]) * p
//				if w >= float64(wage[j]) {
//					qualified = append(qualified, w)
//				}
//			}
//		}
//		if len(qualified) >= k-1 {
//			sort.Float64s(qualified)
//			cost := float64(wage[i])
//			for j := 0; j < k-1; j++ {
//				cost += qualified[j]
//			}
//			if cost < res {
//				res = cost
//			}
//		}
//	}
//	return res
//}

func main() {
	//quality := []int{10, 20, 5}
	//wage := []int{70, 50, 30}
	//k := 2
	//quality := []int{14, 56, 59, 89, 39, 26, 86, 76, 3, 36}
	//wage := []int{90, 217, 301, 202, 294, 445, 473, 245, 415, 487}
	//k := 2
	//quality := []int{3, 1, 10, 10, 1}
	//wage := []int{4, 8, 2, 2, 7}
	//k := 3
	quality := []int{32, 43, 66, 9, 94, 57, 25, 44, 99, 19}
	wage := []int{187, 366, 117, 363, 121, 494, 348, 382, 385, 262}
	k := 4
	fmt.Println(mincostToHireWorkers(quality, wage, k))
}

// k = 1
// 30 (5)
// k = 2
// 70 + 35 (10, 5)
// k = 3
// 70 + 35 + 100 (10, 20, 5)

// 3, 1, 10, 10, 1
// 4, 8, 2, 2, 7
// k = 1
// 2 (10)
// k = 2
// 2 + 2 (10, 10)
// k = 3
// 4 + 13.333 + 13.333 (3, 10, 10)

// 3, 10, 1, 1, 10
// 4, 8, 2, 2, 7
// k = 1
// 2 (1)
// k = 2
// 2 + 2 (1, 1)
// k = 3
// 12 + 2 + 2 = 16

// 3, 1, 10, 10, 1
// 8, 4, 2, 2, 7
// k = 1
// 2 (10)
// k = 2
// 2 + 2 (10, 10)
// k = 3
// 8 + 26.666 + 26.666 (3, 10, 10) = 60.XXX
// 12 + 4 + 40 (3, 1 10) = 56
