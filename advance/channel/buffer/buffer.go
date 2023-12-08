package main

import (
	"log"
)

var done chan struct{}

func init() {
	done = make(chan struct{}, 0)
}

func main() {
	ch := make(chan int, 5)
	log.Printf("ch:%+v\n", ch)
	<-done
}
