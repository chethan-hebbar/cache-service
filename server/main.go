package main

// packages
import(
	"fmt"
	"log"
	"context"
	"proto"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
    "github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
)

type server struct{
	...
}

func main(){
	// setting up the server, chose some arbitrary port
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
  		panic(err)
	}

	// parameters for server creation
	var opts []grpc.ServerOption
	...
	server := grpc.NewServer(opts..., grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor), grpc_prometheus.EnableHandlingTimeHistogram())

	proto.RegisterIncrementServer(server, newServer())
	server.Serve(lis)

	// logging the metrics in the form of a histogram
	log.Print(grpc_server_handling_seconds)
}

// increment procedure with custom metrics
func (s * server) Increment(ctx context.Context, req *proto.Request) (*proto.Response)error{
	name := req.GetName()
	val := req.GetVal()

	// creating a new counter if not created already
	if val == 0{
		val = 1
	}

	else{
		updateCounter()
		val = val + 1
	}

	// functionality given in the docs
	str := "not implemented yet. " + name + " will implement me"

	proto.Response{result : str}

	return &proto.Response{result : str}, nil
}

// counter 
var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
			Name: "custom metric",
			Help: "keeps track of the counter",
	})
)

// updating the counter
func updateCounter() {
	go func() {
			for {
					opsProcessed.Inc()
					time.Sleep(1 * time.Second)
			}
	}()
}
