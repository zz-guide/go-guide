package main

import "fmt"

/**
最大公约数：
最大公因数，也称最大公约数、最大公因子，指两个或多个整数共有约数中最大的一个。
a，b的最大公约数记为（a，b），同样的，a，b，c的最大公约数记为（a，b，c），多个整数的最大公约数也有同样的记号。
求最大公约数有多种方法，常见的有质因数分解法、短除法、辗转相除法、更相减损法。
与最大公约数相对应的概念是最小公倍数，a，b的最小公倍数记为[a，b]。
*/
func main() {
	a, b := 12, 15
	fmt.Println("最大公约数-穷举法:", gcd(a, b))
	fmt.Println("最大公约数-更相减损术:", gcd1(a, b))
	fmt.Println("最大公约数-辗转相除法:", gcd2(a, b))
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// gcd 穷举法
func gcd(a int, b int) int {
	var n int

	if a > b {
		n = b
	} else {
		n = a
	}

	for i := n; i >= 1; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}

	return 1
}

// 更相减损术
//辗转相除法如果 a 和 b 都很大的时候，a % b 性能会较低。在中国，
//《九章算术》中提到了一种类似辗转相减法的 更相减损术。
//它的原理是：两个正整数 a 和 b（a>b），它们的最大公约数等于 a-b 的差值 c 和较小数 b 的最大公约数。
func gcd1(a int, b int) int {
	if a == b {
		return a
	}

	if a < b {
		return gcd(b-a, a)
	}

	return gcd(a-b, b)
}

// 辗转相处法
//如果我们需要计算 a 和 b 的最大公约数，运用辗转相除法的话。
//首先，我们先计算出 a 除以 b 的余数 c，把问题转化成求出 b 和 c 的最大公约数；
//然后计算出 b 除以 c 的余数 d，把问题转化成求出 c 和 d 的最大公约数；
//再然后计算出 c 除以 d 的余数 e，把问题转化成求出 d 和 e 的最大公约数。..... 以此类推，
//逐渐把两个较大整数之间的运算转化为两个较小整数之间的运算，直到两个数可以整除为止。
func gcd2(a, b int) int {
	var tmp int
	for {
		tmp = a % b
		if tmp > 0 {
			a = b
			b = tmp
		} else {
			return b
		}
	}
}
