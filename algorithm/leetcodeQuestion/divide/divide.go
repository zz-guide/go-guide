package main

import (
	"log"
	"math"
)

/**
题目：https://leetcode-cn.com/problems/divide-two-integers/

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

被除数和除数均为 32 位有符号整数。
除数不为0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31, 2^31− 1]。本题中，如果除法结果溢出，则返回 2^31− 1。

思路：
	1.不能使用*，/，%等符号
	2.考虑溢出
	3.使用移位每次扩大步长，然后使用减法,如果减法也不让用，那么就使用位运算模拟
	4.考虑异常情况，被除数最小，除数为1，-1
	5.除数最小，被除数为最小或者其他
	6.除数为0
*/

func main() {
	dividend := -2147483648
	divisor := -1
	log.Println("两数相除:", divide(dividend, divisor))
}

func divide(dividend int, divisor int) int {
	// 提前把异常情况过滤
	if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
		if divisor == 1 {
			return math.MinInt32
		}

		// 题目有说明，-1的时候除法结果溢出，返回最大值
		if divisor == -1 {
			return math.MaxInt32
		}
	}

	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}

		return 0
	}

	if dividend == 0 { // 考虑被除数为 0 的情况
		return 0
	}

	var res int
	// 第一步确定最终结果的符号
	sign := -1
	if (dividend ^ divisor) >= 0 {
		sign = 1
	}

	// 第二步骤：两个数字都转成负数,此处不能求绝对值，因为转为正数的过程中可能会溢出，所以都转换为负数
	var dividendTemp int
	if dividend < 0 {
		dividendTemp = dividend
	} else {
		dividendTemp = -dividend
	}

	var divisorTemp int
	if divisor < 0 {
		divisorTemp = divisor
	} else {
		divisorTemp = -divisor
	}

	// 第三步骤：阈值
	threshold := math.MinInt >> 1
	// 第四步骤：因为此处都是负数，所以正好相反
	for dividendTemp <= divisorTemp {
		tmp := divisorTemp
		times := 1 //除数divisor的倍数
		// tmp移位之后还是负数
		for tmp >= threshold && dividendTemp <= (tmp<<1) {
			tmp <<= 1
			times <<= 1
		}

		// 更新被除数
		dividendTemp -= tmp
		// 累加次数
		res += times
	}

	if sign < 0 {
		res = -res
	}

	if res > math.MaxInt {
		res = math.MaxInt
	}

	return res
}

// https://blog.csdn.net/jjwwwww/article/details/82745855
// Subtraction 位运算模拟减法
func Subtraction(num1, num2 int) int {
	x := num1 ^ num2
	y := x & num2

	for y != 0 {
		y = y << 1
		x = x ^ y
		y = x & y
	}

	return x
}

// Addition 位运算实现加法
func Addition(num1, num2 int) int {
	x := num1 ^ num2 // a^b得到原位和（相当于按位相加没有进位）
	y := num1 & num2

	for y != 0 {
		y = y << 1
		temp := x
		x = x ^ y
		y = temp & y
	}

	return x
}
