package user

import (
	"go-guide/liblearn/rpc/pb/user"
	"io"
	"log"
)

type UserService struct {
	user.UnimplementedUserServiceServer
}

func NewUserServer() *UserService {
	return &UserService{}
}

// GetUserByServerStream 服务端流模式
func (UserService) GetUserByServerStream(in *user.UserRequest, stream user.UserService_GetUserByServerStreamServer) error {
	users := make([]*user.UserInfo, 0)
	for _, item := range in.Users {
		users = append(users, item)

		//每隔两条发送
		if len(users) == 2 {
			err := stream.Send(&user.UserResponse{Users: users})
			log.Println("---已发送---")
			if err != nil {
				return err
			}
			//清空切片
			users = (users)[0:0]

		}

		//time.Sleep(2 * time.Second)
	}

	if len(users) > 0 {
		return stream.Send(&user.UserResponse{Users: users})
	}

	return nil
}

// GetUserByClientStream 客户端流模式
func (UserService) GetUserByClientStream(stream user.UserService_GetUserByClientStreamServer) error {
	users := make([]*user.UserInfo, 0)
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF { //接收完毕
				return stream.SendAndClose(&user.UserResponse{Users: users})
			}

			return err
		}

		for _, item := range req.Users {
			item.Name += "--处理" //这里好比是服务端做的业务处理
			users = append(users, item)
		}
	}

	return nil
}
