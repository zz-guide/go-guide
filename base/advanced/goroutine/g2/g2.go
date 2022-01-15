package main

import "fmt"

/*
 go的抢占式调度
*/
func main() {
	go func(n int) {
		for {
			n++
			fmt.Println("n:", n)
		}
	}(0)

	for {

	}
}
