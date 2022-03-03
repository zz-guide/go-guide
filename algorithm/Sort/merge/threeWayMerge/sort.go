package main

import (
	"log"
)

/**
三路归并排序（Merge Sort）
时间复杂度：O(n log n)
空间复杂度：T（n)
稳定性：稳定
*/
func main() {
	nums := []int{8, 4, 5, 7, 1, 3, 6, 2, 10, 20, 17, 25, 24, 28}
	log.Println("归并排序-递归:", sortArray(nums))
}

func sortArray(nums []int) []int {
	mergeSort1(nums, 0, len(nums)-1)
	return nums
}

// mergeSort 递归实现
// 递归实现先分组，然后自顶向下
func mergeSort1(nums []int, left, right int) {
	if left < right {
		mid := left + (right-left)/3
		midMid := right - (right-left)/3 - 1
		log.Println("left, mid, midMid, right:", left, mid, midMid, right)
		// 递归过程的2次mergeSort1调用可以再迭代过程中去掉优化
		mergeSort1(nums, left, mid)
		mergeSort1(nums, mid+1, midMid)
		mergeSort1(nums, midMid+1, right)
		merge(nums, left, mid, midMid+1, right)
	}
}
func merge(nums []int, left, mid, midMid, right int) {
	if mid > midMid {
		mid = midMid
	}

	twoArrayMerge(nums, left, mid, midMid)

	if midMid > right {
		midMid = right
	}
	twoArrayMerge(nums, left, midMid, right)
}

func twoArrayMerge(nums []int, left, mid, right int) {
	arr := make([]int, right-left+1)
	i, j, k := left, mid+1, 0 // k代表最新的结果的位置
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			arr[k] = nums[i]
			i++
		} else {
			arr[k] = nums[j]
			j++
		}

		k++
	}

	// 检查是左边是否有剩余元素
	for i <= mid {
		arr[k] = nums[i]
		k++
		i++
	}

	// 检查是右边边是否有剩余元素
	for j <= right {
		arr[k] = nums[j]
		k++
		j++
	}

	// 重新赋值nums
	for p := 0; p < len(arr); p++ {
		nums[left+p] = arr[p]
	}
}
