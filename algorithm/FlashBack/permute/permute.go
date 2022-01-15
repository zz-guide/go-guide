package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/permutations/
全排列：给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同

注意：1.什么叫全排列？从n个不同元素中任取m（m≤n）个元素，按照一定的顺序排列起来，叫做从n个不同元素中取出m个元素的一个排列。当m=n时所有的排列情况叫全排列。
2.数全部不重复
3.回溯的时候需要删除当前元素，把剩余元素集合传进去
*/

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("全排列-回溯:", permute(nums))
}

func permute(nums []int) [][]int {
	var res [][]int
	length := len(nums)

	var temp []int
	var backtracking func(track []int)
	backtracking = func(track []int) {
		if len(temp) == length {
			res = append(res, append([]int{}, temp...))
			return
		}

		for i := 0; i < len(track); i++ {
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
