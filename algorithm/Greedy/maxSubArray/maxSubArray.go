package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/maximum-subarray/

最大子数组和

给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。


示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组[4,-1,2,1] 的和最大，为6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [5,4,-1,7,8]
输出：23

提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104

进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。

*/
func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("最大子数组和:", maxSubArray(nums))
	//fmt.Println("最大子数组和:", maxSubArray1(nums))
	//fmt.Println("最大子数组和:", maxSubArray2(nums))
	//fmt.Println("最大子数组和:", maxSubArray3(nums))
}

// maxSubArray 贪心
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// maxSubArray1 动态规划
func maxSubArray1(nums []int) int {
	MaxSum := nums[0]
	f := nums[0]
	for i := 1; i < len(nums); i++ {
		f = max(f+nums[i], nums[i])
		if f > MaxSum {
			MaxSum = f
		}
	}
	return MaxSum
}

// 分治
func maxSubArray2(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Status struct {
	lSum, rSum, mSum, iSum int
}

// 暴力法
func maxSubArray3(nums []int) int {
	MaxSum := nums[0]
	temMax := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			temMax += nums[j]
			if temMax > MaxSum {
				MaxSum = temMax
			}
		}
		temMax = 0
	}
	return MaxSum
}
