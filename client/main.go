package main

// packages
import(
	"fmt"
	"context"
	"google.golang.org/grpc"
	"proto"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
)

func main(){
	// parameters for client stub creation
	var opts []grpc.DialOption
	...

	// creating a connection on port 9090
	conn, err := grpc.Dial("localhost:9090", opts..., grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor))
	if err != nil {
  		panic(err)
	}
	defer conn.Close()

	client := proto.NewRouteGuideClient(conn)

	// passing the unary requests and logging it 
	feature, err := client.Increment(context.Background(), &proto.Request{"chethan", 0})
	if err != nil {
  		panic(err)
	}

	else{
		log.println(feature)
	}
}