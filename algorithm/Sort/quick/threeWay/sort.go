package main

import "log"

func main() {
	arr := []int{10, 5, 7, 3, 4, 2, 8}
	log.Println("快速排序-递归法(3数取中):", QuickSortV1(arr))
}

func QuickSortV1(arr []int) []int {
	SortV1(arr, 0, len(arr)-1)
	return arr
}

// SortV1 三数取中
func SortV1(arr []int, left, right int) {

	if left >= right {
		return
	}

	pivot := arr[left]
	lo, gt, cur := left, right+1, left+1

	for cur < gt {
		if arr[cur] < pivot {
			arr[cur], arr[lo+1] = arr[lo+1], arr[cur]
			lo++
			cur++
		} else if arr[cur] > pivot {
			arr[cur], arr[gt-1] = arr[gt-1], arr[cur]
			gt--
		} else {
			cur++
		}
	}

	arr[left], arr[lo] = arr[lo], arr[left]
	SortV1(arr, left, lo-1)
	SortV1(arr, gt, right)
}
