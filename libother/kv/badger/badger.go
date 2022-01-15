package badger

import (
	"fmt"
	"kv/badger/cacheObj"
	"log"
)

func Work() {
	db := cacheObj.NewBadgerDB()
	err := db.Open("./data/badger", true)
	if err != nil {
		log.Fatal("err:", err)
	}
	defer db.Close()

	priKey := []byte("user1")
	writeData := []byte("你好")

	_ = db.Set(priKey, writeData)
	fmt.Println("Set:", string(writeData))

	res, _ := db.Get(priKey)
	fmt.Println("Get:", string(res))
}
