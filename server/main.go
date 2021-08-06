package main

import(
	"fmt"
	"context"
	"proto"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	server := grpc.NewServer(opts...)
	pb.RegisterIncrementServer(server, newServer())
	server.Serve(lis)
}

func (s * server) Increment(ctx context.Context, req *proto.Request) (*proto.Response)error{
	name := req.GetName()

	str := "not implemented yet. " + name + " will implement me"

	proto.Response{result : str}

	return &proto.Response{result : str}, nil
}

