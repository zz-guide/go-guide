package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// 转换可能溢出，溢出了的话就表示的不对了，不同的a经过BytesToInt(IntToBytes(a))可能值一样
	a := 255
	b := 256
	c := 257
	d := 1
	e := 258
	fmt.Println("a:", a, IntToBytes(a), BytesToInt(IntToBytes(a)))
	fmt.Println("b:", b, IntToBytes(b), BytesToInt(IntToBytes(b)))
	fmt.Println("c:", c, IntToBytes(c), BytesToInt(IntToBytes(c)))
	fmt.Println("d:", d, IntToBytes(d), BytesToInt(IntToBytes(d)))
	fmt.Println("e:", e, IntToBytes(e), BytesToInt(IntToBytes(e)))
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := uint8(n)
	var bytesBuffer = new(bytes.Buffer)
	err := binary.Write(bytesBuffer, binary.BigEndian, x)
	if err != nil {
		return nil
	}

	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data uint8
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}
