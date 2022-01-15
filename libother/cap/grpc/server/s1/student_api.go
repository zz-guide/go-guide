package main

import (
	"context"
	pb "go-guide/lib/grpc/pb/student"
	"log"
)

type StudentRpcServer struct {
	pb.UnimplementedStudentServer
}

func (s *StudentRpcServer) Detail(ctx context.Context, in *pb.StudentReq) (*pb.StudentRes, error) {
	log.Printf("Received: %v", in.GetId())
	return &pb.StudentRes{Name: "许磊 ", Age: 22}, nil
}
