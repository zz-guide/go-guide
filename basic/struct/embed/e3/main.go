package main

import "log"

/**
测试结构体嵌入字段

*/

type S struct {
	t
}

type t struct {
	name string
}

func (_t t) GetName() string {
	return _t.name
}

func main() {
	T1()
}

func T1() {
	s := S{t{name: "ttt"}}
	log.Printf("s:%+v\n", s)
	log.Println(s.GetName())
	log.Println(s.name)
}
