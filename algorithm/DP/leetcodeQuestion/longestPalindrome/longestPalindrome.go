package main

import (
	"fmt"
)

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
	fmt.Println("最长回文子串-Manacher算法:", longestPalindrome2(s))
}

// longestPalindrome 动态规划,O(n^2)
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	res := ""
	// 留意一下i和j的遍历顺序，i和j同时开始，i<j,因为遍历i的时候会访问到i+1，
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

			// 当前i~j是回文子串，则比较maxLength
			if dp[i][j] && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}

	return res
}

// longestPalindrome1 双指针中心扩展,O(n^2)
func longestPalindrome1(s string) string {
	n := len(s)
	// 传入回文中心坐标，向两边扩散查找，直到不是回文，然后返回坐标
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
		// 依次比较哪个大
		if r1-l1 >= end-start {
			start, end = l1, r1
		}

		if r2-l2 >= end-start {
			start, end = l2, r2
		}
	}

	return s[start : end+1]
}

// longestPalindrome2 Manacher 算法 O(n) 不做要求
func longestPalindrome2(s string) string {
	start, end := 0, -1
	t := "#"
	for i := 0; i < len(s); i++ {
		t += string(s[i]) + "#"
	}

	t += "#"
	s = t
	var arm_len []int
	right, j := -1, -1
	for i := 0; i < len(s); i++ {
		var cur_arm_len int
		if right >= i {
			i_sym := j*2 - i
			min_arm_len := min(arm_len[i_sym], right-i)
			cur_arm_len = expand(s, i-min_arm_len, i+min_arm_len)
		} else {
			cur_arm_len = expand(s, i, i)
		}
		arm_len = append(arm_len, cur_arm_len)
		if i+cur_arm_len > right {
			j = i
			right = i + cur_arm_len
		}
		if cur_arm_len*2+1 > end-start {
			start = i - cur_arm_len
			end = i + cur_arm_len
		}
	}
	ans := ""
	for i := start; i <= end; i++ {
		if s[i] != '#' {
			ans += string(s[i])
		}
	}
	return ans
}

func expand(s string, left, right int) int {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return (right - left - 2) / 2
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
