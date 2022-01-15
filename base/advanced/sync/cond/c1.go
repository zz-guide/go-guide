package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var done = false

func main() {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)

	time.Sleep(time.Second * 20)
}

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		fmt.Println("xxxxxx")
		c.Wait()
	}
	log.Println(name, "starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	fmt.Println(name, "starts writing")
	time.Sleep(time.Second * 3)
	c.L.Lock()
	//done = true
	c.L.Unlock()
	log.Println(name, "wakes all")
	c.Broadcast()
}