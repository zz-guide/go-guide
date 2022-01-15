package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/next-greater-element-i/
下一个更大元素  II

给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。

示例 1:
输入: [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数；
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。


注意: 输入数组的长度不会超过 10000。

注意：1.数组变成环以后，长度变为原来的2倍就可以继续遍历了

*/
func main() {
	nums := []int{4, 1, 2}
	fmt.Println("下一个更大元素 II-单调栈:", nextGreaterElements(nums))
}

// nextGreaterElements 单调栈+哈希
func nextGreaterElements(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	for i := range res {
		res[i] = -1
	}

	// 类似于温度那道题,只是长度变成2倍而已，取余就是相对位置
	var stack []int
	for i := 0; i < length*2; i++ {
		for len(stack) > 0 && nums[i%length] > nums[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res[index] = nums[i%length]
		}

		stack = append(stack, i%length)
	}

	return res
}
