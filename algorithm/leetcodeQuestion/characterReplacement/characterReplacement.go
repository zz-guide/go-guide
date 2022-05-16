package main

import "log"

/**
题目：https://leetcode-cn.com/problems/longest-repeating-character-replacement/

替换后的最长重复字符

给你一个字符串 s 和一个整数 k 。你可以选择字符串中的任一字符，并将其更改为任何其他大写英文字符。该操作最多可执行 k 次。

在执行上述操作后，返回包含相同字母的最长子字符串的长度。

示例 1：

输入：s = "ABAB", k = 2
输出：4
解释：用两个'A'替换为两个'B',反之亦然。
示例 2：

输入：s = "AABABBA", k = 1
输出：4
解释：
将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
子串 "BBBB" 有最长重复字母, 答案为 4。

提示：

1 <= s.length <= 105
s 仅由大写英文字母组成
0 <= k <= s.length

*/
func main() {
	s := "AAAAAAA"
	k := 0
	log.Println("替换后的最长重复字符-双指针:", characterReplacement(s, k))
}

// characterReplacement 双指针
// 这道题只要求返回长度，不要求返回具体的字符串，所以只要求长度即可，不需要关心具体构建成什么字符串
func characterReplacement(s string, k int) int {
	cnt := [26]int{}     // 记录字符串中字符出现次数
	maxCnt, left := 0, 0 // maxCnt 代表最大长度(不管是哪个字符)
	for right, ch := range s {
		// 当前字符出现次数+1
		cnt[ch-'A']++
		// 记录当前字符出现了多少次
		maxCnt = max(maxCnt, cnt[ch-'A'])
		// right-left+1 代表left~right区间元素个数，因为left,right都从0开始
		areaLength := right - left + 1

		// 如果当前区间内可替换的长度大于k，代表需要缩小范围
		// 替换的过程如下，left向前移动，left位置的元素次数减一，left++
		if areaLength-maxCnt > k {
			// left位置的字符次数减一
			cnt[s[left]-'A']--
			// left向前移动
			left++
		}
	}

	// left~right区间就是最长长度，因为right最终要遍历完，所以用总长度减去left即可
	return len(s) - left
}

// 改版 字节跳动考题，会把大写字母换成只允许1，0
func characterReplacement2(s string, k int) int {
	// 其实也可以使用数组，为了通用，使用map
	mp := make(map[rune]int) // 记录字符串中字符出现次数
	maxCnt, left := 0, 0     // maxCnt 代表最大长度(不管是哪个字符)
	for right, ch := range s {
		// 当前字符出现次数+1
		mp[ch]++
		// 记录当前字符出现了多少次
		maxCnt = max(maxCnt, mp[ch])

		// right-left+1 代表left~right区间元素个数，因为left,right都从0开始
		areaLength := right - left + 1

		// 如果当前区间内可替换的长度大于k，代表需要缩小范围
		// 替换的过程如下，left向前移动，left位置的元素次数减一，left++
		if areaLength-maxCnt > k {
			// left位置的字符次数减一
			mp[rune(s[left])]--
			// left向前移动
			left++
		}
	}

	// left~right区间就是最长长度，因为right最终要遍历完，所以用总长度减去left即可
	return len(s) - left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
