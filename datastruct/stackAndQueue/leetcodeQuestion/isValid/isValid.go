package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/valid-parentheses/

有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成

*/
func main() {
	s := "{[]}"
	fmt.Println("有效的括号-栈:", isValid(s))
}

// isValid 用一个栈就够了，也可以使用map暂存括号信息，减少==判断
func isValid(s string) bool {
	var stack []rune
	for _, v := range s {
		if v == '{' || v == '(' || v == '[' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if (top == '(' && v == ')') || (top == '{' && v == '}') || (top == '[' && v == ']') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
