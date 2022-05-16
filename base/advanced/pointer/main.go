package main

import (
	"fmt"
	"log"
)

/**

1.不能对nil使用*
2.当struct没有定义tag时，json之后的key还是原字段类型
3.xss

*/

func main() {
	//_t1()
	//_t2()
	_t3()
}

func _t() {
	// a是一个nil
	var a *int
	log.Println("a:", a)
	// 不能对nil进行赋值
	*a = 100
	fmt.Println("a:", a)
}

func _t2() {
	// a是一个nil
	var a int
	log.Println("a:", a)
	// 不能对nil进行赋值
	a = 100
	fmt.Println("a:", a)
}

func _t3() {
	type Student struct {
		Name string
	}

	var stu *Student
	log.Println("stu:", stu)
	// 可以
	//stu = &Student{Name: "许磊1"}
	// 可以
	*stu = Student{Name: "许磊2"} // invalid memory address or nil pointer dereference,不能对nil用*
	log.Println(stu)
}
