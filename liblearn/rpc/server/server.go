package main

import (
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	"go-guide/liblearn/rpc/pb/my"
	"go-guide/liblearn/rpc/pb/user"

	my2 "go-guide/liblearn/rpc/server/my"
	user2 "go-guide/liblearn/rpc/server/user"
)

func main() {
	_grpc()
	//_http()
}

func _grpc() {
	// 1.监听端口
	listener, err := net.Listen("tcp", ":9555")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("../keys/server.crt", "../keys/ca.key")
	if err != nil {
		log.Fatal(err)
	}

	// 2.创建grpc服务,提供了很多选项
	rpcServer := grpc.NewServer(grpc.Creds(creds))
	// 3.向上一步的rpcServer注册自己的服务
	//my.RegisterMyRPCServiceServer(rpcServer, new(my.UnimplementedMyRPCServiceServer))
	my.RegisterMyRPCServiceServer(rpcServer, my2.NewMyServer())
	user.RegisterUserServiceServer(rpcServer, user2.NewUserServer())

	log.Println("RPC服务启动,127.0.0.1:9555")
	if err := rpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func _http() {
	// 创建安全链接
	creds, err := credentials.NewServerTLSFromFile("../keys/server.crt", "../keys/ca.key")
	if err != nil {
		log.Fatal(err)
	}

	// 创建grpc服务,提供了很多选项
	rpcServer := grpc.NewServer(grpc.Creds(creds))
	// 3.向上一步的rpcServer注册自己的服务
	//my.RegisterMyRPCServiceServer(rpcServer, new(my.UnimplementedMyRPCServiceServer))
	my.RegisterMyRPCServiceServer(rpcServer, my2.NewMyServer())
	user.RegisterUserServiceServer(rpcServer, user2.NewUserServer())

	log.Println("RPC服务启动,127.0.0.1:9555")
	httpServer := &http.Server{Addr: ":9555"}
	// 为grpc提供http访问能力
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("http:", request)
		rpcServer.ServeHTTP(writer, request)
	})

	if err := httpServer.ListenAndServeTLS("../keys/server.crt", "../keys/ca.key"); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
