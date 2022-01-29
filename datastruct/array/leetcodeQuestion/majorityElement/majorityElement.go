package main

import "log"

/**
题目：https://leetcode-cn.com/problems/majority-element/solution/duo-shu-yuan-su-by-leetcode-solution/
多数元素

给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于⌊ n/2 ⌋的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

进阶：

尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。

*/
func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	log.Println("多数元素:", majorityElement(nums))
}

// majorityElement 众数一定比其他数出现次数多，所以count不会减到0 时间复杂度O(n),空间复杂度O(1)
func majorityElement(nums []int) int {
	candidate := -1
	count := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
