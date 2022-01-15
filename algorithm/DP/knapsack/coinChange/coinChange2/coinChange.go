package main

import (
	"fmt"
	"math"
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

//带备忘录的递归
var memo []int

func coinChange(coins []int, amount int) int {
	memo = make([]int, amount+1)
	//初始化memo
	for i := 0; i < amount+1; i++ {
		memo[i] = -666
	}
	return dp(coins, amount)
}

// 备忘录版本
func dp(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	//查备忘录 防止重复计算
	if memo[amount] != -666 {
		return memo[amount]
	}
	res := math.MaxInt32
	for _, coin := range coins {
		//計算子问题的结果
		subP := dp(coins, amount-coin)
		//子问题无解则跳过
		if subP == -1 {
			continue
		}
		//在子问题中选择最优解 然后加一
		res = min(res, subP+1)
	}
	// 把计算结果存入备忘录 不然无法跳出子问题的递归调用
	if res == math.MaxInt32 {
		memo[amount] = -1
	} else {
		memo[amount] = res
	}
	return memo[amount]
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
