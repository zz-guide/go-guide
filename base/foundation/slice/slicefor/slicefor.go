package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}

	// 结论：for循环append的话，arr会死循环，因为len每次读取的都是最新的长度
	// 如果提前定义length放到外边，就不会了

	length := len(arr)
	for i := 0; i < length; i++ {
		fmt.Println("长度:", len(arr))
		arr = append(arr, 4)
	}

	fmt.Println("arr:", arr, len(arr))
}
