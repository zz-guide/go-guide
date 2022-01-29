package main

import (
	"log"
)

func main() {
	nums := []int{1, 2, 2, 2, 3}
	target := -2

	log.Println("最小边界：", binarySearchMinBound(nums, target))
	log.Println("最大边界：", binarySearchMaxBound(nums, target))
}

// binarySearchMinBound
// 寻找左侧边界的二分查找:假设存在重复元素，且数组按照升序排序,返回相同元素的最小边界,-1表示不存在
//**
func binarySearchMinBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// 1.首先定义左右边界index
	low, high := 0, len(nums)-1

	// 2.low=high的时候也可能满足条件，所以需要查找
	for low <= high {
		mid := (high-low)/2 + low

		if nums[mid] == target { // 找到了则收缩右边界
			high = mid - 1
		} else if nums[mid] > target { // 发现mid比目标值大，那么肯定在low~mid之间，往左查找
			high = mid - 1
		} else if nums[mid] < target { // 相反则在mid~high之间，往右查找
			low = mid + 1
		}
	}

	// 当遍历一遍的时候，如果target>nums中最大值，则low会超出high,此时high=len(nums)-1
	// 当遍历一遍的时候，如果target<nums中最大值，则high会小于low,此时low=0
	//log.Println("low:", low)
	//log.Println("high:", high)

	// 因为是最小边界，肯定是返回low
	// 因为nums[low]可能越界。所以先判断low和len(nums)的关系
	if low >= len(nums) || nums[low] != target {
		return -1
	}

	return low
}

// binarySearchMaxBound
// 寻找右侧边界的二分查找：假设存在重复元素，且数组按照升序排序,返回相同元素的最大边界,-1表示不存在
//**
func binarySearchMaxBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// 1.首先定义左右边界index
	low, high := 0, len(nums)-1

	// 2.low=high的时候也可能满足条件，所以需要查找
	for low <= high {
		mid := (high-low)/2 + low

		if nums[mid] == target { // 找到了则收缩左边界
			low = mid + 1
		} else if nums[mid] > target { // 发现mid比目标值大，那么肯定在low~mid之间，往左查找
			high = mid - 1
		} else if nums[mid] < target { // 相反则在mid~high之间，往右查找
			low = mid + 1
		}
	}

	// 当遍历一遍的时候，如果target>nums中最大值，则low会超出high,此时high=len(nums)-1
	// 当遍历一遍的时候，如果target<nums中最大值，则high会小于low,此时low=0
	//log.Println("low:", low)
	//log.Println("high:", high)

	// 因为是最小边界，肯定是返回low
	// 因为nums[high]可能越界。所以先判断low和len(nums)的关系
	if high < 0 || nums[high] != target {
		return -1
	}

	return high
}
