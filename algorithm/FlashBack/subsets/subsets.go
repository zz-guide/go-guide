package main

import (
	"fmt"
	"math"
)

/**
题目：https://leetcode-cn.com/problems/subsets/
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

提示：

1 <= nums.length <= 10
-10 <= nums[i] <= 10
nums 中的所有元素 互不相同


注意：1.数组中不能有重复子集


*/
func main() {
	nums := []int{1, 2, 3}
	fmt.Println("子集-迭代法：", subsets(nums))
	fmt.Println("子集-回溯：", subsets1(nums))
}

// subsets 迭代法
func subsets(nums []int) [][]int {
	// 长度平方+1
	capacity := int(math.Pow(2, float64(len(nums)))) + 1
	res := make([][]int, 1, capacity)
	// res = [ [] ]
	res[0] = []int{}
	/**
	 	[[]],
		[[],[1]],
		[[],[1], [2], [1,2]]
	*/
	for _, cur := range nums { // 1,2,3
		for _, arr := range res { // 之前的结果集

			// 把当前值添加到结果集的每一项中
			temp := make([]int, len(arr), len(arr)+1)
			copy(temp, arr)
			temp = append(temp, cur)

			res = append(res, temp)
		}
	}

	return res
}

// subsets1 回溯
func subsets1(nums []int) [][]int {
	var res [][]int
	length := len(nums)

	var track []int
	var backtracking func(start int)
	backtracking = func(start int) {
		if start == length {
			return
		}

		for i := start; i < length; i++ {
			track = append(track, nums[i])
			res = append(res, append([]int{}, track...))
			backtracking(i + 1)
			track = track[:len(track)-1]
		}
	}

	backtracking(0)
	res = append(res, []int{})
	return res
}
