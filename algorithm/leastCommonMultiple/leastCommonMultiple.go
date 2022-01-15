package main

import "fmt"

/**
最小公倍数：
两个或多个整数公有的倍数叫做它们的公倍数，其中除0以外最小的一个公倍数就叫做这几个整数的最小公倍数。整数a，b的最小公倍数记为[a，b]，同样的，a，b，c的最小公倍数记为[a，b，c]，多个整数的最小公倍数也有同样的记号。
与最小公倍数相对应的概念是最大公约数，a，b的最大公约数记为（a，b）。关于最小公倍数与最大公约数，我们有这样的定理：(a,b)x[a,b]=ab(a,b均为整数)。
*/
func main() {
	a, b := 12, 15
	fmt.Println("最小公倍数-穷举法:", lcm(a, b))
	fmt.Println("最小公倍数-辗转相除法:", lcm1(a, b))
}

/*
*穷举写法：最小公倍数
 */
func lcm(x, y int) int {
	var top int = x * y
	var i = x
	if x < y {
		i = y
	}
	for ; i <= top; i++ {
		if i%x == 0 && i%y == 0 {
			return i
		}
	}
	return top
}

/*
*公式解法：最小公倍数=两数之积/最大公约数
 */
func lcm1(x, y int) int {
	return x * y / gcd(x, y)
}

func gcd(x, y int) int {
	var tmp int
	for {
		tmp = x % y
		if tmp > 0 {
			x = y
			y = tmp
		} else {
			return y
		}
	}
}
