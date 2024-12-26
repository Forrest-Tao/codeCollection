package main

import (
	"context"
	"fmt"
	pb "forrest/codeCollection/grpc/grpc-on-k8s/payment/pb"

	"google.golang.org/grpc"
	"log"
	"net"
)

type PaymentServiceServer struct {
	pb.UnimplementedPaymentServiceServer
}

func (s *PaymentServiceServer) GetPayments(context.Context, *pb.PaymentRequest) (*pb.Payments, error) {
	return &pb.Payments{
		PaymentList: []string{"pay-1", "pay-2"},
	}, nil
}

func main() {
	port := ":50003"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPaymentServiceServer(s, &PaymentServiceServer{})
	fmt.Printf("PaymentService listening on port %s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
