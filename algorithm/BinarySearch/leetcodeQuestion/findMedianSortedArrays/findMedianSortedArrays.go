package main

import "log"

/**
题目：https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-2/
寻找两个正序数组的中位数

给定两个大小分别为 m 和 n 的正序（从小到大）数组nums1 和nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106

*/
func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	log.Println("寻找两个正序数组的中位数:", findMedianSortedArrays(nums1, nums2))
}

// FindMedianSortedArrays 归并排序
func findMedianSortedArrays(nums1, nums2 []int) float64 {
	var res []int
	// 归并过程，因为两个数组都是有序的，所以直接合并
	m, n := len(nums1), len(nums2)
	l1, l2 := 0, 0
	for l1 < m && l2 < n {
		if nums1[l1] < nums2[l2] {
			res = append(res, nums1[l1])
			l1++
		} else {
			res = append(res, nums2[l2])
			l2++
		}
	}

	res = append(res, nums1[l1:]...)
	res = append(res, nums2[l2:]...)

	length := m + n
	// 奇数
	if length%2 == 1 {
		return float64(res[length/2])
	}

	// 偶数
	mid1 := res[length/2]
	mid2 := res[length/2-1]
	return float64(mid1+mid2) / 2.0
}

// 二分查找 O(log(m+n)) O(1)
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	// 奇数
	if length%2 == 1 {
		return float64(getKthElement(nums1, nums2, length/2+1))
	}

	// 偶数
	k1, k2 := length/2, length/2+1
	return float64(getKthElement(nums1, nums2, k1)+getKthElement(nums1, nums2, k2)) / 2.0
}

func getKthElement(nums1, nums2 []int, k int) int {
	m, n := len(nums1), len(nums2)
	index1, index2 := 0, 0
	for {
		if m == index1 {
			return nums2[index2+k-1]
		}

		if n == index2 {
			return nums1[index1+k-1]
		}

		if k == 1 {
			return Min(nums1[0], nums2[0])
		}

		half := k / 2
		l1 := Min(index1+half, m) - 1
		l2 := Min(index2+half, n) - 1
		pivot1, pivot2 := nums1[l1], nums2[l2]
		if pivot1 <= pivot2 {
			k -= l1 - index1 + 1
			index1 = l1 + 1
		} else {
			k -= l2 - index2 + 1
			index2 = l2 + 1
		}
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// findMedianSortedArrays 双指针 O(logmin(m,n))) O(1)
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	length := m + n
	l1, l2 := 0, 0
	left, right := 0, 0 // 代表中位数的2个数

	// 只需要循环左半部分就可以
	for i := 0; i <= length/2; i++ {
		// 每次更新这2个值
		left = right
		// 这三个条件缺一不可
		if l1 < m && (l2 >= n || nums1[l1] < nums2[l2]) {
			right = nums1[l1]
			l1++
		} else {
			right = nums2[l2]
			l2++
		}
	}

	// 奇数
	if length%2 == 1 {
		return float64(right)
	}

	// 偶数
	return float64(left+right) / 2.0
}
