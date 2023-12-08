package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"runtime"
)

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		fmt.Println("------for------")
		reader := bufio.NewReader(conn)
		buf := make([]byte, 3)
		n, err := reader.Read(buf) // 读取数据
		fmt.Println("----读取完毕-----")
		if err == io.EOF {
			fmt.Println("没数据了，退出")
			runtime.Goexit()
		}

		if err != nil {
			fmt.Println("read from producer failed, err:", err)
			break
		}

		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据:", n, recvStr)
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
