package my

import (
	"context"
	"go-guide/liblearn/rpc/pb/my"
)

type MyRPCService struct {
	// 新版本的RPC server都必须嵌入默认生成的未实现的结构体
	my.UnimplementedMyRPCServiceServer
}

func NewMyServer() *MyRPCService {
	return &MyRPCService{}
}

func (MyRPCService) GetMy(ctx context.Context, req *my.MyRequest) (*my.MyResponse, error) {
	return &my.MyResponse{
		Res: req,
	}, nil
}
