package main

import (
	"log"
	"reflect"
	"unsafe"
)

func main() {
	//TTT()
	//T1()
	TCheck()
}

func TTT() {
	/**
	i超出范围会报越界错误 runtime error: slice bounds out of range
	1.nums[:i]		---// 冒号在左边，从左往右保留i个
	2.nums[i:]		---// 冒号在右边，从左往右删除i个元素
	3.nums[:], nums[0:]	等价于复制整个数组元素
	*/

	nums := []int{0, 1, 2, 3, 4, 5}
	// i表示要删除索引的位置
	i := 2
	log.Println("nums[:i]:", nums[:i])   //删除索引的前半部分
	log.Println("nums[:i]:", nums[i+1:]) //删除索引的后半部分
	log.Println("nums[i:]:", nums[:], nums[0:])
	// slice删除i位置元素
}

func T1() {
	// 变量a与底层的数组不是一个起始地址
	a := make([]int, 5)
	log.Printf("&a:%p, %d\n", &a, uintptr(unsafe.Pointer(&a)))
	log.Printf("a:%p\n", a)
	log.Printf("a: %p %+v \n", &a, *(*reflect.SliceHeader)(unsafe.Pointer(&a)))
}

func TCheck() {
	// 判断变量是数组还是切片
	arr := [2]int{1, 2}
	t := reflect.TypeOf(arr)
	log.Println("arr类型：", t.Kind() == reflect.Array)

	slice := []int{1, 2}
	t1 := reflect.TypeOf(slice)
	log.Println("slice类型：", t1.Kind() == reflect.Slice)
}
