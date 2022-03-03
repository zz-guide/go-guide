package main

import (
	"log"
	"math"
	"sort"
)

/**
题目：https://leetcode-cn.com/problems/maximum-product-of-three-numbers/

三个数的最大乘积

给你一个整型数组 nums ，在数组中找出由三个数组成的最大乘积，并输出这个乘积。

示例 1：

输入：nums = [1,2,3]
输出：6
示例 2：

输入：nums = [1,2,3,4]
输出：24
示例 3：

输入：nums = [-1,-2,-3]
输出：-6

提示：

3 <= nums.length <=104
-1000 <= nums[i] <= 1000


*/

func main() {
	nums := []int{1, 2, 3, 4}
	log.Println("三个数的最大乘积-排序:", maximumProduct(nums))
	log.Println("三个数的最大乘积-线性扫描:", maximumProduct1(nums))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// maximumProduct1 排序，三个最大正数的乘积，以及两个最小负数与最大正数的乘积
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return max(nums[0]*nums[1]*nums[n-1], nums[n-3]*nums[n-2]*nums[n-1])
}

// maximumProduct 线性扫描
func maximumProduct1(nums []int) int {
	// 最小的和第二小的
	min1, min2 := math.MaxInt64, math.MaxInt64
	// 最大的、第二大的和第三大的
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64

	for _, x := range nums {
		if x < min1 {
			min2 = min1
			min1 = x
		} else if x < min2 {
			min2 = x
		}

		if x > max1 {
			max3 = max2
			max2 = max1
			max1 = x
		} else if x > max2 {
			max3 = max2
			max2 = x
		} else if x > max3 {
			max3 = x
		}
	}

	return max(min1*min2*max1, max1*max2*max3)
}
