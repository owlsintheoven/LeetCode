package main

import (
	"fmt"
	"sort"
)

func isAnagram(s string, t string) bool {
	counts := make([]int, 26)
	for _, c := range s {
		counts[c-97]++
	}
	for _, c := range t {
		counts[c-97]--
	}
	valid := true
	for i := 0; i < 26; i++ {
		if counts[i] != 0 {
			valid = false
			break
		}
	}
	return valid
	//sortedS := SortString(s)
	//sortedT := SortString(t)
	//return sortedS == sortedT
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
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}
