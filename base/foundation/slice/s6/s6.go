package main

import (
	"fmt"
	"log"
)

func main() {
	pointer()
}

func references() {
	x := []int{1, 2, 3}

	y := x[:2]
	y = append(y, 50)
	fmt.Println("x:", x, ";y:", y)
	y = append(y, 60)
	fmt.Println("x:", x, ";y:", y)
}

func pointer() {
	arr := make([]int, 3, 5)
	arr[0] = 1
	arr[1] = 2
	//arr = append(arr, 1, 2)
	log.Printf("arr 长度:%d 容量:%d\n", len(arr), cap(arr))
	arr1 := _fff(arr)
	log.Printf("arr: %p, %v \n", arr, arr)
	log.Printf("arr: %p \n", &arr)
	log.Printf("arr1: %p  %v \n", arr1, arr1)
	log.Printf("arr1: %p \n", &arr1)
}

func _fff(nums []int) []int {
	//nums[0] = 5
	nums = append(nums, 4)
	log.Printf("nums: %p \n", nums)
	log.Printf("nums: %p \n", &nums)
	return nums
}
