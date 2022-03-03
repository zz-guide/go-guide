package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/delete-operation-for-two-strings/
两个字符串的删除操作

给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。

提示：

给定单词的长度不超过500。
给定单词中的字符只含有小写字母。

*/
func main() {
	word1 := "sea"
	word2 := "eat"
	fmt.Println("两个字符串的删除操作-动态规划:", minDistance(word1, word2))
	fmt.Println("两个字符串的删除操作-求最长公共序列:", minDistance1(word1, word2))
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	// dp[i][j] 代表删除最小步数
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}

	for j := 0; j < n+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1]+2, min(dp[i][j-1]+1, dp[i-1][j]+1))
			}
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// minDistance1
func minDistance1(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i, c1 := range word1 {
		for j, c2 := range word2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 长度和 - 最长公共序列长度就是最少次数
	lcs := dp[m][n]
	return m + n - lcs*2
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
