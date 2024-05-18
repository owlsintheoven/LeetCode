package main

import (
	"fmt"
	"sort"
)

func topKFrequent(words []string, k int) []string {
	// counts[i] contains string array consisting strings of frequency i
	frequency := make([][]string, len(words)+1)
	counts := make(map[string]int)
	for _, w := range words {
		c, ok := counts[w]
		if !ok {
			counts[w] = 1
			frequency[1] = append(frequency[1], w)
		} else {
			counts[w] = c + 1
			var idx int
			for idx = 0; idx < len(frequency[c]); idx++ {
				if frequency[c][idx] == w {
					break
				}
			}
			frequency[c] = append(frequency[c][:idx], frequency[c][idx+1:]...)
			frequency[c+1] = append(frequency[c+1], w)
		}
	}
	var res []string
	for i := len(words); i >= 0; i-- {
		end := false
		if len(frequency[i]) != 0 {
			sort.Strings(frequency[i])
			for _, w := range frequency[i] {
				res = append(res, w)
				k -= 1
				if k == 0 {
					end = true
					break
				}
			}
		}
		if end {
			break
		}
	}
	return res
}

func main() {
	//words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	//k := 2
	words := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	k := 4
	//words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	//k := 1
	fmt.Println(topKFrequent(words, k))
}
