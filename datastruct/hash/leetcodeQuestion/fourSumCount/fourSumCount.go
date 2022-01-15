package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/4sum-ii/

四数相加 II

给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：

0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0

 提示：

n == nums1.length
n == nums2.length
n == nums3.length
n == nums4.length
1 <= n <= 200
-228 <= nums1[i], nums2[i], nums3[i], nums4[i] <= 228

注意：1.nums1对应i,nums2对应j，依次
2.这四个数组长度都是n，相等
3.可能有重复的元素

*/
func main() {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}
	fmt.Println("四数相加 II:", fourSumCount(nums1, nums2, nums3, nums4))
}

// fourSumCount 分组+哈希 O(n^2) O(n^2)
//总结，看到形如：A+B....+N=0的式子，要转换为(A+...T)=-((T+1)...+N)再计算，这个T的分割点一般是一半，特殊情况下需要自行判断。定T是解题的关键。
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	target := 0
	var res int
	m := map[int]int{}
	for _, v := range nums1 {
		for _, w := range nums2 {
			m[v+w]++
		}
	}

	for _, v := range nums3 {
		for _, w := range nums4 {
			res += m[target-(v+w)]
		}
	}

	return res
}
