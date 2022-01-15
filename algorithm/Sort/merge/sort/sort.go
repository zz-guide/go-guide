package main

import "fmt"

/**
归并排序（Merge Sort）是建立在归并操作上的一种有效，稳定的排序算法，该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。
若将两个有序表合并成一个有序表，称为二路归并。
时间复杂度：O(n log n)
空间复杂度：T（n)
稳定性：稳定
*/
func main() {
	nums := []int{8, 4, 5, 7, 1, 3, 6, 2}
	fmt.Println("归并排序:", mergeSort(nums))
}

func mergeSort(r []int) []int {
	length := len(r)
	if length <= 1 {
		return r
	}

	num := length / 2
	left := mergeSort(r[:num])
	right := mergeSort(r[num:])
	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	fmt.Println(left, right, "---->", result, left[l:], right[r:])
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}
