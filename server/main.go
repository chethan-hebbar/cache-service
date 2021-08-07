package main

import(
	"fmt"
	"log"
	"context"
	"proto"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
)

type server struct{
	...
}

func main(){
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
  		panic(err)
	}

	var opts []grpc.ServerOption
	...
	server := grpc.NewServer(opts..., grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor), grpc_prometheus.EnableHandlingTimeHistogram())

	proto.RegisterIncrementServer(server, newServer())
	server.Serve(lis)

	log.Print(grpc_server_handling_seconds)
}

func (s * server) Increment(ctx context.Context, req *proto.Request) (*proto.Response)error{
	name := req.GetName()

	str := "not implemented yet. " + name + " will implement me"

	proto.Response{result : str}

	return &proto.Response{result : str}, nil
}

