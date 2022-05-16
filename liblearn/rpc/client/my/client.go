package main

import (
	"context"
	"google.golang.org/grpc/credentials"
	"log"

	"google.golang.org/grpc"

	"go-guide/liblearn/rpc/pb/my"
)

func main() {
	// 创建证书对象
	creds, err := credentials.NewClientTLSFromFile("../../keys/server.crt", "xulei1.com")
	if err != nil {
		log.Fatal(err)
	}

	// 连接grpc服务
	// grpc默认是使用加密传输，grpc.WithInsecure()表示不使用证书
	conn, err := grpc.Dial(":9555", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	// 创建客户端,不加密的
	myClient := my.NewMyRPCServiceClient(conn)

	//获得商品库存
	res, err := myClient.GetMy(context.Background(), &my.MyRequest{Id: 226})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("结果:%+v\n", res)
}
