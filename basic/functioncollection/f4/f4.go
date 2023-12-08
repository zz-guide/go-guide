package main

import (
	"fmt"
	"log"
)

type CoderInterface interface {
	code()
	debug()
}

type Gopher struct {
	language string
	name     string
}

func (p Gopher) code() {
	p.name = "asdasd"
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p *Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)
}

// 如果实现了接收者是值类型的接口方法，会隐含地也实现了接收者是指针类型的方法。
// 编译器会自动为值类型的方法添加一份到指针类型方法里，主要是为了interface调用
func main() {
	c := &Gopher{"Go", "许磊"}
	c.code()
	c.debug()
	log.Printf("c=%+v\n", c)
	//cannot use Gopher{...} (type Gopher) as type coder in assignment:
	//Gopher does not implement coder (debug method has pointer receiver)
	//var c1 CoderInterface = Gopher{"PHP", "许磊"}
	//c1.debug()
	//c1.code()
}
