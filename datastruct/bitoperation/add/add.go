package main

import "log"

/**
题目： https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/
不用加减乘除做加法

写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。

示例:

输入: a = 1, b = 1
输出: 2

提示：

a,b均可能是负数或 0
结果不会溢出 32 位整数

*/
func main() {
	a := 5
	b := 6
	log.Println("不用加减乘除做加法:", add(a, b))
}

func add(a int, b int) int {
	for b != 0 {
		c := a & b << 1
		a ^= b
		b = c
	}
	return a
}
