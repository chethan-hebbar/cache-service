package main

import(
	"fmt"
	"context"
	"google.golang.org/grpc"
	"proto"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
)

func main(){
	var opts []grpc.DialOption
	...
	conn, err := grpc.Dial("localhost:9090", opts..., grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor))
	if err != nil {
  		panic(err)
	}
	defer conn.Close()

	client := proto.NewRouteGuideClient(conn)

	feature, err := client.Increment(context.Background(), &proto.Request{"chethan"})
	if err != nil {
  		panic(err)
	}

	else{
		log.println(feature)
	}
}