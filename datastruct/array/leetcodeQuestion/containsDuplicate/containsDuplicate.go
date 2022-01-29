package main

import (
	"log"
	"sort"
)

/**
题目：https://leetcode-cn.com/problems/contains-duplicate/

存在重复元素

给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。

示例 1：

输入：nums = [1,2,3,1]
输出：true
示例 2：

输入：nums = [1,2,3,4]
输出：false
示例3：

输入：nums = [1,1,1,3,3,4,3,2,4,2]
输出：true


提示：

1 <= nums.length <= 105
-109 <= nums[i] <= 109


注意：1.有重复元素就是true,没有就是false
2.不能利用原地修改，元素可能超出范围

*/

func main() {
	nums := []int{1, 2, 3, 4, 5, 5}
	log.Println("存在重复元素-排序:", containsDuplicate(nums))
	log.Println("存在重复元素-哈希:", containsDuplicate1(nums))
}

// containsDuplicate 先排序，判断相邻元素是不是相等，时间O(NlogN)，空间O(logN)
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

// containsDuplicate1 哈希,时间O(n)，空间O(n)
func containsDuplicate1(nums []int) bool {
	set := map[int]struct{}{}
	for _, v := range nums {
		if _, has := set[v]; has {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}
