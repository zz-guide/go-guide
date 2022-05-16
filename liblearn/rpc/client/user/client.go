package main

import (
	"context"
	"go-guide/liblearn/rpc/pb/user"
	"google.golang.org/grpc/credentials"
	"io"
	"log"
	"strconv"

	"google.golang.org/grpc"
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
	userClient := user.NewUserServiceClient(conn)
	_serverStream(userClient)
}

// 客户端流模式
func _clientStream(userClient user.UserServiceClient) {
	var i int32

	stream, err := userClient.GetUserByClientStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 总共发3次
	for j := 1; j <= 3; j++ {
		req := user.UserRequest{}
		req.Users = make([]*user.UserInfo, 0)
		// 每次5条
		for i = 1; i <= 5; i++ {
			_i := strconv.FormatInt(int64(i), 32)
			req.Users = append(req.Users, &user.UserInfo{Id: i, Name: "许磊" + string(_i)})
		}

		err := stream.Send(&req)
		if err != nil {
			log.Println(err)
		}
	}

	res, err := stream.CloseAndRecv() // 发送EOF
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("结果:%+v\n", res.Users)
}

// 服务端流模式
func _serverStream(userClient user.UserServiceClient) {
	req := user.UserRequest{}
	req.Users = make([]*user.UserInfo, 0)
	var i int32
	for i = 1; i < 20; i++ {
		_i := strconv.FormatInt(int64(i), 32)
		req.Users = append(req.Users, &user.UserInfo{Id: i, Name: "许磊" + string(_i)})
	}

	//获得商品库存
	stream, err := userClient.GetUserByServerStream(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	//需要循环读取，服务器停止才退出
	for {
		res, err := stream.Recv()
		if err != nil {
			//服务端发送结束
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
			}
		}

		log.Printf("结果:%+v\n", res.Users)
	}

	log.Println("----接收结束----")
}
