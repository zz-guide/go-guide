package main

import (
	"bytes"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/zigzag-conversion/

Z 字形变换

将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);

示例 1：

输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"
示例 2：
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I
示例 3：

输入：s = "A", numRows = 1
输出："A"

提示：

1 <= s.length <= 1000
s 由英文字母（小写和大写）、',' 和 '.' 组成
1 <= numRows <= 1000

*/
func main() {
	s := "PAYPALISHIRING"
	numRows := 5
	log.Println("Z 字形变换-直接模拟:", convert(s, numRows))
	log.Println("Z 字形变换-二维数组:", convert2(s, numRows))
	log.Println("Z 字形变换-二维数组优化:", convert3(s, numRows))
}

// convert 直接模拟 O(n) O(1)
// 关键点：1.每个周期最下边一行在s的位置左右两边对称 2.利用二维矩阵寻找周期长度公式
func convert(s string, numRows int) string {
	length, rows := len(s), numRows
	if rows == 1 || rows >= length {
		return s
	}

	// 每一个周期都是6个字符，并且包含了下一个周期的开始字符
	// 第一行和最后一行都只有每个周期的一个字符，其余的行有多个字符

	cycleLength := rows*2 - 2 // 周期长度
	var ans []byte
	for i := 0; i < rows; i++ { // 枚举矩阵的行
		for j := 0; j+i < length; j += cycleLength { // 枚举每个周期的起始下标
			ans = append(ans, s[j+i]) // 当前周期的第一个字符
			// 中间行
			// 每一个周都是以最下边的字符为对称点
			if (i > 0 && i < rows-1) && j+cycleLength-i < length {
				ans = append(ans, s[j+cycleLength-i])
			}
		}
	}

	return string(ans)
}

// 利用二维矩阵模拟
func convert2(s string, numRows int) string {
	n, r := len(s), numRows
	if r == 1 || r >= n {
		return s
	}

	t := r*2 - 2
	c := (n + t - 1) / t * (r - 1)
	mat := make([][]byte, r)
	for i := range mat {
		mat[i] = make([]byte, c)
	}

	x, y := 0, 0
	for i, ch := range s {
		mat[x][y] = byte(ch)
		if i%t < r-1 {
			x++ // 向下移动
		} else {
			x--
			y++ // 向右上移动
		}
	}

	ans := make([]byte, 0, n)
	for _, row := range mat {
		for _, ch := range row {
			if ch > 0 {
				ans = append(ans, ch)
			}
		}
	}

	return string(ans)
}

func convert3(s string, numRows int) string {
	r := numRows
	if r == 1 || r >= len(s) {
		return s
	}
	mat := make([][]byte, r)
	t, x := r*2-2, 0
	for i, ch := range s {
		mat[x] = append(mat[x], byte(ch))
		if i%t < r-1 {
			x++
		} else {
			x--
		}
	}
	return string(bytes.Join(mat, nil))
}
