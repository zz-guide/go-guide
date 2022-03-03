package main

import "fmt"

/**
题目:https://leetcode-cn.com/problems/trapping-rain-water/


接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

提示：

n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105


注意：
1.第一个柱子和最后一个柱子不接水
2.可以看出每一列雨水的高度，取决于，该列 左侧最高的柱子和右侧最高的柱子中最矮的那个柱子的高度。
min(lHeight, rHeight) - height。

3.对于每一列雨水体积=min(lHeight, rHeight) - height。
4.可以纵向计算，也可以横向计算


*/
func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println("接雨水-动态规划:", trap(height))
	fmt.Println("接雨水-双指针:", trap1(height))
	fmt.Println("接雨水-单调栈:", trap2(height))
	fmt.Println("接雨水-暴力法:", trap3(height))
}

// 动态规划
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	leftMax := make([]int, n) // 每一列的左侧最大高度
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n) // 每一列的右侧最大高度
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	var res int
	// 只有中间低，两边高才可以承雨水， min(低，高)-当前高度，当前高度必须小于两边最小的高度才能承水，不然就漏了
	for i, h := range height {
		res += min(leftMax[i], rightMax[i]) - h
	}

	return res
}

// trap1 双指针
func trap1(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	var res int
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])

		// 永远移动高度较小的那一侧
		if height[left] < height[right] {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}

	return res
}

// trap2 单调栈(横向计算)（不会）
func trap2(height []int) int {
	var stack []int
	var res int
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}

			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			res += curWidth * curHeight
		}

		stack = append(stack, i)
	}
	return res
}

// trap3 暴力法
func trap3(height []int) int {
	length := len(height)
	res := 0

	for i := 0; i < length-1; i++ {
		// 向前遍历找最大高度
		leftMax := 0
		for k := 0; k <= i; k++ {
			leftMax = max(height[k], leftMax)
		}

		// 向后遍历找最大高度
		rightMax := 0
		for j := i; j < length; j++ {
			rightMax = max(height[j], rightMax)
		}

		// 求最大体积
		res += min(leftMax, rightMax) - height[i]
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
