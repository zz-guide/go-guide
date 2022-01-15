package main

import (
	"encoding/hex"
	"fmt"
	"hash/fnv"
)

func main() {
	a := fnv.New32()
	a.Write([]byte("hello"))
	fmt.Println(hex.EncodeToString(a.Sum(nil)))
}
