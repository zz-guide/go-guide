package main

import (
	"log"
)

func main() {
	log.Println("值:", deferReturn4())
}

func deferCall() {
	// 结论：先进后出
	defer func() { log.Println("打印前") }()
	defer func() { log.Println("打印中") }()
	defer func() { log.Println("打印后") }()
}

func deferFor() {
	// 先进后出，4，3，2，1，0
	/*for i := 0; i < 5; i++ {
		defer log.Println(i)
	}*/

	/*for i := 0; i < 5; i++ {
		// 打印的时候i已经是最终值了
		defer func() {
			log.Println(i)
		}()
	}*/

	// 与第一种方式等价
	for i := 0; i < 5; i++ {
		// 打印的时候i已经是最终值了
		defer func(a int) {
			log.Println(a)
		}(i)
	}
}

func deferReturn() int {
	// 当函数签名没有声明具体返回值变量的时候，defer无法修改return之后的变量
	t := 5
	defer func() {
		log.Println("t:", t)
		t++
		log.Println("t++:", t)
	}()
	return t
}

func deferReturn1() (r int) {
	log.Println("r1:", r) // 0,此时是零值
	defer func() {
		log.Println("r4:", r) // -1,此时是首次return
		r = 3
		log.Println("r2:", r) // 3,此时是修改return
	}()
	log.Println("r3:", r) // 0,此时是零值
	return -1             // -1,此时是首次return
}

func deferReturn2() (r int) {
	t := 5
	defer func() {
		log.Println("r1:", r) // 5,return的是5
		t = t + 5
		r++                   // 修改的是r的话还会修改，t不影响
		log.Println("r3:", r) // 10，+5之后变成10
	}()
	log.Println("r2:", r) // 0,零值
	return t
}

func deferReturn3() (r int) {
	/**
	此处传值，不会影响返回值，如果是指针则会影响
	*/
	log.Println("r1:", r) // 0, 零值
	// 局部变量同名,局部变量为准，无法影响到外围的r
	defer func(r int) {
		log.Println("r3:", r) // 0
		r = r + 5
		log.Println("r4:", r) // 5
	}(r)
	log.Println("r2:", r) // 0, 零值
	return 1
}

func deferReturn4() (r int) {
	log.Println("r1:", r) // 0, 零值
	// 局部变量不同名，会影响外部变量，无法影响到外围的r
	defer func(t int) {
		log.Println("r3:", r) // 1
		r = r + 5
		log.Println("r4:", r) // 5
	}(r)
	log.Println("r2:", r) // 0, 零值
	return 1
}
