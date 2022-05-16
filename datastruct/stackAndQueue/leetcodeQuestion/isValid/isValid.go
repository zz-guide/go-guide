package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/valid-parentheses/

有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。

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
	fmt.Println("有效的括号-栈:", isValid1(s))
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

// isValid1 时间复杂度O(n) 空间复杂度可以优化
func isValid1(s string) bool {
	// 奇数肯定不匹配
	if len(s)%2 == 1 {
		return false
	}

	// 此处map可以优化，比如用if,else
	mp := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var stack []rune
	for _, v := range s {
		// 栈要想比较的话，必须得有元素，而且至少得有一个
		if len(stack) == 0 {
			stack = append(stack, v)
		} else {
			// 出栈
			top := stack[len(stack)-1]
			// 判断出栈的元素是不是左半部分，如果是的话，并且当前元素是右半部分，那么肯定匹配，就更新栈
			if temp, ok := mp[top]; ok && temp == v {
				stack = stack[:len(stack)-1]
			} else {
				// 否则就加到栈里，因为匹配的话只能是相邻才能匹配
				stack = append(stack, v)
			}
		}
	}

	return len(stack) == 0
}
