package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/n-queens/

N 皇后

n皇后问题 研究的是如何将 n个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的n皇后问题 的解决方案。

每一种解法包含一个不同的n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。


提示：
1 <= n <= 9



*/
func main() {
	n := 4
	fmt.Println("N皇后:", solveNQueens(n))
}

func solveNQueens(n int) [][]string {
	return nil
}
