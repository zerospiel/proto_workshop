//go:generate protoc --proto_path ../helloworld --go_out=paths=source_relative:../helloworld --go-grpc_out=paths=source_relative:../helloworld ../helloworld/helloworld.proto
package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"time"

	pb "github.com/zerospiel/proto_workshop/intro/helloworld"
	"google.golang.org/grpc"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	ports = []string{":10001", ":10002", ":10003"}
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (*server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

type slowServer struct {
	pb.UnimplementedGreeterServer
}

func (*slowServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	time.Sleep(time.Duration(100+rand.Intn(100)) * time.Millisecond)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	var listeners []net.Listener
	var svrs []*grpc.Server
	for i := 0; i < 3; i++ {
		lis, err := net.Listen("tcp", ports[i])
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		listeners = append(listeners, lis)
		s := grpc.NewServer()
		svrs = append(svrs, s)
		if i == 2 {
			pb.RegisterGreeterServer(s, &slowServer{})
		} else {
			pb.RegisterGreeterServer(s, &server{})
		}
		go s.Serve(lis) //nolint
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	for i := 0; i < 3; i++ {
		svrs[i].Stop()
		listeners[i].Close() //nolint
	}
}
