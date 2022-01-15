package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

最佳买卖股票时机含冷冻期

给定一个整数数组，其中第i个元素代表了第i天的股票价格 。

设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

注意：增加了冷冻期，意味着不能连续买入卖出


*/
func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println("最佳买卖股票时机含冷冻期-动态规划：", maxProfit(prices))
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	dp := make([][]int, n)
	status := make([]int, n*4)
	for i := range dp {
		dp[i] = status[:4]
		status = status[4:]
	}
	dp[0][0] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][1]-prices[i], dp[i-1][3]-prices[i]))
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}

	return max(dp[n-1][1], max(dp[n-1][2], dp[n-1][3]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
