package main

import (
	"go-guide/datastruct/redBlackTree/node"
	"log"
	"math/rand"
	"time"
)

func main() {
	const count = 10
	redBlackTree := node.NewRedBlackTree()
	var nums []int
	for i := 0; i < count; i++ {
		nums = append(nums, rand.Intn(count))
	}

	log.Println("数据源:", nums)
	t := time.Now()
	for _, v := range nums {
		redBlackTree.Add(v, v)
	}

	log.Println("redBlackTree:", t.Sub(time.Now()))
	redBlackTree.PrintPreOrder()
	log.Println("节点数量:", redBlackTree.GetSize())
}
