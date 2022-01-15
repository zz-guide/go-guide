package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/min-cost-climbing-stairs/
使用最小花费爬楼梯

数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。

每当你爬上一个阶梯你都要花费对应的体力值，一旦支付了相应的体力值，你就可以选择向上爬一个阶梯或者爬两个阶梯。

请你找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。


示例1：

输入：cost = [10, 15, 20]
输出：15
解释：最低花费是从 cost[1] 开始，然后走两步即可到阶梯顶，一共花费 15 。
示例 2：

输入：cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
输出：6
解释：最低花费方式是从 cost[0] 开始，逐个经过那些 1 ，跳过 cost[3] ，一共花费 6 。

提示：

cost的长度范围是 [2, 1000]。
cost[i] 将会是一个整型数据，范围为 [0, 999] 。


注意：1.题目描述无法理解，跟屎一样的描述
2.cost[n]理解为是楼顶，花费为0，必须要到达楼顶， cost[-1]理解为地面，花费为0，求从地面到达楼顶cost最少
3.每次是1步或者2步，所以有些楼层可以直接跨过去，不需要花费体力
4.地推公式：dp[i]=min(dp[i−1]+cost[i−1],dp[i−2]+cost[i−2])

5.因为只能走1步或者2步，所以当前楼层=i-1楼层的最小花费+i-1的花费与i-2楼层的最小花费+i-2的花费，取最小即可
dp[0,1]=0

*/
func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println("最小花费-dp:", minCostClimbingStairs(cost))
	fmt.Println("最小花费-滚动数组:", minCostClimbingStairs1(cost))
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

func minCostClimbingStairs1(cost []int) int {
	n := len(cost)
	pre, cur := 0, 0
	for i := 2; i <= n; i++ {
		pre, cur = cur, min(cur+cost[i-1], pre+cost[i-2])
	}
	return cur
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
