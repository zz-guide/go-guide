package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/**
slice的源码在runtime.slice包中
结论：
	1.len小于1024扩容2倍
	2.len大于1024扩容1.25倍
	3.当len>cap时发生扩容
	4.append如果超出了原slice大小，则创建新的slice返回；否则返回原slice
*/
func main() {
	//F1()
	append2()
	//F3()
}

func F1() {
	// ptr,len,cap  ptr指向底层数组的开始，len表示长度
	// 结论：通过make创建的切片，初始值是类型的默认值，并且len=cap
	s := make([]int, 5)
	fmt.Println("s6,长度,容量:", s, len(s), cap(s), (*reflect.SliceHeader)(unsafe.Pointer(&s)))

	// 结论：通过make创建的切片，初始值是类型的默认值，并且len为5，cap为10
	// cap不能比len小，否则会编译失败，len larger than cap in make([]int)
	//s1 := make([]int, 3, 4)
	//fmt.Println("s1,长度,容量:", s1, len(s1), cap(s1))

	// 结论：如果append之后，还没有超出原数组的容量，那么，切片中的指针指向的位置，就还是原数组，如果append之后，超过了原数组的容量，那么，Go就会开辟一块新的内存，把原来的值拷贝过来，这种情况丝毫不会影响到原数组。
	a := s[:]
	fmt.Println("a,长度,容量:", a, len(a), cap(a), (*reflect.SliceHeader)(unsafe.Pointer(&a)))
	a = append(a, 1)
	// 结论：超出容量以后创建新的slice，不影响原来的slice
	fmt.Println("----扩容-----")
	fmt.Println("a,长度,容量:", a, len(a), cap(a), (*reflect.SliceHeader)(unsafe.Pointer(&a)))
	fmt.Println("s6,长度,容量:", s, len(s), cap(s), (*reflect.SliceHeader)(unsafe.Pointer(&s)))
}

func append1() {
	// 扩容规则定义再runtime.slice的growslice方法
	s1 := make([]int, 1022, 1022)
	fmt.Println("s1,长度,容量:", len(s1), cap(s1))
	// 结论：长度小于1024扩容为原来cap的2倍
	// 结论：长度大于1024扩容为原来cap的1.25倍
	s1 = append(s1, 1, 2)
	fmt.Println("s1,长度,容量:", len(s1), cap(s1))
}

func append2() {
	// 当len>cap的时候才会扩容
	s1 := make([]int, 1026, 1026)
	fmt.Println("s1,长度,容量:", len(s1), cap(s1))

	// 大于1024扩容1.25倍
	s1 = append(s1, 1)
	fmt.Println("s1,长度,容量:", len(s1), cap(s1))
	// 结论：扩容完之后还需要申请内存，内存对齐，所以不是准确的1.25倍，内存span都是4的倍数内存块
}

func F3() {
	//s6 := make([]struct{}, 0, 2)
	//fmt.Println("s6,长度,容量:", s6, len(s6), cap(s6))
	//a := struct{}{}
	//fmt.Println("a占多少个字节:", unsafe.Sizeof(a))
	//s6 = append(s6, a)
	//fmt.Println("s6,长度,容量:", s6, len(s6), cap(s6))

	var s []int
	s = nil
	fmt.Println("s6,长度,容量:", s, len(s), cap(s), s == nil, unsafe.Pointer(&s))
	s = append(s)
	fmt.Println("s6,长度,容量:", s, len(s), cap(s), s == nil, unsafe.Pointer(&s))
}
