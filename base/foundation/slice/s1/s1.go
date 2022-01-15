package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	t1()
}

func t2() {
	/**
	nil 切片不能直接赋值使用，需要搭配append方法进行初始化，append底层判断容量不够就会进行初始化
	*/
	var s1 []int
	//kmp[0] = 2

	s1 = append(s1, 2)
	fmt.Println("kmp:", s1)
}

func t1() {
	/**
	  nil切片和空切片的区别？
	*/
	var s1 []int // nil 切片

	/**
	所有空切片结构体Data指向的地址是同一个
	*/
	s2 := make([]int, 0) // 空切片
	s4 := make([]int, 0)

	fmt.Printf("kmp pointer:%+v, s2 pointer:%+v, s4 pointer:%+v, \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)), *(*reflect.SliceHeader)(unsafe.Pointer(&s2)), *(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s4))).Data)
}
