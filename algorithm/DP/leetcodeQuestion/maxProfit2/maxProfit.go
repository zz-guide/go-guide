package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/

买卖股票的最佳时机 III

给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成两笔交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

提示：

1 <= prices.length <= 105
0 <= prices[i] <= 105

注意：能交易2次

*/
func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println("买卖股票的最佳时机III-动态规划：", maxProfit(prices))
}

// maxProfit 动态规划
func maxProfit(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
