package main

import (
	"fmt"
	"math"
	"sort"
)

/**
https://leetcode-cn.com/problems/minimum-size-subarray-sum/


长度最小的子数组

给定一个含有n个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

示例 1：

输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组[4,3]是该条件下的长度最小的子数组。
示例 2：

输入：target = 4, nums = [1,4,4]
输出：1
示例 3：

输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0


提示：

1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 105

*/

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println("长度最小的子数组:", minSubArrayLen(target, nums))
	fmt.Println("长度最小的子数组:", minSubArrayLen1(target, nums))
	fmt.Println("长度最小的子数组:", minSubArrayLen2(target, nums))
}

// minSubArrayLen 暴力法
func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	res := math.MaxInt32
	for i := 0; i < length; i++ {
		sum := 0
		for j := i; j < length; j++ {
			sum += nums[j]
			if sum >= target {
				res = min(res, j-i+1)
				break
			}
		}
	}

	if res == math.MaxInt32 {
		return 0
	}

	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 双指针法，right++，sum++, left++，sum--，然后判断此时长度，right = length 退出循环
func minSubArrayLen1(target int, nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	ans := math.MaxInt32
	left, sum := 0, 0

	for right := 0; right < length; right++ {
		sum += nums[right]

		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if ans == math.MaxInt32 {
		return 0
	}

	return ans
}

// minSubArrayLen2 前缀和+二分查找，类似背包问题
func minSubArrayLen2(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	sums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	for i := 1; i <= n; i++ {
		t := target + sums[i-1]
		bound := sort.SearchInts(sums, t)
		if bound < 0 {
			bound = -bound - 1
		}

		if bound <= n {
			ans = min(ans, bound-(i-1))
		}
	}

	if ans == math.MaxInt32 {
		return 0
	}

	return ans
}
