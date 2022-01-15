package main

import (
	"fmt"
	"time"
)

const MAXSLEEP = 128

/**
演示指数退避算法
*/
func main() {
	for numsec := 1; numsec <= MAXSLEEP; numsec <<= 1 {
		// TODO

		if numsec <= MAXSLEEP/2 {
			time.Sleep(time.Second * time.Duration(numsec))
			fmt.Println("slepp time(s):", numsec)
		}
	}
}
