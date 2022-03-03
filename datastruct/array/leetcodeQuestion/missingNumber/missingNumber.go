package main

import (
	"log"
	"sort"
)

/**
题目：https://leetcode-cn.com/problems/missing-number/

丢失的数字

给定一个包含 [0, n]中n个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。

示例 1：

输入：nums = [3,0,1]
输出：2
解释：n = 3，因为有 3 个数字，所以所有的数字都在范围 [0,3] 内。2 是丢失的数字，因为它没有出现在 nums 中。
示例 2：

输入：nums = [0,1]
输出：2
解释：n = 2，因为有 2 个数字，所以所有的数字都在范围 [0,2] 内。2 是丢失的数字，因为它没有出现在 nums 中。
示例 3：

输入：nums = [9,6,4,2,3,5,7,0,1]
输出：8
解释：n = 9，因为有 9 个数字，所以所有的数字都在范围 [0,9] 内。8 是丢失的数字，因为它没有出现在 nums 中。
示例 4：

输入：nums = [0]
输出：1
解释：n = 1，因为有 1 个数字，所以所有的数字都在范围 [0,1] 内。1 是丢失的数字，因为它没有出现在 nums 中。

提示：

n == nums.length
1 <= n <= 104
0 <= nums[i] <= n
nums 中的所有数字都 独一无二

进阶：你能否实现线性时间复杂度、仅使用额外常数空间的算法解决此问题?

*/
func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	log.Println("丢失的数字-排序:", missingNumber(nums))
	log.Println("丢失的数字-哈希:", missingNumber1(nums))
	log.Println("丢失的数字-异或:", missingNumber2(nums))
	log.Println("丢失的数字-高斯求和:", missingNumber3(nums))
}

// missingNumber 先排序，然后比较对应位置的数=index,O(nlogn),O(logn)
func missingNumber(nums []int) int {
	sort.Ints(nums)
	for i, num := range nums {
		if num != i {
			return i
		}
	}
	return len(nums)
}

// missingNumber1 哈希，O(n),O(n)
func missingNumber1(nums []int) int {
	has := map[int]bool{}
	for _, v := range nums {
		has[v] = true
	}
	for i := 0; ; i++ {
		if !has[i] {
			return i
		}
	}
}

// missingNumber2 位运算，异或， O(n),O(1)
func missingNumber2(nums []int) int {
	var res int
	// 首先异或i
	for i := 0; i <= len(nums); i++ {
		res ^= i
	}

	// 异或数组元素
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}

	return res
}

// missingNumber3 高斯求和，O(n),O(1)
func missingNumber3(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	arrSum := 0
	for _, num := range nums {
		arrSum += num
	}
	return total - arrSum
}
