package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/

最长连续递增序列

给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。

连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。

提示：

1 <= nums.length <= 104
-109 <= nums[i] <= 109

注意：1.序列必须是连续的，要求递增
2.一个元素的时候认为自身就是递增序列，那么长度就是1，初始化的时候就是1


思路：
1.dp[i+1] = dp[i] + 1
2.start->i区间是递增区间，一点出现不递增，start设置到i，重新开始计数

*/
func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println("最长连续递增序列-递归:", findLengthOfLCIS(nums))
	fmt.Println("最长连续递增序列-贪心:", findLengthOfLCIS1(nums))
}

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	var result int = 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			dp[i+1] = dp[i] + 1
		}

		result = max(result, dp[i+1])
	}

	return result
}

// findLengthOfLCIS1 贪心
func findLengthOfLCIS1(nums []int) int {
	var res = 1
	start := 0 // 断掉递增的地方
	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] <= nums[i-1] {
			start = i
		}

		res = max(res, i-start+1)
	}

	return res
}
func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
