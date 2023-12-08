package main

import (
	"log"
	"unsafe"
)

func main() {
	byteToString()
}

func byteToString() {
	/**
	byte 转 string
	*/
	b := []byte("hello world 你好")
	s1 := byteToString1(b)
	log.Printf("s1: %s %p\n", s1, &s1)

	s2 := byteToString2(b)
	log.Printf("s2: %s %p\n", s2, &s2)

	s3 := byteToString3(b)
	log.Printf("s3: %s %p\n", s3, &s3)
}

func byteToString1(b []byte) string {
	return unsafe.String(&b[0], len(b))
}

func byteToString2(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func byteToString3(b []byte) string {
	return string(b)
}
