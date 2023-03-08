package main

import "log"

func main() {
	tArray()
}

func tArray() {
	// 第一种方式
	var ages [2]int
	ages = [2]int{1, 2}
	log.Println(ages)

	// 第二种方式
	ages2 := [2]int{1, 2}
	log.Println(ages2)

	// 第三种方式 自动推算数组长度的定义方式
	ages3 := [...]int{1, 2}
	log.Println(ages3)
}
