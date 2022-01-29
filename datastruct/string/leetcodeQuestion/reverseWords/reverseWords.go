package main

import (
	"log"
)

/**
题目：https://leetcode-cn.com/problems/reverse-words-in-a-string

翻转字符串里的单词

给你一个字符串 s ，逐个翻转字符串中的所有 单词 。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

请你返回一个翻转 s 中单词顺序并用单个空格相连的字符串。

说明：

输入字符串 s 可以在前面、后面或者单词间包含多余的空格。
翻转后单词间应当仅用一个空格分隔。
翻转后的字符串中不应包含额外的空格。

提示：

1 <= s.length <= 104
s 包含英文大小写字母、数字和空格 ' '
s 中 至少存在一个 单词

进阶：
请尝试使用O(1) 额外空间复杂度的原地解法。

注意：1.要求单词和单词调换位置，单词本身不反转

所以解题思路如下：
移除多余空格
将整个字符串反转
将每个单词反转
举个例子，源字符串为："the sky is blue "
移除多余空格 : "the sky is blue"
字符串反转："eulb si yks eht"
单词反转："blue is sky the"
这样我们就完成了翻转字符串里的单词。

*/
func main() {
	s := "the sky is blue"
	log.Println("翻转字符串里的单词-双指针:", reverseWords(s))
}

func reverseWords(s string) string {
	ss := []byte(s)
	length := len(ss)
	slow, fast := 0, 0
	// 去掉字符串最左边的冗余空格
	for length > 0 && fast < length && ss[fast] == ' ' {
		fast++
	}
	// 去掉单词之间的冗余空格
	for ; fast < length; fast++ {
		if fast > 1 && ss[fast] == ss[fast-1] && ss[fast] == ' ' {
			continue
		}
		ss[slow] = ss[fast]
		slow++
	}
	// 去掉字符串最右边的冗余空格
	if slow > 1 && ss[slow-1] == ' ' {
		ss = ss[:slow-1]
	} else {
		ss = ss[:slow]
	}
	// 反转整个字符串
	reverse(ss)
	i := 0
	for i < len(ss) {
		// 反转单个单词
		j := i
		// 找到单词的结束位置
		for j < len(ss) && ss[j] != ' ' {
			j++
		}
		// 反转
		reverse(ss[i:j])
		// 更新下一个单词的起始位置，+1是要跳过单词间的空格
		i = j + 1
	}
	return string(ss)
}

func reverse(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		tmp := s[n-1-i]
		s[n-1-i] = s[i]
		s[i] = tmp
	}
}
