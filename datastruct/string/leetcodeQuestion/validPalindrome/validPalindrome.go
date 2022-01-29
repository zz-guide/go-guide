package main

import "log"

/**
题目：https://leetcode-cn.com/problems/valid-palindrome-ii/
验证回文字符串Ⅱ

给定一个非空字符串s，最多删除一个字符。判断是否能成为回文字符串。


示例 1:

输入: s = "aba"
输出: true
示例 2:

输入: s = "abca"
输出: true
解释: 你可以删除c字符。
示例 3:

输入: s = "abc"
输出: false

提示:

1 <= s.length <= 105
s 由小写英文字母组成

*/

func main() {
	s := "abca"
	log.Println("验证回文字符串Ⅱ:", validPalindrome(s))
}

// validPalindrome 时间复杂度O(n),空间复杂度O(1)
func validPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}

	// 验证回文串，左闭右闭
	var isPalindromeII func(s string, begin, end int) bool
	isPalindromeII = func(s string, begin, end int) bool {
		for begin < end {
			if s[begin] != s[end] {
				return false
			}
			begin++
			end--
		}
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		// case-1, 两端回文
		if s[left] == s[right] {
			left++
			right--
			// case-2, 遇上不等字符，验证(left, right]和[left:right)
			// 相当于与删除一个字符，看是否回文，删左或右都可以，所以是||
		} else {
			return isPalindromeII(s, left+1, right) || isPalindromeII(s, left, right-1)
		}
	}
	return true
}
