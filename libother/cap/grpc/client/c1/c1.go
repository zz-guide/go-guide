package main

import (
	"context"
	"log"
	"time"

	pb "go-guide/lib/grpc/pb/student"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/health"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:8888",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(`{
	"loadBalancingPolicy": "round_robin",
	"healthCheckConfig": {
		"serviceName": ""
	}
}`),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStudentClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	r, err := c.Detail(ctx, &pb.StudentReq{Id: 2})
	if err != nil {
		log.Fatalf("could not detail: %v", err)
	}
	log.Printf("接收信息: %s", r.String())
}
