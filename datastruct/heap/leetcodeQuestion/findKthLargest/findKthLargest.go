package main

import (
	"container/heap"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

数组中的第K个最大元素

给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。


示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4

提示：

1 <= k <= nums.length <= 104
-104<= nums[i] <= 104

*/

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	log.Println("数组中的第K个最大元素-堆排序:", findKthLargest(nums, k))

	nums1 := []int{3, 2, 1, 5, 6, 4}
	log.Println("数组中的第K个最大元素-快速排序:", findKthLargest1(nums1, k))
}

type IHeap []int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IHeap) Pop() interface{} {
	length := h.Len()
	x := (*h)[length-1]
	*h = (*h)[0 : length-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	// 初始化一个堆
	h := &IHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)
}

// findKthLargest2 快速排序 挖坑法，也是始终以区间的开始位置为pivot
func findKthLargest1(nums []int, k int) int {
	return SortV3(nums, 0, len(nums)-1, len(nums)-k)
}

func SortV3(nums []int, low, high, k int) int {
	if low >= high {
		return -1
	}

	pivotIndex := partitionV3(nums, low, high)
	if pivotIndex == k {
		return nums[pivotIndex]
	} else if pivotIndex < k {
		return SortV3(nums, pivotIndex+1, high, k)
	}

	return SortV3(nums, low, pivotIndex-1, k)
}

func partitionV3(nums []int, low, high int) int {
	pivot := nums[low]
	// 挖坑法会改变pivot位置的值，所以需要提前保存
	for low < high {
		for low < high && nums[high] >= pivot {
			high--
		}

		if low < high {
			nums[low] = nums[high]
		}

		for low < high && nums[low] <= pivot {
			low++
		}

		if low < high {
			nums[high] = nums[low]
		}
	}

	// 设置pivot的值
	nums[low] = pivot
	return low
}
