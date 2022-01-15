package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/climbing-stairs/
爬楼梯
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。
示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

*/
func main() {
	// 1 1 2 3 5 8 12

	fmt.Println("爬楼梯1:", ClimbStairs(5))
	fmt.Println("爬楼梯2:", ClimbStairs1(5))
}

// ClimbStairs 标准的动态规划
func ClimbStairs(n int) int {
	if n <= 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// ClimbStairs1 节省空间的动态规划
func ClimbStairs1(n int) int {
	if n <= 0 {
		return 0
	}

	pre, next := 1, 1

	var res = 1
	for i := 2; i <= n; i++ {
		res = pre + next
		pre = next
		next = res
	}

	return res
}
