package main

import (
	"errors"
	"fmt"
	"log"
	"runtime/debug"
)

func main() {
	//T1()
	T2()
}

// MyError 只要实现了Error方法就是一个error类型
type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return e.msg
}

func NewMyError(msg string) error {
	return &MyError{msg: msg}
}

func T1() {
	// 1.errors.New可以直接返回一个error类型的值
	msg := "heelo"
	// 2.因为是不同的结构体，所以肯定不相等
	log.Println(errors.New(msg) == errors.New(msg))

	//3.断言error
	myError := NewMyError("world")
	//myError1 := MyError{msg: "world"}
	switch myError.(type) {
	case *MyError:
		log.Println("MyError类型")
	case error:
		log.Println("error类型")
	}
	v, ok := myError.(error)
	log.Printf("值=%s,ok=%t\n", v, ok)
	//
}

func T2() {
	// https://studygolang.com/articles/23346?fr=sidebar
	exampleErr := errors.New("哈哈")
	myError := NewMyError("world")
	// fmt.Errorf 可以包裹一个错误,Unwrap 函数
	w := fmt.Errorf("Wrap了一个错误;%w", myError)
	w1 := fmt.Errorf("Wrap2了一个错误;%w", w)
	log.Println("w1:", w1.Error())
	log.Println("w:", w.Error())
	log.Println("myError:", myError.Error())
	//log.Println("去掉第1层error:", errors.Unwrap(w1))
	log.Println("去掉第2层error:", errors.Unwrap(w))
	// errors.Is()
	//如果err和target是同一个，那么返回true
	//如果err 是一个wrap error,target也包含在这个嵌套error链中的话，那么也返回true。
	//很简单的一个函数，要么咱俩相等，要么err包含target，这两种情况都返回true，其余返回false。
	log.Println("是不是myError类型:", errors.Is(w, myError))
	log.Println("是不是w类型:", errors.Is(w, w))
	log.Println("是不是error类型:", errors.Is(w, exampleErr))
	// errors.As As所做的就是遍历err嵌套链，从里面找到类型符合的error，然后把这个error赋予target,这样我们就可以使用转换后的target了，这里有值得赋予，所以target必须是一个指针
	var newMyErr *MyError
	if errors.As(w, &newMyErr) {
		log.Println("newMyErr:", newMyErr.Error())
	} else {
		log.Println("不符合")
	}

	// debug
	log.Println("debug.Stack:", string(debug.Stack()))

}
