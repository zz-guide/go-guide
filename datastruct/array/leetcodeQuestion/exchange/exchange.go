package main

import "log"

/**
题目：https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/

调整数组顺序使奇数位于偶数前面

输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。

示例：

输入：nums =[1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。

提示：

0 <= nums.length <= 50000
0 <= nums[i] <= 10000

*/

func main() {
	nums := []int{1, 2, 3, 4}
	log.Println("调整数组顺序使奇数位于偶数前面-双指针：", exchange(nums))
}

// exchange 时间复杂度O(n),空间复杂度O(1)
func exchange(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}

	l, r := 0, n-1
	for l < r {
		for l < r && nums[l]%2 == 1 {
			l++
		}

		for l < r && nums[r]%2 == 0 {
			r--
		}

		nums[l], nums[r] = nums[r], nums[l]
	}

	return nums
}
