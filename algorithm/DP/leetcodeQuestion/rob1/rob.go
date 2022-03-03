package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/house-robber-ii/
打家劫舍 II

你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

注意：1.房屋围成一圈，首尾相连
	 2.不能连续偷挨着的2家
     3.偷的钱都是非负的
     4.偷的次数不限制

*/
func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println("打家劫舍II-动态规划:", rob(nums))
}

// rob 动态规划
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	var robRange func(nums []int) int
	robRange = func(nums []int) int {
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

	// 先把第一间屋子排除
	result1 := robRange(nums[1:])
	// 排除第二间屋子
	result2 := robRange(nums[:len(nums)-1])
	return max(result1, result2)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
