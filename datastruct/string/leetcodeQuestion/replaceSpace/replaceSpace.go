package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/

替换空格

请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

限制：

0 <= s 的长度 <= 10000

*/

func main() {
	s := "We are happy."
	fmt.Println("替换空格:", replaceSpace(s))
}

func replaceSpace(s string) string {
	var str string
	for _, v := range s {
		if v == ' ' {
			str += "%20"
		} else {
			str += string(v)
		}
	}
	return str
}
