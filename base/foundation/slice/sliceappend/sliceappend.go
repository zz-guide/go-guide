package main

import "fmt"

func main() {
	//stack()
	//appendd()
	//F2()
	F3()
}

func appendd() {
	arr := []int{1, 2, 3}
	arr = append([]int{5}, arr...)
	fmt.Println("arr--头部添加:", arr)

	arr = append(arr, 6)
	fmt.Println("arr--向后添加:", arr)

	arr = append(arr[:1], append([]int{1212}, arr[1:]...)...)
	fmt.Println("arr--指定位置添加:", arr)
}

func stack() {
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//head := arr[0]
	//arr = arr[1:]
	//fmt.Println("头部出栈head:", head, arr)

	//arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//tail := arr[len(arr)-1]
	//arr = arr[:len(arr)-1]
	//fmt.Println("尾部出栈tail:", tail, arr)

	// 合并到一起的写法,for循环里不要这么写，可能死循环
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//arr, head := arr[1:], arr[0]
	//fmt.Println("头部出栈head:", head, arr)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr, tail := arr[:len(arr)-1], arr[len(arr)-1]
	fmt.Println("尾部出栈tail:", tail, arr)
}

func F1() {
	s := []int{1, 2}
	fmt.Println("s[0:]:", s[0:])
	fmt.Println("s[1:]:", s[1:])
	fmt.Println("s[2:]:", s[2:])
	fmt.Println("s[3:]:", s[3:]) // panic: runtime error: slice bounds out of range [3:2]
	// 最多能取到s[len(s):]
}

func F2() {
	arr := []int{1, 2}
	s := append(arr, 3)
	fmt.Println(arr, s)
}

func F3() {
	arr := []int{1, 2, 3, 4}
	index := 2
	arr = append(arr[:index], arr[index+1:]...)
	fmt.Println("删除指定index位置元素:", arr)
}
