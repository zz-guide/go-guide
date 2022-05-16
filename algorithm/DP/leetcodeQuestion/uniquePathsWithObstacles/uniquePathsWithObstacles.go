package main

import "log"

/**
题目：https://leetcode-cn.com/problems/unique-paths-ii/

不同路径 II

一个机器人位于一个m x n网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

网格中的障碍物和空位置分别用 1 和 0 来表示。

提示：

m ==obstacleGrid.length
n ==obstacleGrid[i].length
1 <= m, n <= 100
obstacleGrid[i][j] 为 0 或 1

*/
func main() {
	obstacleGrid := [][]int{{0, 0}, {1, 1}, {0, 0}}
	log.Println("不同路径 II-动态规划:", uniquePathsWithObstacles(obstacleGrid))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		// 遇到障碍物，后边的都走不通
		if obstacleGrid[i][0] == 1 {
			break
		}

		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		// 遇到障碍物，后边的都走不通
		if obstacleGrid[0][j] == 1 {
			break
		}

		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
