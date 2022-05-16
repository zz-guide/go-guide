package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/n-queens-ii/

n皇后问题 研究的是如何将 n个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。

提示：

1 <= n <= 9

注意：1.此题目要求返回数量，不要求具体项

*/
func main() {
	n := 4
	fmt.Println("N皇后 II-回溯位运算:", totalNQueens(n))
}

func totalNQueens(n int) (ans int) {
	var solve func(row, columns, diagonals1, diagonals2 int)
	solve = func(row, columns, diagonals1, diagonals2 int) {
		if row == n {
			ans++
			return
		}
		availablePositions := (1<<n - 1) &^ (columns | diagonals1 | diagonals2)
		for availablePositions > 0 {
			position := availablePositions & -availablePositions
			solve(row+1, columns|position, (diagonals1|position)<<1, (diagonals2|position)>>1)
			availablePositions &^= position // 移除该比特位
		}
	}
	solve(0, 0, 0, 0)
	return
}
