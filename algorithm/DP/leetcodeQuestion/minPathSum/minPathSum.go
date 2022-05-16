package main

import "log"

/**
题目：https://leetcode-cn.com/problems/minimum-path-sum/

最小路径和

给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
示例 2：

输入：grid = [[1,2,3],[4,5,6]]
输出：12


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 100

*/
func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	log.Println("最小路径和-动态规划 二维:", minPathSum(grid))
	log.Println("最小路径和-动态规划 一维:", minPathSum2(grid))
}

// minPathSum O(mn) O(mn)
func minPathSum(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	dp := make([][]int, rows)
	for i, _ := range dp {
		dp[i] = make([]int, cols)
	}

	dp[0][0] = grid[0][0]
	// 初始化最左侧列,比较特殊只能从上层往下走，没得选
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// 初始化顶层行，只能从左向右，没得选
	for j := 1; j < cols; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	// 剩余表格部分，可以从上方，左侧分别行走，求最小即可
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[rows-1][cols-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minPathSum2(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	dp := make([]int, cols)
	dp[0] = grid[0][0]
	for j := 1; j < cols; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}

	// 剩余表格部分，可以从上方，左侧分别行走，求最小即可
	for i := 1; i < rows; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < cols; j++ {
			dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
		}
	}

	return dp[cols-1]
}
