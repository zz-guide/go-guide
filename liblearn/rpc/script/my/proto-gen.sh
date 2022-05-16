#!/usr/bin/env bash

:<<BLOCK

1.protoc 用来根据proto文件生成对应代码的命令
2.--go_out表示生成go代码的路径
#protoc --go_out=../pb/ *.proto

3.*.proto表示要解析的proto文件
4.--go_out=plugins=grpc:../pb *.proto 添加了grpc代码选项
#protoc --go_out=plugins=grpc:../pb *.proto
--go_out: protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC
新版本不支持了

BLOCK

function protoGen() {
    cd /Users/xulei/jungle/golangworkspace/go-guide/liblearn/rpc/proto && \
    protoc --go_out=../pb/ --go-grpc_out=../pb my.proto
}

protoGen