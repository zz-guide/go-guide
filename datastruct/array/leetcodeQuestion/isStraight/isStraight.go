package main

import (
	"log"
	"sort"
)

/**
题目：https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/solution/pai-xu-bian-li-pai-wu-zhong-fu-shi-jian-lqbj6/
扑克牌中的顺子
从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。


示例1:

输入: [1,2,3,4,5]
输出: True

示例2:

输入: [0,0,1,2,5]
输出: True

限制：

数组长度为 5

数组的数取值为 [0, 13] .

*/

func main() {
	nums := []int{1, 2, 3, 4, 5}
	log.Println("扑克牌中的顺子-排序", isStraight(nums))
	log.Println("扑克牌中的顺子-哈希", isStraight1(nums))
}

// isStraight 排序+遍历 时间复杂度O(logN),空间复杂度O(1)
func isStraight(nums []int) bool {
	joker := 0
	sort.Ints(nums)
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			joker++
		} else if nums[i] == nums[i+1] {
			return false
		}
	}

	// 除了大小王，相差不过5
	return nums[4]-nums[joker] < 5
}

// isStraight1 哈希， O(n)
func isStraight1(nums []int) bool {
	hashMap := make(map[int]bool, 14)
	max := 0
	min := 14
	for _, num := range nums {
		if num == 0 {
			continue
		}
		max = Max(max, num)
		min = Min(min, num)
		if hashMap[num] {
			return false
		}
		hashMap[num] = true
	}

	return max-min < 5
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
