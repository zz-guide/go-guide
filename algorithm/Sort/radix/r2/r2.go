package main

import (
	"log"
	"math"
	"strconv"
)

/**
基数排序（radix sort）属于“分配式排序”（distribution sort），又称“桶子法”（bucket sort）或bin sort，顾名思义，它是透过键值的部份资讯，将要排序的元素分配至某些“桶”中，藉以达到排序的作用，基数排序法是属于稳定性的排序，其时间复杂度为O (nlog(r)m)，其中r为所采取的基数，而m为堆数，在某些时候，基数排序法的效率高于其它的稳定性排序法。
*/
func main() {
	nums := []int{9, 0, 7, 1, 4, 8, 4, 9, 9, 7}
	log.Println("基数排序:", radixsort(nums))
}

func radixsort(arr []int) []int {
	maxValueLen := 0 // 需要循环多少次，以最大数字为准
	for i := 0; i < len(arr); i++ {
		n := len(strconv.Itoa(arr[i])) // 方便起见，数字转字符，再取长度
		if n > maxValueLen {
			maxValueLen = n
		}
	}
	for loc := 1; loc <= maxValueLen; loc++ {
		sort(arr, loc)
	}
	return arr
}

func sort(arr []int, loc int) {
	bucket := make([][]int, 10) // 0~9 总共10个队列

	for i := 0; i < len(arr); i++ {
		ji := digit(arr[i], loc) // 获取对应位的数字
		if bucket[ji] == nil {
			bucket[ji] = make([]int, 0) // 如果 bucket 需要再去初始化
		}

		bucket[ji] = append(bucket[ji], arr[i]) // 将数字 push 进去
	}

	idx := 0
	for i := 0; i <= 9; i++ {
		for j := 0; j < len(bucket[i]); j++ {
			// 遍历二维数组
			arr[idx] = bucket[i][j] //将数字取出来 给原数组重新赋值
			idx++
		}
	}
}

// 数字，右数第几位，从1开始
func digit(num int, loc int) int {
	return num % int(math.Pow10(loc)) / int(math.Pow10(loc-1))
}
