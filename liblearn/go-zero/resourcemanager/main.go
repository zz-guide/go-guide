package main

import (
	"fmt"
	"go-guide/lib/resourcemanager/manager"
)

var connManager = manager.NewResourceManager()

// manager并行
func main() {
	fmt.Printf("manager:%+v\n", connManager)
}
