package main

import "log"

/**
题目：https://leetcode-cn.com/problems/next-permutation/
下一个排列
整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。

必须 原地 修改，只允许使用额外常数空间。
示例 1：

输入：nums = [1,2,3]
输出：[1,3,2]
示例 2：

输入：nums = [3,2,1]
输出：[1,2,3]
示例 3：

输入：nums = [1,1,5]
输出：[1,5,1]


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 100

注意：
	1.如何变大：从低位挑一个大一点的数，交换前面一个小一点的数。
	2.变大的幅度要尽量小。

*/
func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	log.Println("下一个排列:", nums)
}

// nextPermutation 两遍扫描 O(n)
func nextPermutation(nums []int) {
	length := len(nums)
	// 最少也得2个元素
	if length <= 1 {
		return
	}

	// 从末尾向前找较小数，但不能包含最后一位，因为要交换，所以不能包含末尾元素
	smallIndex := length - 1 - 1
	for smallIndex >= 0 && nums[smallIndex] >= nums[smallIndex+1] {
		smallIndex--
	}

	if smallIndex >= 0 {
		// 从末尾向前找较大数与i交换
		largeIndex := length - 1
		for largeIndex >= 0 && nums[smallIndex] >= nums[largeIndex] {
			largeIndex--
		}

		// 交换
		nums[smallIndex], nums[largeIndex] = nums[largeIndex], nums[smallIndex]
	}

	// 反转小数~末尾的数组元素
	l, r := smallIndex+1, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
