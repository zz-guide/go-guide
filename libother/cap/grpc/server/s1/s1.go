package main

import (
	pb "go-guide/lib/grpc/pb/student"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthPb "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

var rpcServer = grpc.NewServer()
var server = &StudentRpcServer{}
var healthCheckServer = health.NewServer()
var HEALTHCHECK_SERVICE = ""

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 注册业务api
	RegisterStudentServer()
	// 注册health
	RegisterHealthCheckServer()

	log.Printf("server listening at %v", lis.Addr())
	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func RegisterStudentServer() {
	pb.RegisterStudentServer(rpcServer, server)
}

func RegisterHealthCheckServer() {
	healthCheckServer.SetServingStatus(HEALTHCHECK_SERVICE, healthPb.HealthCheckResponse_SERVING)
	healthPb.RegisterHealthServer(rpcServer, healthCheckServer)
}
