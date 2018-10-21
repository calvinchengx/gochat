
[![CircleCI](https://circleci.com/gh/calvinchengx/gochat/tree/master.svg?style=shield)](https://circleci.com/gh/calvinchengx/gochat/tree/master)

golang chat with protocol buffers and gRPC
===

```bash
# server
go run server/main.go

# client - golang
go run client/go/main.go

# client - javascript/web
# TODO: add javascript/web usage instructions
```

dependencies for generating protobuf for web js
===

```bash
mkdir -p ~/opensource && cd ~/opensource
git clone https://github.com/grpc/grpc-web
cd grpc-web
sudo make install-plugin
# this will install protoc-gen-grpc-web binary in our /usr/local/bin
```

