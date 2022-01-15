package main

import "fmt"

/***
题目：https://leetcode-cn.com/problems/longest-increasing-subsequence/
最长递增子序列
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
*/
func main() {
	nums := []int{10, 9, 2, 518}
	fmt.Println("最长递增子序列长度1：", lengthOfLIS(nums))
	fmt.Println("最长递增子序列长度2：", lengthOfLIS1(nums))
}

// 动态规划
func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for k := i - 1; k >= 0; k-- {
			if nums[k] < nums[i] {
				dp[i] = max(dp[i], dp[k]+1)
			}
		}
	}
	max := 0
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}

// 单调栈+二分查找
func lengthOfLIS1(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	res := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				res = max(res, dp[i])
			}
		}
	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
