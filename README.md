## 前端
- vue
- webpack

## 后端
- 存储系统 - mongodb
- 爬虫 - colly
- 查询服务 - grpc

## build
####  ProtoBuf 生成
安装[protoc](https://github.com/protocolbuffers/protobuf/releases)

生成go相关文件
```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/golang/protobuf/protoc-gen-go
protoc backend\backend.proto --go_out=plugins=grpc:.
```

#### 爬虫
``` bash
make taohua
```

#### 网关
[grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```

#### rpc服务
``` bash
make server
make client
```

