package main

import (
	"awesomeProject1/metadata/protos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

const (
	timestampFormat = time.StampNano
)

func main() {
	conn,err := grpc.Dial("127.0.0.1:50001",grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := protos.NewGreeterClient(conn)
	md := metadata.Pairs("timestamp",time.Now().Format(timestampFormat))
	md = metadata.New(map[string]string{"key1":"val1","key2":"val2"})
	ctx := metadata.NewOutgoingContext(context.Background(),md)
	resp,err := client.SayHello(ctx,&protos.HelloRequest{Name:"Hello"})
	if err == nil {
		fmt.Printf("Reply is : %s\n",resp.Message)
	} else {
		fmt.Printf("call server error:%s\n",err)
	}

}
