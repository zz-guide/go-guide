package main

import (
	"encoding/json"
	"log"
	"sync"
)

/**
保存和复用临时对象，减少内存分配，降低 GC 压力。
 */

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})


func main() {
	/*stu := &Student{}
	err := json.Unmarshal(buf, stu)
	if err != nil {
		return
	}
	log.Printf("stu:%#v",stu)*/
	T1()
}

/**
sync.Pool 用于存储那些被分配了但是没有被使用，而未来可能会使用的值。这样就可以不用再次经过内存分配，可直接复用已有对象，减轻 GC 的压力，从而提升系统的性能。
sync.Pool 的大小是可伸缩的，高负载时会动态扩容，存放在池中的对象如果不活跃了会被自动清理。
 */

func T1(){
	var studentPool = sync.Pool{
		New: func() interface{} {
			return new(Student)
		},
	}

	stu := studentPool.Get().(*Student)
	log.Printf("stu=%p\n", stu)
	json.Unmarshal(buf, stu)
	studentPool.Put(stu)
	stu1 := studentPool.Get().(*Student)
	log.Printf("stu1=%p\n", stu1)
}