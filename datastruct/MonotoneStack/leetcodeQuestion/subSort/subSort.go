package main

import (
	"log"
	"math"
)

/**
题目：https://leetcode-cn.com/problems/sub-sort-lcci/

部分排序

给定一个整数数组，编写一个函数，找出索引m和n，只要将索引区间[m,n]的元素排好序，整个数组就是有序的。注意：n-m尽量最小，也就是说，找出符合条件的最短序列。函数返回值为[m,n]，若不存在这样的m和n（例如整个数组是有序的），请返回[-1,-1]。

示例：

输入： [1,2,4,7,10,11,7,12,6,7,16,18,19]
输出： [3,9]
提示：

0 <= len(array) <= 1000000

注意：
	1.m左侧的数必须要比n右侧的数小
	2.不能单纯的认为从左找到第一个不满足的就是start,从后向前找第一个不满足的就是end,还需要看end右侧是不是大于start左侧
	3.中心思想就是先找边界，然后向两边扩散，一直扩散到满足条件的地方
*/

func main() {
	nums := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	nums1 := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	nums2 := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	nums3 := []int{1, 3, 5, 7, 9}
	log.Println("部分排序-2个单调栈:", subSort(nums))
	log.Println("部分排序-双指针2个for:", subSort1(nums1))
	log.Println("部分排序-双指针1个for:", subSort2(nums2))
	log.Println("部分排序-中心扩展法(有bug等解决):", subSort3(nums3))
}

// subSort 单调栈,利用2个不同方向的单调栈，O(n),O(n)
func subSort(array []int) []int {
	start, end := -1, -1
	length := len(array)
	if length <= 1 {
		return []int{start, end}
	}

	var leftStack []int
	var rightStack []int

	lMax := math.MinInt
	rMin := math.MaxInt
	for i := 0; i < len(array); i++ {
		if array[i] > lMax {
			lMax = array[i]
		}

		if array[len(array)-1-i] < rMin {
			rMin = array[len(array)-1-i]
		}

		if i == 0 {
			leftStack = append(leftStack, array[i])
			rightStack = append(rightStack, array[len(array)-1-i])
			continue
		}

		// 先入左栈，非递增顺序
		leftStack = append(leftStack, lMax)
		// 再入右栈，非递减顺序
		rightStack = append(rightStack, rMin)
	}

	//输入： [1,  2,  4,  7,  10, 11, 7,  12, 6,  7,  16, 18, 19]
	//左大： [1,  2,  4,  7,  10, 11, 11, 12, 12, 12, 16, 18, 19]
	//右小： [1,  2,  4,  6,  6,  6,  6,  6,  6,  7,  16, 18, 9]

	// 第一个不相等的位置是start,最后一个不相等的位置是end
	left, right := -1, -1
	for i := 0; i < len(array); i++ {
		j := len(array) - 1 - i
		if leftStack[i] == rightStack[j] {
			continue
		} else {
			if left == -1 {
				left = i
			}

			right = i
		}
	}

	return []int{left, right}
}

// subSort1 双指针 O(n),O(1)，2个for
func subSort1(array []int) []int {
	start, end := -1, -1
	length := len(array)
	if length <= 1 {
		return []int{start, end}
	}

	// 假设数列排序结束后递增
	min := math.MaxInt32
	max := math.MinInt32
	for i := 0; i < length; i++ {
		if array[i] >= max {
			max = array[i]
		} else {
			// end会一直向后移动
			end = i
		}
	}

	for i := length - 1; i >= 0; i-- {
		if array[i] <= min {
			min = array[i]
		} else {
			// start会向前移动
			start = i
		}
	}

	return []int{start, end}
}

// subSort2 双指针，一个for, O(n),O(1)
func subSort2(array []int) []int {
	start, end := -1, -1
	length := len(array)
	if length <= 1 {
		return []int{start, end}
	}

	// 假设数列排序结束后递增
	min := math.MaxInt32
	max := math.MinInt32
	for i := 0; i < length; i++ {
		// 开头非递增
		if array[i] >= max {
			max = array[i]
		} else {
			// end会一直向后移动
			end = i
		}

		// 末尾非递减
		j := length - 1 - i
		if array[j] <= min {
			min = array[j]
		} else {
			// start会向前移动
			start = j
		}
	}

	return []int{start, end}
}

// subSort3 先确定边界，然后两边扩散(测试用例没过，等再看)
func subSort3(array []int) []int {
	start, end := -1, -1
	length := len(array)
	if length <= 1 {
		return []int{start, end}
	}

	// 先确定左边界，找到第一个不满足费递增顺序的index
	max := math.MinInt32
	for i := 0; i < length; i++ {
		if array[i] >= max {
			max = array[i]
		} else {
			start = i
			break
		}
	}

	// 再确定右边界，找到第一个不满足非递减顺序的index
	min := math.MaxInt32
	for i := length - 1; i >= 0; i-- {
		if array[i] <= min {
			min = array[i]
		} else {
			end = i
			break
		}
	}

	if start == -1 && end == -1 {
		return []int{start, end}
	}

	// 寻找start~end之间，从左往右的最小值，从右往左的最大值
	max = math.MinInt32
	min = math.MaxInt32
	for i := start; i <= end; i++ {
		if array[i] < min {
			min = array[i]
		}

		j := start + (end - i)
		if array[j] > max {
			max = array[j]
		}
	}

	for start > 0 {
		if array[start-1] >= min {
			start--
		} else {
			break
		}
	}

	for end < length-1 {
		if array[end+1] <= max {
			end++
		} else {
			break
		}
	}

	return []int{start, end}
}
