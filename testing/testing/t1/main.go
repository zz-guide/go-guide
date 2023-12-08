package main

import "log"

func main() {
	a := 1
	b := 2
	res := Max(a, b)
	log.Println("res:", res)
}

func T1() {
	// 测试文件以_test.go结尾
	// go test
	// 忽略_或.开头的文件
	// 忽略testdata子目录
	// 执行go vet
	// 函数名：Test<Name>
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
