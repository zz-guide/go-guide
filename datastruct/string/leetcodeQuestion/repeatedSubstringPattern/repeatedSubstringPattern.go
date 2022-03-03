package main

import (
	"log"
	"strings"
)

/**
题目：https://leetcode-cn.com/problems/repeated-substring-pattern/

重复的子字符串

给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。

*/
func main() {
	s := "abc"
	log.Println("重复的子字符串-枚举暴力:", repeatedSubstringPattern(s))
	log.Println("重复的子字符串-规律性质:", repeatedSubstringPattern1(s))
	log.Println("重复的子字符串-KMP:", repeatedSubstringPattern2(s))
	log.Println("重复的子字符串-KMP优化:", repeatedSubstringPattern3(s))
}

// repeatedSubstringPattern 枚举暴力，时间复杂度O(n^2),空间复杂度O(1)
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	// i从1开始，从0开始没有意义
	// 重复字符串必然是从开头就开始算的，只需要判断前一半的长度是不是满足就行
	for i := 1; i*2 <= n; i++ {
		// 保证i在下一个重复串的开始位置
		if n%i != 0 {
			continue
		}

		// 假设是匹配的
		match := true
		// 逐个比较后边的字符
		for j := i; j < n; j++ {
			if s[j] != s[j-i] {
				match = false
				break
			}
		}

		if match {
			return true
		}
	}

	return false
}

// repeatedSubstringPattern1
func repeatedSubstringPattern1(s string) bool {
	// 1.将2个s拼接起来，去掉第一个和最后一个字符
	// 2.如果此时s是这个字符串的子串，则证明必然可以重复得到
	// 3.使用kmp也可以
	return strings.Contains((s + s)[1:len(s)*2-1], s)
}

// repeatedSubstringPattern2 kmp算法， 时间和空间都是O(n)
func repeatedSubstringPattern2(s string) bool {
	return kmp(s+s, s)
}

func kmp(query, pattern string) bool {
	n, m := len(query), len(pattern)
	fail := make([]int, m)
	for i := 0; i < m; i++ {
		fail[i] = -1
	}

	for i := 1; i < m; i++ {
		j := fail[i-1]
		for j != -1 && pattern[j+1] != pattern[i] {
			j = fail[j]
		}
		if pattern[j+1] == pattern[i] {
			fail[i] = j + 1
		}
	}

	match := -1
	for i := 1; i < n-1; i++ {
		for match != -1 && pattern[match+1] != query[i] {
			match = fail[match]
		}
		if pattern[match+1] == query[i] {
			match++
			if match == m-1 {
				return true
			}
		}
	}
	return false
}

// repeatedSubstringPattern3 优化后的kmp算法，时间和空间都是O(n)
func repeatedSubstringPattern3(s string) bool {
	return kmp3(s)
}

func kmp3(pattern string) bool {
	n := len(pattern)
	fail := make([]int, n)
	for i := 0; i < n; i++ {
		fail[i] = -1
	}
	for i := 1; i < n; i++ {
		j := fail[i-1]
		for j != -1 && pattern[j+1] != pattern[i] {
			j = fail[j]
		}
		if pattern[j+1] == pattern[i] {
			fail[i] = j + 1
		}
	}
	return fail[n-1] != -1 && n%(n-fail[n-1]-1) == 0
}
