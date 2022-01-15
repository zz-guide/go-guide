package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/longest-palindromic-substring/

最长回文子串

给你一个字符串 s，找到 s 中最长的回文子串。

提示：

1 <= s.length <= 1000
s 仅由数字和英文字母（大写和/或小写）组成

*/
func main() {
	s := "abccba"
	fmt.Println("最长回文子串-动态规划:", longestPalindrome(s))
	fmt.Println("最长回文子串-中心扩展法:", longestPalindrome1(s))
}

// longestPalindrome 动态规划
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	ans := ""
	// 留意一下i和j的遍历顺序，i和j同时开始，i<j,因为遍历i的时候回访问到i+1，
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			// 单个字符情况肯定是回文子串
			if i == j {
				dp[i][j] = true
			} else if j-i == 1 && s[i] == s[j] { // 2个字符并且相等，也是回文子串
				dp[i][j] = true
			} else if j-i > 1 && s[i] == s[j] && dp[i+1][j-1] { // 首尾相等，并且区间内是回文串
				dp[i][j] = true
			}

			if dp[i][j] && j-i+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}

	return ans
}

// longestPalindrome1 双指针中心扩展
func longestPalindrome1(s string) string {
	n := len(s)
	var palindrome func(l, r int) (int, int)
	palindrome = func(l, r int) (int, int) {
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		return l + 1, r - 1
	}

	start, end := 0, 0
	for i := 0; i < n; i++ {
		// 奇数
		l1, r1 := palindrome(i, i)
		// 偶数
		l2, r2 := palindrome(i, i+1)

		if r1-l1 >= end-start {
			start, end = l1, r1
		}

		if r2-l2 >= end-start {
			start, end = l2, r2
		}
	}

	return s[start : end+1]
}
