package main

import (
	"fmt"
	"sort"
)

func maximumHappinessSum(happiness []int, k int) int64 {
	var res int64
	sort.Ints(happiness)
	round := 0
	pos := len(happiness) - 1
	for k > 0 {
		var point int
		if happiness[pos] >= round {
			point = happiness[pos] - round
		} else {
			point = 0
		}
		res += int64(point)
		k--
		round++
		pos--
	}
	return res
}

func main() {
	//happiness := []int{1, 2, 3}
	//k := 2
	happiness := []int{1, 1, 1, 1}
	k := 2
	//happiness := []int{2, 3, 4, 5}
	//k := 1
	fmt.Println(maximumHappinessSum(happiness, k))
}
