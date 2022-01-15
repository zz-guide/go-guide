package main

import "fmt"

type Student struct {
	Name string
}

func main() {
	stu := &Student{Name: "许磊"}
	stu1 := stu

	fmt.Printf("stu %p %p %#v\n", stu, &stu, *stu)
	fmt.Printf("stu1 %p %p %#v\n", stu1, &stu1, *stu1)

	stu1 = &Student{Name: "李四"}
	fmt.Println("stu:", stu.Name)
	fmt.Println("stu1:", stu1.Name)
}
