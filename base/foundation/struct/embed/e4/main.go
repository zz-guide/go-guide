package main

import (
	. "go-guide/base/foundation/struct/embed/e4/t"
	"log"
)

/**
1.内嵌接口的目的是什么？
2.因为接口不能定义属性,所以接口变量不能直接访问结构体属性
3.内嵌一个接口，编译器就认为struct实现了这个接口，即便实际上没有实现，运行的时候可能报错
4.内嵌接口相当于拥有一个成员，这个成员可以是实现了这个接口的任何成员。类似于工厂里有很多工人一样，各种各样的工人。
初始化的时候需要传入。如果不满意可以自己实现接口的所有方法。
*/

// 父类，存放公共属性，公共方法
type animal struct {
	name string
	//IAnimal
}

func (_animal animal) FnInterface() string {
	return "父类->FnInterface"
}

type Dog struct {
	feet int
	animal
}

// FnInterface 可以覆盖内嵌结构体的方法，可以认为是"重写"
func (_dog Dog) FnInterface() string {
	return "狗->FnInterface"
}

func (_dog Dog) private() {

}

type Cat struct {
	feet int
	animal
}

func (_cat Cat) FnInterface() string {
	return "猫->FnInterface"
}

func (_cat Cat) private() {

}

func main() {
	//T1()
	T2()
}

func T1() {
	s := Dog{animal: animal{name: "狗"}, feet: 4}
	log.Printf("s:%+v\n", s)
	log.Println(s.FnInterface())
	log.Println(s.name)
}

func T2() {
	var iAnimal1 IAnimal = Dog{animal: animal{name: "狗"}, feet: 4}
	var iAnimal2 IAnimal = Cat{animal: animal{name: "猫"}, feet: 2}
	log.Printf("s:%+v\n", iAnimal1)
	log.Println(iAnimal1.FnInterface())
	log.Println(iAnimal2.FnInterface())

}
