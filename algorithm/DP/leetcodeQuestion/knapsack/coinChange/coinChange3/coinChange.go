package main

import (
	"fmt"
)

/**
题目：https://leetcode-cn.com/problems/coin-change/
零钱兑换:给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的数量无限，再给一个总金额 amount，问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1 。
*/
func main() {
	coins := []int{1, 2, 5}
	amount := 11

	fmt.Println("兑换结果:", coinChange(coins, amount))
}

// dp数组的定义：当目标金额为i时，至少需要dp[i]枚硬币凑出
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	//初始化 数组大小为amount+1， 初始值也为amount+1
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	//base case
	dp[0] = 0
	//外层for循环遍历所有状态的所有取值
	for i := 1; i < len(dp); i++ {
		//内层循环求所有选择的最小值
		for _, coin := range coins {
			//子问题无解 跳过
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
