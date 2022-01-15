package main

import "src/command"

/**
go build 不加参数体积过大，
‘-s’ 相当于strip掉符号表， 但是以后就没办法在gdb里查看行号和文件了。
‘-w’ flag to the linker to omit the debug information 告知连接器放弃所有debug信息

//①打包静态文件到go文件中，go-bindata -o=./asset/asset.go -pkg=asset static/...
//②go build -ldflags "-s -w" -o ddb main.go
//③upx -9 ddb
//④运行：./ddb 按照提示输入参数即可
//⑤因为前端使用ajax请求，并且是file协议，所以需要放置到服务器运行

*/
func main() {
	command.Start()
}
