package main

import (
	"log"
	"reflect"
)

/**
结论：
	1.不管是指针接受者还是值接收者，内部都发生了拷贝，形参不等于实参
	2.指针和值接收者都可以相互调用各自定义的方法，本质上是语法糖，语法糖只在编译阶段生效，换句话说
	3.
*/
func main() {
	//T1()
	//T2()
	T3()
}

type S struct {
	name string
}

func (s *S) FnPointer() {
	log.Printf("指针：%p\n", s)
}

func (s S) FnValue() {
	log.Printf("值：%p\n", &s)
}

func FnValueM(s S) {
	log.Printf("值：%p\n", &s)
}

func T1() {
	ss := &S{}
	ss.FnPointer()
	ss.FnPointer()
	ss.FnPointer()
	ss.FnValue()
	ss.FnValue()
	ss.FnValue()
	ss.FnValue()
}

func T2() {
	ss := S{}
	ss.FnPointer()
	ss.FnPointer()
	ss.FnPointer()
	ss.FnValue()
	ss.FnValue()
	ss.FnValue()
	ss.FnValue()
}

func T3() {
	ss := S{}
	// 隐形第一个参数就是变量本身
	// 以下两种方式等价
	S.FnValue(ss)
	ss.FnValue()
	//(S{name: "许磊"}).FnPointer() // 编译期间无法确定地址，所以不能调用，cannot call pointer method FnPointer on S
	// 本质上只与参数和返回值有关
	log.Println("FnValueM(s S) 等价于 func (s S) FnValue()：", reflect.TypeOf(S.FnValue) == reflect.TypeOf(FnValueM))
}
