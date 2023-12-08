package main

import (
	"fmt"
	"net"
	"runtime"
)

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		fmt.Println("------for------")
		n, _ := conn.Write([]byte("你好!!!"))
		fmt.Println("发送数据到client:", n)
		runtime.Goexit()
	}

	fmt.Println("-----process end----")
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	for {
		fmt.Println("---------start-------")
		conn, err := listen.Accept()
		// 等待客户端连接
		fmt.Println("---------cccccc-------")
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}

		fmt.Println("---------建立连接成功-------")
		go process(conn) // 启动一个goroutine处理连接
	}
}
