package main

import "fmt"

func main() {
	x := []int{1, 2, 3}

	y := x[:2]
	y = append(y, 50)
	fmt.Println("x:", x, ";y:", y)
	y = append(y, 60)
	fmt.Println("x:", x, ";y:", y)
}
