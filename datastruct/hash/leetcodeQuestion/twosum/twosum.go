package main

import (
	"log"
)

/**
题目:https://leetcode-cn.com/problems/two-sum/

两数之和：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例：
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

结论：不能有多个答案，不能有重复数字，且元素只能使用一次，场景比较简单
*/

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	log.Println("两数之和-暴力法:", twoSum(nums, target))
	log.Println("两数之和-哈希map:", twoSum1(nums, target))
}

// twoSum 双指针，时间复杂度O(n^2),空间复杂度：O(1)
func twoSum(nums []int, target int) []int {
	var result []int
	if len(nums) == 0 {
		return result
	}
	for i, v := range nums {
		for k := i + 1; k < len(nums); k++ {
			if v == target-nums[k] {
				return []int{i, k}
			}
		}
	}

	return result
}

// twoSum 利用map，时间复杂度：O(N) 空间复杂度：O(N)
func twoSum1(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	m := make(map[int]int)
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}

		m[v] = i
	}

	return nil
}
