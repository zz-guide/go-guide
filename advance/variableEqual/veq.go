package main

import (
	"fmt"
	"unsafe"
)

type I interface {
	GetName() string
}

type Parent struct {
	Id   int
	Name string
}

type Parent1 struct {
	Id   int
	Name string
}

func (p *Parent) GetName() string {
	p.Name = "李四"
	return p.Name
}

func (p Parent) GetId() int {
	return p.Id
}

func main() {
	//var p I = &Parent{Name: "许磊"}
	//var _ I = new(Parent1)
	//var _ I = (*Parent1)(nil)

	//fmt.Println("实现了接口：", 88)
	//fmt.Println("first:", p.(I))

	p := new([]int)
	fmt.Println("p:", *p)
	fmt.Printf("p的值：%+v\n", p)
	fmt.Printf("p的地址：%p\n", &p)
	fmt.Printf("p的地址：%p\n", p)
	fmt.Printf("p的地址：%d\n", p)
	fmt.Printf("p的大小：%d\n", unsafe.Sizeof(p))

	//v := make([]int, 0)
	//fmt.Printf("v的值：%+v\n", v)
	//fmt.Printf("v的地址：%p\n", v)
	//fmt.Printf("v的地址：%p\n", &v)
	//fmt.Printf("v的大小：%d\n", unsafe.Sizeof(v))
	//fmt.Println("v相等nil:", len(v), cap(v), v == nil)

	//(*p)[0] = 18
	//v[1] = 18

	//p1 := &Parent{Name: "许磊"}
	//fmt.Printf("p存的地址：%p\n", p)
	//fmt.Printf("p自己的地址:%p\n", &p)
	//fmt.Printf("p1自己的地址:%p\n", &p1)
	//
	//c := p.GetName()
	//fmt.Printf("返回值的地址:%p\n", &c)
	//fmt.Println("是否相等:", p == p1)
	//
	//
	//fmt.Println("相等：", reflect.DeepEqual(p, p1))
	//
	//var a int = 1
	//var b int = 1
	//
	//fmt.Printf("a的地址：%p\n", &a)
	//fmt.Printf("b的地址：%p\n", &b)

	//var a int = 11
	//b := &a
	//
	//fmt.Printf("a的地址 = %p\n", &a)
	//fmt.Printf("b的值 = %#x\n", b)
	//
	//fmt.Printf("b的地址 = %p\n", &b)
	//fmt.Printf("b的地址1 = %p\n", b)
	//fmt.Printf("a的值 = %d\n", a)

	// 结论：简单类型的变量值相等， ==认为就相等，reflect.DeepEqual也认为相等
	// 复杂变量首先是类型相等，其次是存的值相等则相等（即引用相等），reflect.DeepEqual只要深度值相等即相等
	// 若p不是一个指针变量，则*p的写法会报错，编译失败
	// 若p不是一个指针变量，则%p打印结果为：%!p(int=2)
}
