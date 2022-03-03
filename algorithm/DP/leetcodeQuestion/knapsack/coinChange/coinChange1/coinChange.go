package main

import (
	"log"
	"math"
)

/**
题目：https://leetcode-cn.com/problems/coin-change/
零钱兑换:给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的数量无限，再给一个总金额 amount，问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1 。
(完全背包)
*/
func main() {
	coins := []int{1, 2, 5}
	amount := 11

	log.Println("兑换结果:", coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {
	return dp(coins, amount)
}

/**
DP方程：
dp(n) = [
	0, n = 0
	-1, n <0
	min{dp(n- coin) + 1 | coin , coins}, n >0
]
*/
func dp(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	res := math.MaxInt // 表示凑不齐
	for _, coin := range coins {
		// 拆分子问题
		count := dp(coins, amount-coin) // 可以优化成DP数组
		if count == -1 {
			continue
		}

		// subProblem表示amount=coin，次数应该加1
		res = min(res, count+1)
	}

	// 说明凑不齐
	if res == math.MaxInt {
		return -1
	}

	return res
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
