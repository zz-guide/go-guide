package main

import (
	"fmt"
)

/**
题目：https://leetcode-cn.com/problems/maximum-subarray/
最大子数组和

给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。

提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104

注意：1.值可能是负的; 2.一个元素的时候也可以认为是连续子数组

*/
func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("最大子数组和-动态规划:", maxSubArray(nums))
	fmt.Println("最大子数组和-贪心算法:", maxSubArray1(nums))
}

// maxSubArray 动态规划
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var res int = nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(res, dp[i])
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// maxSubArray1 贪心算法
func maxSubArray1(nums []int) int {
	MaxSum := nums[0]
	temSum := 0
	for i := 0; i < len(nums); i++ {
		temSum += nums[i]
		if temSum > MaxSum {
			MaxSum = temSum
		}
		//遇到temSum<0后直接抛弃之前的和(负数加下一个数只会变小,此时局部最优)
		if temSum < 0 {
			temSum = 0
		}
	}

	return MaxSum
}
