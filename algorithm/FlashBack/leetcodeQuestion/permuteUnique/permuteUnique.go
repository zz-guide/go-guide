package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/permutations-ii/

全排列 II
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

提示：

1 <= nums.length <= 8
-10 <= nums[i] <= 10

注意：1.数字有重复

*/
func main() {
	nums := []int{1, 1, 2}
	fmt.Println("全排列 II-回溯:", permuteUnique(nums))
}

func permuteUnique(nums []int) [][]int {
	var res [][]int
	length := len(nums)

	var temp []int
	var backtracking func(track []int)
	backtracking = func(track []int) {
		if len(temp) == length {
			res = append(res, append([]int{}, temp...))
			return
		}

		// 唯一的区别就是如果是重复数字，直接跳过
		used := make(map[int]bool) // 使用slice也是一样的，没有map方便

		for i := 0; i < len(track); i++ {
			if _, ok := used[nums[i]]; ok {
				continue
			}

			used[nums[i]] = true
			cur := track[i]
			temp = append(temp, cur)
			track = append(track[:i], track[i+1:]...)
			backtracking(track)
			track = append(track[:i], append([]int{cur}, track[i:]...)...)
			temp = temp[:len(temp)-1]
		}
	}

	backtracking(nums)
	return res
}
