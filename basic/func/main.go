package main

import "log"

func main() {
	l1, l2 := Do()
	log.Println("l1:", l1)
	log.Println("l2:", l2)
}

func Do() (name string, err error) {
	// 1.可以直接赋值，然后返回
	name = "许磊"
	// 2.也可以直接覆盖返回
	return "xxxx", nil
}
