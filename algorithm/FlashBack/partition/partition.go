package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/palindrome-partitioning/
分割回文串

给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

回文串 是正着读和反着读都一样的字符串。

提示：

1 <= s.length <= 16
s 仅由小写英文字母组成


注意：1.分隔方案必须完整，互不影响
2.单独判断各自分隔
3.字符串自身不算子串。
4.分两步，截取子串，确定区间，l,r;然后判断l,r之间是不是回文即可

*/
func main() {
	s := "aab"
	fmt.Println("分割回文串-回溯 + 记忆化搜索:", partition(s))
	fmt.Println("分割回文串-回溯 + 动态规划预处理:", partition1(s))
}

// 回溯 + 记忆化搜索
func partition(s string) [][]string {
	length := len(s)
	var res [][]string
	var isPalindrome func(l, r int) bool
	isPalindrome = func(l, r int) bool {
		for l <= r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	var track []string
	var backtracking func(start int)
	backtracking = func(start int) {
		if start == length {
			res = append(res, append([]string{}, track...))
			return
		}

		// 从当前位置遍历整个字符串
		for i := start; i < length; i++ {
			// 是回文子串，就继续递归，直到遍历到字符串末尾
			if isPalindrome(start, i) {
				track = append(track, s[start:i+1])
				backtracking(i + 1)
				track = track[:len(track)-1]
			}
		}
	}

	backtracking(0)
	return res
}

// partition1 回溯 + 动态规划预处理
func partition1(s string) [][]string {
	var res [][]string
	n := len(s)
	f := make([][]bool, n)
	// 判断回文子串的逻辑使用动态规划暂存起来
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	var splits []string
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			res = append(res, append([]string(nil), splits...))
			return
		}

		for j := i; j < n; j++ {
			if f[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}

	dfs(0)
	return res
}
