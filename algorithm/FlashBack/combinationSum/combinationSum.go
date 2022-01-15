package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/combination-sum/
组合总和

给你一个 无重复元素 的整数数组candidates 和一个目标整数target，
找出candidates中可以使数字和为目标数target 的 所有不同组合 ，并以列表形式返回。
你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

对于给定的输入，保证和为target 的不同组合数少于 150 个。


提示：
1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都 互不相同
1 <= target <= 500


注意: 1.不用考虑溢出的问题，直接用int
2.即使集合元素一样，顺序不一样视为不同的结果
3.没有重复元素

*/
func main() {
	candidates := []int{2, 3, 5}
	target := 8
	fmt.Println("组合总和-回溯:", combinationSum(candidates, target))
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	length := len(candidates)

	var temp []int // 存放每一组的结果
	var backtracking func(start int, sum int)
	backtracking = func(start int, sum int) {
		if sum == target {
			res = append(res, append([]int{}, temp...))
			return
		}

		if sum > target {
			return
		}

		for i := start; i < length; i++ {
			temp = append(temp, candidates[i])
			backtracking(i, sum+candidates[i])
			// 回溯完撤回继续使用
			temp = temp[:len(temp)-1]
		}
	}

	backtracking(0, 0)
	return res
}
