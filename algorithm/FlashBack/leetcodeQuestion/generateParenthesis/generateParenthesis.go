package main

import "log"

/**
题目：https://leetcode-cn.com/problems/generate-parentheses/

括号生成

数字 n代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。


示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]


提示：

1 <= n <= 8

*/

func main() {
	n := 3
	log.Println("括号生成:", generateParenthesis(n))

}

// generateParenthesis 回溯
func generateParenthesis(n int) []string {
	var res []string
	var backtrack func(left, right int, item string)
	backtrack = func(left, right int, item string) {
		if left == right && left == n {
			res = append(res, item)
			return
		}

		if right > left || left > n || right > n { //不合法情况，直接返回
			return
		}

		backtrack(left+1, right, item+"(")
		backtrack(left, right+1, item+")")
	}

	backtrack(0, 0, "")
	return res
}
