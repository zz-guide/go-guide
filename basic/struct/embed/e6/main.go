package main

import "log"

type IA interface {
	Name()
}

type T struct {
	IA
}

func (t T) Name() {
	log.Println("T值接收者方法")
}

//func (t *T) Name() {
//	log.Println("T指针接收者方法")
//}

func main() {
	T1()
	//T2()
}

func T1() {
	t := T{}
	t.Name()

	t1 := &T{}
	t1.Name()

	// 结论：实现接口的方法如果接收者是值，那么值和指针类型都可以赋值给接口类型
	// 如果是指针类型，只有指针类型的值才可以赋值给接口类型
	//var i1 IA = T{} //Type does not implement 'IA' as the 'Name' method has a pointer receiver
	//var i1 IA = &T{}
}

func T2() {
	t := T{}
	t.Name()

	t1 := &T{}
	t1.Name()
}
