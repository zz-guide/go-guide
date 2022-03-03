package main

import (
	"log"
	"math"
)

/**
题目：https://leetcode-cn.com/problems/divide-two-integers

两数相除

给定两个整数，被除数dividend和除数divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数dividend除以除数divisor得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2


示例1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
示例2:

输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = truncate(-2.33333..) = -2

提示：
	1.被除数和除数均为 32 位有符号整数。
	2.除数不为0。
	3.假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31, 2^31− 1]。本题中，如果除法结果溢出，则返回 2^31− 1。

*/

func main() {
	dividend := 10
	divisor := 3
	log.Println("最小值：", math.MinInt32)
	log.Println("两数相除:", divide(dividend, divisor))
}

// 时间复杂度: O(1), O(1)
// 步骤：1、先处理边界 2.决定好结果的符号位 3.
func divide(dividend int, divisor int) int {
	// 除数是-1且a是最小值，结果会溢出，返回最大值即可
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	// 处理边界
	if dividend == divisor {
		return 1
	}

	if dividend == 0 {
		return 0
	}

	// 除数绝对值最大，结果必为 0 或 1
	if divisor == math.MinInt32 {
		if dividend == divisor {
			return 1
		} else {
			return 0
		}
	}

	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}

	res, step, sum := 0, 0, 0
	for {
		// 将divisor成倍增长，并且是正数
		tmp := 0
		if sign == -1 {
			tmp = sum - divisor<<step
		} else {
			tmp = sum + divisor<<step
		}

		if (dividend > 0 && tmp > dividend) || (dividend < 0 && tmp < dividend) {
			if step == 0 {
				break
			}

			step-- //步长减半
			continue
		}

		sum = tmp
		res += 1 << step
		step++ //步长加倍
	}

	// 因为不能使用乘法，需要分开判断符号位
	if sign == -1 {
		return -res
	}

	return res
}
