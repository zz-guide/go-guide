package main

import "fmt"

/**
Sunday算法由Daniel M.Sunday在1990年提出，是一种字符串模式匹配算法，能够以O(n)的时间复杂度下完成子串索引。
Sunday算法从前往后匹配，在匹配失败时关注的是文本串中参加匹配的最末位字符的下一位字符。
如果该字符没有在模式串中出现则直接跳过，即移动位数 = 匹配串长度 + 1；
否则，其移动位数 = 模式串中最右端的该字符到末尾的距离+1。

Sunday算法最巧妙的地方，就在于它发现匹配失败之后可以直接考察文本串中参加匹配的最末尾字符的下一个字符
*/
func main() {
	str := "ababd"
	target := "abd"
	fmt.Println("Sunday算法查找:", SundaySearch(str, target))
}

func SundaySearch(str string, target string) int {
	if len(target) == 0 {
		return -1
	}

	if len(str) == 0 || len(str) < len(target) {
		return -1
	}

	mStr := []rune(str)
	subStr := []rune(target)

	// 记录子串每个字符出现的最右侧位置
	m := make(map[rune]int)
	for i := 0; i < len(subStr); i++ {
		m[subStr[i]] = i
	}

	// 长度变量
	mStrLength, subStrLength := len(mStr), len(subStr)
	mStrStartIndex := 0
	subStrStartIndex := 0
	redirectIndex := 0 // 下一次主串开始遍历的位置

	// 遍历主串
	for mStrStartIndex < mStrLength {
		// 主串和子串字符相等，继续比较
		if mStr[mStrStartIndex] == subStr[subStrStartIndex] {
			// 发现子串已经比较到末尾了,说明完全匹配
			if subStrStartIndex == subStrLength-1 {
				// 此时mStrStartIndex在末尾，减去子串的长度+1刚好是开始的位置
				return mStrStartIndex - subStrLength + 1
			}

			mStrStartIndex++
			subStrStartIndex++
		} else {
			// 若不相等，则开始移动，先重置子串比较的位置
			subStrStartIndex = 0

			next := redirectIndex + subStrLength
			// 已经到了主串的末尾，无法再跳转匹配了
			if next >= mStrLength {
				return -1
			}

			// 此时主串中子串末尾位置的下一个字符如果在 子串中存在，则把子串的开头对应到这个字符的位置
			if existsIndex, ok := m[mStr[next]]; ok {
				redirectIndex += subStrLength - existsIndex
			} else {
				// 不存在就跳过这个字符，子串的开头位置对应到这个不存在字符的下一个位置
				redirectIndex += subStrLength
			}

			mStrStartIndex = redirectIndex
		}
	}

	return -1
}
