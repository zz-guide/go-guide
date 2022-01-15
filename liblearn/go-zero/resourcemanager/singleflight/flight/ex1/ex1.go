package main

import (
	"errors"
	"fmt"
	"sync"
)

var errorNotExist = errors.New("not exist")

/**
	结论：当缓存中值没有的时候，使用常规方法会导致请求都到数据库，给数据库造成过大压力。理想状况下我们希望只有一个请求去
获取数据库的值然后写到缓存，并把结果同步到其他的同类型请求中（同类型请求指的是大家都在等请求数据库并把结果写到缓存）
*/
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data, err := getData("key")
			if err != nil {
				fmt.Println("错误:", err)
				return
			}

			fmt.Println("-------结果:", data)
		}()
	}

	wg.Wait()
	fmt.Println("----执行完毕----")
}

// getData 获取数据
func getData(key string) (string, error) {
	data, err := getDataFromCache(key)
	// 命中缓存直接返回
	if err == nil {
		return data, nil
	}

	// 缓存击穿或者穿透
	if err == errorNotExist {
		// 先从数据库获取
		data, err = getDataFromDB(key)
		if err == nil {
			// 数据库中有值，设置到缓存，并且返回
			return data, nil
		}

		fmt.Println("缓存失效，从数据库获取失败")
		return "", err
	}

	return data, nil
}

// getDataFromCache 模拟从cache中获取值，cache中无该值
func getDataFromCache(key string) (string, error) {
	return "", errorNotExist
}

// getDataFromDB 模拟从数据库中获取值
func getDataFromDB(key string) (string, error) {
	fmt.Println("从数据库中获取")
	return "success", nil
}
