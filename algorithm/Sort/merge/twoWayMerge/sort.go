package main

import (
	"log"
)

/**
两路归并排序（Merge Sort）是建立在归并操作上的一种有效，稳定的排序算法，该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。
若将两个有序表合并成一个有序表，称为二路归并。
时间复杂度：O(n log n)
空间复杂度：T（n)
稳定性：稳定
*/
func main() {
	nums := []int{8, 4, 5, 7, 1, 3, 6, 2, 10, 20, 17, 25, 24, 28}
	log.Println("归并排序-递归:", sortArray(nums))
	log.Println("归并排序-栈:", mergeSort2(nums))
}

func sortArray(nums []int) []int {
	mergeSort1(nums, 0, len(nums)-1)
	return nums
}

// mergeSort 递归实现
// 递归实现先分组，然后自顶向下
func mergeSort1(nums []int, left, right int) {
	if left < right {
		mid := left + (right-left)>>1
		// 递归过程的2次mergeSort1调用可以再迭代过程中去掉优化
		mergeSort1(nums, left, mid)
		mergeSort1(nums, mid+1, right)
		merge1(nums, left, mid, right)
	}
}

func merge1(nums []int, left, mid, right int) {
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

// mergeSort 迭代实现
// 迭代过程是先两两排序分组，最后四个四个排序分组，直至完成排序，从底向上
func mergeSort2(nums []int) []int {
	step := 1 // 步长间隔为1
	length := len(nums)
	for step < length {
		// 每次从头开始合并，每次步长*2
		for i := 0; i+step < length; i += step * 2 {
			left := i
			right := i + step*2 - 1
			mid := i + step - 1

			// 右边界不能超
			if right > length-1 {
				right = length - 1 //整个待排序数组为奇数的情况
			}

			merge1(nums, left, mid, right)
		}

		step *= 2
	}

	return nums
}
