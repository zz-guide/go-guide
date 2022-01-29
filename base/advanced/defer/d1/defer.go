package main

import "fmt"

func main() {
	//deferCall()
	//log.Println("值:", f7())
}

func deferCall() {
	// 结论：先进后出
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
}

func f1() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func f2() {
	for i := 0; i < 5; i++ {
		/**
		结论：i在defer执行的时候已经是4了
		*/
		defer func() {
			fmt.Println("i的值", i)
		}()
	}
}

func f3() {
	for i := 0; i < 5; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
}

func f4() int {
	t := 5
	defer func() {
		fmt.Println("t:", t)
		t++
	}()
	return t
}

func f5() (r int) {
	/**
	结论：r1: 0
		r3: 0
		r4: -1
		r2: 3
		值: 3

	*/
	fmt.Println("r1:", r)
	defer func() {
		fmt.Println("r4:", r)
		r = 3
		fmt.Println("r2:", r)
	}()
	fmt.Println("r3:", r)
	return -1
}

func f6() (r int) {
	/**
	结论：r2: 0
		r1: 5
		值: 5

		一旦return，r的值就是5了，因为修改的是t，r不变还是5

	*/
	t := 5
	defer func() {
		fmt.Println("r1:", r)
		t = t + 5
	}()
	fmt.Println("r2:", r)
	return t
}

func f7() (r int) {
	/**
	此处传值，不会影响返回值，如果是指针则会影响
	*/
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
