package main

import (
	"errors"
	"fmt"
)

func main() {
	Print()
}

// 示例代码
func Oops() error {
	return errors.New("我是Error")
}

func Print() {
	err := Oops()
	fmt.Println(err.Error())
	fmt.Println(err)
	fmt.Println(errors.New("hello error") == errors.New("hello error"))
	errhello := errors.New("hello error")
	fmt.Println(errhello == errhello)
}
