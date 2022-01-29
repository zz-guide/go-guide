package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {

}

func CreateTopic() {
	topic := "bingo"
	nameSrvAddr := []string{"127.0.0.1:9876"}
	brokerAddr := "127.0.0.1:10911"

	testAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)))
	if err != nil {
		panic(err)
	}

	// 创建topic
	err = testAdmin.CreateTopic(
		context.Background(),
		admin.WithTopicCreate(topic),
		admin.WithBrokerAddrCreate(brokerAddr))
	if err != nil {
		fmt.Println("Create topic error:", err)
	}

	err = testAdmin.Close()
	if err != nil {
		fmt.Println("Shutdown admin error:", err)
	}
}

func DeleteTopic() {
	topic := "bingo"
	nameSrvAddr := []string{"127.0.0.1:9876"}
	//brokerAddr := "127.0.0.1:10911"

	testAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)))
	if err != nil {
		panic(err)
	}

	// 删除topic
	err = testAdmin.DeleteTopic(
		context.Background(),
		admin.WithTopicDelete(topic),
		//admin.WithBrokerAddrDelete(brokerAddr),
		//admin.WithNameSrvAddr(nameSrvAddr),
	)

	err = testAdmin.Close()
	if err != nil {
		fmt.Println("Shutdown admin error:", err)
	}
}
