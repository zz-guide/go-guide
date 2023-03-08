package main

import "log"

func main() {
	TFor()
}

func t1() int {
	log.Println("sss")
	return 3
}

func TFor() {
	// 支持所有的迭代方式
	// 结论：循环变量重复使用，只初始化一次
	//for i, c := 0, t1(); i < c; i++ {
	//	log.Printf("i=%d,%p;c=%d,%p\n ", i, &i, c, &c)
	//}

	// 结论：条件表达式中的函数被多次调用，或被内联优化。
	c := 0
	for c < t1() {
		log.Println("c=", c)
		c++
	}
}
