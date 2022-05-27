package main

import (
	"fmt"
	"log"
	"strconv"
)

/**
	1.泛型变量:一般写作T，也可以是其他的标识符
    2.泛型约束:T int | float32 | float64
    3.泛型切片: params []T
    4.泛型map:type M[K string, V any] map[K]V
    5.泛型通道：type Ch[T any] chan T
    6.方法约束：func PrintStr[T Animal](params ...T) {}
    7.类型与方法双重约束：
	8.~表示底层类型约束，可以识别别名的情况
    9.comparable,可比较指的是 可以执行 != == 操作的类型，并没确保这个类型可以执行大小比较（ >,<,<=,>= ）
    10.参考网址：https://segmentfault.com/a/1190000041634906?utm_source=sf-hot-article
    11.https://go.dev/ref/spec
*/

func main() {
	//a := 1.1
	//b := 2.3
	//log.Println(T1(a, b))

	//s := []int{1, 2, 3}
	//log.Println(T2(s))

	//T3()
	//T5()
	T6()
}

func T1[T int | float32 | float64](a T, b T) T {
	if a < b {
		return a
	}

	return b
}

func T2[T int | string](params []T) T {
	var sum T
	for _, elem := range params {
		sum += elem
	}
	return sum
}

func T3() {
	type M[K string, V any] map[K]V
	m := M[string, any]{
		"zx": 123,
		"as": "456",
		"qw": 789,
	}

	for k, v := range m {
		log.Printf("key=%s,val=%#v\n", k, v)
	}

	log.Printf("测试4.1: 类型=%T，val=%+v\n", m, m)
}

func T5() {
	type Ch[T any] chan T
	ch := make(Ch[int], 1)
	ch <- 10

	res := <-ch
	log.Printf("测试5.1: 类型=%T，val=%+v", res, res)
	log.Printf("测试5.2: 类型=%T，val=%+v", ch, ch)
}

// ~表示底层类型

type MyInt int

type Animal interface {
	~int
	ToString() string
}

type Dog int

func (d Dog) ToString() string {
	return "string_" + strconv.Itoa(int(d))
}

type Cat int

func PrintStr[T Animal](params ...T) {
	for _, param := range params {
		fmt.Println(param.ToString())
	}
}

func T6() {
	// 测试6.1 传入实现了方法的类型
	dog := Dog(1)
	PrintStr(dog)

	myint := MyInt(22)
	PrintStr(Dog(myint)) // possibly missing ~ for int in constraint Animal
	// 测试6.2 传入未实现对应方法的类型
	// Cat does not implement Animal (missing ToString method)
	//cat := Cat(100)
	//PrintStr(cat)
}
