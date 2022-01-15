package main

import (
	"fmt"
	"unsafe"
)

/**
类型转换：
	1.普通变量类型int,float,stringSearch 都可以使用 type (a)这种形式来进行强制类型转换,比如
	2.指针的强制类型转换需要用到unsafe包中的函数实现
*/
func main() {
	F2()
}

func F1() {
	var a int32 = 10
	var b int64 = int64(a)
	var c float32 = 12.3
	var d float64 = float64(c)

	fmt.Println("b:", b)
	fmt.Println("d:", d)
}

func F2() {
	var a int = 10
	var b *int = &a
	var c *int64 = (*int64)(unsafe.Pointer(b))
	fmt.Println(*c)
}
