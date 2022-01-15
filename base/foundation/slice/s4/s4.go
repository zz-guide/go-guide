package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	c := a[0:len(a)]
	// 0-长度减1可以把最后一个元素弹出
	b := a[0 : len(a)-1]

	fmt.Println("a:", a)
	fmt.Println("c:", c)
	fmt.Println("b:", b)
}
