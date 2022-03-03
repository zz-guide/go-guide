package main

import "log"

/**
题目：https://leetcode-cn.com/problems/spiral-matrix/

螺旋矩阵

给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100

*/
func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	log.Println("螺旋矩阵-遍历到底", spiralOrder(matrix))
	log.Println("螺旋矩阵-遍历不到底", spiralOrder1(matrix))
}

// 螺旋矩阵 spiralOrder 模拟螺旋顺序遍历
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	var res []int
	top, right, left, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	for top <= bottom && left <= right {
		// 按照先从左到右
		for i := top; i <= right; i++ {
			res = append(res, matrix[top][i])
		}

		top++

		// 从上到下
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}

		right--

		// 判断下是否越界，非常重要
		if top > bottom || left > right {
			break
		}

		// 从右到左
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}

		bottom--

		// 从下到上
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}

		left++
	}

	return res
}

func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	var res []int
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for top < bottom && left < right {
		for i := left; i < right; i++ {
			res = append(res, matrix[top][i])
		}
		for i := top; i < bottom; i++ {
			res = append(res, matrix[i][right])
		}
		for i := right; i > left; i-- {
			res = append(res, matrix[bottom][i])
		}
		for i := bottom; i > top; i-- {
			res = append(res, matrix[i][left])
		}

		right--
		top++
		bottom--
		left++
	}

	if top == bottom {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
	} else if left == right {
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][left])
		}
	}

	return res
}
