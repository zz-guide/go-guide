package main

import (
	"fmt"
	"unsafe"
)

/**
nil是一个预先声明的标识符，
代表指针(pointer)、通道(channel)、函数(func)、接口(interface)、map、切片(slice)。
也可以这么理解：指针、通道、函数、接口、map、切片的零值就是nil，就像布尔类型的零值是false、整型的零值是0。
nil标识符是没有类型的，所以==对于nil来说是一种未定义的操作，不可以进行比较
*/
func main() {
	T8()
}

func T1() {
	// 结论: nil不能直接和nil比较
	//fmt.Println(nil == nil) //invalid operation: nil == nil (operator == not defined on nil)
}

func T2() {
	const val1 = iota
	// int
	fmt.Printf("%T\n", val1)
	// 结论：use of untyped nil
	// nil是没有默认类型的，它的类型具有不确定性，我们在使用它时必须要提供足够的信息能够让编译器推断nil期望的类型。
	//var val2 = nil
	//fmt.Printf("%T\n", val2)
}

func T3() {
	// 结论：指针类型nil、channel类型的nil、interface类型可以相互比较，而func类型、map类型、slice类型只能与nil标识符比较，两个类型相互比较是不合法的。

	// 同一类型的nil值比较
	// 指针类型的nil比较
	//fmt.Println((*int64)(nil) == (*int64)(nil))
	// channel 类型的nil比较
	//fmt.Println((chan int)(nil) == (chan int)(nil))
	// func类型的nil比较
	//fmt.Println((func())(nil) == (func())(nil)) // func() 只能与nil进行比较，(func())(nil) == (func())(nil) (func can only be compared to nil)
	// interface类型的nil比较
	//fmt.Println((interface{})(nil) == (interface{})(nil))
	// map类型的nil比较
	//fmt.Println((map[string]int)(nil) == (map[string]int)(nil)) // map 只能与nil进行比较，(map[string]int)(nil) == (map[string]int)(nil) (map can only be compared to nil)
	// slice类型的nil比较
	//fmt.Println(([]int)(nil) == ([]int)(nil)) // slice 只能与nil进行比较，([]int)(nil) == ([]int)(nil) (slice can only be compared to nil)
}

func T4() {
	// 结论：只有指针类型和channel类型与interface类型可以比较，其他类型的之间是不可以相互比较的
	/*var ptr *int64 = nil
	var cha chan int64 = nil
	var fun func() = nil
	var inter interface{} = nil
	var ma map[string]string = nil
	var slice []int64 = nil
	fmt.Println(ptr == cha)
	fmt.Println(ptr == fun)
	fmt.Println(ptr == inter)
	fmt.Println(ptr == ma)
	fmt.Println(ptr == slice)

	fmt.Println(cha == fun)
	fmt.Println(cha == inter)
	fmt.Println(cha == ma)
	fmt.Println(cha == slice)

	fmt.Println(fun == inter)
	fmt.Println(fun == ma)
	fmt.Println(fun == slice)

	fmt.Println(inter == ma)
	fmt.Println(inter == slice)

	fmt.Println(ma == slice)*/
}

type Err interface {
}
type err struct {
	Code int64
	Msg  string
}

func T5() {
	Todo := func() Err {
		var res *err
		// 指针类型的零值是nil
		println("res:", res, res == nil)
		// interface 不是单纯的值，而是分为类型和值。所以必须要类型和值同时都为 nil 的情况下，interface 的 nil 判断才会为 true
		return res
	}
	err := Todo()
	println("err:", err)
	fmt.Println(err == nil)
}

func T6() {
	var m map[string]string
	fmt.Println(m["asoong"])
	// 结论：nil map可读不可写
	m["asong"] = "Golang梦工厂" //panic: assignment to entry in nil map
}

func T7() {
	// 结论：close一个nil chan 直接panic;一个nil的channel读写数据都会造成永远阻塞
	var cha chan int
	close(cha) //panic: close of nil channel
}

func T8() {
	// 结论：一个为nil的索引，不可以进行索引，否则会引发panic，其他操作是可以。
	var slice []int64 = nil
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	for range slice {

	}
	fmt.Println(slice[0]) //panic: runtime error: index out of range [0] with length 0
}

type man struct {
}

func (m *man) GetName() string {
	// return m.Name
	return "asong"
}

func T9() {
	// 结论：方法接收者为nil时，我们仍然可以访问对应的方法，但是要注意方法内的写法，否则也会引发panic
	var m *man
	fmt.Println(m.GetName())
}

func T10() {
	// 空指针就是一个没有指向任何值的指针
	var a = (*int64)(unsafe.Pointer(uintptr(0x0)))
	fmt.Println(a == nil) //true
}
