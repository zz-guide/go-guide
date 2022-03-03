package main

import (
	"fmt"
	"math/big"
)

/**
题目：https://leetcode-cn.com/problems/unique-paths/
不同路径

一个机器人位于一个 m x n网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

提示：

1 <= m, n <= 100
题目数据保证答案小于等于 2 * 109

*/
func main() {
	m := 3
	n := 2
	fmt.Println("不同的路径-dp:", uniquePaths(m, n))
	fmt.Println("不同的路径-数论:", uniquePaths1(m, n))
}

// uniquePaths dp
func uniquePaths(m int, n int) int {
	// 1.定义递推公式，dp[i][j],i代表行，j代表列，dp[i][j]就代表到第i行，j列的不同路径
	// 2.从dp[0][0]=0

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// uniquePaths1 数论
func uniquePaths1(m, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}
