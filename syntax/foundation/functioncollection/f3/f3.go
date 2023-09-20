package main

import (
	"fmt"
)

type S struct {
	T
}

type T struct {
	int
}

func (t T) testT() {
	fmt.Println("如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T 方法")
}
func (t *T) testP() {
	fmt.Println("如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 *T 方法")
}

/*
如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T + *T 方法。
这条规则说的是当我们嵌入一个类型的指针，嵌入类型的接受者为值类型或指针类型的方法将被提升，可以被外部类型的值或者指针调用。
*/
func main() {
	s1 := S{T{1}}
	s2 := &s1
	fmt.Printf("s1 is : %v\n", s1)
	s1.testT()
	s1.testP()
	fmt.Printf("s2 is : %v\n", s2)
	s2.testT()
	s2.testP()
}
