package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/is-subsequence/
判断子序列

给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

进阶：

如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？

提示：

0 <= s.length <= 100
0 <= t.length <= 10^4
两个字符串都只由小写字符组成。

*/
func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println("判断子序列-贪心算法:", isSubsequence(s, t))
	fmt.Println("判断子序列-动态规划:", isSubsequence1(s, t))
	fmt.Println("判断子序列-大数据dp:", isSubsequence2(s, t))
}

// isSubsequence 贪心算法
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s) > len(t) {
		return false
	}

	sStart, sLength := 0, len(s)
	for i := 0; i < len(t); i++ {
		if sStart < sLength && t[i] == s[sStart] {
			sStart++
		}
	}

	if sStart != sLength {
		return false
	}

	return true
}

// isSubsequence1 动态规划
func isSubsequence1(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s) > len(t) {
		return false
	}

	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(s)][len(t)] == len(s)
}

// dp预处理母串，dp[i][j]表示第i个字符之后出现j+'a'字母的位置
//判断时直接通过dp数组跳转到下一位置，为0则找不到对应字母，返回false
// isSubsequence2
func isSubsequence2(s string, t string) bool {
	n, m := len(s), len(t)
	// 预处理 T
	dp := make([][]int, m+1)
	dp[m] = make([]int, 26)
	// 从后向前扫描 T，找出i位置之后字母出现的位置
	for i := m - 1; i >= 0; i-- {
		dp[i] = make([]int, 26)
		copy(dp[i], dp[i+1])
		dp[i][t[i]-'a'] = i + 1
	}

	i, j := 0, 0
	for i < m && j < n {
		p := dp[i][s[j]-'a']
		if p != 0 {
			j++
			i = p
		} else {
			return false
		}
	}

	if j == n {
		return true
	}

	return false
}
