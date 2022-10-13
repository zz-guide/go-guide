package main

import "log"

func main() {
	a := 1
	b := 2
	res := Max(a, b)
	log.Println("res:", res)
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
