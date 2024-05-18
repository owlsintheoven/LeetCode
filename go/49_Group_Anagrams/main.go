package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	m := make(map[string]int)
	for _, str := range strs {
		sorted := SortString(str)
		idx, ok := m[sorted]
		if !ok {
			m[sorted] = len(res)
			res = append(res, []string{str})
		} else {
			res[idx] = append(res[idx], str)
		}
	}
	return res
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
