package main

import (
	"fmt"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/palindromic-substrings/

回文子串

给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。

回文字符串 是正着读和倒过来读一样的字符串。

子字符串 是字符串中的由连续字符组成的一个序列。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。


提示：
1 <= s.length <= 1000
s 由小写英文字母组成


注意：1.单个字符也是回文子串
2.子串必须是连续的
3.相同的子串可能是不同的位置，此时认为是不同的子串

算法思路分几类：
	1.固定一个中心字符，向左右两边扩散判断是不是回文
	2.暴力法穷举所有的子串，然后单独判断是不是回文
	3.dp[i][j]分析可知三种情况是回文
	4.回文串长度是奇数的话，中间是一个字符，偶数的话，中间是2个相等的字符
	5.假设字符串都是单个字节的字符
*/
func main() {
	s := "abccba"
	fmt.Println("回文子串数目-中心拓展:", countSubstrings(s))
	fmt.Println("回文子串数目-Manacher 算法:", countSubstrings1(s))
	fmt.Println("回文子串数目-暴力法:", countSubstrings2(s))
	fmt.Println("回文子串数目-动态规划:", countSubstrings3(s))
	fmt.Println("回文子串数目-降维动态规划:", countSubstrings4(s))
	fmt.Println("回文子串数目-双指针:", countSubstrings5(s))
}

// countSubstrings O(n^2) 中心拓展法
func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	// 统一了奇数和偶数的情况的写法
	// 回文中心可能是一个字符或者两个字符
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}

	return ans
}

// countSubstrings5 双指针(中心拓展)
func countSubstrings5(s string) int {
	n := len(s)
	var palindrome func(l, r int) int
	palindrome = func(l, r int) (count int) {
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			count++
		}
		return
	}
	ans := 0
	for i := 0; i < n; i++ {
		// 奇数
		ans += palindrome(i, i)
		// 偶数
		ans += palindrome(i, i+1)
	}
	return ans
}

// countSubstrings1 O(n) Manacher 算法
// 通过插入特殊字符，把字符串固定成奇数个数，就不需要判断偶数的情况了
func countSubstrings1(s string) int {
	n := len(s)
	t := "$#"
	for i := 0; i < n; i++ {
		t += string(s[i]) + "#"
	}

	n = len(t)
	t += "!"
	// 将原始字符串转换格式，添加#,头添加$#,尾部添加!
	log.Println("s:", t)

	f := make([]int, n) //表示以 s 的第 i 位为回文中心，可以拓展出的最大回文半径
	iMax, rMax, ans := 0, 0, 0
	for i := 1; i < n; i++ {
		// 初始化 f[i]
		if i <= rMax {
			f[i] = min(rMax-i+1, f[2*iMax-i])
		} else {
			f[i] = 1
		}

		// i中心拓展
		for t[i+f[i]] == t[i-f[i]] {
			f[i]++
		}

		// 动态维护 iMax 和 rMax
		if i+f[i]-1 > rMax {
			iMax = i
			rMax = i + f[i] - 1
		}

		// 统计答案, 当前贡献为 (f[i] - 1) / 2 上取整
		ans += f[i] / 2
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// countSubstrings2 暴力法
func countSubstrings2(s string) int {
	count := 0

	// 判断是不是回文字符串
	var isPalindrome func(s string) bool
	isPalindrome = func(s string) bool {
		i, j := 0, len(s)-1
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindrome(s[i : j+1]) {
				count++
			}
		}
	}

	return count
}

// countSubstrings3 动态规划
func countSubstrings3(s string) int {
	count := 0
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	// 留意一下i和j的遍历顺序，i和j同时开始，i<j,因为遍历i的时候回访问到i+1，
	// i 每次从头开始，j为右边界，查看i~j之间是不是回文
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			// 单个字符情况肯定是回文子串
			if i == j {
				dp[i][j] = true
				count++
			} else if j-i == 1 && s[i] == s[j] { // 2个字符并且相等，也是回文子串
				dp[i][j] = true
				count++
			} else if j-i > 1 && s[i] == s[j] && dp[i+1][j-1] { // 首尾相等，并且区间内是回文串
				dp[i][j] = true
				count++
			}
		}
	}

	return count
}

// countSubstrings4 降维动态规划
func countSubstrings4(s string) int {
	count := 0
	dp := make([]bool, len(s))

	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if i == j {
				dp[i] = true
				count++
			} else if j-i == 1 && s[i] == s[j] {
				dp[i] = true
				count++
			} else if j-i > 1 && s[i] == s[j] && dp[i+1] {
				dp[i] = true
				count++
			} else {
				dp[i] = false
			}
		}
	}
	return count
}
