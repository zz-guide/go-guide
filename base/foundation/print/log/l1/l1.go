package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

/**
log的使用
1.Fatal 和 Panic 相关的会输出日志的同时，还会退出和产生panic。
*/
func main() {
	//设置前缀
	log.SetPrefix("[许磊ERROR]")
	//设置要打印的内容：日期，时间，长文件名
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	fileName := "/a/debug.log"
	//打开文件，并且设置了文件打开的模式
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(logFile.Name())
	//设置输出方式为：文件
	log.SetOutput(io.MultiWriter(logFile))
	//输出
	log.Println(123)
	log.Printf("%v, %T", []int{1, 2, 3, 4}, []int{1, 2, 3, 4})
}
