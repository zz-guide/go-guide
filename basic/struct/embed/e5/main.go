package main

import "log"

type S struct {
	T
}

type SP struct {
	*T
}

type T struct{}

func (t T) GetNameValue() {
	log.Println("T值接收者方法")
}

func (t *T) GetNamePointer() {
	log.Println("T指针接收者方法")
}

func main() {
	T1()
	//T2()
}

func T1() {
	// 嵌入一个值类型

	// S值类型
	//s := S{T{}}
	//s.GetNameValue()
	//s.GetNamePointer()

	// S指针类型
	sP := &S{T{}}
	sP.GetNameValue()
	sP.GetNamePointer()
}

func T2() {
	// 嵌入一个指针类型

	// SP值类型
	s := SP{&T{}}
	s.GetNameValue()
	s.GetNamePointer()

	// SP指针类型
	sP := &SP{&T{}}
	sP.GetNameValue()
	sP.GetNamePointer()
}
