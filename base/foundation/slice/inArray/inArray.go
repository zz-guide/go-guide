package main

import (
	"fmt"
	"unsafe"
)

// struct{}{}不占内存
func main() {
	fmt.Println(unsafe.Sizeof(struct{}{})) // 结果为 0
}

func inArrayOrSlice(arrOrSlice []int, ele int) bool {
	if len(arrOrSlice) == 0 {
		return false
	}

	for _, item := range arrOrSlice {
		if item == ele {
			return true
		}
	}

	return false
}
