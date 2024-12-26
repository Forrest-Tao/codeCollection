package main

import (
	"context"
	"fmt"

	pb "forrest/codeCollection/grpc/grpc-on-k8s/order/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type OrderServiceServer struct {
	pb.UnimplementedOrderServiceServer
}

func (o *OrderServiceServer) GetOrders(context.Context, *pb.OrderRequest) (*pb.Orders, error) {
	return &pb.Orders{
		OrderList: []string{"order-1", "order-2"},
	}, nil
}

func main() {
	port := ":50002"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &OrderServiceServer{})
	fmt.Printf("OrderServiceServer listening on port %s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
