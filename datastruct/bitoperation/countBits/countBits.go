package main

import "log"

/**
题目:https://leetcode-cn.com/problems/counting-bits/
比特位计数
给你一个整数 n ，对于0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。


示例 1：

输入：n = 2
输出：[0,1,1]
解释：
0 --> 0
1 --> 1
2 --> 10
示例 2：

输入：n = 5
输出：[0,1,1,2,1,2]
解释：
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101


提示：

0 <= n <= 105


进阶：

很容易就能实现时间复杂度为 O(n log n) 的解决方案，你可以在线性时间复杂度 O(n) 内用一趟扫描解决此问题吗？
你能不使用任何内置函数解决此问题吗？（如，C++ 中的 __builtin_popcount ）


*/
func main() {
	n := 9
	log.Println("比特位计数-Brian Kernighan 算法:", countBits(n))
	log.Println("比特位计数-动态规划1:", countBits1(n))
	log.Println("比特位计数-动态规划2:", countBits2(n))
	log.Println("比特位计数-动态规划3:", countBits3(n))
}

// countBits Brian Kernighan 算法,令 x=x&(x-1)，该运算将 x 的二进制表示的最后一个 11 变成 0。因此，对 x 重复该操作，直到 x 变成 0，则操作次数即为 x 的「一比特数」。
// 时间复杂度O(nlogn)，空间复杂度O(1)
func countBits(n int) []int {
	var onesCount func(x int) int
	onesCount = func(x int) int {
		var count int
		for ; x > 0; x &= x - 1 {
			count++
		}
		return count
	}

	bits := make([]int, n+1)
	for i := range bits {
		bits[i] = onesCount(i)
	}

	return bits
}

// countBits1 动态规划，最高有效位，时间复杂度O(n)，空间复杂度O(1)
// x=x&(x-1) 可以判断x是不是2的整数次幂
func countBits1(n int) []int {
	bits := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}

		bits[i] = bits[i-highBit] + 1
	}

	return bits
}

// countBits2 动态规划，最低有效位，时间复杂度O(n)，空间复杂度O(1)
// x=x&(x-1) 可以判断x是不是2的整数次幂
func countBits2(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i>>1] + i&1
	}
	return bits
}

// countBits3 动态规划，对最低设置为，时间复杂度O(n)，空间复杂度O(1)
// x=x&(x-1) 可以判断x是不是2的整数次幂
func countBits3(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i&(i-1)] + 1
	}
	return bits
}
