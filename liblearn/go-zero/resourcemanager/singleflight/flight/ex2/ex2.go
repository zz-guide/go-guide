package main

import (
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/singleflight"
)

var errorNotExist = errors.New("not exist")

var g singleflight.Group

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
		// 并行获取
		v, err, _ := g.Do(key, func() (interface{}, error) {
			data, err = getDataFromDB(key)
			if err == nil {
				// TODO::数据库中有值，设置到缓存，并且返回
				return data, nil
			}

			fmt.Println("缓存失效，从数据库获取失败")
			return "", err
		})

		if err == nil {
			data = v.(string)
			return data, err
		}

		return "", nil
	}

	return "", nil
}

// getDataFromCache 模拟从cache中获取值，cache中无该值
func getDataFromCache(key string) (string, error) {
	fmt.Println("从缓存中获取")
	return "", errorNotExist
}

// getDataFromDB 模拟从数据库中获取值
func getDataFromDB(key string) (string, error) {
	fmt.Println("从数据库中获取")
	return "success", nil
}
