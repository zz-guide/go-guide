package main

import (
	"log"
)

func main() {
	remove()
}

func remove() {
	arr := []int{1, 2, 3, 4, 5}
	log.Printf("arr %d,地址=%p", arr, &arr)

	// 结论:根据索引位置删除数组元素，不包含右边元素，含左不含右
	//arr = arr[:len(arr)-1]
	//log.Printf("arr %d,地址=%p", arr, &arr)

	//arr = arr[1:]
	//log.Printf("arr %d,地址=%p", arr, &arr)

	// 结论：[索引:]，索引最大可以是切片长度，不会报错，返回空切片
	// [:0] 等价于 [len(arr):] 等价于[]
	arr3 := arr[len(arr):]
	log.Printf("arr3 %d,地址=%p", arr3, &arr3)

	//deleteIndex := 4
	//arr = append(arr[:deleteIndex], arr[(deleteIndex+1):]...)
	//log.Printf("arr %d,地址=%p", arr, &arr)

	// 结论：[:]浅拷贝数组，底层指向同一个数组，会影响原数据
	//arr1 := arr[:]
	//log.Printf("arr1[:](全复制) %d,地址=%p", arr1, &arr1)
	//arr1[0] = 3
	//log.Printf("arr %d,地址=%p", arr, &arr)

	// 结论：[:0]代表删除数组,并且不是nil
	//arr2 := arr[:0]
	//arr = nil
	//log.Printf("arr %d,地址=%p", arr, &arr)
	//log.Printf("arr[:0](删除数组) %d,地址=%p, %v", arr2, &arr2, arr2 == nil)
	//log.Printf("arr %d,地址=%p", arr, &arr)

	// empty slice的底层数组地址是一样的
	//var arr3 []int
	//var arr4 []int
	//emptySlice1 := make([]int, 0) // 空切片
	//emptySlice2 := make([]int, 0)
	//log.Printf("arr3 %d,地址=%p, %+v, %p", arr3, &arr3, arr3 == nil, (*reflect.SliceHeader)(unsafe.Pointer(&arr3)))
	//log.Printf("arr4 %d,地址=%p, %+v, %p", arr4, &arr4, arr4 == nil, (*reflect.SliceHeader)(unsafe.Pointer(&arr4)))
	//log.Printf("emptySlice1 %d,地址=%p, %+v, %+v", emptySlice1, &emptySlice1, emptySlice1 == nil, *(*reflect.SliceHeader)(unsafe.Pointer(&emptySlice1)))
	//log.Printf("emptySlice2 %d,地址=%p, %+v, %+v", emptySlice2, &emptySlice2, emptySlice2 == nil, *(*reflect.SliceHeader)(unsafe.Pointer(&emptySlice2)))
}
