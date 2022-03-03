package main

import (
	"log"
)

/**
题目：https://leetcode-cn.com/problems/combinations/

组合
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。

输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]


提示：

1 <= n <= 20
1 <= k <= n


注意:
1.暴力法需要写k层for循环 2.一层递归就相当于是一层for循环

回溯的模板：
result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return
    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择


*/
func main() {
	n := 4
	k := 2
	log.Println("组合-回溯递归:", combine(n, k))
	log.Println("组合-概率论:", combine1(n, k))
	log.Println("组合-迭代:", combine2(n, k))
}

// combine 回溯
func combine(n int, k int) [][]int {
	var res [][]int
	if n <= 0 || k <= 0 || k > n {
		return res
	}

	var track []int
	var backtrack func(start int)
	backtrack = func(start int) {
		// 判断是不是已经有k个结果了
		if len(track) == k {
			res = append(res, append([]int{}, track...))
			return
		}

		// 此判断用来进行剪枝，减少循环次数
		// 不满足k个数的话也就没必要再继续循环了
		if len(track)+(n-start+1) < k {
			return
		}

		for i := start; i <= n; i++ {
			track = append(track, i)
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}

	// 因为是1~n，所以从1开始
	backtrack(1)
	return res
}

// combine2 非递归，理解不了
func combine2(n int, k int) [][]int {
	var res [][]int
	// 初始化
	// 将 temp 中 [0, k - 1] 每个位置 i 设置为 i + 1，即 [0, k - 1] 存 [1, k]
	// 末尾加一位 n + 1 作为哨兵
	var temp []int
	for i := 1; i <= k; i++ {
		temp = append(temp, i)
	}

	temp = append(temp, n+1)

	for j := 0; j < k; {
		comb := make([]int, k)
		copy(comb, temp[:k])
		res = append(res, comb)
		// 寻找第一个 temp[j] + 1 != temp[j + 1] 的位置 t
		// 我们需要把 [0, t - 1] 区间内的每个位置重置成 [1, t]
		for j = 0; j < k && temp[j]+1 == temp[j+1]; j++ {
			temp[j] = j + 1
		}
		// j 是第一个 temp[j] + 1 != temp[j + 1] 的位置
		temp[j]++
	}
	return res
}

// combine1 概率论 概率论中有个公式
func combine1(n int, k int) [][]int {
	var res [][]int
	var helper func(n, k int, path []int)
	helper = func(n, k int, path []int) {
		if n < k || k == 0 {
			if k == 0 {
				temp := make([]int, len(path))
				copy(temp, path)
				res = append(res, temp)
			}
			return
		}
		helper(n-1, k-1, append(path, n))
		helper(n-1, k, path)
	}
	helper(n, k, []int{})
	return res
}
