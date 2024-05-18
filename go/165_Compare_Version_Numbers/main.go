package main

import (
	"fmt"
	"strconv"
)

func compareVersion(version1 string, version2 string) int {
	var pos1, pos2 int
	for {
		if pos1 == len(version1) && pos2 == len(version2) {
			break
		}
		var revision1 int
		if pos1 == len(version1) {
			revision1 = 0
		} else {
			found := false
			for i := pos1; i < len(version1); i++ {
				if version1[i] == '.' {
					revision1 = convert(version1[pos1:i])
					pos1 = i + 1
					found = true
					break
				}
			}
			if !found {
				revision1 = convert(version1[pos1:])
				pos1 = len(version1)
			}
		}
		var revision2 int
		if pos2 == len(version2) {
			revision2 = 0
		} else {
			found := false
			for i := pos2; i < len(version2); i++ {
				if version2[i] == '.' {
					revision2 = convert(version2[pos2:i])
					pos2 = i + 1
					found = true
					break
				}
			}
			if !found {
				revision2 = convert(version2[pos2:])
				pos2 = len(version2)
			}
		}
		if revision1 < revision2 {
			return -1
		} else if revision1 > revision2 {
			return 1
		}
	}
	return 0
}

func convert(s string) int {
	// trim leading 0s
	trimmed := "0"
	for i := 0; i < len(s); i++ {
		if s[i] != '0' {
			trimmed = s[i:]
			break
		}
	}
	// convert to int
	res, _ := strconv.Atoi(trimmed)
	return res
}

func main() {
	//version1 := "1.01"
	//version2 := "1.001"
	//version1 := "1.0"
	//version2 := "1.0.0"
	//version1 := "0.1"
	//version2 := "1.1"
	version1 := "11"
	version2 := "10"
	fmt.Println(compareVersion(version1, version2))
}
