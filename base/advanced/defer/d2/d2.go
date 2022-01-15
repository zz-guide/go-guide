package main

import (
	"fmt"
	"os"
)

func main() {
	deferExit()
}

// 结论：1.os.Exit 可导致defer不执行
// defer执行的时机：1.包裹defer的函数return 2.包裹defer的函数执行到末尾  3.goroutine panic的时候会执行本goroutine的所有defer
func deferExit() {
	defer func() {
		fmt.Println("defer")
	}()
	os.Exit(0)
}
