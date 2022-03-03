package main

import "log"

/**
题目: https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

无重复字符的最长子串

给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。


示例1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。

请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。

提示：

0 <= s.length <= 5 * 104
s由英文字母、数字、符号和空格组成

*/

func main() {
	s := "abcabcbb"
	log.Println("无重复字符的最长子串-双指针:", lengthOfLongestSubstring1(s))
	log.Println("无重复字符的最长子串-滑动窗口+哈希:", lengthOfLongestSubstring2(s))
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

// lengthOfLongestSubstring1 双指针
func lengthOfLongestSubstring1(s string) int {
	// 思路：end逐位移动，内层循环比较end位置与start~end-1位置字符是否相同，若相同，start移动置i+1位置，继续向后遍历，直到末尾
	var maxLength int
	start, end := 0, 0
	for ; end < len(s); end++ {
		// 逐个比较end 与 start~end-1 区间的字符
		for i := start; i < end; i++ {
			// 若相同，计算长度然后修改start
			if s[i] == s[end] {
				maxLength = max(maxLength, end-start)
				start = i + 1
				break
			}
		}
	}

	// 最后别忘了再比较一次
	maxLength = max(maxLength, end-start)
	return maxLength
}

// lengthOfLongestSubstring2 滑动窗口+哈希
// 思路：map存每一个字符，最近一次出现的位置，向后遍历的同时
func lengthOfLongestSubstring2(s string) int {
	// 1. 定义用于存放已遍历字符的map
	m := map[rune]int{}
	var startIndex int
	var maxLength int

	// 2. 按序遍历字符串中所有字符
	for i, v := range []rune(s) {
		// 3. 记录是否出现过字符的index
		if lastI, ok := m[v]; ok && lastI >= startIndex {
			// 如果当前字符在之前出现过了，说明不能继续向后遍历，需要把起始遍历位置放到上此位置的下一个位置
			startIndex = lastI + 1
		}

		// 4. 记录遍历过的字符 【字符，最近一次出现的位置】
		m[v] = i

		// 5. 遇到不同的字符就计算长度
		if i-startIndex+1 >= maxLength {
			maxLength = i - startIndex + 1
		}
	}

	return maxLength
}
