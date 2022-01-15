package main

import (
	"errors"
	"fmt"
)

func main() {
	As()
}

// Is Go1.13版本为fmt.Errorf函数新加了一个%w占位符用来生成一个可以包裹Error的Wrapping Error。
// Is用来判断error链是否满足某一个error类型，严格判断
func Is() {
	err1 := errors.New("new error")
	err2 := fmt.Errorf("err2: [%w]", err1)
	err3 := fmt.Errorf("err3: [%w]", err2)

	fmt.Println(errors.Is(err3, err2))
	fmt.Println(errors.Is(err3, err1))
}

type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
}

// As 则是判断类型是否相同，并提取第一个符合目标类型的错误，用来统一处理某一类错误。
func As() {
	var targetErr *ErrorString
	err := fmt.Errorf("new error:[%w]", &ErrorString{s: "target err"})
	fmt.Println(errors.As(err, &targetErr))
}
