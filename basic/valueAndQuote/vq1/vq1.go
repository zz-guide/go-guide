package main

import "log"

/**
go语言中只有值传递

1.如果是基本类型string,数组，int等，传递的是参数的副本。
2.如果是slice,channel,map等引用类型，传递的指针的副本，形参和实参指向同一个地址，但形参和实参是2个变量。
对于引用类型，修改形参，外部实参也会被修改。
3.c语言中的引用传递，形参等同于实参自身，修改形参等于修改实参。
*/
func main() {
	TQuote()
}

func TQuote() {
	f := func(arr [3]int) {
		arr[0] = 123
		log.Printf("func-arr=%+v,%p\n", arr, &arr)
	}

	arr := [3]int{1, 2, 3}
	log.Printf("arr1=%+v,%p\n", arr, &arr)
	f(arr)
	log.Printf("arr2=%+v,%p\n", arr, &arr)
}
