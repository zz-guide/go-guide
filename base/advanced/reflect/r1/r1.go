package main

import (
	"fmt"
	"reflect"
)

/**
reflect 实现了运行时的反射能力，能够让程序操作不同类型的对象。反射包中有两对非常重要的函数和类型，两个函数分别是：

reflect.TypeOf 能获取类型信息
reflect.ValueOf 能获取数据的运行时表示

两个类型是 reflect.Type 和 reflect.Value，它们与上边两个函数是一一对应的关系

三大反射定律：
	1.反射第一定律：反射可以将"接口类型变量"转换为"反射类型对象"。
	2.反射可以将"反射类型对象"转换为"接口类型变量"。
	3.如果要修改“反射类型对象”，其值必须是“可写的”（settable）。

参考网址：https://www.jb51.net/article/90021.htm
https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-reflect/#432-%E7%B1%BB%E5%9E%8B%E5%92%8C%E5%80%BC
*/
func main() {
	F2()
}

func F1() {
	author := "draven"
	t := reflect.TypeOf(author)
	v := reflect.ValueOf(author)
	s := v.Interface().(string)
	fmt.Println("TypeOf author:", t)
	fmt.Println("ValueOf author:", v)
	fmt.Println("s6:", s)
}

func F2() {
	i := 1
	v := reflect.ValueOf(&i)
	//v.SetInt(10) // 不起作用
	v.Elem().SetInt(10)
	fmt.Println(i)
}
