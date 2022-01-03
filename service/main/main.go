package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-easy/TLS"
	"grpc-easy/config"
	pb "grpc-easy/ecommerce"
	"grpc-easy/service"
	"log"
	"net"
)

func main() {
	config.Init()
	opt := TLS.CreateOp()
	lis, err := net.Listen("tcp", config.Con.Port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer(opt...)
	pb.RegisterManageServer(s, &service.Manage{})

	log.Printf("Starting grpc listener on port" + config.Con.Port)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
