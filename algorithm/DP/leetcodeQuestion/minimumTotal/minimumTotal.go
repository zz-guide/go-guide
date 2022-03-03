package main

import (
	"log"
)

/**
题目：https://leetcode-cn.com/problems/triangle/

三角形最小路径和

给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。


示例 1：

输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：
   2
  3 4
 6 5 7
4 1 8 3
自顶向下的最小路径和为11（即，2+3+5+1= 11）。
示例 2：

输入：triangle = [[-10]]
输出：-10


提示：

1 <= triangle.length <= 200
triangle[0].length == 1
triangle[i].length == triangle[i - 1].length + 1
-104 <= triangle[i][j] <= 104

进阶：

你可以只使用 O(n)的额外空间（n 为三角形的总行数）来解决这个问题吗？

*/
func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	log.Println("三角形最小路径和-动态规划 二维:", minimumTotal(triangle))
	log.Println("三角形最小路径和-动态规划 一维:", minimumTotal2(triangle))
}

// minimumTotal 从底部向上开始dp  O(n^2) O(n^2)
func minimumTotal(triangle [][]int) int {
	rows := len(triangle)
	dp := make([][]int, rows) // dp[i][j]表示从底部向上选到第i行，第j位的最小路径
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, rows)
	}

	for i := rows - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			// 最后一行
			if i == rows-1 {
				dp[i][j] = triangle[i][j] // 最后一行先选，最小路径就等于选择自己
			} else {
				// 从上一行的j,j+1位置取最小+三角形当前位置的值
				dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
			}
		}
	}

	return dp[0][0]
}

// minimumTotal2 优化成一维 O(n^2) O(n)
func minimumTotal2(triangle [][]int) int {
	rows := len(triangle)
	bottom := triangle[rows-1] // 三角形最底层的元素
	dp := make([]int, len(bottom))
	for i := range dp {
		dp[i] = bottom[i]
	}

	// 最底层的值已经加到dp了，所以直接从rows-2开始
	for i := rows - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
