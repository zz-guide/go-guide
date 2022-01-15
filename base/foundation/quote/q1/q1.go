package main

import "fmt"

func main() {
	F1()
	F2()
	F3()
}

func F1() {
	// 单字符 等同于rune
	a := '世'
	b := '界'
	c := a + b
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}

func F2() {
	// 可以转义
	a := "世界\n"
	b := "你好\n"
	c := a + b
	fmt.Printf("a: %s6", a)
	fmt.Printf("b: %s6", b)
	fmt.Printf("c: %s6", c)
}

func F3() {
	// 原样输出
	a := `世界\n`
	b := `世界\n`
	c := a + b
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}
