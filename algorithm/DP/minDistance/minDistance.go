package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/edit-distance/
编辑距离


给你两个单词word1 和word2，请你计算出将word1转换成word2 所使用的最少操作数。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符

提示：

0 <= word1.length, word2.length <= 500
word1 和 word2 由小写英文字母组成

注意：1.每次只能操作一个字符
2.新增，删除，替换

*/
func main() {
	word1 := "intention"
	word2 := "execution"
	fmt.Println("编辑距离-动态规划:", minDistance(word1, word2))
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	//base case
	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}

	for j := 0; j < n+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				// 字符相同，保持状态
				dp[i][j] = dp[i-1][j-1]
			} else {
				//dp[i-1][j] 表示删除word1字符
				//dp[i][j-1]表示在word1插入字符
				//dp[i-1][j-1]表示替换word1一个字符
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
			}
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
