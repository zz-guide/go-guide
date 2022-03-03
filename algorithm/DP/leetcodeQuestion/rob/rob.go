package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/house-robber/
打家劫舍

你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

注意：1.一排房屋，看成是一个数组，首尾不相连
	 2.不能连续偷挨着的2家
     3.偷的钱都是非负的
     4.偷的次数不限制

*/
func main() {
	nums := []int{2, 7, 9, 3, 1}
	nums = []int{2, 1, 1, 2}
	fmt.Println("打家劫舍-动态规划:", rob(nums))
}

// rob 动态规划
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	// 代表直到第i间屋子为止最高偷窃金额，还不知道偷不偷呢
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		// 要么偷，要么不偷，因为不能连续，所以是减2
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[len(nums)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
