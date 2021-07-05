package main

// Tags:
// Dynamic Programming
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := range text1 {
		for j := range text2 {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = Max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[m][n]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
