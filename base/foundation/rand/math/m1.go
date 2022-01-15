package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
“math/rand” 包实现了伪随机数生成器。也就是生成 整形和浮点型。

该包中根据生成伪随机数是是否有种子(可以理解为初始化伪随机数)，可以分为两类：

1、有种子。通常以时钟，输入输出等特殊节点作为参数，初始化。该类型生成的随机数相比无种子时重复概率较低。
2、无种子。可以理解为此时种子为1， Seek(1)。


按类型随机：
func (r *Rand) Int() int
func (r *Rand) Int31() int32
func (r *Rand) Int63() int64
func (r *Rand) Uint32() uint32
func (r *Rand) Float32() float32  // 返回一个取值范围在[0.0, 1.0)的伪随机float32值
func (r *Rand) Float64() float64  // 返回一个取值范围在[0.0, 1.0)的伪随机float64值


指定随机范围：
func (r *Rand) Intn(n int) int
func (r *Rand) Int31n(n int32) int32
func (r *Rand) Int63n(n int64) int64


*/
func main() {
	for i := 0; i < 10; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		fmt.Printf("%d ", r.Int31())
	}

	fmt.Println("")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", rand.Int31())
	}
}
