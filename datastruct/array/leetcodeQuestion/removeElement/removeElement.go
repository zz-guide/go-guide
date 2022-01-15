package main

import "fmt"

/**
https://leetcode-cn.com/problems/remove-element/
移除元素
给你一个数组 nums和一个值 val，你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。


注意：原地删除的意思就会把要删除的值移动到数组的末尾，然后返回新的长度，超出长度的部分当做是不存在


*/
func main() {
	arr := []int{3, 2, 2, 3}
	arr1 := []int{3, 2, 2, 3}
	val := 3
	fmt.Println("原地删除后长度-双指针1：", removeElement(arr, val))
	fmt.Println("原地删除后长度-双指针2：", removeElement1(arr1, val))
}

// removeElement1 思路是2个指针，一个从头部开始遍历，一个从尾部开始遍历，然后交换元素位置，返回长度
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, len(nums)-1
	for left <= right {
		// 左边界始终是要删除的元素
		for nums[left] != val && left < right {
			left++
		}

		// 右边界始终是不删除的元素
		for nums[right] == val && right > left {
			right--
		}

		// 当左和右相等，判断是不是要删除，如果不删除则直接退出
		if left == right && nums[left] != val {
			break
		}

		// 删除就是换位置，然后各自向前移动
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

	return right + 1
}

// removeElement1 2个指针，同时从头部出发，快的找要删除的元素，慢的
func removeElement1(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right, length := 0, 0, len(nums)
	for right < length {
		if nums[right] == val {
			right++
			continue
		}

		nums[left] = nums[right]
		left++
		right++
	}

	return left
}
