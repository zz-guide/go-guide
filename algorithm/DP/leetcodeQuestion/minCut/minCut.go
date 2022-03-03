package main

import (
	"log"
	"math"
)

/**
题目：https://leetcode-cn.com/problems/palindrome-partitioning-ii/

分割回文串 II

给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文。

返回符合要求的 最少分割次数 。

示例 1：

输入：s = "aab"
输出：1
解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
示例 2：

输入：s = "a"
输出：0
示例 3：

输入：s = "ab"
输出：1

提示：

1 <= s.length <= 2000
s 仅由小写英文字母组成

*/

func main() {
	s := "ebaabab"
	log.Println("(最少分割次数)分割回文串II-两次动态规划", minCut(s))
}

// 动态规划 O(n^2), 需要2次
func minCut(s string) int {
	isPalindrome := make([][]bool, len(s))
	for i := 0; i < len(isPalindrome); i++ {
		// 默认都不是回文串
		isPalindrome[i] = make([]bool, len(s))
		// dp[i][i]代表单个字符，肯定是回文串
		isPalindrome[i][i] = true
	}

	// 第一次动态规划标记i~j之内哪些是回文串
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if i == j {
				isPalindrome[i][j] = true
			} else if j-i == 1 && s[i] == s[j] { // 2个字符并且相等，也是回文子串
				isPalindrome[i][j] = true
			} else if j-i > 1 && s[i] == s[j] && isPalindrome[i+1][j-1] { // 首尾相等，并且区间内是回文串
				isPalindrome[i][j] = true
			}
		}
	}

	dp := make([]int, len(s)+1)
	for i := 1; i < len(s)+1; i++ {
		dp[i] = math.MaxInt64
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if isPalindrome[i][j] {
				dp[j+1] = min(dp[j+1], dp[i]+1)
			}
		}
	}

	return dp[len(s)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
