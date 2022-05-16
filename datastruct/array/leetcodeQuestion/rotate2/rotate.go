package main

import "log"

/**
题目：https://leetcode-cn.com/problems/rotate-image/

旋转图像

给定一个 n×n 的二维矩阵matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

示例 1：


输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[[7,4,1],[8,5,2],[9,6,3]]
示例 2：


输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

提示：

n == matrix.length == matrix[i].length
1 <= n <= 20
-1000 <= matrix[i][j] <= 1000


思路：
	1.要求直接修改原数组，不能返回新数组
*/
func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	log.Println("旋转图像：", matrix)
	matrix1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate1(matrix1)
	log.Println("旋转图像：", matrix1)
	matrix2 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate2(matrix2)
	log.Println("旋转图像：", matrix2)
}

// rotate 用翻转代替旋转 O(N^2) O(1)
func rotate(matrix [][]int) {
	// 水平+对角线=反转90度
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		// j必须小于i
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// 原地旋转 rotate1 O(N^2) O(1)
func rotate1(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j],
				matrix[n-j-1][i],
				matrix[n-i-1][n-j-1],
				matrix[j][n-i-1] =
				matrix[n-j-1][i],
				matrix[n-i-1][n-j-1],
				matrix[j][n-i-1],
				matrix[i][j]
		}
	}
}

// rotate2 使用辅助数组 O(N^2) O(N^2)
func rotate2(matrix [][]int) {
	n := len(matrix)
	tmp := make([][]int, n)
	for i := range tmp {
		tmp[i] = make([]int, n)
	}

	for i, row := range matrix {
		for j, v := range row {
			tmp[j][n-1-i] = v
		}
	}

	copy(matrix, tmp)
}
