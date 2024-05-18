package main

import "fmt"

// find a group of tasks requiring the least interval to wait
// if there are multiple, select the one with mose replicas
func leastInterval(tasks []byte, n int) int {
	intervals, counts := make([]int, 26), make([]int, 26)
	for _, t := range tasks {
		counts[t-'A'] += 1
	}

	res := 0
	done := 0
	for done != len(tasks) {
		var candidates []int
		minI := n + 1
		for i := 0; i < 26; i++ {
			if counts[i] > 0 {
				if intervals[i] < minI {
					candidates = []int{i}
					minI = intervals[i]
				} else if intervals[i] == minI {
					candidates = append(candidates, i)
				}
			}
		}
		maxC := -1
		var chosen int
		for _, c := range candidates {
			if counts[c] > 0 {
				if counts[c] > maxC {
					chosen = c
					maxC = counts[c]
				}
			}
		}
		waiting := intervals[chosen]
		for i := 0; i < 26; i++ {
			if counts[i] > 0 {
				if intervals[i] >= waiting+1 {
					intervals[i] -= waiting + 1
				} else {
					intervals[i] = 0
				}
			}
		}
		counts[chosen] -= 1
		if counts[chosen] > 0 {
			intervals[chosen] = n
		}
		res += waiting + 1
		done += 1
	}
	return res
}

func main() {
	//tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'E'}
	//n := 2
	//tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	//n := 2
	//tasks := []byte{'A', 'C', 'A', 'B', 'D', 'B'}
	//n := 1
	//tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	//n := 3
	//tasks := []byte{'A', 'B', 'C', 'A', 'B'}
	//n := 2
	tasks := []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	n := 1
	fmt.Println(leastInterval(tasks, n))
}
