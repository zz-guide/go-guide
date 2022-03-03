package main

import (
	"fmt"
)

/**
题目：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

买卖股票的最佳时机
给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
提示：

1 <= prices.length <= 105
0 <= prices[i] <= 104


注意：只能交易一次

*/
func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("买卖股票的最佳时机-贪心算法：", maxProfit(prices))
	fmt.Println("买卖股票的最佳时机-暴力解法：", maxProfit1(prices))
	fmt.Println("买卖股票的最佳时机-动态规划：", maxProfit2(prices))
}

// maxProfit 贪心算法，遍历过程中计算之前的最小值和当前值的差，依次求最大就是最大利润
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	intSize := 32 << (^uint(0) >> 63)
	var res int
	var low int = 1<<(intSize-1) - 1 // 最大值
	for i := 0; i < len(prices); i++ {
		low = min(low, prices[i])
		if prices[i]-low < 0 {
			continue
		}
		res = max(prices[i]-low, res)
	}

	return res
}

// maxProfit1 暴力法
func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var res int
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] < 0 {
				continue
			}

			res = max(res, prices[j]-prices[i])
		}
	}

	return res
}

// maxProfit2 动态规划
// dp[i][0]表示第i天没卖出股票的最大利润，dp[i][1]表示第i天卖出股票的最大利润
// 持有股票的意思就是股票还在手上没卖，可能今天花钱了，也可能没花钱（老早之前买的）
// 今天股票在手上的最大利润=max((昨天股票在手上的最大利润，今天不买也不卖), (今天买入花钱))
// 今天卖出股票的最大利润=max(昨天卖出股票的最大利润(今天不买不卖), 今天卖出股票)
// dp[i][0] 是给下一天计算最大利润使用的
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
		dp[i][0] = max(dp[i-1][0], -prices[i])
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
