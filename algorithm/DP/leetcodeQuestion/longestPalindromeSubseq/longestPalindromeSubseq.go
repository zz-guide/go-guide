package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/longest-palindromic-subsequence/

最长回文子序列

给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。

子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。


提示：

1 <= s.length <= 1000
s 仅由小写英文字母组成


*/
func main() {
	s := "bbbab"
	fmt.Println("最长回文子序列-动态规划:", longestPalindromeSubseq1(s))
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// longestPalindromeSubseq1 动态规划
func longestPalindromeSubseq1(s string) int {
	length := len(s)
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
	}

	for i := length - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < length; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][length-1]
}
