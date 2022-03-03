package main

import "log"

/**
题目：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

在排序数组中查找元素的第一个和最后一个位置

给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回[-1, -1]。

进阶：

你可以设计并实现时间复杂度为O(log n)的算法解决此问题吗？


示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]

提示：

0 <= nums.length <= 105
-109<= nums[i]<= 109
nums是一个非递减数组
-109<= target<= 109

*/
func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	log.Println("在排序数组中查找元素的第一个和最后一个位置:", searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// 先找最左边的target
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if nums[low] != target {
		return []int{-1, -1}
	}

	// 找右边比target大的第一个数
	left := low
	high = len(nums)
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}

	right := low - 1
	return []int{left, right}
}
