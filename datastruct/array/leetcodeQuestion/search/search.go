package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/binary-search/
二分查找

给定一个n个元素有序的（升序）整型数组nums 和一个目标值target ，写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。

提示：

你可以假设 nums 中的所有元素是不重复的。
n 将在 [1, 10000]之间。
nums 的每个元素都将在 [-9999, 9999]之间。

*/
func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println("二分查找:", search(nums, target))
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		pivot := (r-l)/2 + l
		if nums[pivot] == target {
			return pivot
		} else if nums[pivot] > target {
			r = pivot - 1
		} else if nums[pivot] < target {
			l = pivot + 1
		}
	}

	return -1
}
