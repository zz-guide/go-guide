package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	// 创建
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s6\n", u1)

	// 解析
	u2, err := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
	if err != nil {
		fmt.Printf("Something gone wrong: %s6", err)
		return
	}
	fmt.Printf("Successfully parsed: %s6", u2)
}
