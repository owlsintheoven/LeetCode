package main

import (
	"fmt"
	"slices"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	slices.Reverse(people)
	l := 0
	r := len(people) - 1
	res := 0
	for l < r {
		if people[l]+people[r] <= limit {
			r -= 1
		}
		l += 1
		res += 1
	}
	if l == r {
		res += 1
	}
	return res
}

func main() {
	people := []int{1, 2}
	limit := 3
	fmt.Println(numRescueBoats(people, limit))
}
