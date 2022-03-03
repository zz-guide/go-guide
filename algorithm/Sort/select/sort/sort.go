package main

import (
	"log"
	"math/rand"
)

/**
选择排序（Selection sort）是一种简单直观的排序算法。它的工作原理是：
第一次从待排序的数据元素中选出最小（或最大）的一个元素，存放在序列的起始位置，然后再从剩余的未排序元素中寻找到最小（大）元素，
然后放到已排序的序列的末尾。以此类推，直到全部待排序的数据元素的个数为零。选择排序是不稳定的排序方法

时间复杂度：O(n^2)
空间复杂度：O(1)
不稳定
*/
func main() {
	nums := rand.Perm(10)
	log.Println("选择排序:", selectSort(nums))
}

func selectSort(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return nums
	}

	// 1.i=0,从i=1开始遍历依次找最小值，与i=0交换
	// 2.i=1 重复步骤

	// 精髓：每次挑选一个最小值，放在当前位置，最后整体有序
	for i := 0; i < length-1; i++ {
		minIdx := i
		for j := i + 1; j <= length-1; j++ {
			// 找最小数或者最大数并保存索引
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}

		// 当前元素和最小元素位置交换
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}

	return nums
}
