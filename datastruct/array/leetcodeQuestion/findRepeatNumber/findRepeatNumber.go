package main

import "log"

/**

数组中重复的数字

题目：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/solution/pythonjavajavascriptgo-ha-xi-by-himymben-j25s/

找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
请找出数组中任意一个重复的数字。

限制：
2 <= n <= 100000

注意：1.只需要找出第一个重复的数字即可


*/
func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	log.Println("数组中重复的数字-哈希:", findRepeatNumber(nums))
	log.Println("数组中重复的数字-交换坐标:", findRepeatNumber1(nums))
}

// findRepeatNumber 时间复杂度O(n)，空间复杂度O(n), 哈希
func findRepeatNumber(nums []int) int {
	s := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := s[num]; ok {
			return num
		}

		s[num] = struct{}{}
	}

	return -1
}

// findRepeatNumber1 时间复杂度O(n)，空间复杂度O(1), 对应坐标位置交换数字
func findRepeatNumber1(nums []int) int {
	// nums := []int{2, 3, 1, 0, 2, 5, 3}
	for i := 0; i < len(nums); i++ {
		// 判断 i == nums[i]
		for nums[i] != i {
			// 判断nums[i] == nums[nums[i]]
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}

			// 交换
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	return -1
}
