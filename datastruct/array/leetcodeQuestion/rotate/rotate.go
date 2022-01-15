package main

import "fmt"

/**
https://leetcode-cn.com/problems/rotate-array/
轮转数组
给你一个数组，将数组中的元素向右轮转 k个位置，其中k是非负数。


示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
示例2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]

提示：

1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
0 <= k <= 105



注意：原地数组元素移动位置，O(1)
*/
func main() {
	arr := []int{1, 2, 3, 4, 5}
	k := 1
	fmt.Println("arr--前:", arr)
	//rotate(arr, k)
	//rotate1(arr, k)
	//rotate2(arr, k)
	rotate3(arr, k)
}

// rotate 拷贝一个新数组移动，规律是 （i+k）% len 就是新的相对位置，时间复杂度O(n)，空间复杂度O(n)
func rotate(nums []int, k int) {
	if k == 0 {
		return
	}

	length := len(nums)
	newNums := make([]int, length)
	for i, val := range nums {
		newNums[(i+k)%length] = val
	}

	nums = newNums
	fmt.Println("nums:", nums)
}

// rotate1 环状替换，，空间复杂度O(1)，最大公约数，最小公倍数
func rotate1(nums []int, k int) {
	n := len(nums)
	k %= n
	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// rotate2 反转数组，，时间复杂度O(2n)=O(n)，空间复杂度O(1)
func rotate2(nums []int, k int) {
	var reverse func(a []int)
	reverse = func(a []int) {
		for i, n := 0, len(a); i < n/2; i++ {
			a[i], a[n-1-i] = a[n-1-i], a[i]
		}
	}

	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

// rotate3 拆分成若干次右移一次 时间复杂度O(N * (K%N))，空间复杂度O(1)
func rotate3(nums []int, k int) {
	length := len(nums)
	k %= length

	for j := 0; j < k; j++ {
		// 保存最后一位值
		temp := nums[length-1]

		// 从尾部开始覆盖
		for i := length - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}

		//最后一位值赋值给第一个数
		nums[0] = temp
	}

	fmt.Println("nums:", nums)
}
