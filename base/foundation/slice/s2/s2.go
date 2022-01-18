package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	ex := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// ex[:], ex[0:], ex[:0] 等同于浅拷贝自身
	newEx := ex[:]
	fmt.Printf("ex: %p %#v \n", &ex, ex)
	fmt.Printf("newEx: %p %#v \n", &newEx, newEx)
	fmt.Println("-------------------------")

	// 会同步修改原来的slice
	newEx[0] = 20
	fmt.Printf("ex: %p %#v \n", &ex, ex)
	fmt.Printf("newEx: %p %#v \n", &newEx, newEx)
	fmt.Println("-------------------------")

	// :左右必须是none negative即非负数，两边的数字表示0-capacity,且左边的必须小于右边的
	// 从index=left开始数，左开右必，含左不含右
	// [:2]等同于[0:2]
	newEx2 := ex[0:1]
	fmt.Printf("ex: %p %#v \n", &ex, ex)
	fmt.Printf("newEx2: %p %#v \n", &newEx2, newEx2)
	fmt.Println("-------------------------")

	// ex[0:0]，ex[1:1]，ex[2:2] 都表示重置切片，清空拥有的元素，数字不能大于capacity
	//newEx1 := ex[0:0]
	//fmt.Printf("ex: %p %#v \n", &ex, ex)
	//fmt.Printf("newEx1: %p %#v \n", &newEx1, newEx1)
	fmt.Println("-------------------------")

	// nil slice,nil 切片不能直接赋值使用，需要搭配append方法进行初始化，append底层判断容量不够就会进行初始化
	// 2个nil slice不能直接比较
	var a []int
	var b []int

	// empty slice, 所有空切片结构体Data指向的地址是同一个
	c := make([]int, 0)
	var d = make([]int, 0)
	var e = []int{}
	f := []int{}

	fmt.Printf("a: %p %+v \n", &a, *(*reflect.SliceHeader)(unsafe.Pointer(&a)))
	fmt.Printf("b: %p %+v \n", &b, *(*reflect.SliceHeader)(unsafe.Pointer(&b)))
	fmt.Printf("c: %p %#v \n", &c, *(*reflect.SliceHeader)(unsafe.Pointer(&c)))
	fmt.Printf("d: %p %#v \n", &d, *(*reflect.SliceHeader)(unsafe.Pointer(&d)))
	fmt.Printf("e: %p %#v \n", &e, *(*reflect.SliceHeader)(unsafe.Pointer(&e)))
	fmt.Printf("f: %p %#v \n", &f, *(*reflect.SliceHeader)(unsafe.Pointer(&f)))
	fmt.Println("a nil slice", a == nil)
	fmt.Println("b nil slice", b == nil)
	fmt.Println("f empty slice", f == nil)

	// 检查slice是否为空，通过len()检测
	//fmt.Println("slice is empty:", len(a), cap(a), len(f), cap(f))

	// copy 深拷贝
	t1 := []int{1, 2, 3}
	t2 := make([]int, 3)
	copy(t2, t1)
	fmt.Printf("t1: %p %#v \n", &t1, t1)
	fmt.Printf("t2: %p %#v \n", &t2, t2)
	fmt.Println("-------------------------")

	t2[2] = 45
	fmt.Printf("t1: %p %#v \n", &t1, t1)
	fmt.Printf("t2: %p %#v \n", &t2, t2)
}

func Pop() {
	a := []int{1, 2, 3, 4, 5}
	c := a[0:len(a)]
	// 0-长度减1可以把最后一个元素弹出
	b := a[0 : len(a)-1]

	fmt.Println("a:", a)
	fmt.Println("c:", c)
	fmt.Println("b:", b)
}
