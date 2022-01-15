package main

import (
	"fmt"
	"sort"
)

/**
题目:https://leetcode-cn.com/problems/4sum/

四数之和

给你一个由 n 个整数组成的数组nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组
[nums[a], nums[b], nums[c], nums[d]]（若两个四元组元素一一对应，则认为两个四元组重复）：

0 <= a, b, c, d< n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。

提示：

1 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109

*/
func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println("四数之和:", fourSum(nums, target))
	fmt.Println("四数之和:", fourSum1(nums, target))
}

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	if length < 4 {
		return nil
	}

	sort.Ints(nums)
	var res [][]int

	for i := 0; i < length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			l, r := j+1, length-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum < target {
					l++
				} else if sum > target {
					r--
				} else if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})

					for l < r && nums[l] == nums[l+1] {
						l++
					}

					for l < r && nums[r] == nums[r-1] {
						r--
					}

					l++
					r--
				}
			}
		}
	}
	return res
}

func fourSum1(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		tmp := threeSum(nums[i+1:], target-nums[i])
		if tmp != nil {
			for _, t := range tmp {
				t = append(t, nums[i])
				res = append(res, t)
			}
		}
	}
	return res
}

func threeSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i-1 >= 0 && nums[i] == nums[i-1] {
			continue
		}
		tmp := twoSum(nums[i+1:], target-nums[i])
		if tmp != nil {
			for _, t := range tmp {
				t = append(t, nums[i])
				res = append(res, t)
			}
		}
	}
	return res
}

func twoSum(nums []int, target int) [][]int {
	has := make(map[int]bool)
	drop := make(map[int]bool)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		if has[target-nums[i]] {
			if drop[target-nums[i]] == false {
				res = append(res, []int{target - nums[i], nums[i]})
				drop[target-nums[i]] = true
			}
		} else {
			has[nums[i]] = true
		}
	}
	return res
}
