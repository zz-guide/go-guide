package main

import "fmt"

/**
	结论：
		1.panic只会调用当前协程的defer，从而也只会调用当前协程的recover。
		2.panic之前按照出栈顺序调用defer。
 */
func main() {
	defer defer1()
	defer defer2()
	fmt.Println("in main")
	panic("panic崩溃了")
}

func defer1() {
	fmt.Println("defer1")
}

func defer2() {
	fmt.Println("defer2")
}