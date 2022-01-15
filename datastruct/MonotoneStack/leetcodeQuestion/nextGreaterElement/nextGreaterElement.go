package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/next-greater-element-i/
下一个更大元素 I

nums1中数字x的 下一个更大元素 是指x在nums2 中对应位置 右侧 的 第一个 比x大的元素。

给你两个 没有重复元素 的数组nums1 和nums2 ，下标从 0 开始计数，其中nums1是nums2的子集。

对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。

返回一个长度为nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。


示例 1：

输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
输出：[-1,3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 4 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
- 1 ，用加粗斜体标识，nums2 = [1,3,4,2]。下一个更大元素是 3 。
- 2 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
示例 2：

输入：nums1 = [2,4], nums2 = [1,2,3,4].
输出：[3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 2 ，用加粗斜体标识，nums2 = [1,2,3,4]。下一个更大元素是 3 。
- 4 ，用加粗斜体标识，nums2 = [1,2,3,4]。不存在下一个更大元素，所以答案是 -1 。


提示：

1 <= nums1.length <= nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 104
nums1和nums2中所有整数 互不相同
nums1 中的所有整数同样出现在 nums2 中

进阶：你可以设计一个时间复杂度为 O(nums1.length + nums2.length) 的解决方案吗？


注意：
1.res 数组含义，-1表示没有更大的元素。如果有就把nums2中对应位置的值赋值过来
2.对应位置的含义表示的是nums1中的值在nums2中出现，并且右边有比这个值大的第一个值


*/
func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println("下一个更大元素 I-暴力:", nextGreaterElement(nums1, nums2))
	fmt.Println("下一个更大元素 I-单调栈:", nextGreaterElement1(nums1, nums2))
}

// nextGreaterElement 暴力法 先找元素存在不存在，然后在找nums2第一个大于的元素
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	res := make([]int, m)
	for i, num := range nums1 {
		j := 0
		for j < n && nums2[j] != num {
			j++
		}

		k := j + 1
		for k < n && nums2[k] < nums2[j] {
			k++
		}

		if k < n {
			res[i] = nums2[k]
		} else {
			res[i] = -1
		}
	}

	return res
}

// nextGreaterElement1 单调栈+哈希
func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i := range res {
		res[i] = -1
	}

	// 需要遍历一次num1存到map，用来判断值是不是存在
	mp := map[int]int{}
	for i, v := range nums1 {
		mp[v] = i
	}

	// 类似于温度那道题
	var stack []int
	for i, v := range nums2 {
		for len(stack) > 0 && v > nums2[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if _, ok := mp[nums2[top]]; ok {
				index := mp[nums2[top]]
				res[index] = v
			}
		}

		stack = append(stack, i)
	}

	return res
}
