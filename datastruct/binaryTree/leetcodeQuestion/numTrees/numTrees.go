package main

import "log"

/**
题目：https://leetcode-cn.com/problems/unique-binary-search-trees/

给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。

不同的二叉搜索树

给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。

提示：

1 <= n <= 19

注意：
	1. G(n)=f(1)+f(2)+f(3)+f(4)+...+f(n)， 假设 n 个节点存在二叉排序树的个数是 G (n)
	2. 令 f(i) 为以 i 为根的二叉搜索树的个数 f(i)=G(i−1)∗G(n−i)
	综合两个公式可以得到 卡特兰数 公式
	3.G(n)=G(0)∗G(n−1)+G(1)∗G(n−2)+...+G(n−1)∗G(0)

*/
func main() {
	n := 5
	log.Println("不同的二叉搜索树-动态规划:", numTrees(n))
	log.Println("不同的二叉搜索树-卡塔兰数:", numTrees2(n))
}

// numTrees1  O(n^2) O(n)
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1 // 表示0个，1个结点任意选择一个结点作为root的不同BST的种类

	for end := 2; end <= n; end++ {
		for i := 1; i <= end; i++ {
			// dp[m] = dp[m-1]*dp[n-m]
			dp[end] += dp[i-1] * dp[end-i]
		}
	}

	return dp[n]
}

// numTrees2 卡塔兰数 O(n) O(1)
func numTrees2(n int) int {
	C := 1
	for i := 0; i < n; i++ {
		C = C * 2 * (2*i + 1) / (i + 2)
	}

	return C
}
