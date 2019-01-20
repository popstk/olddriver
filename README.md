## 前端
- vue
- webpack

## 后端
- 存储系统 - mongodb
- 爬虫 - colly
- 查询服务 - grpc

## build
####  ProtoBuf 生成
生成grpc-web相关的文件

```bash
# 生成protobuf message classes
protoc backend\backend.proto --js_out=import_style=commonjs:.  
# 生成gRPC-Web service client stub
protoc backend\backend.proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. 
```

生成go相关文件

```bash
protoc backend\backend.proto --go_out=plugins=grpc:.
```

#### 爬虫
``` bash
make taohua
```

#### rpc服务
``` bash
make server
make client
```

