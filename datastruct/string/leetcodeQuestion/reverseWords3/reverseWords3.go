package main

import (
	"log"
)

/**
题目：https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/

反转字符串中的单词 III

给定一个字符串s，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。


示例 1：

输入：s = "Let's take LeetCode contest"
输出："s'teL ekat edoCteeL tsetnoc"
示例 2:

输入： s = "God Ding"
输出："doG gniD"

提示：

1 <= s.length <= 5 * 104
s包含可打印的 ASCII 字符。
s不包含任何开头或结尾空格。
s里 至少 有一个词。
s中的所有单词都用一个空格隔开。

注意：1.要求反转单词中的字符


*/
func main() {
	s := "the sky is blue"
	log.Println("翻转字符串里的单词III-原地交换:", reverseWords(s))
	log.Println("翻转字符串里的单词III-额外空间:", reverseWords1(s))
}

// reverseWords 原地交换 空间复杂度O(1)，时间复杂度O(n)
func reverseWords(s string) string {
	str := []rune(s)
	f := 0
	for f < len(str) {
		l := f
		for f < len(str) && str[f] != ' ' {
			f++
		}
		r := f - 1

		for l < r {
			str[l], str[r] = str[r], str[l]
			l++
			r--
		}
		f++
	}
	return string(str)
}

// reverseWords1 额外开辟空间，空间复杂度O(1)，时间复杂度O(n)
func reverseWords1(s string) string {
	length := len(s)
	var ret []byte
	for i := 0; i < length; {
		start := i
		for i < length && s[i] != ' ' {
			i++
		}
		for p := start; p < i; p++ {
			ret = append(ret, s[start+i-1-p])
		}
		for i < length && s[i] == ' ' {
			i++
			ret = append(ret, ' ')
		}
	}
	return string(ret)
}
