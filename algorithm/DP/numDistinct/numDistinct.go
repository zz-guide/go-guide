package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/distinct-subsequences/

不同的子序列

给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。

字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。
（例如，"ACE"是"ABCDE"的一个子序列，而"AEC"不是）

题目数据保证答案符合 32 位带符号整数范围。

提示：

0 <= s.length, t.length <= 1000
s 和 t 由英文字母组成



注意：1.这道题目如果不是子序列，而是要求连续序列的，那就可以考虑用KMP。
2.只有删除，不用考虑替换。增加之类的操作
3.dp[i][j]：以i-1为结尾的s子序列中出现以j-1为结尾的t的个数为dp[i][j]。


*/
func main() {
	s := "rabbbit"
	t := "rabbit"
	fmt.Println("不同的子序列个数-动态规划:", numDistinct(s, t))
}

// numDistinct 动态规划
func numDistinct(s, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[m][n]
}
