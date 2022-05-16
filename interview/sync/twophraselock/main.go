package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

const WorkerRepeat = 10000
const M = 2
const N = 100000

var S [N]int

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 准备四把锁
	rwLockI := sync.RWMutex{}
	rwLockI1 := sync.RWMutex{}
	rwLockI2 := sync.RWMutex{}
	rwLockJ := sync.RWMutex{}

	var wg sync.WaitGroup
	for i := 0; i < M; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < WorkerRepeat; i++ {
				// 1.第一步骤，获取要操作的所有元素索引
				tempI := rand.Intn(N)
				tempJ := rand.Intn(N)
				tempI1 := tempI + 1
				tempI2 := tempI + 2
				if tempI1 >= N {
					tempI1 = (tempI + 1) % N
				}

				if tempI2 >= N {
					tempI2 = (tempI + 2) % N
				}

				var tem = -1

				// 2.读取i，i+1，i+2三个位置的值
				rwLockI.RLock()
				tem += S[tempI]
				rwLockI.RUnlock()

				rwLockI1.RLock()
				tem += S[tempI1]
				rwLockI1.RUnlock()

				rwLockI2.RLock()
				tem += S[tempI2]
				rwLockI2.RUnlock()
				// 3.加锁修改j位置的值
				rwLockJ.Lock()
				S[tempJ] = tem
				rwLockJ.Unlock()
			}
		}()
	}

	wg.Wait()
	log.Println("完毕")
}
