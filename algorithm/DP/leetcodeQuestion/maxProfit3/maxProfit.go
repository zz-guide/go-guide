package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/

买卖股票的最佳时机 IV

给定一个整数数组prices ，它的第 i 个元素prices[i] 是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

提示：

0 <= k <= 100
0 <= prices.length <= 1000
0 <= prices[i] <= 1000


注意：可以交易k次


*/
func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	k := 3
	fmt.Println("买卖股票的最佳时机IV-动态规划：", maxProfit(k, prices))
}

func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))
	status := make([]int, (2*k+1)*len(prices))
	for i := range dp {
		dp[i] = status[:2*k+1]
		status = status[2*k+1:]
	}
	for j := 1; j < 2*k; j += 2 {
		dp[0][j] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k; j += 2 {
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}
	return dp[len(prices)-1][2*k]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
