package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

买卖股票的最佳时机 II
给定一个数组 prices ，其中prices[i] 是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

提示：

1 <= prices.length <= 3 * 104
0 <= prices[i] <= 104



注意：能交易多次，不限制

*/
func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("买卖股票的最佳时机II-贪心算法：", maxProfit(prices))
	fmt.Println("买卖股票的最佳时机II-动态规划：", maxProfit2(prices))
}

// maxProfit 贪心算法
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var res int
	for i := 1; i < len(prices); i++ {
		res += max(0, prices[i]-prices[i-1])
	}

	return res
}

// maxProfit2 动态规划
// dp[i][0]表示第i天没卖出股票的最大利润，dp[i][1]表示第i天卖出股票的最大利润
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	length := len(prices)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		// 与I不同的地方，  dp[i-1][1] - prices[i]
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[length-1][1]
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
