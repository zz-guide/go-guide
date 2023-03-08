package main

import (
	"log"
	"reflect"
)

// Sayer 接口
type Sayer interface {
	say()
	toString(a string) string
}

// Mover 接口
type Mover interface {
	move()
	toString(string) string
}

// 接口嵌套
type Animal interface {
	Sayer
	Mover
	toString(string) string
}

type Cat struct {
	name string
}

func (c *Cat) getName() string {
	return c.name
}

func (c Cat) say() {
	log.Println("喵喵喵")
}

func (c Cat) move() {
	log.Println("猫会动")
}

func (c Cat) toString(s string) string {
	log.Println(s + ",toString")
	return "toString"
}

// 接口嵌套
func main() {
	T1()
	T2()
}

func T1() {
	// 结论：嵌入接口，允许有相同签名的方法（方法名，参数列表，返回值），也可以重新覆盖，最终生效的只有一个
	// 不能嵌入自身，不能循环嵌入
	var x Animal
	x = Cat{name: "花花"}
	x.move()
	x.say()
}

func T2() {
	// 结论：接口类型的变量只能用var定义
	dog := &Cat{}
	var animal Animal = dog
	animal.say()

	// 结论：接口由类型和值组成，只有两个字段均为nil才是nil
	var animal1 Animal
	log.Println("类型和值都为nil才等于nil:", animal1 == nil || reflect.ValueOf(animal1).IsNil())
}
