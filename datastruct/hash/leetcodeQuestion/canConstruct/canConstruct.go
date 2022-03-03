package main

import (
	"log"
)

/**
题目：https://leetcode-cn.com/problems/ransom-note/

赎金信

给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。

提示：

1 <= ransomNote.length, magazine.length <= 105
ransomNote 和 magazine 由小写英文字母组成

注意：1.准备一个map一次判断即可，减少重复遍历
2.magazine 中的每个字符只能在 ransomNote 中使用一次。

*/
func main() {
	ransomNote := "aa"
	magazine := "ab"
	log.Println("赎金信-哈希map:", canConstruct(ransomNote, magazine))
	log.Println("赎金信-数组:", canConstruct1(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, v := range magazine {
		m[v]++
	}

	for _, v := range ransomNote {
		if _, ok := m[v]; !ok {
			return false
		}

		m[v]--
		if m[v] < 0 {
			return false
		}
	}

	return true
}

func canConstruct1(ransomNote, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	cnt := [26]int{}
	for _, ch := range magazine {
		cnt[ch-'a']++
	}

	for _, ch := range ransomNote {
		cnt[ch-'a']--
		if cnt[ch-'a'] < 0 {
			return false
		}
	}

	return true
}
