package main

import (
	"log"
)

/**
题目：https://leetcode-cn.com/problems/check-permutation-lcci/

判定是否互为字符重排

给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

示例 1：

输入: s1 = "abc", s2 = "bca"
输出: true
示例 2：

输入: s1 = "abc", s2 = "bad"
输出: false
说明：

0 <= len(s1) <= 100
0 <= len(s2) <= 100

*/
func main() {
	s1 := "aab"
	s2 := "abb"
	log.Println("判定是否互为字符重排-位运算:", CheckPermutation(s1, s2))
	log.Println("判定是否互为字符重排-哈希:", CheckPermutation1(s1, s2))
}

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var (
		res  byte = 0
		bit1 byte = 0
		bit2 byte = 0
	)

	for i := 0; i < len(s1); i++ {
		// 可以保证数量是一样的,偶数次数出现，但会有特殊情况
		res = res ^ s1[i] ^ s2[i]
		// 只能保证出现过相同的字符，无法保证数量
		bit1 = bit1 | 1<<(s1[i]-'a')
		bit2 = bit2 | 1<<(s2[i]-'a')
	}

	return res == 0 && bit1 == bit2
}

func CheckPermutation1(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	mapList := make(map[rune]int)
	for _, char := range s1 {
		mapList[char]++
	}

	for _, char := range s2 {
		mapList[char]--
	}

	for _, value := range mapList {
		if value != 0 {
			return false
		}
	}

	return true
}
