package main

import "log"

/**
三数取中法可以保证每次选取的pivot不是最大值和最小值
避免排序区间一边倒的情况
*/

func main() {
	arr := []int{10, 5, 7, 3, 4, 2, 8}
	arr1 := []int{10, 5, 7, 3, 4, 2, 8}
	log.Println("快速排序-递归法(三数取中):", QuickSortV1(arr))
	log.Println("快速排序-递归法(三数取中):", QuickSortV2(arr1))
}

func QuickSortV1(arr []int) []int {
	SortV1(arr, 0, len(arr)-1)
	return arr
}

func SortV1(arr []int, low, high int) {
	if low >= high {
		return
	}

	pivotIndex := partitionV1(arr, low, high)
	SortV1(arr, low, pivotIndex-1)
	SortV1(arr, pivotIndex+1, high)
}

func partitionV1(arr []int, low, high int) int {
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
	pivotIndex := low

	for low < high {
		// high挪动到一个比基准值小的位置，准备交换
		for low < high && arr[high] >= pivot {
			high--
		}

		// low挪动到一个比基准值大的位置，准备交换，相等的值不需要挪动位置
		for low < high && arr[low] <= pivot {
			low++
		}

		// 其实不可能大于，最多是等于
		if low >= high {
			break
		}

		arr[low], arr[high] = arr[high], arr[low]
	}

	arr[pivotIndex], arr[low] = arr[low], arr[pivotIndex]
	return pivotIndex
}

func QuickSortV2(arr []int) []int {
	SortV1(arr, 0, len(arr)-1)
	return arr
}

func SortV2(arr []int, low, high int) {
	if low >= high {
		return
	}

	pivotIndex := partitionV2(arr, low, high)
	SortV2(arr, low, pivotIndex-1)
	SortV2(arr, pivotIndex+1, high)
}

func partitionV2(arr []int, low, high int) int {

	// 三数取中，将中间大小的数字交换到low的位置作为pivot
	midIndex := ((high - low) >> 1) + low
	if arr[low] > arr[high] {
		arr[low], arr[high] = arr[high], arr[low]
	}

	if arr[midIndex] > arr[high] {
		arr[midIndex], arr[high] = arr[high], arr[midIndex]
	}

	if arr[midIndex] < arr[low] {
		arr[midIndex], arr[low] = arr[low], arr[midIndex]
	}

	pivot := arr[midIndex]
	for low < high {
		for low < high && arr[low] < pivot {
			low++
		}

		for low < high && arr[high] > pivot {
			high--
		}

		if low >= high {
			break
		}

		arr[low], arr[high] = arr[high], arr[low]
	}

	// 返回left,就是left-1
	// 返回right,就是right+1
	// 此时right <= left
	return high
}
