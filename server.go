package main

import (
	"awesomeProject1/metadata/protos"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
)

var host = "127.0.0.1"

var (
	ServiceName = flag.String("ServiceName","hello_service","service name")
	Port = flag.Int("Port",50001,"listening port")
)

type server struct {
}

func main() {
	flag.Parse()

	lis,err := net.Listen("tcp",fmt.Sprintf("127.0.0.1: %d",*Port))
	if err != nil {
		log.Fatalf("failed to listen:%s",err)
	} else {
		fmt.Printf("listen at :%d\n",*Port)
	}

	defer lis.Close()

	s := grpc.NewServer()
	defer s.GracefulStop()

	protos.RegisterGreeterServer(s, &server{})
	addr := fmt.Sprintf("%s:%d",host,*Port)
	fmt.Printf("server add:%s\n",addr)

	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to server: %s",err)
	}
}


func (s *server) SayHello (ctx context.Context,in *protos.HelloRequest) (*protos.HelloReply, error) {
	md,ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Printf("get metadata error")
	}
	if t,ok := md["timestamp"]; ok {
		fmt.Printf("timestamp from metadata:\n")
		for i,e := range t {
			fmt.Printf(" %d. %s\n",i,e)
		}
	}
	if t1,ok1 := md["key1"]; ok1 {
		fmt.Printf("key1 from metadata:\n")
		for i,e := range t1 {
			fmt.Printf(" %d . %s\n",i,e)
		}
	}
	if len(md) > 0 {
		for k,v := range md {
			fmt.Printf("%v:%v\n",k,v)
		}
	}
	return &protos.HelloReply{Message:"server: "+ in.Name},nil
}