package main

import "log"

/**
题目：https://leetcode-cn.com/problems/jump-game-ii/

跳跃游戏 II

给你一个非负整数数组nums ，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

假设你总是可以到达数组的最后一个位置。


示例 1:

输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
从下标为 0 跳到下标为 1 的位置，跳1步，然后跳3步到达数组的最后一个位置。
示例 2:

输入: nums = [2,3,0,1,4]
输出: 2


提示:

1 <= nums.length <= 104
0 <= nums[i] <= 1000

*/
func main() {
	nums := []int{2, 3, 1, 1, 4}
	log.Println("跳跃游戏 II:", jump(nums))
	log.Println("跳跃游戏 II:", jump1(nums))
}

func jump(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}

	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func jump1(nums []int) int {
	position := len(nums) - 1
	steps := 0
	for position > 0 {
		for i := 0; i < position; i++ {
			if i+nums[i] >= position {
				position = i
				steps++
				break
			}
		}
	}
	return steps
}
