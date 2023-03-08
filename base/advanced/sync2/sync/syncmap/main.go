package main

import (
	"log"
	"sync"
	"time"
)

var mp = sync.Map{}

func main() {
	T1()
}

/**
https://www.cnblogs.com/ricklz/p/13659397.html
1.如果read字段有的话，先覆盖read
2.同时往dirty字段写
3.如果misses累加器数量大于了dirty的长度，那么把dirty赋值到read
4.大致的原理就是read不加锁，dirty加锁，当频繁读取dirty的时候就该把dirty清空了
 */


func T1() {
	go func() {
		mp.Store("name", "许磊")
	}()

	go func() {
		mp.Store("name", "李四")
	}()

	go func() {
		time.Sleep(time.Second * 2)
		v,ok := mp.Load("name")
		log.Println("v:", v)
		log.Println("ok:", ok)
	}()

	go func() {
		time.Sleep(time.Second * 3)
		// 读取或写入,存在就读取，不存在就写入
		v,ok := mp.LoadOrStore("name", "王五")
		log.Println("v:", v)
		log.Println("ok:", ok)
	}()

	time.Sleep(time.Second * 5)
}
