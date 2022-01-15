package main

import (
	"fmt"
	"unsafe"
)

func main() {

}

func TestMap() {
	//1.map的创建
	var map1 map[string]string //默认map是nil,是一个指针，指向nil，需要初始化
	var map2 map[string]string = map[string]string{}
	var map3 map[string]string = make(map[string]string, 10)
	// panic: assignment to entry in nil map
	//map1["name"] = "zhangfei"		//可读不可写
	map2["name"] = "lisi"
	map3["name"] = "xulei"

	for key, value := range map1 {
		fmt.Println("Key:", key, "Value:", value)
	}
	for key, value := range map2 {
		fmt.Println("Key:", key, "Value:", value)
	}
	for key, value := range map3 {
		fmt.Println("Key:", key, "Value:", value)
	}

	s1 := map1["1"]
	s2 := map2["2"]
	s3 := map3["3"]

	fmt.Printf("val=%s,%s,%s\n", s1, s2, s3)
	fmt.Printf("len=%d,%d,%d\n", len(map1), len(map2), len(map3))
	fmt.Printf("size=%d,%d,%d\n", unsafe.Sizeof(map1), unsafe.Sizeof(map2), unsafe.Sizeof(map3))

	fmt.Println(unsafe.Pointer(&map1), 8)
	fmt.Println(unsafe.Pointer(&map2), 8)
	fmt.Println(unsafe.Pointer(&map3), 8)

}
