package main

import (
	"log"
	"reflect"
)

func main() {
	//TGlobalVariable()
	//TVariableShortStatement()
}

// 全局变量声明
var a int = 123

// a2 := 33 // 简短式声明不能用在函数外部

func TGlobalVariable() {
	// 结论：函数内部声明的变量，类型，常量，只能在本函数内部使用
	// 结论：未使用变量会引发编译错误
	//var b int = 22

	log.Printf("a: 值=%d,类型=%s\n", a, reflect.TypeOf(a))

	// 结论：简短式声明，只能用在函数内部,并且不能指定类型
	a1 := 123
	log.Printf("a1: 值=%d,类型=%s\n", a1, reflect.TypeOf(a1))
}

func TVariableShortStatement() {
	s := "aaa"
	log.Printf("s1: 值=%s,类型=%s,地址=%p\n", s, reflect.TypeOf(s), &s)

	// 结论：要想重复使用简短式声明同一个变量，左边必须有新的变量产生
	// 结论：如果是用一个作用域内，s重复使用，内存地址不变
	sxx, s := "我是", "bbb"
	log.Printf("s2: 值=%s,类型=%s,地址=%p\n", s, reflect.TypeOf(s), &s)
	log.Printf("sxx: 值=%s,类型=%s,地址=%p\n", sxx, reflect.TypeOf(sxx), &sxx)

	if true {
		// 结论：当作用域变化时，即使是用一个变量重复声明，内存地址会发生变化
		s := "ccc"
		log.Printf("s3: 值=%s,类型=%s,地址=%p\n", s, reflect.TypeOf(s), &s)
	}

	{
		// 不同的作用域，内存地址会发生变化
		s := "ddd"
		log.Printf("s4: 值=%s,类型=%s,地址=%p\n", s, reflect.TypeOf(s), &s)
	}
}

func TVariablesAssign() {
	// 结论：多变量赋值，先计算出全部右值，然后再依次赋值。
	x, y := 1, 2
	x, y = y+3, x+2
	log.Println(x, y) // 5, 3
}

func TConst() {
	// 结论：未使用常量不会引发编译错误。
}
