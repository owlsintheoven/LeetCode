package main

import (
	"fmt"
)

// numBusesToDestination organizes which routes pass which stops first.
// Start from source, find every stops every routes pass by source, and check if target is one of them
// Iterate again but from all possible stops from previous step, again check if target exists in this iteration
// If no possible stops in this iteration returns -1
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	stopHasRoutes := make([][]int, 1000000)
	var maxS int
	var foundSource, foundTarget bool
	for i, route := range routes {
		for _, stop := range route {
			if stop == source {
				foundSource = true
			} else if stop == target {
				foundTarget = true
			}
			stopHasRoutes[stop] = append(stopHasRoutes[stop], i)
			if maxS < stop {
				maxS = stop
			}
		}
	}
	if !(foundSource && foundTarget) {
		return -1
	}
	stops := []int{source}
	prevStops := make([]bool, maxS+1)
	prevStops[source] = true
	prevRoutes := make([]bool, len(routes))
	num := 1
	for {
		var next_stops []int
		for _, stop := range stops {
			// go on buses that pass these stops
			for _, r := range stopHasRoutes[stop] {
				if !prevRoutes[r] {
					for _, s := range routes[r] {
						if s == target {
							return num
						}
						if !prevStops[s] {
							next_stops = append(next_stops, s)
							prevStops[s] = true
						}
					}
					prevRoutes[r] = true
				}
			}
		}
		// check all possible stops
		if len(next_stops) == 0 {
			break
		}
		stops = next_stops
		num += 1
	}
	return -1
}

func main() {
	//routes := [][]int{{1, 2, 7}, {3, 6, 7}}
	//source := 1
	//target := 6
	routes := [][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}
	source := 15
	target := 12
	fmt.Println(numBusesToDestination(routes, source, target))
}
