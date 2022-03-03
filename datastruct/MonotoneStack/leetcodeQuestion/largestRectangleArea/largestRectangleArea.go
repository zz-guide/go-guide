package main

import (
	"fmt"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
柱状图中最大的矩形

给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。

输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10

提示：

1 <= heights.length <=105
0 <= heights[i] <= 104

注意:1.柱子之间彼此相邻，高度不排除是0的情况
2.宽度为1，单个柱子的面积就是高度
3.找到当前柱子左右两边第一个比自己矮的位置，res = (r - l - 1) * height 就是最大体积


*/
func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	log.Println("柱状图中最大的矩形-暴力法:", largestRectangleArea(heights))
	log.Println("柱状图中最大的矩形-单调栈:", largestRectangleArea1(heights))
}

// largestRectangleArea 暴力解法
func largestRectangleArea(heights []int) int {
	var res int

	for i, height := range heights {
		//找到左右两边第一根比当前柱子矮的柱子
		l, r := i-1, i+1
		for l >= 0 && heights[l] >= height {
			l--
		}
		for r < len(heights) && heights[r] >= height {
			r++
		}
		//计算当前柱子能构成的矩形面积(不包含左右两根比当前柱子矮的柱子)
		//(r-l+1-2)*height
		if (r-l-1)*height > res {
			res = (r - l - 1) * height
		}
	}

	return res
}

// largestRectangleArea1 单调栈+场数优化
func largestRectangleArea1(heights []int) int {
	var res int
	//单调栈（单调递增）
	var stack []int
	stack = append(stack, -1)
	heights = append(heights, 0)

	for i := 0; i < len(heights); i++ {
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			l := stack[len(stack)-1]
			res = max(res, (i-l-1)*heights[top])
		}

		stack = append(stack, i)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
