package main

import (
	"log"
	"reflect"
	"sync"
	"unsafe"
)

type singleton struct {
}

var instance *singleton
var once sync.Once

/*
*
结论：

	1.sync.Once是在代码运行中需要的时候执行，且只执行一次。
	2.sync.Once 不能被复制
*/
func main() {
	s1 := GetInstance()
	s2 := GetInstance()
	s3 := GetInstance()
	s4 := s3
	log.Printf("s1: %p %+v %+v \n", &s1, s1, unsafe.Pointer(s1))
	log.Printf("s2: %p %+v %+v \n", &s2, s2, unsafe.Pointer(s2))
	log.Printf("s3: %p %+v %+v \n", &s3, s3, unsafe.Pointer(s3))
	log.Printf("s4: %p %+v %+v \n", &s4, s4, unsafe.Pointer(s4))
	log.Println("s1 == s2", reflect.DeepEqual(s1, s2))
	log.Println("s2 == s3", reflect.DeepEqual(s2, s3))
	log.Println("s3 == s4", reflect.DeepEqual(s3, s4))
}

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}
