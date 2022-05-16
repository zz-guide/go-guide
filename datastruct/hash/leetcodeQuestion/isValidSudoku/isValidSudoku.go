package main

import "log"

/**
题目：https://leetcode-cn.com/problems/valid-sudoku/

有效的数独

请你判断一个9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字1-9在每一行只能出现一次。
数字1-9在每一列只能出现一次。
数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。（请参考示例图）

注意：

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
空白格用'.'表示。

提示：

board.length == 9
board[i].length == 9
board[i][j] 是一位数字（1-9）或者 '.'

*/

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	log.Println("有效的数独:", isValidSudoku(board))
}

// isValidSudoku map保存 行，列，小宫格 的元素是否出现
func isValidSudoku(board [][]byte) bool {
	rm, cm, nm := make(map[int]bool), make(map[int]bool), make(map[int]bool)
	for i, row := range board {
		for j, col := range row {
			if col == '.' {
				continue
			}

			tmp := int(col - '0')
			// 计算第几个单元格
			n := (i/3)*3 + j/3
			// 节省后续计算耗时
			a, b, c := i*10+tmp, j*10+tmp, n*10+tmp
			if rm[a] || cm[b] || nm[c] {
				return false
			}
			rm[a] = true
			cm[b] = true
			nm[c] = true
		}
	}
	return true
}
