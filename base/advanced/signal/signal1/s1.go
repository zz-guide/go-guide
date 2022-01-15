package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("Start Work...")

	//watchSignal()
}

var (
	ch = make(chan os.Signal, 1)
)

func watchSignal() {
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	// SIGHUP  1  终端控制进程结束(终端连接断开)
	// SIGINT  2  用户发送INTR字符(Ctrl+C)触发
	// SIGQUIT 3  用户发送QUIT字符(Ctrl+/)触发
	// SIGKILL 9  无条件结束程序(不能被捕获、阻塞或忽略)
	// SIGTERM 15  结束程序(可以被捕获、阻塞或忽略)
	// SIGSTOP 17,19,23  停止进程(不能被捕获、阻塞或忽略)
	for {
		fmt.Println("asdasd")
		sl := <-ch
		switch sl {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT, syscall.SIGHUP:
			quit(sl)
			return
		default:
			return
		}
	}
}

func quit(sl os.Signal) {
	log.Printf("get a signal %s, 程序停止", sl.String())
	close(ch)
	log.Println("Work End...")
}
