# metadataGrpcDemo

在http请求当中我们可以设置header用来传递数据，grpc底层采用http2协议也是支持传递数据的，采用的是metadata。 
Metadata 对于 gRPC 本身来说透明， 它使得 client 和 server 能为对方提供本次调用的信息。
就像一次 http 请求的 RequestHeader 和 ResponseHeader，http header 的生命周期是一次 http 请求， 
Metadata 的生命周期则是一次 RPC 调用。



1.创建proto文件；
2.执行 protoc --go_out=plugins=grpc:. demo.proto  生成 demo.pb.go 文件
3.编写server.go
4.编写client.go
进入到文件夹目录
5.执行 go run server.go
6.执行 go run client.go
