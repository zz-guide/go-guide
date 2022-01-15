package main

import (
	"crypto/rand"
	"fmt"
)

/**
crypto/rand包实现了用于加解密的更安全的随机数生成器。
该包中常用的是 func Read(b []byte) (n int, err error) 这个方法， 将随机的byte值填充到b 数组中，以供b使用。
*/
func main() {
	b := make([]byte, 20)
	fmt.Println(b) //

	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(b)
}
