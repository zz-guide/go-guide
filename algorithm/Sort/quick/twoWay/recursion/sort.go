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

*/

func main() {
	arr := []int{10, 5, 7, 3, 4, 2, 8}
	log.Println("快速排序-递归法(2指针):", QuickSortV1(arr))
	log.Println("快速排序-递归法(1指针):", QuickSortV2(arr))
}

func QuickSortV1(arr []int) []int {
	SortV1(arr, 0, len(arr)-1)
	return arr
}

// SortV1 2个指针
func SortV1(arr []int, low, hight int) {
	if low >= hight {
		return
	}

	left, right := low, hight
	pivot := arr[(low+hight)/2] // 这里的经验值取的是中间数，经过 Benchmark 测试，确实比较优秀

	for left <= right {
		// 从左边开始迭代

		// 左边的数如果比 pivot 小，那么就应该将他放在左边，继续向右滑动，遇到一个比他大的为止
		for arr[left] < pivot {
			left++
		}

		// 右边的数如果比 pivot 大，那么就应该将他放在右边，继续向左滑动，遇到一个比他小的为止
		for arr[right] > pivot {
			right--
		}

		// 这里进行一次交换，将上面碰到的大数和小数交换一次
		//left 继续右走，right 继续左走 注意这里还不一定相遇，去继续执行上面的逻辑
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	// 最终left会大于right
	SortV1(arr, low, right)
	SortV1(arr, left, hight)
}

func QuickSortV2(arr []int) []int {
	SortV2(arr, 0, len(arr)-1)
	return arr
}

// SortV2 使用一个指针
func SortV2(arr []int, left, right int) {
	if left >= right {
		return
	}
	cur, lo := left+1, left
	for cur <= right {
		if arr[cur] <= arr[left] {
			arr[lo+1], arr[cur] = arr[cur], arr[lo+1]
			lo++
		}
		cur++
	}
	arr[left], arr[lo] = arr[lo], arr[left]
	SortV2(arr, left, lo-1)
	SortV2(arr, lo+1, right)
}
