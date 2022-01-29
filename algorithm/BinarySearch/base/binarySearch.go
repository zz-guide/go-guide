package main

import (
	"log"
)

/**
题目：给定一个n个元素有序的（升序）整型数组nums 和一个目标值target，写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。
示例 1:

输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
示例2:

输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1
提示：

你可以假设 nums中的所有元素是不重复的。
n将在[1, 10000]之间。
nums的每个元素都将在[-9999, 9999]之间。
本质是考察二分查找
*/
func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	log.Println("值:", binarySearch(nums, target))
}

// F1 时间复杂度：O(logN)，其中 n 是数组的长度。 空间复杂度：O(1)。原地查找
// 最简单的二分查找，假设不存在重复元素，且数组按照升序排序
// 缺点：当存在重复元素的时候，无法准确返回其他相等值的index
//**
func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// 1.首先定义左右边界index
	low, high := 0, len(nums)-1

	// 2.low=high的时候也可能满足条件，所以需要查找
	for low <= high {
		// 3.此处按照正常理解应该写成 (low+high)/2,但是越靠近右部分区间越大，可能溢出
		// (low+high)/2始终偏向于较小的数，因为是向下取整
		// (low+high)/2 = (high-low)/2 + low
		// mid := (left + right + rand.Intn(2)) / 2  随机取平均位置，偶数的时候使用
		mid := (high-low)/2 + low

		if nums[mid] == target { // 找到了直接返回
			return mid
		} else if nums[mid] > target { // 发现mid比目标值大，那么肯定在low~mid之间，往左查找
			high = mid - 1 // mid本身没被查找过，所以要mid-1=high
		} else if nums[mid] < target { // 相反则在mid~high之间，往右查找
			low = mid + 1 // mid本身没被查找过，所以要mid+1=low
		}
	}
	// 没找到返回-1
	return -1
}
