package main

import (
	"log"
	"sort"
)

/**
题目：https://leetcode-cn.com/problems/3sum/
三数之和
给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]


注意：1.要求和为0，且元素不重复
2.只需要a<b<c就可以保证不重复
3.这三个元素中肯定得有负数，然后从小到大才能满足可能是0

*/

func main() {
	nums := []int{-1, 0, 1, 2}
	log.Println("三数之和:", threeSum(nums))
}

// threeSum 排序+双指针
// 时间复杂度：O(N^2)，空间复杂度：O(logN)。我们忽略存储答案的空间，额外的排序的空间复杂度为 O(logN)
func threeSum(nums []int) [][]int {
	target := 0
	length := len(nums)
	if length < 3 {
		return nil
	}

	sort.Ints(nums)
	var res [][]int

	for i := 0; i < length; i++ {
		// 重复值第二次遍历直接跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, length-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				res = append(res, []int{nums[i], nums[l], nums[r]})

				// 挪动到下一个不等于自身值的位置
				for l < r && nums[l] == nums[l+1] {
					l++
				}

				for l < r && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--
			} else if sum > target {
				r--
			} else if sum < target {
				l++
			}
		}
	}

	return res
}
