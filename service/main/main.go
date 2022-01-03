package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-easy/TLS"
	pb "grpc-easy/ecommerce"
	"grpc-easy/service"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {
	opt := TLS.CreateOp()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer(opt...)
	pb.RegisterManageServer(s, &service.Manage{})

	log.Printf("Starting grpc listener on port" + port)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
