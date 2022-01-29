package main

import "log"

/**
题目：https://leetcode-cn.com/problems/search-insert-position/solution/sou-suo-cha-ru-wei-zhi-by-leetcode-solution/

搜索插入位置

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。

提示:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 为无重复元素的升序排列数组
-104 <= target <= 104

注意：1.经典的二分查找

*/
func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 3
	log.Println("搜索插入位置-二分查找", searchInsert(nums, target))
}

// searchInsert 时间复杂度：O(logn)，空间复杂度O(1)
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := len(nums) // 不存在的话，长度刚好就是要插入的位置

	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}
