package main

import "log"

type Mover interface {
	move()
}

type Dog struct{}

func (d Dog) move() {
	log.Println("狗会动")
}

func main() {
	T2()
}

func T1() {
	// 如果实现接口的是值类型的
	var x Mover

	var wangCai = Dog{} // 旺财是dog类型
	x = wangCai         // x可以接收dog类型

	// 指针类型变量，编译器会进行解引用，类似语法糖
	var fuGui = &Dog{} // 富贵是*dog类型
	x = fuGui          // x可以接收*dog类型
	x.move()
}

type Mover2 interface {
	move2()
}

func (d *Dog) move2() {
	log.Println("狗会动")
}

func T2() {
	// 实现接口类型是指针类型，只能指针赋值
	var x Mover2
	//var wangCai = Dog{} // 旺财是dog类型
	//x = wangCai         // x不可以接收dog类型
	var fuGui = &Dog{} // 富贵是*dog类型
	x = fuGui          // x可以接收*dog类型
	x.move2()
}
