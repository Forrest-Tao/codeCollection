package main

import (
	"context"
	"fmt"
	pb "forrest/codeCollection/grpc/grpc-on-k8s/user/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (u *UserServiceServer) GetUserInfo(context.Context, *pb.UserRequest) (*pb.UserInfo, error) {
	return &pb.UserInfo{
		UserId:   10,
		UserName: "Forresr",
	}, nil
}

func main() {
	port := ":50001"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &UserServiceServer{})
	fmt.Printf("UserServiceServer listening on port %s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
