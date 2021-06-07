//go:generate protoc --proto_path ../helloworld --go_out=paths=source_relative:../helloworld --go-grpc_out=paths=source_relative:../helloworld ../helloworld/helloworld.proto
//go:generate mkdir -p ../helloworld/third_party/github.com/envoyproxy/protoc-gen-validate/validate/
//go:generate curl -sSL0 https://raw.githubusercontent.com/envoyproxy/protoc-gen-validate/b212fd4217dfe09d85069051d9c71fcc2089d486/validate/validate.proto -o ../helloworld/third_party/github.com/envoyproxy/protoc-gen-validate/validate/validate.proto
package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	pb "github.com/zerospiel/proto_workshop/reflection/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (*server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp4", ":55001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close() //nolint

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)

	go s.Serve(lis) //nolint

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	s.Stop() //nolint
}
