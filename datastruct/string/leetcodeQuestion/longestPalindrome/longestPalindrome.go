package main

import "log"

/**
题目：https://leetcode-cn.com/problems/longest-palindrome/
最长回文串

给定一个包含大写字母和小写字母的字符串s，返回通过这些字母构造成的 最长的回文串。

在构造过程中，请注意 区分大小写 。比如"Aa"不能当做一个回文字符串。


示例 1:

输入:s = "abccccdd"
输出:7
解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
示例 2:

输入:s = "a"
输入:1
示例 3:

输入:s = "bb"
输入: 2

提示:

1 <= s.length <= 2000
s只能由小写和/或大写英文字母组成


注意：1.字符串字符只包含字母大小写，一共52个
2.

*/

func main() {
	s := "abccccdd"
	log.Println("最长回文串:", longestPalindrome(s))
}

// longestPalindrome 贪心算法，时间复杂度:O(n)，空间复杂度O(S)
func longestPalindrome(s string) int {
	charCnt := make(map[rune]int, 52)
	for _, i2 := range s {
		charCnt[i2]++
	}

	count := 0
	for _, cnt := range charCnt {
		// 偶数不变，奇数，最多用奇数-1
		count += cnt / 2 * 2
		// 奇数并且长度目前是偶数，那么可以使用
		// 如果已经有一个奇数被计算了，那么现在这个就不能计数
		if cnt%2 == 1 && count%2 == 0 {
			count++
		}
	}

	return count
}
