package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/fibonacci-number/
斐波那契数列
*/
func main() {
	// 1 1 2 3 5 8 12

	fmt.Println("fib1结果:", fib1(16))
	fmt.Println("fib2结果:", fib2(16))
}

// DP TABLE
func fib1(n int) int {
	if n < 2 {
		return n
	}

	if n == 2 {
		return 1
	}

	// 从优化递归暴力解法中受到启发，变成DP TABLE
	// 创建DP集合，map或者slice都行
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 继续优化DP，上一个方法的缺点是DP数组存储了从0~n的全部值，实际上只需要知道（n-1）和（n-2）的值就可以了，
// 进一步优化空间利用
func fib2(n int) int {
	if n < 2 {
		return n
	}

	if n == 2 {
		return 1
	}

	pre, cur := 0, 1
	for i := 2; i <= n; i++ {
		// 状态转移方程
		sum := pre + cur
		// 先把pre前移,(n-1)赋值给(n-2)
		pre = cur
		// 再把相加的值给了n赋值给cur(n-1)
		cur = sum
	}

	return cur
}
