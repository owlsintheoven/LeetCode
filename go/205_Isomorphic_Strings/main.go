package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	st, ts := make([]uint8, 256), make([]uint8, 256)
	for i := 0; i < len(s); i++ {
		if st[s[i]] == 0 {
			st[s[i]] = t[i] + 1
		} else if st[s[i]] != t[i]+1 {
			return false
		}
		if ts[t[i]] == 0 {
			ts[t[i]] = s[i] + 1
		} else if ts[t[i]] != s[i]+1 {
			return false
		}
	}
	return true
}

func main() {
	s := "badc"
	t := "baba"
	fmt.Println(isIsomorphic(s, t))
}

//func isIsomorphicOrdered(s string, t string) bool {
//	replaced := make(map[int32]uint8)
//	for i, c := range s {
//		r, ok := replaced[c]
//		if ok && r != t[i] {
//			return false
//		} else {
//			replaced[c] = t[i]
//		}
//	}
//	return true
//}
//
//func isIsomorphic(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//
//	return isIsomorphicOrdered(s, t) && isIsomorphicOrdered(t, s)
//}
