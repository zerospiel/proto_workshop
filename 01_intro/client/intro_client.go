package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	pb "github.com/zerospiel/proto_workshop/intro/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
	"google.golang.org/grpc/status"
)

const (
	defaultName = "world"
)

func main() {
	const schemeName = "workshop"
	r := manual.NewBuilderWithScheme(schemeName)
	resolver.Register(r)
	defer resolver.UnregisterForTesting(schemeName)

	conn, err := grpc.Dial(
		r.Scheme()+":///test.server",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [ { "round_robin": {} } ]}`),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	r.UpdateState(resolver.State{
		Addresses: []resolver.Address{{Addr: ":10001"}, {Addr: ":10002"}, {Addr: ":10003"}},
	})

	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	for i := 0; i < 100; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
		defer cancel()
		p := peer.Peer{}
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name}, grpc.Peer(&p))
		if err != nil {
			serr := status.Convert(err)
			if serr.Code() != codes.DeadlineExceeded {
				log.Fatalf("unexpected error code %s: %v", serr.Code(), err)
			}
			log.Printf("could not get response: %v", serr.Err())
		} else {
			log.Printf("Response from %s (%s): %s", conn.Target(), p.Addr.String(), r.Message)
		}
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
