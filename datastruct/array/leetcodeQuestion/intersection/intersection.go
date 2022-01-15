package main

import (
	"fmt"
	"sort"
)

/**
题目：https://leetcode-cn.com/problems/intersection-of-two-arrays/
两个数组的交集

给定两个数组，编写一个函数来计算它们的交集。

说明：

输出结果中的每个元素一定是唯一的。
我们可以不考虑输出结果的顺序。

注意：1.如果有重复的值相等，只能加入一次结果集

*/

func main() {
	a := []int{1, 2, 2, 1}
	b := []int{2, 2}
	fmt.Println("两个数组的交集-哈希:", intersection(a, b))
	fmt.Println("两个数组的交集-排序:", intersection1(a, b))
	fmt.Println("两个数组的交集-排序2(降低空间复杂度):", intersection2(a, b))
	fmt.Println("两个数组的交集-暴力法:", intersection3(a, b))
}

// intersection 哈希
func intersection(nums1 []int, nums2 []int) []int {
	var res []int

	m := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		if _, ok := m[nums1[i]]; ok {
			continue
		}
		m[nums1[i]] = 1
	}

	for i := 0; i < len(nums2); i++ {
		if n, ok := m[nums2[i]]; ok {
			// 去重
			if n > 0 {
				res = append(res, nums2[i])
				m[nums2[i]] = n - 1
			}
		}
	}

	return res
}

// intersection1 双指针+排序
func intersection1(nums1 []int, nums2 []int) []int {
	var res []int
	sort.Ints(nums1)
	sort.Ints(nums2)
	length1, length2 := len(nums1), len(nums2)
	index1, index2 := 0, 0

	for index1 < length1 && index2 < length2 {
		if nums1[index1] < nums2[index2] {
			index1++
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else {
			// res为空，或者不为空但是和上一个值相等，说明是重复了，不需要加入结果集
			if len(res) == 0 || (len(res) > 0 && nums1[index1] != res[len(res)-1]) {
				res = append(res, nums1[index1])
			}

			index1++
			index2++
		}
	}
	return res
}

// intersection2 排序，重复利用nums1，降低空间复杂度
func intersection2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	// 打算利用nums1返回结果，0~k-1位置是重复元素的值
	var i, j, k int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			v := nums1[i]
			nums1[k] = v // 把重复的元素覆盖到肯定不重复值的位置
			k++
			// i和j挪动到不等于刚才值的位置，因为有可能重复
			for i < len(nums1) && nums1[i] == v {
				i++
			}

			for j < len(nums2) && nums2[j] == v {
				j++
			}
		}
	}

	return nums1[:k]
}

// intersection3 暴力解法
func intersection3(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	var res []int
	for _, v := range nums1 {
		for _, vv := range nums2 {
			if v == vv {
				m[vv]++
			}
		}
	}

	for k, _ := range m {
		res = append(res, k)
	}

	return res
}
