# 安装grpc 

## 1. 先安装 protoc
```
 https://github.com/protocolbuffers/protobuf/releases
```
### 2. 安装  protoc-gen-go
```shell
go get -u github.com/golang/protobuf/protoc-gen-go
```

### 3. 安装grpc
```shell
go get -u google.golang.org/grpc
```

### 4. 安装 grpc-gateway
```shell
go get -u \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
### 5. buf 安装
```shell
brew tap bufbuild/buf
brew install buf
```
运行
```
buf generate protos --debug
```