package main

import (
	"client/TLS"
	pb "client/ecommerce"
	"client/question"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	opts := TLS.CreateOp()
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}
	defer conn.Close()
	c := pb.NewManageClient(conn)

	clientDeadline := time.Now().Add(time.Duration(2 * time.Second))
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	question.Operation(ctx, c)
}
