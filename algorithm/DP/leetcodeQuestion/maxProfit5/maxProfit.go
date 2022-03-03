package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

买卖股票的最佳时机含手续费

给定一个整数数组prices，其中第i个元素代表了第i天的股票价格 ；整数fee 代表了交易股票的手续费用。

你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。

返回获得利润的最大值。

注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。

提示：

1 <= prices.length <= 5 * 104
1 <= prices[i] < 5 * 104
0 <= fee < 5 * 104

注意：增加了手续费


*/
func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fee := 2
	fmt.Println("买卖股票的最佳时机含手续费-动态规划：", maxProfit(prices, fee))
}

func maxProfit(prices []int, fee int) int {
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
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}

	return dp[length-1][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
