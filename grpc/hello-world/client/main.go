package main

import (
	"context"
	"flag"
	pb "forrest/codeCollection/grpc/hello-world/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "127.0.0.1:8900", "the address to connect to")
	name = flag.String("name", "world", "Name to greet")
)

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		panic(err)
	}
	log.Printf("greeting from server with message %q\n", r.GetMessage())

	r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		panic(err)
	}
	log.Printf("greeting from server again with message %q", r.GetMessage())
}
