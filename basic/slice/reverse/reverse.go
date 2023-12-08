package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%v\n", x)
	fmt.Printf("结果：%v\n", ReverseSlice2(x))
	//fmt.Printf("结果：%v\n", ReverseSlice1(x))
}

func ReverseSlice1(x []int) []int {
	if len(x) <= 1 {
		return x
	}

	// i < j的时候互换，=的时候指向同一个位置
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}

	return x
}

func ReverseSlice2(x []int) []int {
	if len(x) <= 1 {
		return x
	}

	for i := 0; i < len(x)/2; i++ {
		x[i], x[len(x)-i-1] = x[len(x)-i-1], x[i]
	}

	return x
}
