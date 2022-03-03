package main

import (
	"log"
)

/**
冒泡排序：冒泡排序（Bubble Sort），是一种计算机科学领域的较简单的排序算法。
它重复地走访过要排序的元素列，依次比较两个相邻的元素，如果顺序（如从大到小、首字母从Z到A）错误就把他们交换过来。走访元素的工作是重复地进行直到没有相邻元素需要交换，也就是说该元素列已经排序完成。
这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端（升序或降序排列），就如同碳酸饮料中二氧化碳的气泡最终会上浮到顶端一样，故名“冒泡排序”。

时间复杂度：O(N)~O(N2)
稳定性排序

golang没有while
*/
func main() {
	nums := []int{11, 8, 2, 5, 7, 10, 3, 6}
	nums1 := []int{11, 8, 2, 5, 7, 10, 3, 6}
	nums2 := []int{11, 8, 2, 5, 7, 10, 3, 6}
	log.Println("冒泡排序:", bubbleSort(nums))
	log.Println("冒泡排序-优化1:", bubbleSort2(nums1))
	log.Println("冒泡排序-优化2:", bubbleSort3(nums2))
}

func bubbleSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}

	// 1.N个元素代表N趟排序，每次经过一趟排序都可以确定一个元素的最终位置，如果是从小到大，最右边的元素最先确定，其他的依次
	// 2.每经过一趟排序，右边的边界就少一位，最右边的边界index=len(nums)- i - 1,i从0开始
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			// 用来控制是从小到大还是从大到小
			if nums[j+1] < nums[j] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}

		log.Println("nums:", nums)
	}
	return nums
}

func bubbleSort2(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}

	// 优化的写法：对于123456这种的如果发现经过一趟排序之后，没有发生过交换位置行为，则表明已经有序，可以提前中断
	isSwap := true
	for i := 0; i < length-1; i++ {
		if !isSwap {
			break
		}

		for j := 0; j < length-i-1; j++ {
			// 用来控制是从小到大还是从大到小
			if nums[j+1] < nums[j] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isSwap = true
			}
		}
	}

	return nums
}

func bubbleSort3(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}

	// 更优化的写法：
	isSwap := true
	// 最后一个没有经过交换的元素的下标
	indexOfLastUnsortedElement := length - 1
	// 上次发生交换的位置
	swappedIndex := -1

	for isSwap {
		// 本轮无交换表示有序，直接退出
		isSwap = false
		// 更加优化的写法是：每次经过一趟排序以后，indexOfLastUnsortedElement~length-1位置的元素是有序的，只需要遍历0~indexOfLastUnsortedElement的元素交换即可
		for j := 0; j < indexOfLastUnsortedElement; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				// 说明本趟排序发生了交换，则还需要下一趟排序
				isSwap = true
				// 标记当前发生交换行为的元素index
				swappedIndex = j
			}
		}

		// 设置最后一个排序交换的index
		indexOfLastUnsortedElement = swappedIndex
	}

	return nums
}
