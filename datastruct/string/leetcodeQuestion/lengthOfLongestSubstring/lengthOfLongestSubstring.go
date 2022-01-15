package main

import (
	"fmt"
)

/**
题目：无重复字符串的最长子串

示例1:
输入: s6 = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s6 = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s6 = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
示例 4:

输入: s6 = ""
输出: 0
提示：

0 <= s6.length <= 5 * 104
s由英文字母、数字、符号和空格组成
*/
func main() {
	s := "pwwkew"
	//s6 := "qwertyu"
	count := F1(s)
	fmt.Println("无重复字符最长子串的长度：", count)
}

//F1
//时间复杂度：O(N)O(N)，其中 NN 是字符串的长度。左指针和右指针分别会遍历整个字符串一次。
//空间复杂度：O(|\Sigma|)O(∣Σ∣)，其中 \SigmaΣ 表示字符集（即字符串中可以出现的字符），|\Sigma|∣Σ∣ 表示字符集的大小。在本题中没有明确说明字符集，因此可以默认为所有 ASCII 码在 [0, 128)[0,128) 内的字符，即 |\Sigma| = 128∣Σ∣=128。我们需要用到哈希集合来存储出现过的字符，而字符最多有 |\Sigma|∣Σ∣ 个，因此空间复杂度为 O(|\Sigma|)O(∣Σ∣)。
//**
func F1(s string) int {
	// 1. 定义用于存放已遍历字符的map
	m := map[rune]int{}
	// 2. 定义起点索引,定义最大长度
	left, maxLength := 0, 0

	// 4. 按序遍历字符串中所有字符
	for right, v := range []rune(s) {
		// 5. 发现当前字符已经出现，则移动left到old的下一个位置
		if oldRuneIndex, ok := m[v]; ok && oldRuneIndex >= left {
			// 如果出现过就将开始位置往后挪一个
			left = oldRuneIndex + 1
		}

		// 6. 计算窗口长度,原理：right-left+1=长度，跟求数组长度是一个意思
		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}

		// 7. 记录遍历过的字符 【字符，index】
		m[v] = right
	}

	return maxLength
}
