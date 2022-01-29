package main

import "log"

/**
题目：https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/solution/lian-xu-zi-shu-zu-de-zui-da-he-by-leetco-tiui/

连续子数组的最大和

输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。


示例1:

输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释:连续子数组[4,-1,2,1] 的和最大，为6。


提示：

1 <=arr.length <= 10^5
-100 <= arr[i] <= 100

*/
func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	log.Println("连续子数组的最大和-动态规划:", maxSubArray(nums))
	nums1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	log.Println("连续子数组的最大和-线段树分治:", maxSubArray1(nums1))
}

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

// maxSubArray1 利用线段树求
func maxSubArray1(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Status struct {
	//对于一个区间 [l,r]，我们可以维护四个量：
	//lSum 表示 [l,r] 内以 ll 为左端点的最大子段和
	//rSum 表示 [l,r] 内以 rr 为右端点的最大子段和
	//mSum 表示 [l,r] 内的最大子段和
	//iSum 表示 [l,r] 的区间和

	lSum, rSum, mSum, iSum int
}
