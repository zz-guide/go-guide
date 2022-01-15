#/bin/sh

protoc ./proto/*.proto \
--go_out=. \
--go-grpc_out=.