package main

import "log"

/**
题目：https://leetcode-cn.com/problems/longest-valid-parentheses/

最长有效括号
给你一个只包含 '('和 ')'的字符串，找出最长有效（格式正确且连续）括号子串的长度。

示例 1：

输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"
示例 2：

输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"
示例 3：

输入：s = ""
输出：0

提示：

0 <= s.length <= 3 * 104
s[i] 为 '(' 或 ')'


注意：
	1.括号只有圆括号,没有其他字符
	2.可以思考一下如果同时有其他括号呢?
	3."()(())" 也认为是正确的
*/
func main() {
	s := "()(())"
	log.Println("最长有效括号-动态规划:", longestValidParentheses(s))
	log.Println("最长有效括号-栈:", longestValidParentheses1(s))
	log.Println("最长有效括号-双指针正逆序:", longestValidParentheses2(s))
}

// longestValidParentheses 动态规划 O(n) O(n)
func longestValidParentheses(s string) int {
	// 4种场景：https://leetcode-cn.com/problems/longest-valid-parentheses/solution/32-zui-chang-you-xiao-gua-hao-by-chen-we-dp1n/
	maxAns := 0
	dp := make([]int, len(s)) // dp[i]表示以i结尾的最长有效括号的长度
	dp[0] = 0                 // 初始化
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				// ()() 场景
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					// () 场景
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				// ()(())场景,末尾表示i
				// () 末尾表示dp[i-dp[i-1]-2]
				// ()( 末尾表示s[i-dp[i-1]-1]
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					// (()) 场景，末尾表示i
					dp[i] = dp[i-1] + 2
				}
			}
			maxAns = max(maxAns, dp[i])
		}
	}
	return maxAns
}

// longestValidParentheses1 栈 动态规划 O(n) O(n)
func longestValidParentheses1(s string) int {
	maxAns := 0
	var stack []int
	stack = append(stack, -1) // 栈里存放括号对应的索引位置，用来计算长度，而非括号自身
	for i := 0; i < len(s); i++ {
		// 左括号直接入栈
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			// 右括号出栈
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				// 说明没有匹配的
				stack = append(stack, i)
			} else {
				// 说明有左边界
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}
	return maxAns
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// longestValidParentheses2 正向逆向结合 优化 O(n) O(1)
func longestValidParentheses2(s string) int {
	// 双指针正向， "(()"的情况无法求得
	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}

	// 双指针逆向， "())"的情况无法求得,"()",")(()"
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}

	return maxLength
}
