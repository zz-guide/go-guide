package main

import (
	"fmt"
)

type T struct {
	int
}

func (t T) testT() {
	fmt.Println("类型 *T 方法集包含全部 receiver T 方法。")
}

func (t *T) testP() {
	fmt.Println("类型 *T 方法集包含全部 receiver *T 方法。")
}

/**
类型 *T 方法集包含全部 receiver T + *T 方法。
*/
func main() {
	t1 := T{1}
	t2 := &t1
	fmt.Printf("t2 is : %v\n", t2)
	t2.testT()
	t2.testP()
}
