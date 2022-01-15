package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func IntToBytes(n int) []byte {
	data := int8(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int8
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}

func main() {
	a := 127
	fmt.Println(IntToBytes(a))
	fmt.Println(BytesToInt(IntToBytes(a)))
}
