package main

import (
	"fmt"
	"log"
	"time"
)

var mp = make(map[string]string)

func main() {
	T2()
}

func T1() {
	m := make(map[string]string, 2)
	m["name"] = "许磊"
	fmt.Println(m)

}

func T2() {
	// fatal error: concurrent map writes
	// 不支持并发写,不管是不是用一个key
	go func() {
		mp["name"] = "许磊"
	}()

	go func() {
		//mp["name"] = "李四"
		mp["age"] = "23"
	}()

	time.Sleep(time.Second * 5)
}

func T3() {
	// 并发读支持
	mp["name"] = "许磊"
	go func() {
		log.Println("name1:", mp["name"])
	}()

	go func() {
		log.Println("name2:", mp["name"])
	}()

	time.Sleep(time.Second * 5)
}

func T4() {
	//只要不是并发写就行
	go func() {
		log.Println("name1:", mp["name"])
	}()

	go func() {
		mp["name"] = "许磊"
		log.Println("name2:", mp["name"])
	}()

	time.Sleep(time.Second * 5)
}
