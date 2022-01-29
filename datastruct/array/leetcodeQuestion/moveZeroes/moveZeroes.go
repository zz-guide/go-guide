package main

import "log"

/**
题目：https://leetcode-cn.com/problems/move-zeroes/solution/yi-dong-ling-by-leetcode-solution/

移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意，必须在不复制数组的情况下原地对数组进行操作。

示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]

示例 2:
输入: nums = [0]
输出: [0]

提示:
1 <= nums.length <= 104
-231<= nums[i] <= 231- 1

进阶：你能尽量减少完成的操作次数吗？

*/
func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	log.Println("移动零:", nums)
}

// moveZeroes 双指针，left指向0，right指向不为0的数，然后交换，时间复杂度O(n),空间复杂度O(1)
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
