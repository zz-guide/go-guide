package main

import (
	"log"
	"sync/atomic"
	"unsafe"
)

type myCopy struct {
	noCopy noCopy
	checker copyChecker
	Name   string
}

func(my *myCopy) Run() {
	my.checker.check()
	log.Println("myCopy Run")
}

func(my *myCopy) Copy() *myCopy {
	return NewMyCopy()
}

func(my myCopy) Copy1() myCopy {
	return myCopy{}
}

func NewMyCopy() *myCopy {
	return &myCopy{}
}

// 编译时检查 go vet检查
type noCopy struct{}
func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

// 运行时检查
type copyChecker uintptr
func (c *copyChecker) check() {
	if uintptr(*c) != uintptr(unsafe.Pointer(c)) &&
		!atomic.CompareAndSwapUintptr((*uintptr)(c), 0, uintptr(unsafe.Pointer(c))) &&
		uintptr(*c) != uintptr(unsafe.Pointer(c)) {
		panic("myCopy is copied")
	}
}


/**
注意：

1. 该noCopy并没有导出来，我们不可使用

2. 想要实现NoCopy，需要自己定义一个这样的结构体实现其Lock()接口即可，结构体名字随意

3. 即使包含了NoCopy的结构体，也不是真正的就不可复制了，实际上毫无影响，无论是编译，还是运行都毫不影响

4. 只有在go vet命令下，才会提示出其中的nocopy问题

5. goland，vscode编译器中可以直接提示出来

*/

func main() {
	// 因为go语言是传值，如果不使用指针的话，每次都是拷贝
	//t1 := myCopy{}
	//t1.Name = "许磊"
	//log.Printf("t1:%#v \n", t1)

	/// 正确做法：一定要使用指针，然后使用go vet xxx.go来检查
	t3 := NewMyCopy()
	t3.Name = "李四"
	log.Printf("t3:%v \n", t3)

	t4 := t3 // t4指向t3所指向的对象,所以是同一个对象
	t4.Name = "王五"
	log.Printf("t3:%v,t4:%v \n", t3, t4)

	t5 := myCopy{}
	t5.Run()

	// 报错，防止运行时拷贝。
	t6 := t5
	t6.Run()
}
