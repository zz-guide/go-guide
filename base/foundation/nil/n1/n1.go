package main

import (
	"fmt"
)

/**
nil实质上是一个变量，不是关键词。定义在 buildin/buildin.go 中
*/

// nil 可以被全局覆盖
// 局部覆盖的话只在当前作用域起作用
var nil = "sss"

func main() {
	F3()
}

func F1() {
	fmt.Println("nil:", nil)
}

func F2() {
	fmt.Println("nil:", nil)
}

func F3() {
	var b bool
	fmt.Println("bool:", b)

	var i int
	fmt.Println("int:", i)

	var str string
	fmt.Println("string:", str)

	var m map[string]string
	fmt.Println("map[string]string:", m)

	var slice []int
	var slice1 = make([]int, 0)
	fmt.Println("slice:", slice)
	fmt.Println("slice:", slice1)

	var err error
	fmt.Println("error:", err)

	type Person struct {
		AgeYears int
		Name     string
		Friends  []Person
	}

	var p Person
	fmt.Printf("Person:%#v\n", p)

	type A interface{}
	var a A
	println("interface:", a, a == nil)
}
