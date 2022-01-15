package main

import (
	"fmt"
	"sort"
)

/**
题目：https://leetcode-cn.com/problems/subsets-ii/

子集 II

给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。



提示：

1 <= nums.length <= 10
-10 <= nums[i] <= 10


*/
func main() {
	nums := []int{1, 2, 2}
	fmt.Println("子集 II-回溯:", subsetsWithDup(nums))
	fmt.Println("子集 II-迭代法:", subsetsWithDup1(nums))
}

func subsetsWithDup(nums []int) [][]int {
	// 先排序让重复元素靠在一起
	sort.Ints(nums)
	var res [][]int
	length := len(nums)

	var track []int
	var backtracking func(start int)
	backtracking = func(start int) {
		if start == length {
			return
		}

		for i := start; i < length; i++ {
			// 判断是重复元素，直接跳过，判断后边的
			if i > start && nums[i-1] == nums[i] {
				continue
			}

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

// subsetsWithDup1 二进制序列

/*记原序列中元素的总数为 nn。原序列中的每个数字 a_ia
i
的状态可能有两种，即「在子集中」和「不在子集中」。我们用 11 表示「在子集中」，00 表示不在子集中，那么每一个子集可以对应一个长度为 nn 的 0/10/1 序列，第 ii 位表示 a_ia
i
是否在子集中。例如，n = 3n=3 ，a = \{ 5, 2, 9 \}a={5,2,9} 时：

0/10/1 序列	子集	0/10/1 序列对应的二进制数
000000	\{ \}{}	00
001001	\{ 9 \}{9}	11
010010	\{ 2 \}{2}	22
011011	\{ 2, 9 \}{2,9}	33
100100	\{ 5 \}{5}	44
101101	\{ 5, 9 \}{5,9}	55
110110	\{ 5, 2 \}{5,2}	66
111111	\{ 5, 2, 9 \}{5,2,9}	77
可以发现 0/10/1 序列对应的二进制数正好从 00 到 2^n - 12
n
−1。我们可以枚举 \textit{mask} \in [0, 2^n - 1]mask∈[0,2
n
−1]，\textit{mask}mask 的二进制表示是一个 0/10/1 序列，我们可以按照这个 0/10/1 序列在原集合当中取数。当我们枚举完所有 2^n2
n
个 \textit{mask}mask，我们也就能构造出所有的子集。*/

func subsetsWithDup1(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	n := len(nums)
outer:
	for mask := 0; mask < 1<<n; mask++ {
		var t []int
		for i, v := range nums {
			if mask>>i&1 > 0 {
				if i > 0 && mask>>(i-1)&1 == 0 && v == nums[i-1] {
					continue outer
				}
				t = append(t, v)
			}
		}
		res = append(res, append([]int(nil), t...))
	}
	return res
}
