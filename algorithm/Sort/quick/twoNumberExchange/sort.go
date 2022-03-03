package main

import "log"

/**
题目：https://leetcode-cn.com/problems/sort-an-array/solution/kuai-pai-vs-gui-bing-pai-xu-vs-dui-pai-x-44hf/
快速排序：
快速排序算法通过多次比较和交换来实现排序，其排序流程如下： [2]
(1)首先设定一个分界值，通过该分界值将数组分成左右两部分。 [2]
(2)将大于或等于分界值的数据集中到数组右边，小于分界值的数据集中到数组的左边。此时，左边部分中各元素都小于或等于分界值，而右边部分中各元素都大于或等于分界值。 [2]
(3)然后，左边和右边的数据可以独立排序。对于左侧的数组数据，又可以取一个分界值，将该部分数据分成左右两部分，同样在左边放置较小值，右边放置较大值。右侧的数组数据也可以做类似处理。 [2]
(4)重复上述过程，可以看出，这是一个递归定义。通过递归将左侧部分排好序后，再递归排好右侧部分的顺序。当左、右两个部分各数据排序完成后，整个数组的排序也就完成了。

注意：随机选一个数作为pivot，降低原始排序对算法时间的干扰，比如当原始数据本身就是一个降序数组的时候，如果每次都选最左边的数作为pivot，则算法时间接近于O(n^2)。

最好时间复杂度:O(nlogn)
最坏时间复杂度:O(n^2)
平均时间复杂度:O(nlogn)
最好空间复杂度:O(logn)
最坏空间复杂度:O(n)
平均空间复杂度:O(logn)
是否稳定:不稳定

常规的快速排序无法解决：
1.pivot选取最大值或者最小值
2.pivot是数组中重复值较多的那一项

*/

func main() {
	arr := []int{10, 5, 7, 11, 9, 2, 8}
	arr1 := []int{10, 5, 7, 11, 9, 2, 8}
	arr2 := []int{10, 5, 7, 11, 9, 2, 8}
	arr3 := []int{10, 5, 7, 11, 9, 2, 8}
	arr4 := []int{10, 5, 7, 11, 9, 2, 8}
	log.Println("快速排序-递归法(双指针交换法，中间到两边):", QuickSortV1(arr))
	log.Println("快速排序-递归法(单边单指针):", QuickSortV2(arr1))
	log.Println("快速排序-递归法(挖坑法):", QuickSortV3(arr2))
	log.Println("快速排序-递归法(双指针交换法):", QuickSortV4(arr3))
	log.Println("快速排序-迭代:", QuickSortV5(arr4))
}

func QuickSortV1(nums []int) []int {
	SortV1(nums, 0, len(nums)-1)
	return nums
}

func SortV1(nums []int, low, high int) {
	if low >= high {
		return
	}

	pivotIndex := partitionV1(nums, low, high)
	SortV1(nums, low, pivotIndex)
	SortV1(nums, pivotIndex+1, high)
}

// SortV1 双指针，交换法
func partitionV1(nums []int, low, high int) int {
	// 选取中间靠左部分pivot
	pivotIndex := (high-low)>>1 + low
	pivot := nums[pivotIndex]

	// 因为选取的pivot在中间部分，也是需要移动位置的，所以=pivot值的时候也是需要交换的
	for low < high {
		for low < high && nums[low] < pivot {
			low++
		}

		for low < high && nums[high] > pivot {
			high--
		}

		if low >= high {
			break
		}

		nums[low], nums[high] = nums[high], nums[low]
	}

	// 返回left,就是left-1
	// 返回right,就是right+1
	// 此时right <= left
	return high
}

func QuickSortV2(nums []int) []int {
	SortV2(nums, 0, len(nums)-1)
	return nums
}

// SortV2 使用一个指针
func SortV2(nums []int, low, high int) {
	if low >= high {
		return
	}

	cur, pivot := low+1, low
	for cur <= high {
		if nums[cur] <= nums[low] {
			nums[pivot+1], nums[cur] = nums[cur], nums[pivot+1]
			pivot++
		}
		cur++
	}

	nums[low], nums[pivot] = nums[pivot], nums[low]
	SortV2(nums, low, pivot-1)
	SortV2(nums, pivot+1, high)
}

// QuickSortV3 挖坑法，也是始终以区间的开始位置为pivot
func QuickSortV3(nums []int) []int {
	SortV3(nums, 0, len(nums)-1)
	return nums
}

func SortV3(nums []int, low, high int) {
	if low >= high {
		return
	}

	pivotIndex := partitionV3(nums, low, high)
	SortV3(nums, low, pivotIndex-1)
	SortV3(nums, pivotIndex+1, high)
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

// QuickSortV4 双指针法，始终选择区间开始位置为pivot
func QuickSortV4(nums []int) []int {
	SortV4(nums, 0, len(nums)-1)
	return nums
}

func SortV4(nums []int, low, high int) {
	if low >= high {
		return
	}

	pivotIndex := partitionV4(nums, low, high)
	SortV4(nums, low, pivotIndex-1)
	SortV4(nums, pivotIndex+1, high)
}

func partitionV4(nums []int, low, high int) int {
	// 双指针法不会改变pivotIndex位置的值，为了保持统一写法，还这么写
	pivot := nums[low]
	pivotIndex := low

	for low < high {
		// high挪动到一个比基准值小的位置，准备交换
		for low < high && nums[high] >= pivot {
			high--
		}

		// low挪动到一个比基准值大的位置，准备交换，相等的值不需要挪动位置
		for low < high && nums[low] <= pivot {
			low++
		}

		// 其实不可能大于，最多是等于
		if low >= high {
			break
		}

		nums[low], nums[high] = nums[high], nums[low]
	}

	// low之前的数字必然是小于等于pivot的，所以需要和low位置的数做交换
	nums[pivotIndex], nums[low] = nums[low], nums[pivotIndex]
	return pivotIndex
}

// QuickSortV5 非递归写法
func QuickSortV5(nums []int) []int {
	var stack []int
	stack = append(stack, len(nums)-1)
	stack = append(stack, 0)
	for len(stack) > 0 {
		low := stack[len(stack)-1]
		high := stack[len(stack)-2]

		stack = stack[:len(stack)-2]
		if low < high {
			pivotIndex := partitionV3(nums, low, high)
			log.Println("pivotIndex:", pivotIndex, nums[pivotIndex])
			stack = append(stack, pivotIndex-1, low, high, pivotIndex+1)
		}
	}

	return nums
}
