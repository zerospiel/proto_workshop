package main

import (
	"context"
	"fmt"
	"net"

	foobar "github.com/zerospiel/proto_workshop/reflect_package/pkg/foobar"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

func _() {
	var (
		_ foobar.BarServiceServer = (*srv)(nil)
		_ foobar.FooServiceServer = (*srv)(nil)
	)
}

type (
	srv struct {
		foobar.UnimplementedBarServiceServer
		foobar.UnimplementedFooServiceServer
	}
)

func (s *srv) BarMethod(ctx context.Context, req *emptypb.Empty) (*foobar.MethodNameResponse, error) {
	return &foobar.MethodNameResponse{
		F: foobar.BarService_ServiceDesc.ServiceName + ": " + foobar.BarService_ServiceDesc.Methods[0].MethodName,
		T: map[string]float64{},
	}, nil
}

func (s *srv) FooMethod(ctx context.Context, req *foobar.Empty) (*foobar.MethodNameResponse, error) {
	return &foobar.MethodNameResponse{
		F: foobar.FooService_ServiceDesc.ServiceName + ": " + foobar.FooService_ServiceDesc.Methods[0].MethodName,
		T: map[string]float64{},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:55500")
	if err != nil {
		panic(err)
	}

	impl := &srv{}
	server := grpc.NewServer()
	foobar.RegisterFooServiceServer(server, impl)
	foobar.RegisterBarServiceServer(server, impl)

	reflection.Register(server)

	fmt.Println("started to listen", lis.Addr().String())

	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}
