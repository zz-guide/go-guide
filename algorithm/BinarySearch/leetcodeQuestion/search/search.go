package main

import "log"

/**
题目：https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

搜索旋转排序数组

整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

示例 1：

输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例2：

输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：

输入：nums = [1], target = 0
输出：-1

提示：

1 <= nums.length <= 5000
-10^4 <= nums[i] <= 10^4
nums 中的每个值都 独一无二
题目数据保证 nums 在预先未知的某个下标上进行了旋转
-10^4 <= target <= 10^4

进阶：你可以设计一个时间复杂度为 O(log n) 的解决方案吗？

思路：
	1.原数组是升序的，从某一个位置经过旋转以后，数组变成了2部分有序数组
	2.使用二分查找从有序的部分进行搜索即可
	3.可以直接拆分，肯定有一半是有序的


*/
func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	log.Println("搜索旋转排序数组:", search(nums, target))
}

// search 二分查找 O(logN)
func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	l, r := 0, length-1
	for l <= r {
		mid := (l + r) / 2
		// 中间值即为target,直接返回mid
		if target == nums[mid] {
			return mid
		}

		// 此时0~mid部分有序
		if nums[0] <= nums[mid] {
			// 此时target落在前半部分有序区间内,mid=的情况已经提前判断了
			if nums[0] <= target && target < nums[mid] {
				// 收缩r = mid - 1
				r = mid - 1
			} else {
				// 此时target落在后半部分无序区间内
				l = mid + 1
			}
		} else {
			// 此时后半部分有序
			// 此时target落在后半部分有序区间内
			if nums[mid] < target && target <= nums[length-1] {
				l = mid + 1
			} else {
				// 此时target落在前半部分无序区间内
				r = mid - 1
			}
		}
	}

	// 说明没找到
	return -1
}
