package main

import (
	"log"
)

/**
题目：https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

数组中的第K个最大元素

给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

提示：

1 <= k <= nums.length <= 104
-104 <= nums[i] <= 104

*/
func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	log.Println("数组中的第K个最大元素：", findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	return 0
}
