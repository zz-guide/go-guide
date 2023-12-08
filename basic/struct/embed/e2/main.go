package main

import (
	"log"
)

/**
测试接口嵌入
1.如果多个接口有相同名字的方法，则参数和返回值相同，是可以的。参数，返回值不同，编译报错
2.var _ IC = (*S)(nil) 这种写法在内嵌接口的时候判断有误。
*/

var _ IC = (*S)(nil)

type S struct {
	//IC
}

func (s *S) Name() string {
	return "S"
}

type IName1 interface {
	Name() string
}

type IName2 interface {
	Name() string
}

type IC interface {
	IName1
	IName2
}

func main() {
	T1()
}

func T1() {
	s := S{}

	var cc IC = &s
	log.Println(cc.Name())

}
