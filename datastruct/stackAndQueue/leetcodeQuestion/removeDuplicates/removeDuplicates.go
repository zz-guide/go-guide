package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/valid-parentheses/

删除字符串中的所有相邻重复项

给出由小写字母组成的字符串S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。

在 S 上反复执行重复项删除操作，直到无法继续删除。

在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

提示：

1 <= S.length <= 20000
S 仅由小写英文字母组成。

*/
func main() {
	s := "abbaca"
	fmt.Println("删除字符串中的所有相邻重复项-栈:", removeDuplicates(s))
}

func removeDuplicates(s string) string {
	var stack []rune
	for _, v := range s {
		if len(stack) > 0 && stack[len(stack)-1] == v {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return string(stack)
}
