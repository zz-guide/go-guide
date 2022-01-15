GRPC学习
官网：https://grpc.io/docs/languages/go/quickstart/

#1.安装依赖
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
export PATH="$PATH:$(go env GOPATH)/bin"
go get -u google.golang.org/grpc


2.编写好proto文件以后，生成代码,进入到项目根目录下执行
protoc ./proto/*.proto \
--go_out=. \
--go-grpc_out=. 


–proto_path或者-I ：指定 import 路径，可以指定多个参数，编译时按顺序查找，不指定时默认查找当前目录。
.proto 文件中也可以引入其他 .proto 文件，这里主要用于指定被引入文件的位置。
–go_out：golang编译支持，指定输出文件路径
其他语言则替换即可，比如 --java_out 等等
–go_opt：指定参数，比如--go_opt=paths=source_relative就是表明生成文件输出使用相对路径。
path/to/file.proto ：被编译的 .proto 文件放在最后面

3.当main包中有多个文件时,需要变化执行方式
go run main.go a.go ....
go run *.go

4.proto3 语法文档地址：https://developers.google.cn/protocol-buffers/docs/proto3

5.ssl
.openssl
进入命令行
openssl genrsa -des3 -out server.key 2048 (会生成server.key私钥文件)
密码：xulei.com(输入2次，一次验证)

创建证书请求文件(生成server.csr)：
openssl req -new -key server.key -out server.csr
(全部为空就行)
删除密码
openssl rsa -in server.key -out server_no_passwd.key

生成服务端证书
openssl x509 -req -days 365 -in server.csr -signkey server_no_passwd.key -out server.crt

拷贝server.crt和server_no_passwd.key到server/keys目录

