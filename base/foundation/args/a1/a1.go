package main

import "fmt"

func main() {
	c := []int{2, 3, 4}
	MultiParams(1, c...)
	//MultiParams(1, 2, "ss")
}

func MultiParams(a int, b ...int) {
	fmt.Println("a:", a)
	fmt.Printf("v:%+v\n", b)
	for i, v := range b {
		fmt.Println("i:", i, "  v:", v)
	}
}
