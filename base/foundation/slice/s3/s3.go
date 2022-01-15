package main

import (
	"container/list"
	"fmt"
)

func main() {
	queueExample()
}

// stackExample
// 模拟栈的先进后出
//**
func stackExample(){
	stack := list.New()

	stack.PushBack(1)
	stack.PushBack(2)
	stack.PushBack(3)

	num1 := stack.Remove(stack.Back())
	fmt.Println("num1:", num1)

	num2 := stack.Remove(stack.Back())
	fmt.Println("num2:", num2)
}

// queueExample
// 模拟队列的先进先出
//**
func queueExample(){
	stack := list.New()

	stack.PushBack(1)
	stack.PushBack(2)
	stack.PushBack(3)

	num1 := stack.Remove(stack.Front())
	fmt.Println("num1:", num1)

	num2 := stack.Remove(stack.Front())
	fmt.Println("num2:", num2)
}