package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/**
能说说uintptr和unsafe.Pointer的区别吗？
答案：
1.unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；
2.而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象， uintptr 类型的目标会被回收；
3.unsafe.Pointer 可以和 普通指针 进行相互转换；
4.unsafe.Pointer 可以和 uintptr 进行相互转换。


Golang指针
*类型:普通指针类型，用于传递对象地址，不能进行指针运算。
unsafe.Pointer:通用指针类型，用于转换不同类型的指针，不能进行指针运算，不能读取内存存储的值（必须转换到某一类型的普通指针）。
uintptr:用于指针运算，GC 不把 uintptr 当指针，uintptr 无法持有对象。uintptr 类型的目标会被回收。
unsafe.Pointer 是桥梁，可以让任意类型的指针实现相互转换，也可以将任意类型的指针转换为 uintptr 进行指针运算。
unsafe.Pointer 不能参与指针运算，比如你要在某个指针地址上加上一个偏移量，Pointer是不能做这个运算的，那么谁可以呢?

就是uintptr类型了，只要将Pointer类型转换成uintptr类型，做完加减法后，转换成Pointer，通过*操作，取值，修改值，随意。

 总结：unsafe.Pointer 可以让你的变量在不同的普通指针类型转来转去，也就是表示为任意可寻址的指针类型。而 uintptr 常用于与 unsafe.Pointer 打配合，用于做指针运算。


unsafe.Pointer
unsafe.Pointer称为通用指针，官方文档对该类型有四个重要描述：
（1）任何类型的指针都可以被转化为Pointer
（2）Pointer可以被转化为任何类型的指针
（3）uintptr可以被转化为Pointer
（4）Pointer可以被转化为uintptr
unsafe.Pointer是特别定义的一种指针类型（译注：类似C语言中的void类型的指针），在golang中是用于各种指针相互转换的桥梁，它可以包含任意类型变量的地址。
当然，我们不可以直接通过*p来获取unsafe.Pointer指针指向的真实变量的值，因为我们并不知道变量的具体类型。
和普通指针一样，unsafe.Pointer指针也是可以比较的，并且支持和nil常量比较判断是否为空指针。


链接：https://www.cnblogs.com/-wenli/p/12682477.html


*/

func main() {
	//F1()
	F2()
}

func F2() {
	var w *W = new(W)
	//这时w的变量打印出来都是默认值0，0
	fmt.Println(w.b, w.c)

	//现在我们通过指针运算给b变量赋值为10
	b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))
	fmt.Println("b:", b, reflect.TypeOf(b))
	*((*int)(b)) = 10
	//此时结果就变成了10，0
	fmt.Println(w.b, w.c)
}

type W struct {
	b int32
	c int64
}

type Student struct {
	Name string
}

func F1() {
	s1 := &Student{Name: "许磊"}
	s3 := &Student{Name: "许磊"}
	s2 := s1

	fmt.Println("s1 == s2:", s1 == s2) // true
	fmt.Println("s1 == s3:", s1 == s3) // false
	fmt.Println("s2 == s3:", s2 == s3) // false
}
