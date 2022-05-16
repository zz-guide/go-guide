package main

import (
	"log"
	"strconv"
)

/**
题目：https://leetcode-cn.com/problems/count-and-say/solution/wai-guan-shu-lie-by-leetcode-solution-9rt8/

外观数列

给定一个正整数 n ，输出外观数列的第 n 项。

「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。

你可以将其视作是由递归公式定义的数字字符串序列：

countAndSay(1) = "1"
countAndSay(n) 是对 countAndSay(n-1) 的描述，然后转换成另一个数字字符串。

提示：

1 <= n <= 30

*/
func main() {
	n := 4
	log.Println("外观数列-递归:", countAndSay(n))
	log.Println("外观数列-迭代:", countAndSay1(n))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	s := countAndSay(n - 1)

	i, res := 0, "" // i表示重复字符
	for j, c := range s {
		// 发现当前字符和s[i]不重复的时候，说明j-i区间都是i
		if c != rune(s[i]) { //要注意 rune 与 byte 的格式的数据不能直接进行比较
			res += strconv.Itoa(j-i) + string(s[i])
			i = j
		}
	}

	// 此时i是最后一个字符，还没有算到结果里边，需要最后再处理
	res += strconv.Itoa(len(s)-i) + string(s[i])
	return res
}

func countAndSay1(n int) string {
	s := "1"
	for i := 0; i < n-1; i++ {
		j, tmp := 0, ""
		for k, c := range s {
			if c != rune(s[j]) {
				tmp += strconv.Itoa(k-j) + string(s[j])
				j = k
			}
		}
		s = tmp + strconv.Itoa(len(s)-j) + string(s[j])
	}

	return s
}
