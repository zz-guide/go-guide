1.rpc服务定义
2.流模式
3.验证proto文件
4.tls连接配置

#grpc-gateway使用
1.安装 buf
buf 用于代替 protoc 进行生成代码，可以避免使用复杂的 protoc 命令，避免 protoc 各种失败问题

brew tap bufbuild/buf
brew install buf

2.安装 grpc-gateway
go install \
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
google.golang.org/protobuf/cmd/protoc-gen-go \
google.golang.org/grpc/cmd/protoc-gen-go-grpc

This will place four binaries in your $GOBIN;

protoc-gen-grpc-gateway
protoc-gen-openapiv2
protoc-gen-go
protoc-gen-go-grpc
Make sure that your $GOBIN is in your $PATH.

3.添加 buf 配置文件 buf.gen.yaml
version: v1beta1
plugins:
- name: go
  out: proto
  opt: paths=source_relative
- name: go-grpc
  out: proto
  opt: paths=source_relative,require_unimplemented_servers=false
  
4.添加配置文件 buf.yaml
  version: v1beta1
  build:
  roots:
    - proto


#流模式
1.直接模式
一次性发送。在数据量大并且查询速度慢的情况下，如果一次性查完在返回，耗时太长。

2.客户端流

3.服务端流
分批次返回结果给客户端


4.双向流



