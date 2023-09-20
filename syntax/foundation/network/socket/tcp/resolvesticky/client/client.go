package main

import (
	"fmt"
	"go-guide/base/type/network/socket/tcp/resolvesticky/protocol"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 2; i++ {
		msg := `HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!HelloWorld!` // header 4 byte+body 12 byte = 16byte * 2 = 32 byte
		data, err := protocol.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
