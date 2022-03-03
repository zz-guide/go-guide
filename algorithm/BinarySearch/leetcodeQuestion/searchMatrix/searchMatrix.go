package main

import (
	"log"
)

/**
题目：https://leetcode-cn.com/problems/search-a-2d-matrix/

搜索二维矩阵

编写一个高效的算法来判断m x n矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。

提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-104 <= matrix[i][j], target <= 104

注意：
	1.矩阵的每一行都是递增顺序的
	2.每一行的开头都比上一行的结尾大，也就是一条龙递增顺序

*/
func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	matrix1 := [][]int{{12}}
	target := 13
	log.Println("搜索二维矩阵-一次二分查找:", searchMatrix(matrix, target))
	log.Println("搜索二维矩阵-两次二分查找:", searchMatrix1(matrix1, target))
}

// searchMatrix 二维数组转一维数组做法 O(logmn)
func searchMatrix(matrix [][]int, target int) bool {
	// 此方法适用于每一行元素个数相同的场景，不同的话无法求出列
	rows, cols := len(matrix), len(matrix[0])
	l, r := 0, rows*cols-1 // 转一维
	for l <= r {
		mid := (r-l)>>1 + l // mid是一维的索引

		curRow := mid / cols        // 整除，得二维的当前行索引
		curCol := mid - curRow*cols // 一维mid减去它头顶上行的元素个数，得二维的当前列索引

		if matrix[curRow][curCol] == target {
			return true
		} else if matrix[curRow][curCol] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}

// searchMatrix1 两次二分查找，先找第一列最后一个比target小的元素，然后确定在第几行，然后在行里找target
func searchMatrix1(matrix [][]int, target int) bool {
	// 先找所在行，重点步骤
	l, r := 0, len(matrix)
	for l < r {
		mid := l + (r-l)>>1
		if matrix[mid][0] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	// 判断没有元素或者一个元素的情况
	rows := l - 1
	if rows == -1 {
		return false
	}

	// 再找所在列
	l, r = 0, len(matrix[rows])-1
	for l <= r {
		mid := l + (r-l)>>1
		if matrix[rows][mid] == target {
			return true
		} else if matrix[rows][mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
