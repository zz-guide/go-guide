#!/usr/bin/env bash

:<<BLOCK


BLOCK

function protoGen() {
    cd /Users/xulei/jungle/golangworkspace/go-guide/liblearn/rpc/proto && \
    protoc --go_out=../pb/ --go-grpc_out=../pb user-server-stream.proto
}

protoGen