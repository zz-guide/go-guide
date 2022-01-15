package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/fibonacci-number/
斐波那契数列--递归解法（暴力解法）
*/
func main() {
	// 1 1 2 3 5 8 12

	fmt.Println("fib1结果:", fib1(10))
	fmt.Println("fib2结果:", fib2(10))
}

// 暴力解法
// 问题:存在重复计算，重叠子问题
func fib1(n int) int {
	if n < 2 {
		return n
	}

	if n == 2 {
		return 1
	}

	return fib1(n-1) + fib1(n-2)
}

// 优化暴力解法，使用数组或者map记录已经计算过的值
func fib2(n int) int {
	m := make([]int, n+1) // 第0位不存数据，所以多分配一个位置
	return helper(m, n)
}

func helper(m []int, n int) int {
	if n < 2 {
		return n
	}

	if n == 2 {
		return 1
	}

	if m[n] != 0 { // 说明已经计算过了
		return m[n]
	}

	m[n] = helper(m, n-1) + helper(m, n-2) // 此处递归等同于dp[i] = dp[i-1] + dp[i-2]，可以换成数组
	return m[n]
}
