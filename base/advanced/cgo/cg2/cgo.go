package Cgo

/* #include <stdlib.h>
 */
import "C"
import "fmt"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i)) //强转
}

func Ma() {
	Seed(100) //设置随机数种子
	fmt.Println("Random", Random())
}
