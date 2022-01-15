package main

import (
	"fmt"
	"math"
)

/**
求 1+2+ ... +n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句。
*/
func main() {
	result := sumNums(3)
	fmt.Println("result:", result)
}

func plus(a *int, b int) bool {
	*a += b
	fmt.Println("a:", a, ";b:", b)
	return true
}

// 方法1
func sumNums(n int) int {
	_ = n > 0 && plus(&n, sumNums(n-1))
	return n
}

// 方法2
func sumNums1(n int) int {
	return (int(math.Pow(float64(n), float64(2))) + n) >> 1
}
