package main

import "log"

/**
题目：https://leetcode.cn/problems/powx-n/solution/powx-n-by-leetcode-solution/
Pow(x, n)

实现pow(x, n)，即计算 x 的 n 次幂函数（即，xn ）。


示例 1：

输入：x = 2.00000, n = 10
输出：1024.00000
示例 2：

输入：x = 2.10000, n = 3
输出：9.26100
示例 3：

输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25

提示：

-100.0 < x < 100.0
-231 <= n <= 231-1
-104 <= xn <= 104

*/
func main() {
	x := 2.00000
	n := -2
	log.Println("实现pow()函数:", myPow(x, n))
}

// myPow 迭代法 O(logN)，O(1)
func myPow(x float64, n int) float64 {
	// 正数的情况
	if n >= 0 {
		return quickMul(x, n)
	}
	// 负数的情况
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, N int) float64 {
	if N == 1 {
		return x
	}

	if N == 0 {
		return 1.0
	}

	// 任何数的0次幂都是1
	res := 1.0
	tmp := x

	// 这个循环通过N/2降低循环次数，时间复杂度变为log级别
	for N > 0 {
		// 如果不够平分，比如N是奇数的时候
		if N%2 == 1 {
			res *= tmp
		}

		// N为偶数直接翻倍
		tmp *= tmp
		N /= 2
	}

	return res
}
