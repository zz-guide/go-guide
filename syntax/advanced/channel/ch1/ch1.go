package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	readWrite()
}

func write() {
	/*
		写未初始化的 chan，会死锁：all goroutines are asleep - deadlock!chan send (nil chan)
	*/

	var c chan int
	c <- 1
	log.Println("ssssssss")
}

func readWrite() {
	/**
	all goroutines are asleep - deadlock![chan receive (nil chan)]
	*/
	var c chan int
	go func() {
		c <- 1
	}()

	num, ok := <-c
	fmt.Printf("读chan的协程结束, num=%v, ok=%v\n", num, ok)
	time.Sleep(5)
}
