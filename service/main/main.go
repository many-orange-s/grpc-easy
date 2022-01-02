package main

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	pb "grpc-easy/ecommerce"
	"grpc-easy/service"
	"io/ioutil"
	"log"
	"net"
)

const (
	port    = ":50051"
	crtFile = "D:\\go_project\\keys\\server.crt"
	keyFile = "D:\\go_project\\keys\\server.key"
	caFile  = "D:\\go_project\\keys\\ca.crt"
)

func main() {
	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load key part : %s", err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalf("could not read ca certificate :%s", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("faild to append ca certificate")
	}

	opt := []grpc.ServerOption{
		grpc.Creds(credentials.NewTLS(&tls.Config{
			ClientAuth:   tls.RequireAnyClientCert,
			Certificates: []tls.Certificate{cert},
			ClientCAs:    certPool,
		})),
	}

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
