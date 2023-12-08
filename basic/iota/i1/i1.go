package main

import "fmt"

/**
iota是常量的计数器，可以理解为const定义常量的行数的索引，注意是行数。
const中每新增一行常量声明将使iota计数一次，当然前提是iota在const中出现。
主要应用场景是在需要枚举的地方
*/
func main() {
	fmt.Println("a1:", a1)
	fmt.Println("b2:", a2)
	fmt.Println("a3:", a3)
	fmt.Println("------------")
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)
	fmt.Println("b4:", b4)
}

const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
)

// iota位于第三行，值就是n-1,从0开始计算
const (
	b1 = 1
	b2 = 4
	b3 = iota
	b4 = iota
)
