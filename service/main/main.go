package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grpc-easy/ecommerce"
	"grpc-easy/interceptor"
	"grpc-easy/service"
	"log"
	"net"
)

const (
	port    = ":50051"
	crtFile = "D:\\go_project\\grpc-easy\\service\\server.crt"
	keyFile = "D:\\go_project\\grpc-easy\\service\\server.key"
)

func main() {
	/*cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load key part : %s", err)
	}
	opt := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	*/
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(interceptor.MangeInterceptor))
	pb.RegisterManageServer(s, &service.Manage{})

	log.Printf("Starting grpc listener on port" + port)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
