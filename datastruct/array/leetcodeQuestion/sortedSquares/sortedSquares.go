package main

import (
	"fmt"
	"sort"
)

/**
题目：https://leetcode-cn.com/problems/squares-of-a-sorted-array/
有序数组的平方

给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1：

输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]
示例 2：

输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]

提示：
1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 已按 非递减顺序 排序

进阶：
请你设计时间复杂度为 O(n) 的算法解决本问题


注意：
什么叫非递减？
1,2,3,4,5,........ : 递增排列
9,8,7,6,5......... : 递减排列
1，2，3，3，4，5，8，8，.............. : 非递减排列
9，8，7，7，6，5，5，2，1，........ : 非递增排列

*/
func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println("平方和排序：", sortedSquares(nums))
	fmt.Println("平方和排序：", sortedSquares1(nums))
}

// sortedSquares 先求平方，再排序
//时间复杂度：O(nlogn)，其中 nn 是数组 {nums}nums 的长度。空间复杂度：O(logn)。除了存储答案的数组以外，我们需要 O(logn) 的栈空间进行排序。
func sortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	res := make([]int, len(nums))
	for i, num := range nums {
		res[i] = num * num
	}

	sort.Ints(res)
	return res
}

// sortedSquares1 双指针法, O(n)时间复杂度，O(1)空间复杂度
func sortedSquares1(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	res := make([]int, len(nums))
	start, end := 0, len(nums)-1

	// 先找最大的值
	for i := len(nums) - 1; i >= 0; i-- {
		startSquare := nums[start] * nums[start]
		endSquare := nums[end] * nums[end]

		// 负数的平方和可能大，先放
		if startSquare > endSquare {
			res[i] = startSquare
			start++
		} else {
			res[i] = endSquare
			end--
		}
	}

	return res
}

// sortedSquares2 双指针，负数部分平方之后单调递减，正数部分单调递增，然后归并排序，O(n)时间复杂度，O(1)空间复杂度
func sortedSquares2(nums []int) []int {
	n := len(nums)
	lastNegIndex := -1
	for i := 0; i < n && nums[i] < 0; i++ {
		lastNegIndex = i
	}

	ans := make([]int, 0, n)
	for i, j := lastNegIndex, lastNegIndex+1; i >= 0 || j < n; {
		// 没有负数
		if i < 0 {
			ans = append(ans, nums[j]*nums[j])
			j++
		} else if j == n { // 就一个元素，
			ans = append(ans, nums[i]*nums[i])
			i--
		} else if nums[i]*nums[i] < nums[j]*nums[j] { // 多个元素
			ans = append(ans, nums[i]*nums[i])
			i--
		} else {
			ans = append(ans, nums[j]*nums[j])
			j++
		}
	}

	return ans
}
