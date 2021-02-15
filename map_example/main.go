package main

import (
	"context"
	"fmt"
	"net"

	foobar "github.com/zerospiel/proto_workshop/map_example/external"
	baz "github.com/zerospiel/proto_workshop/map_example/pkg/baz"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

func _() {
	var _ baz.BazServiceServer = (*srv)(nil)
}

type (
	srv struct {
		baz.UnimplementedBazServiceServer
	}
)

func (s *srv) MethodName(ctx context.Context, req *emptypb.Empty) (*foobar.FooResponse, error) {
	return &foobar.FooResponse{Echo: "Server Echo"}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:55500")
	if err != nil {
		panic(err)
	}

	impl := &srv{}
	server := grpc.NewServer()
	baz.RegisterBazServiceServer(server, impl)

	reflection.Register(server)

	fmt.Println("started to listen", lis.Addr().String())

	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}
