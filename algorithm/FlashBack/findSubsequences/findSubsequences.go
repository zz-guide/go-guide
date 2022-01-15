package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/increasing-subsequences/
递增子序列

给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。

数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。

提示：

1 <= nums.length <= 15
-100 <= nums[i] <= 100


注意：1.子序列不是子串，不需要连续，只需要递增
2.因为有重复数字，所以是>=
3.子串是子序列的一种
4.子序列中元素个数>=2

*/
func main() {
	nums := []int{4, 4, 7, 7}
	fmt.Println("递增子序列:", findSubsequences(nums))
}

func findSubsequences(nums []int) [][]int {
	length := len(nums)
	var res [][]int
	var temp []int

	var backtrack func(start int)
	backtrack = func(start int) {
		if start > length {
			return
		}

		if len(temp) >= 2 {
			res = append(res, append([]int{}, temp...))
		}

		// 主要用来去重，前边用过，后边就不能再用了
		used := make(map[int]bool)

		for i := start; i < length; i++ {
			// 判断是不是递增和重复使用
			if used[nums[i]] || (len(temp) > 0 && temp[len(temp)-1] > nums[i]) {
				continue
			}

			used[nums[i]] = true

			temp = append(temp, nums[i])
			backtrack(i + 1)
			temp = temp[:len(temp)-1]
		}
	}

	backtrack(0)
	return res
}
