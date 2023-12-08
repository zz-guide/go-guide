package main

import (
	"fmt"
	"math"
	"strconv"
	"unsafe"
)

/**
结论：原码，反码，补码的产生过程，就是为了解决，计算机做减法和引入符号位（正号和负号）的问题。
原码：是最简单的机器数表示法。用最高位表示符号位，‘1’表示负号，‘0’表示正号。其他位存放该数的二进制的绝对值。
*/
func main() {
	//T1()
	//T2()
	//T3()
	//T4()
	//divide()
	overflow()
}

func T2() {
	// 1 - 2相当于1 + （-2），换成补码进行计算丢弃符号位就是255
	var a uint8 = 1
	var b uint8 = 2
	fmt.Println("a - b:", a-b) // 结果是最大值255

	// // 255 + 1换成补码计算全部位数为0，并且超出8位，丢弃，最终就是0000 0000 ，可不就是0么
	var c uint8 = 255
	var d uint8 = 1
	fmt.Println("c + d:", c+d) // 结果是0
}

func T3() {
	var a int8 = -128
	var b int8 = 1
	fmt.Println("a - b:", a-b) // 结果是最大值255
}

func T1() {
	a := 255
	fmt.Printf("a:%b\n", a)
}

func T4() {
	fmt.Println(bInt8(-1))
}

func bInt8(n int8) string {
	// fmt内部做了格式化，没法打印正确的负数，只能通过unsafe打印补码
	return strconv.FormatUint(uint64(*(*uint8)(unsafe.Pointer(&n))), 2)
}

func divide() {
	fmt.Println("5/2:", 7/8)             // 向下取整
	fmt.Println("5/2:", math.Ceil(2.5))  // 向上取整
	fmt.Println("5/2:", math.Floor(2.5)) // 向下取整
	fmt.Println("5/2:", math.Round(2.5)) // 四舍五入
}

func overflow() {
	var a uint8 = 255
	var b uint8 = 2
	fmt.Println("res:", (a+b)/2) // 溢出以后结果不对
}
