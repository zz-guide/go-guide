package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/coin-change-2/
给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
请你计算并返回可以凑成总金额的硬币组合数。
如果任何硬币组合都无法凑出总金额，返回 0 。
假设每一种面额的硬币有无限个。
题目数据保证结果符合 32 位带符号整数。

提示：

1 <= coins.length <= 300
1 <= coins[i] <= 5000
coins 中的所有值 互不相同
0 <= amount <= 5000
*/
func main() {
	amount := 0
	coins := []int{1, 2, 5}
	fmt.Println("组合数(二维数组):", coinCombinationNum1(amount, coins))
	fmt.Println("组合数(一维数组):", coinCombinationNum2(amount, coins))
}

// coinCombinationNum1 二维数组
func coinCombinationNum1(amount int, coins []int) int {
	//无限个
	n := len(coins)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, amount+1)
		f[i][0] = 1
	}
	//f[i][j] 前i个硬币,组成j的方法数
	for i := 1; i <= n; i++ {
		a := coins[i-1]
		for j := 0; j <= amount; j++ {
			f[i][j] = f[i-1][j]
			if j-a >= 0 {
				f[i][j] += f[i][j-a]
			}
		}
	}
	return f[n][amount]
}

// coinCombinationNum2 一维数组
// https://leetcode-cn.com/problems/coin-change-2/solution/hua-tu-li-jie-cong-chang-gui-er-wei-dpda-4gfy/
func coinCombinationNum2(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1 // 剩余钱的组合数，amount=0是1种组合，只有当使用coin和不使用coin都凑不出来的时候就返回0

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	fmt.Println("dp:", dp)
	return dp[amount]
}
