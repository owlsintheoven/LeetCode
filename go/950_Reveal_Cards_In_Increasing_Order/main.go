package main

import (
	"fmt"
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	res := make([]int, len(deck))
	var indexes []int
	for i := 0; i < len(deck); i++ {
		indexes = append(indexes, i)
	}
	reveal := true
	for len(indexes) != 0 {
		if reveal {
			res[indexes[0]] = deck[0]
			indexes = indexes[1:]
			deck = deck[1:]
		} else {
			indexes = append(indexes[1:], indexes[0])
		}
		reveal = !reveal
	}
	return res
}

func main() {
	deck := []int{17, 13, 11, 2, 3, 5, 7}
	fmt.Println(deckRevealedIncreasing(deck))
}
