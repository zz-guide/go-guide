#!/usr/bin/env bash

:<<BLOCK

BLOCK


function rpcStart() {
   cd /Users/xulei/jungle/golangworkspace/go-guide/liblearn/rpc/server && \
   go run server.go
}

rpcStart