package main

import (
	"fmt"
	"sort"
)

/**
题目：https://leetcode-cn.com/problems/combination-sum-ii/
组合总和 II

给定一个数组candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。

candidates中的每个数字在每个组合中只能使用一次。

注意：解集不能包含重复的组合。

提示:

1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30

注意：1.一个数字只能使用一次，但是本身可能有重复数字
2.相同的元素集合不同顺序视为不同的结果
3.重复的元素视为一个元素，只能用一次

*/
func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println("组合总和II-回溯:", combinationSum(candidates, target))
}

func combinationSum(candidates []int, target int) [][]int {
	// 先排序，把相同的元素放到一起
	sort.Ints(candidates)
	var res [][]int
	var temp []int
	var backtracking func(start int, sum int)
	backtracking = func(start int, sum int) {
		if sum == target {
			res = append(res, append([]int{}, temp...))
			return
		}

		if sum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			// 当前元素如果是重复元素直接跳过
			if i > start && candidates[i-1] == candidates[i] {
				continue
			}

			temp = append(temp, candidates[i])
			backtracking(i+1, sum+candidates[i])
			temp = temp[:len(temp)-1]
		}
	}

	backtracking(0, 0)
	return res
}
