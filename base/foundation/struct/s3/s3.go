package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	//s := Student{Name: "许磊"}
	//s = nil
	s1 := []int{1, 2}
	// struct不能使用range遍历
	for i, i2 := range s1 {
		fmt.Println("i:", i)
		fmt.Println("i:", i2)
	}
}
