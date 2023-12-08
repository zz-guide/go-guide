package main

import (
	"errors"
	"fmt"
)

type EasyError struct {
	Msg  string // 错误文字信息
	Code int64  // 错误码
	Nest error  // 嵌套的错误
}

func (me *EasyError) Unwrap() error {
	return me.Nest
}

func (me *EasyError) Error() string {
	return me.Msg
}

func DoSomething1() error {
	// ...
	err := DoSomething2()
	if err != nil {
		return &EasyError{"from DoSomething1", 1, err}
	}

	return nil
}

func DoSomething2() error {
	err := DoSomething3()
	if err != nil {
		return &EasyError{"from DoSomething2", 2, err}
	}

	return nil
}

func DoSomething3() error {
	return &EasyError{"from DoSomething3", 3, nil}
}

/**
Unwrap()可以透传error
*/
func main() {
	err := DoSomething1()
	for err != nil {
		e := err.(*EasyError)
		fmt.Printf("code %d, msg %s6\n", e.Code, e.Msg)
		err = errors.Unwrap(err) // errors.Unwrap中调用EasyError的Unwrap返回子error
	}
}
