package main

import "log"

/**
三项切分用来解决pivot重复值的问题
将数组切分成了三部分，小于基准数的左区间，大于基准数的右区间，等于基准数的中间区间。
*/

const InsertionSortMaxLength = 7

func main() {
	arr := []int{10, 5, 7, 3, 4, 2, 8}
	log.Println("快速排序-递归法(三数取中+三向切分+插入排序):", QuickSortV1(arr))
}

func QuickSortV1(arr []int) []int {
	SortV1(arr, 0, len(arr)-1)
	return arr
}

func SortV1(arr []int, low, high int) {
	if low >= high {
		return
	}

	if high-low <= InsertionSortMaxLength {
		insertSort(arr, low, high)
		return
	}

	left, right := partitionV1(arr, low, high)
	SortV1(arr, low, left-1)
	SortV1(arr, right+1, high)
}

func partitionV1(arr []int, low, high int) (int, int) {
	// 三数取中，将中间大小的数字交换到low的位置作为pivot
	midIndex := ((high - low) >> 1) + low
	if arr[low] > arr[high] {
		arr[low], arr[high] = arr[high], arr[low]
	}

	if arr[midIndex] > arr[high] {
		arr[midIndex], arr[high] = arr[high], arr[midIndex]
	}

	if arr[midIndex] > arr[low] {
		arr[midIndex], arr[low] = arr[low], arr[midIndex]
	}

	pivot := arr[low]
	left, i, right := low, low+1, high
	for i < right {
		if pivot < arr[i] {
			arr[i], arr[right] = arr[right], arr[i]
			right--
		} else if pivot == arr[i] {
			i++
		} else {
			arr[left], arr[i] = arr[i], arr[left]
			left++
			i++
		}
	}

	return left, right
}

func insertSort(nums []int, low, high int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	for i := low + 1; i <= high; i++ {
		// 利用二分查找的特性缩小区间，最终i只需要把low~i之前的数字移动一次即可
		low, high := 0, i
		for low <= high {
			middle := low + (high-low)>>1
			if nums[middle] > nums[i] {
				high = middle - 1
			} else {
				low = middle + 1
			}
		}

		for j := i - 1; j >= low; j-- {
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}
	}

	return nums
}
