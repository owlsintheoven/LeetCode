package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	// dp[i][j] represents the least actions to transform word1[0...i] to word2[0...j]
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	// delete word1[0...i] to make an empty string
	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	// insert an empty string to make word2[0...j]
	for j := 0; j < n+1; j++ {
		dp[0][j] = j
	}
	// if word1[i] == word2[j] dp[i][j] = dp[i-1][j-1]
	// or the minimum of
	// replace dp[i-1][j-1] + 1
	// insert dp[i-1][j] + 1
	// delete dp[i][j-1] + 1
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				rep := dp[i-1][j-1] + 1
				ins := dp[i-1][j] + 1
				del := dp[i][j-1] + 1
				dp[i][j] = rep
				for _, e := range []int{ins, del} {
					if dp[i][j] > e {
						dp[i][j] = e
					}
				}
			}

		}
	}
	return dp[m][n]
}

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}
