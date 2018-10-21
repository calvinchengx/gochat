#!/bin/bash

# build our stringer tool
go get -v golang.org/x/tools/cmd/stringer
go build -v golang.org/x/tools/cmd/stringer

# build our mockgen tool
go get github.com/golang/mock/gomock
go build -v github.com/golang/mock/mockgen

# build our protoc-gen-go tool
go get -v github.com/golang/protobuf/protoc-gen-go
go build -v github.com/golang/protobuf/protoc-gen-go

# set path so that we can invoke our installed tools
export PATH=$PATH:$PWD

# prepare the protobuf directory - golang
mkdir -p pb
protoc -I=proto \
    --go_out=plugins=grpc,import_path=pb:pb/ \
    proto/*.proto

# prepare the protobuf directory - web/javascript
mkdir -p client/web/pb
protoc -I=proto \
    --js_out=import_style=commonjs:client/web/pb \
    --grpc-web_out=import_style=commonjs,mode=grpcwebtext:client/web/pb \
    proto/*.proto

# remove our protoc-gen-go tool
rm protoc-gen-go

# invoke our mockgen script and stringer generation commands defined in hedera.go
go generate -v ./...

# remove our stringer tool and mockgen tools
rm stringer mockgen
