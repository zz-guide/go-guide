package main

import (
	"log"
)

/*
 go的抢占式调度
*/
func main() {
	go func(n int) {
		for {
			n++
			log.Println("n:", n)
		}
	}(0)

	for {

	}
}
