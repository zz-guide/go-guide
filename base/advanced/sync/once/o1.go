package main

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

/**
结论：
	1.init函数是在文件包首次被加载的时候执行，且只执行一次。
	sync.Once是在代码运行中需要的时候执行，且只执行一次。

	2.sync.Once 不能被复制
*/
func main() {
	s1 := GetInstance()
	s2 := GetInstance()
	s3 := GetInstance()
	s4 := s3
	fmt.Printf("s1: %p %+v %+v \n", &s1, s1, unsafe.Pointer(s1))
	fmt.Printf("s2: %p %+v %+v \n", &s2, s2, unsafe.Pointer(s2))
	fmt.Printf("s3: %p %+v %+v \n", &s3, s3, unsafe.Pointer(s3))
	fmt.Printf("s4: %p %+v %+v \n", &s4, s4, unsafe.Pointer(s4))
	fmt.Println("s1 == s2", reflect.DeepEqual(s1, s2))
	fmt.Println("s2 == s3", reflect.DeepEqual(s2, s3))
	fmt.Println("s3 == s4", reflect.DeepEqual(s3, s4))
}

type singleton struct {
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}
