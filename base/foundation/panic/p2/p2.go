package main

import (
	"fmt"
	"time"
)

/**
	结论：
		1.panic之后的defer不执行
		2.当前goroutine panic，当前定义的defer会执行
		3.recover可以随意低啊用，返回nil
		4.recover必须配合defer使用
		5.函数没有return也会执行defer
		6.得先注册defer，再return
		7.只会执行当前panic协程中的defer,不会执行调方的defer
		8.
 */
func main() {
	fmt.Println("----开始----")
	defer defer1()
	defer func () {
		if r := recover(); r != nil {
			fmt.Println("recover value is", r)
		}
	}()

	go goroutine()
	//fmt.Println(recover())
	time.Sleep(10 * time.Second)
	fmt.Println("----结束----")
}

func defer1() {
	fmt.Println("defer1")
}

func defer2() {
	fmt.Println("defer2")
}

func defer3() {
	fmt.Println("defer3")
}

func defer4() {
	fmt.Println("defer4")
}

func goroutine() {
	defer defer2()
	panic("产生panic")
	defer defer4()
}