package main

import (
	"fmt"
	"sort"
)

/**
题目：https://leetcode-cn.com/problems/valid-anagram/

有效的字母异位词

给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。

提示:

1 <= s.length, t.length <= 5 * 104
s 和 t仅包含小写字母


进阶:如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

注意:1.两个字符串必须包含相同的字母，且字母出现次数一致
2.等价于一个字符串打乱字母顺序之后互为异位词

*/
func main() {
	s := "ab"
	t := "a"
	fmt.Println("有效的字母异位词-2个哈希:", isAnagram(s, t))
	fmt.Println("有效的字母异位词-排序:", isAnagram(s, t))
	fmt.Println("有效的字母异位词-数组模拟哈希:", isAnagram2(s, t))
	fmt.Println("有效的字母异位词-一个哈希:", isAnagram3(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := make(map[string]int)
	for _, v := range s {
		v1 := string(v)
		if _, ok := m1[v1]; ok {
			m1[v1] += 1
		} else {
			m1[v1] = 1
		}
	}

	m2 := make(map[string]int)
	for _, v := range t {
		v1 := string(v)
		if _, ok := m2[v1]; ok {
			m2[v1] += 1
		} else {
			m2[v1] = 1
		}
	}

	for k, v := range m1 {
		if _, ok := m2[k]; ok && m2[k] == v {
			continue
		}

		return false
	}

	return true
}

// 排序
func isAnagram1(s, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	// 相同的规则排序之后相等就是
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}

// isAnagram2 如果只有小写字母的话可以这么写
func isAnagram2(s, t string) bool {
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
}

func isAnagram3(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}

	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}

	return true
}
