package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	F1()
}

func F1() {
	m := sync.Map{}
	m.Store(1, 1)
	go do(&m)
	go do(&m)

	time.Sleep(100 * time.Nanosecond)
	fmt.Println(m.Load(1))
}

func do(m *sync.Map) {
	i := 0
	for i < 1000000 {
		m.Store(1, i)
		i++
	}
}
