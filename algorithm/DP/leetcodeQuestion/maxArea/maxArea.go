package main

import "log"

/**
题目：https://leetcode-cn.com/problems/container-with-most-water/

盛最多水的容器

给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])。

找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为49。
示例 2：

输入：height = [1,1]
输出：1


提示：

n == height.length
2 <= n <= 105
0 <= height[i] <= 104

*/

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	log.Println("盛最多水的容器:", maxArea(height))
}

// 双指针 O(N)
func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	maxVolume := 0
	// 最大体积公式 min(左边界，右边界)*距离，
	for i < j {
		temp := (j - i) * min(height[i], height[j])
		// 比较哪个面积大
		if maxVolume < temp {
			maxVolume = temp
		}

		// 移动较小高度的那一端
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxVolume
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
