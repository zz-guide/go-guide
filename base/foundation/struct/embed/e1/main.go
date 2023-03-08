package main

import "log"

/**
测试结构体嵌入字段
1.内嵌结构体，初始化的时候不是nil,而是一个empty struct
内嵌普通匿名字段，使用对应类型进行访问。不能同时嵌入多个同类型的匿名字段
2.内嵌指针，初始化的时候是nil,访问的时候也是使用类型名去访问
3.内嵌interface相当于需要实现抽象方法
4.对于结构体自身，匿名字段无法控制对外暴露，而具名字段则通过名字大小写来暴露
5.当继承了多个结构体，冲突的方法无法直接访问，需要明确指出是访问谁的方法。包括字段也是一样的。
*/

type S struct {
	Q
	T
}

func (s S) GetName() string {
	return "s"
}

type T struct {
	name string
}

func (t T) GetName1() string {
	return t.name
}

type Q struct {
	name string
}

func (q Q) GetName() string {
	return q.name
}

type S1 struct {
	*Q
	*T
}

func main() {
	T1()
}

func T1() {
	s := S{Q{name: "xxx"}, T{name: "ccc"}}
	log.Printf("s:%+v\n", s)
	log.Println(s.T)
	log.Println(s.GetName())
	log.Println(s.T.GetName1())
	log.Println(s.Q.GetName())

	//s1 := S1{&Q{name: "vvv"}, &T{name: "bbb"}}
	//log.Printf("s1:%+v\n", s1)
	//log.Println(s1.GetName())
	//log.Println(s1.T)
	//log.Println(s1.T.GetName())
	//log.Println(s1.Q.GetName())
}
