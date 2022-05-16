package main

import (
	"fmt"
	"strings"
)

/**
题目：https://leetcode-cn.com/problems/n-queens/

N 皇后

n皇后问题 研究的是如何将 n个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的n皇后问题 的解决方案。

每一种解法包含一个不同的n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。


提示：
1 <= n <= 9


注意：
	1.不能同行，不能同列，不能在一条斜线

步骤：
	1.先构建n*n的二维数组，值默认都是.
	2.helper函数开始回溯，从第0行开始
	3.从第0列开始，一直到n
	4.isValid函数逻辑，[i][j] 检查列，2个对角是不是冲突了
	5.边界条件，返回最终数据

*/
func main() {
	n := 4
	fmt.Println("N皇后-回溯:", solveNQueens(n))
	fmt.Println("N皇后-回溯优化空间:", solveNQueens2(n))
}

// solveNQueens 回溯 时间复杂度：O(N!) O(N)
func solveNQueens(n int) [][]string {
	// 先处理成[[....],[....]]，全是.
	bd := make([][]string, n)
	for i := range bd {
		bd[i] = make([]string, n)
		for j := range bd[i] {
			bd[i][j] = "."
		}
	}

	// 此方法每次都得从开始遍历寻找出现过的皇后位置，可以优化，用数组或者map记录
	var isValid func(rows, cols int, bd [][]string) bool
	isValid = func(rows, cols int, bd [][]string) bool {
		for i := 0; i < rows; i++ { // 之前的行
			for j := 0; j < n; j++ { // 所有列
				// 发现了皇后，并且和自己同列/对角线,对角线有2个，45度和135度
				// 不需要检查同行，因为每次遍历都不可能是同行
				if bd[i][j] == "Q" && (j == cols || i+j == rows+cols || i-j == rows-cols) {
					return false
				}
			}
		}
		return true
	}

	var res [][]string
	var helper func(start int, bd [][]string, n int)
	helper = func(start int, bd [][]string, n int) {
		// 退出条件
		if start == n {
			// 把对应行转成string
			temp := make([]string, len(bd))
			for i := 0; i < n; i++ {
				temp[i] = strings.Join(bd[i], "") // 将每一行拼成字符串，bd[i]
			}

			// 加入结果集
			res = append(res, temp)
			return
		}

		// i表示列，start表示行
		for i := 0; i < n; i++ { // 枚举出所有选择
			// 假设当前位置要放皇后，看之前的位置满足不满足条件
			if isValid(start, i, bd) { // 剪掉无效的选择
				bd[start][i] = "Q"     // 作出选择，放置皇后
				helper(start+1, bd, n) // 继续选择，往下递归
				bd[start][i] = "."     // 撤销当前选择
			}
		}
	}

	helper(0, bd, n)
	return res
}

// solveNQueens2 优化后的
func solveNQueens2(n int) [][]string {
	bd := make([][]string, n)
	for i := range bd {
		bd[i] = make([]string, n)
		for j := range bd[i] {
			bd[i][j] = "."
		}
	}

	// 记录出现过的皇后位置
	cols := map[int]bool{}
	diag1 := map[int]bool{} // start+i
	diag2 := map[int]bool{} // start-i

	var res [][]string
	var helper2 func(r int, bd [][]string)
	helper2 = func(start int, bd [][]string) {
		if start == n {
			temp := make([]string, len(bd))
			for i := 0; i < n; i++ {
				temp[i] = strings.Join(bd[i], "")
			}
			res = append(res, temp)
			return
		}

		for i := 0; i < n; i++ {
			// 剪枝
			if !cols[i] && !diag1[start+i] && !diag2[start-i] {
				bd[start][i] = "Q"

				cols[i] = true
				diag1[start+i] = true
				diag2[start-i] = true

				helper2(start+1, bd)
				// 撤销
				bd[start][i] = "."
				cols[i] = false
				diag1[start+i] = false
				diag2[start-i] = false
			}
		}
	}

	helper2(0, bd)
	return res
}
