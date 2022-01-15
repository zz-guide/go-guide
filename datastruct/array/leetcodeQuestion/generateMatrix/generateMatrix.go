package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/spiral-matrix-ii/

螺旋矩阵 II
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。


输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]

提示：
1 <= n <= 20

*/
func main() {
	n := 3
	fmt.Println("螺旋矩阵 II:", generateMatrix(n))
}

// generateMatrix 四个变量对应四个方向，画一下对应坐标的关系就知道了
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	num := 1
	tar := n * n
	left, right, top, bottom := 0, n-1, 0, n-1
	for num <= tar {
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++

		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++

	}

	return matrix
}
